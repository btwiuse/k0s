{ pkgs ? import (fetchTarball "https://github.com/NixOS/nixpkgs/archive/master.tar.gz") { } }:

pkgs.mkShell {
  # nativeBuildInputs is usually what you want -- tools you need to run
  nativeBuildInputs = with pkgs; [
    nixpkgs-fmt
    rnix-lsp
    docker-client
    gnumake

    # go development
    go_1_19
    go-outline
    gopls
    gopkgs
    go-tools
    delve
  ];

  hardeningDisable = [ "all" ];
}
