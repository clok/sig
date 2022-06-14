% sig 8
# NAME
sig - Statistics in Go - CLI tool for quick statistical analysis of data streams
# SYNOPSIS
sig


# COMMAND TREE

- [simple](#simple)
- [stream](#stream)
- [version, v](#version-v)

**Usage**:
```
sig [GLOBAL OPTIONS] command [COMMAND OPTIONS] [ARGUMENTS...]
```

# COMMANDS

## simple

simple statistics - one time batch process

**--no-header**: do not print out header

**--path, -p**="": File path to files to stream, can be a glob. If not set, a pipe is assumed.

**--transpose, -t**: transpose table output

## stream

stream process 

**--cap, -c**="": max value of refresh rate for updates (default: 100000000)

**--factor, -f**="": rate of growth of refresh value (default: 10)

**--path, -p**="": File path to files to stream, can be a glob. If not set, a pipe is assumed.

**--refresh, -r**="": how many rows of data between updates (default: 100)

## version, v

Print version info

