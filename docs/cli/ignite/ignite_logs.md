## ignite logs

Get the logs for a running VM

### Synopsis


Show the logs for the given VM. The VM needs to be running (its backing
container needs to exist). The VM is matched by prefix based on its ID and name.


```
ignite logs <vm> [flags]
```

### Options

```
  -h, --help   help for logs
```

### Options inherited from parent commands

```
      --log-level loglevel      Specify the loglevel for the program (default info)
      --network-plugin plugin   Network plugin to use. Available options are: [cni docker-bridge] (default docker-bridge)
  -q, --quiet                   The quiet mode allows for machine-parsable output by printing only IDs
```

### SEE ALSO

* [ignite](ignite.md)	 - ignite: easily run Firecracker VMs

