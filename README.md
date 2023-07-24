# nixpkgs-url

For my personal use.

I'm a new to the Nix ecosystem.\
(If you know a better way, please let me know!)

I have `flake.nix` and `default.nix` in many repositories.\
They use different nipkgs url from the created timing.\
Personally, I use the latest [nixpkgs](https://github.com/NixOS/nixpkgs) ref. But I avoid to specify `unstable`.\
When I want to bump it, I always visit the nixpkgs repository and copy and paste. It is a tedious task.

## Installation

[Prebuilt binaries](https://github.com/kachick/nixpkgs-url/releases)

```console
> curl -L https://github.com/kachick/nixpkgs-url/releases/latest/download/nixpkgs-url_Linux_x86_64.tar.gz | tar xvz -C ./ nixpkgs-url
> ./nixpkgs-url --version
...
```

In [Nix](https://nixos.org/), you can skip installation steps

```console
> nix run github:kachick/nixpkgs-url -- --version
nixpkgs-url dev (rev) # unknown
> nix run github:kachick/nixpkgs-url/v0.2.4 -- detect --current
...(With specific version)
```

`go install`

```console
> go install github.com/kachick/nixpkgs-url/cmd/nixpkgs-url@latest
go: downloading...
> ${GOPATH:-"$HOME/go"}/bin/nixpkgs-url --version
nixpkgs-url dev (rev) # unknown
```

## Usage

Providing two subcommands. I'm using `detect` in CI and `bump` in local.

```console
> nixpkgs-url detect --current
e57b65abbbf7a2d5786acc86fdf56cde060ed026

> nixpkgs-url bump && git commit -m 'Bump nixpkgs to latest' *.nix
[main 213d1bf] Bump nixpkgs to latest
 1 file changed, 1 insertion(+), 1 deletion(-)
```

## NOTE

- I guess there are many other syntax patterns in Nix files that I have not used. This code will not care about them.
- I don't know [nix-community/go-nix](https://github.com/nix-community/go-nix) will fit or not.
- I don't know if Nix provides this feature with the CLI or not.
