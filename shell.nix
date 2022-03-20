{ pkgs ? import <nixpkgs> {} }:

let goapi-gen = pkgs.buildGoModule {
		name = "goapi-gen";
		version = "081d60b";

		src = pkgs.fetchFromGitHub {
			owner  = "discord-gophers";
			repo   = "goapi-gen";
			rev    = "4be5ffd30c4c1fb32bfaeff8ab8923a95a91120f";
			sha256 = "0kcr1rfgdhlsgbjdw23v1zx13w2gcd2zvmgfamwgk9z1p6if4y4c";
		};
		# src = pkgs.fetchgit {
		# 	url    = "https://github.com/diamondburned/goapi-gen";
		# 	rev    = "49e462fafc1d82572218bdec3917d50c98ebed2e";
		# 	sha256 = "0jgvjf51bzfm620gy6r9fxnyq9yi54vvif8jzfrrn4rj13zqvhc3";
		# };

		patches =
			with builtins;
			let commit = user: c: fetchurl "https://github.com/${user}/goapi-gen/commit/${c}.patch";
				pr     = user: c: fetchurl "https://github.com/${user}/goapi-gen/pull/${toString c}.patch";
			in [
				(pr "discord-gophers" 80)
				(pr "discord-gophers" 82)
				(pr "discord-gophers" 83)
				(pr "discord-gophers" 85)
			];

		vendorSha256 = "1aq24cx5qirgzjcahzqjkzc50687xj2vqz623f56q5j5m2x8cj73";
	};

	sqlc = pkgs.buildGoModule {
		name = "sqlc";
		version = "1.12.0";

		src = pkgs.fetchFromGitHub {
			owner  = "kyleconroy";
			repo   = "sqlc";
			rev    = "45bd150";
			sha256 = "1np2xd9q0aaqfbcv3zcxjrfd1im9xr22g2jz5whywdr1m67a8lv2";
		};

		proxyVendor = true;
		vendorSha256 = "1qr8nymhcwmqj3b8f50fknl2rbnrnvfl0yj5vfs0l1jd0r4rz2d0";
	};

	moq = pkgs.buildGoModule {
		name = "moq";
		version = "0.2.6";

		src = pkgs.fetchFromGitHub {
			owner  = "matryer";
			repo   = "moq";
			rev    = "5d3d962614e152b11aa8080d6de7b12445bf09a1";
			sha256 = "0zsr466iaxzb24kjq82g00765hhw0lgikdva2nkxhrrgijczp8hk";
		};

		vendorSha256 = "02kb11pjcrjjsqaafj07fmvzzk03mmy74kmh004rd3ddkkdbjdsx";
		subPackages = [ "." ];
	};

in pkgs.mkShell {
	buildInputs = with pkgs; [
		go_1_17
		goapi-gen
		sqlc
		moq

		nodejs
	];

	shellHook = ''
		PATH="$PWD/frontend/node_modules/.bin:$PATH"
	'';
}
