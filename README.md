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
   sig [global options] command [command options] [arguments...]

AUTHOR:
   Derek Smith <derek@clokwork.net>

COMMANDS:
   simple      simple statistics - one time batch process
   stream      stream process
   version, v  Print version info
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)

COPYRIGHT:
   (c) 2022 Derek Smith
```

- [Documentation](./docs/sig.md)
- [Use Cases](#use-cases)
    - [Simple](#simple)
    - [Stream](#stream)
- [Installation](#installation)
    - [Homebrew](#homebrewhttpsbrewsh-for-macos-users)
    - [curl binary](#curl-binary)
    - [docker](#dockerhttpswwwdockercom)
- [Development](#development)
- [Versioning](#versioning)
- [Authors](#authors)
- [License](#license)

## Use Cases

### Simple

```text
$ cat tmp/random.log | sig simple 
N       Min     Max     Mean    Mode    Median  Sum     Std Dev Variance        p50     p75     p90     p95     p99     Q1      Q2      Q3      Outliers        Mild    Extreme
17021   0       255     127.84  70      127     2.176007e+06    74.0224 5479.3108       127     192     231     243     253     64      127     192     0       0       0
```

```text
$ sig simple -t -p 'tmp/*.log'
N       17021
Min     0
Max     255
Mean    127.84
Mode    70
Median  127
Sum     2.176007e+06
Std Dev 74.0224
Variance        5479.3108
p50     127
p75     192
p90     231
p95     243
p99     253
Q1      64
Q2      127
Q3      192
Outliers        0
Mild    0
Extreme 0
```

### Stream

```text
$ cat tmp/random.log | sig stream
N	5000000
Min	1258
Max	2084
Mean	2.28595
Mode	239
Median	224.55
Sum	1.139274e+07
Std Dev	4.717073
Variance	22.24969850
p50	2245
p75	2245.5
p90	4445
p95	5145
p99	117.55
Q1	107.5
Q2	224.55
Q3	224.55
Outliers	538195
Mild	341450
Extreme	196745

[22] next refresh at N modulo 1,000,000 == 0

... THEN ...

N	11388121
Min	1
Max	2084
Mean	1.56
Mode	1
Median	1
Sum	1.7780861e+07
Std Dev	3.1893
Variance	10.1714
p50	1
p75	2
p90	2
p95	3
p99	7
Q1	1
Q2	1
Q3	2
Outliers	538195
Mild	341450
Extreme	196745

Done. Processed 11,388,121 rows
```

> More to come ...

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