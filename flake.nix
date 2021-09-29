{
  description = "Open Credit Score Chain";
  inputs = {
    flake-utils.url = "github:numtide/flake-utils";
    flake-compat = {
      url = "github:edolstra/flake-compat";
      flake = false;
    };
  };
  outputs = inputs@{ self, nixpkgs, flake-utils, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let pkgs = nixpkgs.legacyPackages.${system};
      in {
        devShell = pkgs.mkShell {
          buildInputs = with pkgs; [
            nodePackages.ganache-cli # A tool for creating a local blockchain for fast Ethereum development
            go # The Go Programming language
            gopls # Official language server for the Go language
            solc # Compiler for Ethereum smart contract language Solidity
            go-ethereum # Official golang implementation of the Ethereum protocol
          ];
        };
      });
}
