{
  description = "foo";
  inputs = {
    pkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils, ... } @ inputs: 
    (flake-utils.lib.eachDefaultSystem (system:
      let
        # add the Zig overlay pkgs
        pkgs = import nixpkgs {
          inherit system;
        };
      in {
        # Nix shell
        devShells.default = pkgs.mkShell {
          packages = (with pkgs; [
            lua5_5
            lua55Packages.luasocket
          ]);
        };
      })
    );
}
/*
nixpkgs.overlays = [
  (final: prev: {
    lua = prev.lua.overrideAttrs (oldAttrs: rec {
      version = "5.4.7"; # Or whatever the latest is
      src = fetchurl {
        url = "https://lua.org{version}.tar.gz";
        sha256 = "your-new-sha256-here";
      };
    });
  })
];
  */
