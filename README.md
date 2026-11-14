# randpw

Random password generator. Characters are picked with `crypto/rand`, and every
character in the chosen set is equally likely.

## Install

```sh
go install github.com/mikemadden42/randpw@latest
```

Requires Go 1.26 or later. Prebuilt binaries for Linux, macOS and Windows are
on the [releases page](https://github.com/mikemadden42/randpw/releases).

## Usage

```sh
randpw                      # 16-character alphanumeric password
randpw -length 32           # longer password
randpw -type number         # digits only (alpha, alphanum or number)
randpw -symbols             # add the symbols !#$%&*+-=?@^_
randpw -all                 # one labeled password of each type
randpw | pbcopy             # copy to the clipboard on macOS
```

| Flag | Default | Description |
|---|---|---|
| `-length` | `16` | Password length (at least 1) |
| `-type` | `alphanum` | Character set: `alphanum`, `alpha` or `number` |
| `-symbols` | `false` | Also use the symbols `!#$%&*+-=?@^_` |
| `-all` | `false` | Print one labeled password of each type |

## License

MIT
