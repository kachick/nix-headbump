# nix-headbump

For my personal use.

I'm a new to the Nix ecosystem.
(If you know a better way, please let me know!)

I have `default.nix` and `shell.nix` in many repositories. They have different nixpath(?) in the ref from the created timing.
Personally, I use the latest [nixpkgs](https://github.com/NixOS/nixpkgs) ref.
When I want to bump it, I always visit the nixpkgs repository and copy and paste. It is a tedious task.

## Usage

```console
> go install .
> cd repository-that-using-nix
> ${GOPATH:-"$HOME/go"}/bin/nix-headbump && git commit -m 'Bump nixpkgs to latest' *.nix
[main 213d1bf] Bump nixpkgs to latest
 1 file changed, 1 insertion(+), 1 deletion(-)
```

## NOTE

- I guess there are many other syntax patterns in Nix files that I have not used. This code will not care about them.
- I don't know [nix-community/go-nix](https://github.com/nix-community/go-nix) will fit or not.
- I don't know if Nix provides this feature with the CLI or not.
