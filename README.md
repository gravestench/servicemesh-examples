## About

This module contains examples for the [Service Mesh](https://github.com/gravestench/servicemesh) library.

There are example services in the `services` directory. Some of 
these services depend on each other, namely the config file manager.

There are runnable example applications in the `cmd` directory. Some commands
also contain small, purpose-built service declarations.

Each command registers its complete service set before waiting for
initialization. This allows services with cyclic dependencies to discover each
other while ensuring initialization has completed before the mesh run loop
starts.

Run an example from the repository root:

```sh
go run ./cmd/bare_minimum
```

Some examples require platform facilities or credentials, including audio,
desktop graphics, Twitch, or TLS certificate configuration. Compile and vet the
complete collection with:

```sh
go test -p 1 ./...
```
