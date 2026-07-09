{
  description = "foo";

  inputs = {
    pkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils, ... } @ inputs: 
    (flake-utils.lib.eachDefaultSystem (system:
      let
        repo_root = builtins.toString ./.;

        pkgs = import nixpkgs {
          inherit system;
        };

        packages = (with pkgs; [
          mesa
          glibc
          libXi
          libXcursor
          libXrandr
          raylib
          libglvnd
          libXinerama
          wayland
          libxkbcommon
          odin
        ]) ++ [  ];
      in {
        devShells.default = pkgs.mkShell {
          buildInputs = packages;
          packages = packages;
        };
      })
    );
}
