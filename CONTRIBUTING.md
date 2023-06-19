# How to develop

## Setup

1. Install [Nix](https://nixos.org/) package manager
2. Run `nix-shell` or `nix-shell --command 'zsh'`
3. You can use development tools

```console
> nix-shell
(prepared bash)

> task fmt
task: [fmt] dprint fmt
task: [fmt] go fmt

> task
task: [build] ..."
task: [test] go test
task: [lint] dprint check
task: [lint] go vet
task: [lint] actionlint
PASS
ok      nix-headbump    0.313s

> find dist
dist
dist/metadata.json
dist/config.yaml
dist/nix-headbump_linux_amd64_v1
dist/nix-headbump_linux_amd64_v1/nix-headbump
dist/artifacts.json

> ./dist/nix-headbump_linux_amd64_v1/nix-headbump --version
nix-headbump 0.1.1-next (906924b) # 2023-06-19T09:33:14Z
```
