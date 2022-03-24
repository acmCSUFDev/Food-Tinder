# configuration.nix contains the services required to run the backend process in production. It
# currently only contains PostgreSQL.
{ config, pkgs, lib, ...}:

let shell = import ./shell.nix {};

	backend = pkgs.buildGoModule {
		name = "food-tinder-backend";
		src  = ./backend;
		subPackages = [ "." ];
		vendorSha256 = "02llr4ns3yp13rw816dgn7619knmiq2jrv1c5r9z40my0065b0zr";
	};

in {
	programs.bash.promptInit = ''
		localIP=$(ip -br addr show eth0 | sed 's/.*\(10\..*\)\/32.*/\1/')
		function is-active() { systemctl --quiet is-active "$1"; }

		echo "***"

		is-active postgresql \
			&& echo "postgresql server started at $localIP:5432." \
			|| echo "postgresql failed!"

		is-active food-tinder-backend \
			&& echo "food-tinder-backend started at $localIP:3001." \
			|| echo "food-tinder-backend failed!"

		echo -n "***"
	'';

	systemd.services.food-tinder-backend = {
		description = "Food Tinder backend service";
		environment = {
			HTTP_ADDRESS = "0.0.0.0:3001";
			# DB_ADDRESS   = "mock://${./backend/dataset/mockdb.json}";
			DB_ADDRESS   = "postgres://foodtinder@localhost:5432/foodtinder";
		};
		after    = [ "postgresql.service" ];
		wantedBy = [ "multi-user.target"  ];
		serviceConfig = {
			ExecStart   = "${backend}/bin/backend";
			Restart     = "on-failure";
			DynamicUser = true;
		};
	};

	services.postgresql = {
		enable = true;
		enableTCPIP = true;
		authentication = ''
			local foodtinder foodtinder           trust
			host  foodtinder foodtinder 0.0.0.0/1 trust
		'';
		ensureDatabases = [ "foodtinder" ];
		ensureUsers = [
			{
				name = "foodtinder";
				ensurePermissions = { "DATABASE foodtinder" = "ALL PRIVILEGES"; };
			}
		];
	};
}
