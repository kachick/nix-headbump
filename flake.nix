{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/7fd307937db70af23b956c4539033542809ae263";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
      in
      rec {
        devShells.default = with pkgs;
          mkShell {
            buildInputs = [
              go_1_20
              nil
              nixpkgs-fmt
              dprint
              actionlint
              go-task
              goreleaser
              typos
              go-tools
            ];
          };

        packages.nixpkgs-path = pkgs.stdenv.mkDerivation
          {
            name = "nixpkgs-path";
            src = self;
            buildInputs = with pkgs; [
              go_1_20
              go-task
            ];
            buildPhase = ''
              # https://github.com/NixOS/nix/issues/670#issuecomment-1211700127
              export HOME=$(pwd)
              task build
            '';
            installPhase = ''
              mkdir -p $out/bin
              install -t $out/bin dist/nixpkgs-path
            '';
          };

        packages.default = packages.nixpkgs-path;

        # `nix run`
        apps.default = {
          type = "app";
          program = "${packages.nixpkgs-path}/bin/nixpkgs-path";
        };
      }
    );
}
