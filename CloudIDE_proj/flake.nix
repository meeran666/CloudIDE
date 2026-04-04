{
  description = "Go development environment";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  };

  outputs = {
    self,
    nixpkgs,
  }: let
    system = "x86_64-linux";
    pkgs = import nixpkgs {inherit system;};
  in {
    devShells.${system}.default = pkgs.mkShell {
      buildInputs = with pkgs; [
        go
        gopls
        git

        # optional tooling
        golangci-lint
        delve

        # if you use bun / node tooling
        bun
        nodejs

        # fix for libstdc++
        stdenv.cc.cc.lib
      ];

      shellHook = ''
        echo "Go dev environment ready"
        go version
      '';
    };
  };
}
