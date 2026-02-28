{
  description = "foo bar baz";

  inputs.nixpkgs.url = "https://flakehub.com/f/NixOS/nixpkgs/0.1";

  outputs = { self, ... }@inputs:
    let
      supportedSystems = [
        "x86_64-linux" # 64-bit Intel/AMD Linux
        "aarch64-linux" # 64-bit ARM Linux
        "aarch64-darwin" # 64-bit ARM macOS
      ];

      forEachSupportedSystem =
        f:
        inputs.nixpkgs.lib.genAttrs supportedSystems (
          system:
          f {
            inherit system;
            pkgs = import inputs.nixpkgs {
              inherit system;
              config.allowUnfree = true;
            };
          }
        );
    in {
      devShells = forEachSupportedSystem (
        { pkgs, system }: {
          default = pkgs.mkShellNoCC {
            packages = with pkgs; [
              gcc
              vala
              gtk4
              glib
              glibc.dev
              libadwaita
              pkg-config
              gobject-introspection
            ];

            shellHook = ''
              # may put something here
            '';
          };
        }
      );
    };
}
