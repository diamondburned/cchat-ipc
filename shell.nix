let pkgs = import <nixpkgs> {};

	flatbuffers = pkgs.flatbuffers.overrideAttrs(old: {
		# Use master.
		src = pkgs.fetchFromGitHub {
			owner  = "google";
			repo   = "flatbuffers";
			rev    = "39e115fdb468d614accc09d177f1b46900473481";
			sha256 = "0dqn5vi8gfwa43r68hdnwqdqzsmll23kxvyr9zjvs62qjnj40s7f";
		};
	});

in pkgs.stdenv.mkDerivation rec {
	name = "cchat-ipc";
	version = "0.0.0-tip";

	buildInputs = [
		flatbuffers
	];
}
