# Food-Tinder Backend

This folder contains the backend code for Food-Tinder. API documentation is at
[./openapi](./openapi).

## Running

The backend can run without any database server. It does this by using an
existing JSON file containing the initial state of the database, and all further
writes will be stored in memory until the program exits.

To run the backend, you'll need Go 1.17 or later. Run the command:

```sh
go run .
```

within the `./backend` (this) directory. If not, point `DB_ADDRESS` to the right
`./dataset/mockdb.json`.

By default, the process listens to http://localhost:3001. To change this, see
`.env`.

If `mockdb.json` is used, then the example login credentials are
`food_tinder_user:password`.

## Developing

It is recommended to develop the backend inside the supplied
[shell.nix](../shell.nix) environment. The environment contains the needed tools
to generate the Go code in this repository.

To set this up, first follow the [Install Nix](https://nixos.org/download.html)
instructions, then run `nix-shell` inside the *project* root directory.

To regenerate all Go files, use `go generate ./...`.

To generate all Go *and* JS files, `cd` back to the project root directory and
run `./generate.sh`.

To run the backend with a proper PostgreSQL development instance, use `sudo
nixos-shell` on the project root directory.

## Resources

- [Common Anti-Patterns in Go Web Applications](https://threedots.tech/post/common-anti-patterns-in-go-web-applications/)
- [OpenAPI Specification](https://swagger.io/specification/)
- [discord-gophers/goapi-gen](https://github.com/discord-gophers/goapi-gen)
- [kyleconroy/sqlc](https://github.com/kyleconroy/sqlc)
