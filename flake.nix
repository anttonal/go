{
  description = "Go dev shell";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-26.05";

  outputs = { self, nixpkgs }:
    let
      system = "x86_64-linux";
      pkgs = import nixpkgs { inherit system; };
    in {
      devShells.${system}.default = pkgs.mkShell {
        packages = with pkgs; [
          go
          gopls
          golangci-lint
          delve
        ];

        # Without this, the `go` directive in go.mod makes the toolchain fetch
        # and run its own version, silently ignoring the go pinned above.
        shellHook = ''
          export GOTOOLCHAIN=local
        '';
      };
    };
}
