# nix-headbump

For my personal use.

I'm a new to the Nix ecosystem.
(If you know a better way, please let me know!)

I have `default.nix` and `shell.nix` in many repositories. They have different nixpath(?) in the ref from the created timing.
Personally, I use the latest [nixpkgs](https://github.com/NixOS/nixpkgs) ref.
When I want to bump it, I always visit the nixpkgs repository and copy and paste. It is a tedious task.

## Installation

`go install` is also okay, or use [prebuilt binaries](https://github.com/kachick/nix-headbump/releases)

```console
> curl -L https://github.com/kachick/nix-headbump/releases/latest/download/nix-headbump_Linux_x86_64.tar.gz | tar xvz -C ./ nix-headbump
> ./nix-headbump --version
nix-headbump 0.1.0 (d8e9da7) # 2023-06-19T08:55:33Z
```

## Usage

```console
> nix-headbump && git commit -m 'Bump nixpkgs to latest' *.nix
[main 213d1bf] Bump nixpkgs to latest
 1 file changed, 1 insertion(+), 1 deletion(-)
```

## NOTE

- I guess there are many other syntax patterns in Nix files that I have not used. This code will not care about them.
- I don't know [nix-community/go-nix](https://github.com/nix-community/go-nix) will fit or not.
- I don't know if Nix provides this feature with the CLI or not.
