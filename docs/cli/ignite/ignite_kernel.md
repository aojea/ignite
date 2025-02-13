## ignite kernel

Manage VM kernels

### Synopsis


Groups together functionality for managing VM kernels.
Calling this command alone lists all available kernels.


```
ignite kernel [flags]
```

### Options

```
  -h, --help   help for kernel
```

### Options inherited from parent commands

```
      --log-level loglevel      Specify the loglevel for the program (default info)
      --network-plugin plugin   Network plugin to use. Available options are: [cni docker-bridge] (default docker-bridge)
  -q, --quiet                   The quiet mode allows for machine-parsable output by printing only IDs
```

### SEE ALSO

* [ignite](ignite.md)	 - ignite: easily run Firecracker VMs
* [ignite kernel import](ignite_kernel_import.md)	 - Import a kernel image from an OCI image
* [ignite kernel ls](ignite_kernel_ls.md)	 - List available VM kernels
* [ignite kernel rm](ignite_kernel_rm.md)	 - Remove kernels

