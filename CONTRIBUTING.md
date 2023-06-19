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
task: [lint] dprint check
task: [lint] go vet
task: [lint] actionlint

> task fmt
task: [fmt] dprint fmt
task: [fmt] go fmt

> ./nix-headbump --version
0.1.0(ceaa32d)
```
