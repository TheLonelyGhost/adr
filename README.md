# ADR Go

A minimalist command line tool written in Go to work with
[Architecture Decision Records](http://thinkrelevance.com/blog/2011/11/15/documenting-architecture-decisions)
(ADRs).

Greatly inspired by the [adr-tools](https://github.com/npryce/adr-tools)
with the portability of a single Go binary.

## Quick start

## Installation

```bash
~/workspace $ go install github.com/thelonelyghost/adr@latest
```

This will place a command `adr` in your go binaries directory
(`~/go/bin`, by default). It is recommended to either add that
to your PATH or copy it to a location that is already within
your PATH, such as...

```bash
~/workspace $ install -m0755 ~/go/bin/adr ~/.local/bin/adr
```

## Initializing an ADR store

Before creating any new architecture decision record, a location on your filesystem must be
configured to host your ADRs. Use the `init` sub-command to initialize the repository:

```bash
~/workspace $ adr init .
```

Or, if you would like to track decision records globally:

```bash
~/worksapce $ adr init --global
```

It is recommended to version your records (and the repository metadata in `.adr/`) with a
version control system like `git`.

## Creating a new ADR

```bash
~/workspace $ adr new my awesome proposition
```

> [!note] Add `--global` if adding to the global store

This will create a new ADR in your decision directory: `xxx-my-new-awesome-proposition.md`.

Next, open the file in your preferred markdown editor and document the details of
the decision being made.
