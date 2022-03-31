# goworkspace

## Tutorial

We want to work on https://github.com/udhos/goworkspace-lib

1. Test first:

    $ go run app/main.go 
    lib.Modify("fooBAR") => "FOOBAR"

2. Download module:

    git clone https://github.com/udhos/goworkspace-lib
    go work init goworkspace-lib
    $ more go.work
    go 1.18

    use ./goworkspace-lib

3. Change the module, replacing `ToUpper` with `ToLower`:

    vi goworkspace-lib/lib/lib.go

4. Test again:

    $ go run app/main.go 
    lib.Modify("fooBAR") => "foobar"

## Hints

### Do NOT commit `go.work`

Do NOT add `go.work` to source control, it is intended for local work.

It might be handy to add to `.gitignore`:

    go.work
    goworkspace-lib

### Enhanced `replace` directive

`go.work` supports an enhanced `replace` directive. Example:

```
replace golang.org/x/net v1.2.3 => example.com/fork/net v1.4.5

replace (
    golang.org/x/net v1.2.3 => example.com/fork/net v1.4.5
    golang.org/x/net => example.com/fork/net v1.4.5
    golang.org/x/net v1.2.3 => ./fork/net
    golang.org/x/net => ./fork/net
)
```

See: https://go.dev/ref/mod#workspaces