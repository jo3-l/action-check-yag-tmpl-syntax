# action-check-yag-tmpl-syntax

> Action to check syntax for YAGPDB template files.

## Inputs

### `include`

**Required:** A glob for files to run the syntax checker on. The glob syntax is that of the Go standard library function [filepath.Match](https://golang.org/pkg/path/filepath/#Match).

## Outputs

### `output`

Output from the syntax checker. Each line will contain an error, formatted like such:

```
<filepath>: template: :<line>: <error message>
```

## Example usage

```yml
uses: jo3-l/action-check-yag-tmpl-syntax@v1.0.1
with:
  include: "**/*.go.tmpl"
```

## Maintenance

Adding new functions when they're added in YAGPDB is relatively painless, thanks to how this repository works. Please open an issue if you find that the list of functions has become outdated.<br />

For the curious, this is how it works internally:<br />
We have two files that describe all functions available to YAGPDB templates. Those are:

- [`funcs.json`](./funcs.json) - YAGPDB-specific functions, can be generated by the `genfuncdata` command [here](https://github.com/jo3-l/yagpdb/blob/master/stdcommands/genfuncdata/genfuncdata.go) on my fork of YAGPDB.
- [`builtin_funcs.json`](./builtin_funcs.json) - Functions built-in to the `text/template` package. Unfortunately, these need to be updated manually. On the bright side, they don't change very often.

Based off those two files, a [script](./gen/gen_funcs) generates a Go file exporting a map of functions that accept the correct number of arguments. To regenerate this file, simply run:

```sh
$ go generate
```

## Author

**action-check-yag-tmpl-syntax** © [Joe L.](https://github.com/jo3-l) under the MIT license. Authored and maintained by Joe L.

> GitHub [@jo3-l](https://github.com/jo3-l)
