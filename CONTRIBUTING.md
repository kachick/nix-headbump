# How to develop

## Setup

1. Install [Nix](https://nixos.org/) package manager
2. Run `nix develop` or `direnv allow` in project root
3. You can use development tools

```console
> nix develop
(prepared shell)

> task fmt
task: [fmt] dprint fmt
task: [fmt] go fmt

> task
task: [build] ..."
task: [test] go test
task: [lint] dprint check
task: [lint] go vet
PASS
ok      nix-headbump    0.313s

> ./dist/nix-headbump --version
nix-headbump dev (rev) # unknown
```
