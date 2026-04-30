# SOAP Go

[![PkgGoDev](https://pkg.go.dev/badge/github.com/justinclift-prvidr/soap-go)](https://pkg.go.dev/github.com/justinclift-prvidr/soap-go)
[![GoReportCard](https://goreportcard.com/badge/github.com/justinclift-prvidr/soap-go)](https://goreportcard.com/report/github.com/justinclift-prvidr/soap-go)
[![CI](https://github.com/justinclift-prvidr/soap-go/actions/workflows/ci.yaml/badge.svg)](https://github.com/justinclift-prvidr/soap-go/actions/workflows/ci.yaml)

A Go SDK and CLI tool for SOAP web services.

## Features

- Support for SOAP 1.1, WSDL 1.1, and XSD 1.0
- Code generation from WSDL files
- Documentation generation

## Developing

```bash
go test ./...                # run unit tests
./tools/mage build           # full CI pipeline: generate, lint, test, tidy
./tools/mage integrationtest # integration tests
```

See [AGENTS.md](./AGENTS.md) for architecture and design notes.

## CLI tool

The `soap` CLI tool can generate code, documentation, and call SOAP APIs on the fly.

```bash
$ soap

  Multi-tool for SOAP APIs

  USAGE

    soap [command] [--flags]  

  CODE GENERATION

    gen [--flags]         Generate code for a SOAP API

  DOCUMENTATION

    doc [--flags]         Display documentation for a SOAP API

  NETWORK OPERATIONS

    call [--flags]        Call a SOAP action

  UTILS

    completion [command]  Generate the autocompletion script for the specified shell
    help [command]        Help about any command

  FLAGS

    -h --help             Help for soap
    -v --version          Version for soap
```

### Installing

The CLI is distributed as a Go module. Install it with the Go toolchain:

```bash
go install github.com/justinclift-prvidr/soap-go/cmd/soap@latest
```

The resulting `soap` binary lands in `$(go env GOBIN)` (or `$(go env GOPATH)/bin` if `GOBIN` is unset).

## License

This SDK is published under the [MIT License](./LICENSE).

## Security

Security researchers, please open a private advisory via the [Security tab](https://github.com/justinclift-prvidr/soap-go/security/advisories/new).

## Code of Conduct

Be nice.
