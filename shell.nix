{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
	buildInputs = with pkgs; [
		go_1_17
		nodejs
	];

	shellHook = ''
		PATH="$PWD/node_modules/.bin:$PATH"
	'';
}
