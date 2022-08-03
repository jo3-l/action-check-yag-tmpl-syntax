# action-check-yag-tmpl-syntax

> Action to check syntax for YAGPDB template files.

## Inputs

### `include`

**Required:** A glob for files to run the syntax checker on. The glob syntax is described by the documentation for the [`Match`](https://pkg.go.dev/github.com/bmatcuk/doublestar/v4#Match) function from the Go library [doublestar](https://github.com/bmatcuk/doublestar).

## Outputs

### `output`

Output from the syntax checker. Each line will contain an error, formatted like such:

```
<filepath>: template: :<line>: <error message>
```

## Example usage

```yml
uses: jo3-l/action-check-yag-tmpl-syntax@v2.1.1
with:
  include: '**/*.go.tmpl'
```

## Maintenance

New template functions are queried for on a weekly basis using [`yagfuncdata`](https://github.com/jo3-l/yagfuncdata) through a cron-based [GitHub action](./.github/workflows/regenerate-funcs.yml). If changes are detected, a PR is automatically issued.

Changes to the template executor itself (e.g., addition of new keywords) are more involved and require manual intervention, though this is rare. (Specifically, the `template` folder needs to be synchronized with its upstream counterpart, [`yagpdb/lib/template`](https://github.com/botlabs-gg/yagpdb/tree/master/lib/template).)

## Author

**action-check-yag-tmpl-syntax** © [Joe L.](https://github.com/jo3-l) under the MIT license. Authored and maintained by Joe L.

> GitHub [@jo3-l](https://github.com/jo3-l)
