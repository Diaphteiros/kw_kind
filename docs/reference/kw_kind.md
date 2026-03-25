## kw_kind

Switch to a local kind cluster

### Synopsis

Switch to a local kind cluster.

This command sets the kubeconfig to the output of 'kind get kubeconfg --name <name>'.
No-op if the current kubeconfig is already set to the kubeconfig of the specified kind cluster.

The '--reload' flag causes the kubeconfig to be loaded even if it is the currently targeted cluster.
If the flag is set, the cluster name can be omitted to reload the current cluster's kubeconfig. This results in an error if the currently targeted kubeconfig was not set via this command.

If neither a cluster name, nor the '--reload' flag is set, the user will be prompted to select one of the known kind clusters via a fuzzy selector.

```
kw_kind [<name>] [flags]
```

### Options

```
  -h, --help     help for kw_kind
  -r, --reload   Load the cluster kubeconfig even if it is the same cluster as the currently selected one. Can be used without cluster name to reload the current cluster's kubeconfig.
```

### SEE ALSO

* [kw_kind version](kw_kind_version.md)	 - Print the version

