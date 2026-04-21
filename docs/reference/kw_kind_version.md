## kw_kind version

Print the version

### Synopsis

Output the version of the CLI.

```
kw_kind version [flags]
```

### Examples

```
  > kw kind version
  v1.2.3

  > kw kind version -o json
  {"version":"v1.2.3-dev-4516e7f4dee0861b3d1a31b53d3a8aabbd084f48","gitTreeState":"dirty","gitCommit":"4516e7f4dee0861b3d1a31b53d3a8aabbd084f48","buildDate":"2026-04-21T13:14:36Z","major":1,"minor":2,"patch":3,"suffix":"dev-4516e7f4dee0861b3d1a31b53d3a8aabbd084f48"}

  > kw kind version -o yaml
	buildDate: "2026-04-21T13:14:36Z"
	gitCommit: 4516e7f4dee0861b3d1a31b53d3a8aabbd084f48
	gitTreeState: dirty
	major: 1
	minor: 2
	patch: 3
	suffix: dev-4516e7f4dee0861b3d1a31b53d3a8aabbd084f48
	version: v1.2.3-dev-4516e7f4dee0861b3d1a31b53d3a8aabbd084f48
```

### Options

```
  -h, --help            help for version
  -o, --output string   Output format. Valid formats are [json, text, yaml]. (default "text")
```

### SEE ALSO

* [kw_kind](kw_kind.md)	 - Switch to a local kind cluster

