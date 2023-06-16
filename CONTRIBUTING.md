# How to develop

## Setup

1. Install [Nix](https://nixos.org/) package manager
2. Run `nix-shell` or `nix-shell --command 'zsh'`
3. You can use development tools

```console
> nix-shell
(prepared bash)

> go version
go version go1.20.4 linux/amd64

> dprint --version
dprint 0.36.1

> go build -ldflags "-X main.revision=$(git rev-parse --short HEAD)"
```
