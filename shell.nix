let systemPkgs = import <nixpkgs> {};

	pkgs = import (systemPkgs.fetchFromGitHub {
		owner  = "NixOS";
		repo   = "nixpkgs";
		rev    = "c30ad096b2cd5f64d479a70e656b315ea5e96ae9"; # release-20.09
		sha256 = "1hfgriz50z8f2h8b9fasv6f39cqwd7qj80gs6n42rmsg9xxsf5bx";
	}) {};

in pkgs.stdenv.mkDerivation rec {
	name = "cchat-ipc";
	version = "0.0.0-tip";

	buildInputs = with pkgs; [
		flatbuffers
	];
}
