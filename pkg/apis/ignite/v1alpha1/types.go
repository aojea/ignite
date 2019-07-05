package v1alpha1

import (
	meta "github.com/weaveworks/ignite/pkg/apis/meta/v1alpha1"
)

// Image represents a cached OCI image ready to be used with Ignite
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Image struct {
	meta.TypeMeta `json:",inline"`
	// meta.ObjectMeta is also embedded into the struct, and defines the human-readable name, and the machine-readable ID
	// Name is available at the .metadata.name JSON path
	// ID is available at the .metadata.uid JSON path (the Go type is k8s.io/apimachinery/pkg/types.UID, which is only a typed string)
	meta.ObjectMeta `json:"metadata"`

	Spec ImageSpec `json:"spec"`
	//Status ImageStatus `json:"status"`
}

// ImageSpec declares what the image contains
type ImageSpec struct {
	Source ImageSource `json:"source"`
}

// ImageSourceType is an enum of different supported Image Source Types
type ImageSourceType string

const (
	// ImageSourceTypeDocker defines that the image is imported from Docker
	ImageSourceTypeDocker ImageSourceType = "Docker"
)

// ImageSource defines where the image was imported from
type ImageSource struct {
	// Type defines how the image was imported
	Type ImageSourceType `json:"type"`
	// ID defines the source's ID (e.g. the Docker image ID)
	ID string `json:"id"`
	// Name defines the user-friendly name of the imported source
	Name string `json:"name"`
	// Size defines the size of the source in bytes
	Size meta.Size `json:"size"`
}

// ImageStatus defines the status of the image
//type ImageStatus struct {
//	// LayerID points to the index of the device in the DM pool
//	LayerID meta.DMID `json:"layerID"`
//}

// Pool defines device mapper pool database
// This file is managed by the snapshotter part of Ignite, and the file (existing as a singleton)
// is present at /var/lib/firecracker/snapshotter/pool.json
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Pool struct {
	meta.TypeMeta `json:",inline"`
	// Not needed (yet)
	// meta.ObjectMeta `json:"metadata"`

	Spec   PoolSpec   `json:"spec"`
	Status PoolStatus `json:"status"`
}

// PoolSpec defines the Pool's specification
type PoolSpec struct {
	// MetadataSize specifies the size of the pool's metadata
	MetadataSize meta.Size `json:"metadataSize"`
	// DataSize specifies the size of the pool's data
	DataSize meta.Size `json:"dataSize"`
	// AllocationSize specifies the smallest size that can be allocated at a time
	AllocationSize meta.Size `json:"allocationSize"`
	// MetadataPath points to the file where device mapper stores all metadata information
	// Defaults to constants.SNAPSHOTTER_METADATA_PATH
	MetadataPath string `json:"metadataPath"`
	// DataPath points to the backing physical device or sparse file (to be loop mounted) for the pool
	// Defaults to constants.SNAPSHOTTER_DATA_PATH
	DataPath string `json:"dataPath"`
}

// PoolStatus defines the Pool's current status
type PoolStatus struct {
	// The Devices array needs to contain pointers to accommodate "holes" in the mapping
	// Where devices have been deleted, the pointer is nil
	Devices []*PoolDevice `json:"devices"`
}

type PoolDeviceType string

const (
	PoolDeviceTypeImage  PoolDeviceType = "Image"
	PoolDeviceTypeResize PoolDeviceType = "Resize"
	PoolDeviceTypeKernel PoolDeviceType = "Kernel"
	PoolDeviceTypeVM     PoolDeviceType = "VM"
)

// PoolDevice defines one device in the pool
type PoolDevice struct {
	Size   meta.Size `json:"size"`
	Parent meta.DMID `json:"parent"`
	// Type specifies the type of the contents of the device
	Type PoolDeviceType `json:"type"`
	// MetadataPath points to the JSON/YAML file with metadata about this device
	// This is most often of the format /var/lib/firecracker/{type}/{id}/metadata.json
	MetadataPath string `json:"metadataPath"`
}

// Kernel is a serializable object that caches information about imported kernels
// This file is stored in /var/lib/firecracker/kernels/{oci-image-digest}/metadata.json
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Kernel struct {
	meta.TypeMeta `json:",inline"`
	// meta.ObjectMeta is also embedded into the struct, and defines the human-readable name, and the machine-readable ID
	// Name is available at the .metadata.name JSON path
	// ID is available at the .metadata.uid JSON path (the Go type is k8s.io/apimachinery/pkg/types.UID, which is only a typed string)
	meta.ObjectMeta `json:"metadata"`

	Spec KernelSpec `json:"spec"`
	//Status KernelStatus `json:"status"`
}

// KernelSpec describes the properties of a kernel
type KernelSpec struct {
	Version string      `json:"version"`
	Source  ImageSource `json:"source"`
	// Optional future feature, support per-kernel specific default command lines
	// DefaultCmdLine string
}

// VM represents a virtual machine run by Firecracker
// These files are stored in /var/lib/firecracker/vm/{vm-id}/metadata.json
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type VM struct {
	meta.TypeMeta `json:",inline"`
	// meta.ObjectMeta is also embedded into the struct, and defines the human-readable name, and the machine-readable ID
	// Name is available at the .metadata.name JSON path
	// ID is available at the .metadata.uid JSON path (the Go type is k8s.io/apimachinery/pkg/types.UID, which is only a typed string)
	meta.ObjectMeta `json:"metadata"`

	Spec   VMSpec   `json:"spec"`
	Status VMStatus `json:"status"`
}

// VMSpec describes the configuration of a VM
type VMSpec struct {
	Image *ImageClaim `json:"image"`
	// TODO: Temporary ID for the old metadata handling
	Kernel   *KernelClaim      `json:"kernel"`
	CPUs     uint64            `json:"cpus"`
	Memory   meta.Size         `json:"memory"`
	DiskSize meta.Size         `json:"diskSize"`
	Ports    meta.PortMappings `json:"ports,omitempty"`
	// This will be done at either "ignite start" or "ignite create" time
	// TODO: We might to revisit this later
	CopyFiles []FileMapping `json:"copyFiles,omitempty"`
	// SSH specifies how the SSH setup should be done
	// SSH appends to CopyFiles when active
	// nil here means "don't do anything special"
	// An empty struct means "generate a new SSH key and copy it in"
	// Specifying a path means "use this public key"
	SSH *SSH `json:"ssh,omitempty"`
}

// ImageClaim specifies a claim to import an image
type ImageClaim struct {
	Type ImageSourceType `json:"type"`
	Ref  string          `json:"ref"`
	// TODO: Temporary ID for the old metadata handling
	UID meta.UID `json:"uid"`
}

// TODO: Temporary helper for the old metadata handling
type KernelClaim struct {
	UID     meta.UID `json:"uid"`
	CmdLine string   `json:"cmdline"`
}

// FileMapping defines mappings between files on the host and VM
type FileMapping struct {
	HostPath string `json:"hostPath"`
	VMPath   string `json:"vmPath"`
}

// SSH specifies different ways to connect via SSH to the VM
type SSH struct {
	PublicKey string `json:"publicKey,omitempty"`
}

// VMState defines different states a VM can be in
type VMState string

const (
	VMStateCreated VMState = "Created"
	VMStateRunning VMState = "Running"
	VMStateStopped VMState = "Stopped"
)

// VMStatus defines the status of a VM
type VMStatus struct {
	State       VMState          `json:"state"`
	IPAddresses meta.IPAddresses `json:"ipAddresses"`
}
