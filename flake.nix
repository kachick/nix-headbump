{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/e57b65abbbf7a2d5786acc86fdf56cde060ed026";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
      in
      {
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
            ];
          };
      });
}
