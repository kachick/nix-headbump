{ pkgs ? import (fetchTarball "https://github.com/NixOS/nixpkgs/archive/e57b65abbbf7a2d5786acc86fdf56cde060ed026.tar.gz") { } }:

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
