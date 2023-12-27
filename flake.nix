{
  description = "Build diet3 using maven and java with this flake";

  #inputs
  inputs = {
    nixpkgs.url = "nixpkgs/nixos-23.11"; 
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let pkgs = nixpkgs.legacyPackages.x86_64-linux;
      in {
        devShell =
        pkgs.mkShell { 
          buildInputs = with pkgs; [ jdk19 maven ]; 

          shellHook = ''
            export PS1="dev > "
          '';
        };

      }
    );
}
