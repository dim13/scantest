# scantest

A simple, responsive (a word which here means 'snappy') test runner. Like GoConvey, but smarter about what package to run, and with a much simpler interface, and at a fraction of the LOC.

## Features

- Runs `go test ./...` or any command you supply whenever a .go file in any package under the current directory changes.
- Provides colorful output according to exit status of tests (green=passed, red=failed).

### Installation

```
go get github.com/dim13/scantest
```

### Console Runner Execution

```
cd my-project
scantest
```

Results of your tests will display in the terminal until you enter `<ctrl>+c`.
