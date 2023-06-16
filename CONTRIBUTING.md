# How to develop

## Setup

1. Install [Nix](https://nixos.org/) package manager
2. Run `nix-shell` or `nix-shell --command 'zsh'`
3. You can use development tools

```console
> nix-shell
(prepared bash)

> task
task: [build] go build -ldflags "-X main.revision=$(git rev-parse --short HEAD)"
task: [fmt] dprint fmt
task: [lint] dprint check
task: [lint] go vet
task: [fmt] go fmt
task: [lint] actionlint

> ./nix-headbump --version
0.1.0(ceaa32d)
```
