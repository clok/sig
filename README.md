# sig

[![License: MIT](https://img.shields.io/badge/License-MIT-brightgreen.svg)](https://github.com/clok/sig/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/clok/sig)](https://goreportcard.com/report/clok/sig)
[![Coverage Status](https://coveralls.io/repos/github/clok/sig/badge.svg)](https://coveralls.io/github/clok/sig)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/clok/sig?tab=overview)

Statistics in Go - CLI tool for quick statistical analysis of data streams

> Please see [the docs for details on the commands.](./docs/sig.md)

```text
$ sig --help
NAME:
   sig - Statistics in Go - CLI tool for quick statistical analysis of data streams

USAGE:
   main [global options] command [command options] [arguments...]

AUTHOR:
   Derek Smith <derek@clokwork.net>

COMMANDS:
   hello       say hello
   version, v  Print version info
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)

COPYRIGHT:
   (c) 2022 Derek Smith
```

- [Documentation](./docs/sig.md)
- [Use Cases](#use-cases)
- [Installation](#installation)
    - [Homebrew](#homebrewhttpsbrewsh-for-macos-users)
    - [curl binary](#curl-binary)
    - [docker](#dockerhttpswwwdockercom)
- [Development](#development)
- [Versioning](#versioning)
- [Authors](#authors)
- [License](#license)

## Use Cases

Will be filled out later.

## Installation

### [Homebrew](https://brew.sh) (for macOS users)

```
brew tap clok/sig
brew install sig
```

### curl binary

```
$ curl https://i.jpillora.com/clok/sig! | bash
```

### [docker](https://www.docker.com/)

The compiled docker images are maintained
on [GitHub Container Registry (ghcr.io)](https://github.com/orgs/clok/packages/container/package/sig). We maintain the
following tags:

- `edge`: Image that is build from the current `HEAD` of the main line branch.
- `latest`: Image that is built from the [latest released version](https://github.com/clok/sig/releases)
- `x.y.z` (versions): Images that are build from the tagged versions within Github.

```bash
docker pull ghcr.io/clok/sig
docker run -v "$PWD":/workdir ghcr.io/clok/sig --version
```

### man page

To install `man` page:

```
$ sig install-manpage
```

## Development

1. Fork the [clok/sig](https://github.com/clok/sig) repo
1. Use `go >= 1.17`
1. Branch & Code
1. Run linters :broom: `golangci-lint run`
    - The project uses [golangci-lint](https://golangci-lint.run/usage/install/#local-installation)
1. Commit with a Conventional Commit
1. Open a PR

## Versioning

We employ [git-chglog](https://github.com/git-chglog/git-chglog) to manage the [CHANGELOG.md](CHANGELOG.md). For the
versions available, see the [tags on this repository](https://github.com/clok/sig/tags).

## Authors

* **Derek Smith** - [@clok](https://github.com/clok)

See also the list of [contributors](https://github.com/clok/sig/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details