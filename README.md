# pfind - Parallel Directory Finder

A fast, concurrent file system traversal utility written in Go.

## Overview

`pfind` is a command-line tool that recursively lists all files and directories starting from a specified root directory. It uses Go's concurrency features and a worker pool to efficiently traverse large directory structures in parallel.

## Features

- Fast parallel directory traversal
- Uses a worker pool to limit concurrency
- Simple command-line interface

## Installation

```bash
go install github.com/yourusername/pfind@latest
```

Or build from source:

```bash
git clone https://github.com/yourusername/pfind.git
cd pfind
go build
```

## Usage

```bash
pfind <start-directory>
```

Example:

```bash
pfind /home/user/documents
```

## Dependencies

- [github.com/alitto/pond](https://github.com/alitto/pond) - Worker pool implementation

## License

See the [LICENSE](LICENSE) file for details.
