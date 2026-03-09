# KubeSwitcher Plugin: Kind

This is a plugin for the [kubeswitcher](https://github.com/Diaphteiros/kw) tool that allows to switch between clusters created via [kind](https://kind.sigs.k8s.io/).

## Installation

To install the KubeSwitcher plugin, simply run the following command
```shell
go install github.com/Diaphteiros/kw_kind@latest
```
or clone the repository and run
```shell
task install
```

> [!NOTE]
> This project uses [task](https://taskfile.dev/) instead of `make`.

## Configuration

The plugin takes a small configuration in the kubeswitcher config. It can be completely defaulted, if missing.
```yaml
<...>
- name: kind # under which kw subcommand this plugin will be reachable
  short: "Switch to kind clusters" # short message for display in 'kw --help'
  binary: kw_kind # name of or path to the plugin binary
  config:
    binary: kind # path to the kind binary (has to be in $PATH if specified without any path separators) (optional, defaults to 'kind')
```

## Usage

Examples (assuming the plugin was registered with name `kind`):
- `kw kind foo`
  - Sets the kubeconfig to target the kind cluster with name `foo`. No-op, if this cluster is already targeted.
- `kw kind foo --reload`
  - Sets the kubeconfig to target the kind cluster with name `foo`. Reloads the kubeconfig even if 'foo' is already targeted.
- `kw kind --reload`
  - Reloads the kubeconfig for the currently targeted kind cluster.
  - Fails if the current kubeconfig was not set via `kw kind`.
