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
