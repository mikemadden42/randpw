# TODO

## Correctness and security (`randpw.go`)

- [x] **Fix modulo bias** in `randStr` (`dictionary[v%byte(len(dictionary))]`).
  256 isn't a multiple of the set sizes, so some characters are 25% more likely:
  - alphanum (62): `0`–`7` favored (5/256 vs 4/256)
  - alpha (52): `A`–`v` (first 48) favored
  - number (10): `0`–`5` favored (26/256 vs 25/256)

  Use `crypto/rand.Int(rand.Reader, big.NewInt(int64(len(dict))))` per character,
  or rejection-sample bytes `>= 256 - 256%n`. This also clears gosec G115.
- [x] **Validate `-length`**: `-length -1` panics with `makeslice: len out of range`.
  Reject values `< 1` after `flag.Parse()` with a clear error.
- [x] **Handle unknown `randType`**: an unknown type leaves `dictionary` empty and
  causes a divide-by-zero panic. Use a `map[string]string` or a `switch` with a
  default that returns an error.
- [x] **Replace `panic` in `checkErr`**: print to stderr and `os.Exit(1)`.

## Usability

- [ ] Add a `-type` flag and print only the bare password, so output can be piped
  (e.g. `randpw | pbcopy`).
- [ ] Fix the double space in output (`Println("Alphanum : ", x)`).
- [ ] Consider a `-symbols` option.

## Tooling and configuration

- [x] **Fix `.golangci.yml` for v2** (`golangci-lint config verify` fails):
  - `version: 2` → `version: "2"`
  - `disable-all: true` → `linters.default: none`
  - `linters-settings` → `linters.settings`; goimports settings go under
    `formatters.settings.goimports.local-prefixes` (a list)
  - `issues.exclude-dirs` → `linters.exclusions.paths`
  - After fixing, address the lint findings: G115 (above) and missing package comment.
- [x] **Fix `.goreleaser.yaml`** (`goreleaser check` fails):
  - Add `version: 2`
  - `snapshot.name_template` → `snapshot.version_template`
  - Drop the unused `go generate ./...` hook and leftover example comments.
- [ ] **Makefile**: the `release` target uses `--snapshot`, so it never publishes.
  Rename it or drop the flag.

## CI

- [ ] Raise the `go` line in `go.mod` from 1.18 to the oldest version CI tests (1.23).
- [x] Add `randpw_test.go`: length, character-set membership, and a rough
  distribution check (would have caught the modulo bias).
- [ ] Add a lint step (`golangci/golangci-lint-action`) once the config is fixed.
- [ ] Bump `actions/setup-go` and `actions/checkout` to current major versions.

## Docs and minor cleanup

- [ ] Expand README with install and usage instructions, including flags.
- [x] Remove stray blank lines at the start of `randStr` and the end of `main`.

## Suggested order

1. Fix the bias
2. Validate input
3. Add tests
4. Fix the golangci and goreleaser configs
5. Add the lint step to CI
