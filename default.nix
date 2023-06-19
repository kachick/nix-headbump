{ pkgs ? import (fetchTarball "https://github.com/NixOS/nixpkgs/archive/d50f95c6e2a8f58a9e883d918d1e184a6b512900.tar.gz") { } }:

pkgs.mkShell {
  buildInputs = [
    pkgs.go_1_20
    pkgs.nil
    pkgs.nixpkgs-fmt
    pkgs.dprint
    pkgs.actionlint
    pkgs.go-task
  ];
}
