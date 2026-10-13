# Testing and release gates

[Documentation index](README.md)

Start standalone development with `go test ./...`: deterministic tests, native
fuzz seeds and saved corpus regressions do not need sibling checkouts or an
ecosystem ledger. With the checksum-pinned public verifier, `make check` and
`make ci` select bounded local verification, not a full campaign. Ordinary CI
uses the released verifier rather than source bootstrapping and retains
applicable security and fail-closed Required checks.

`make check-full` retains the original full operations and online specification
refresh: formatting, documentation validation, vet, tests, coverage, race,
lint/static checks, embedded-list integrity, examples, dependency verification,
vulnerability/license review, API compatibility, benchmarks and all four native
10s fuzz budgets. Broader assurance is selected for a named risk or scheduled/
release boundary; mutation is optional local/scheduled, never a required PR
status. See [contributor verification](../CONTRIBUTING.md#verification) for a
single-target 2s fuzz command and optional single-package mutation.

A future release requires the approved
[independent review record](security-review.md) for its source, plus applicable
CI and release gates, including `golib release dry-run --module .`. The existing
`v1.0.0` tag predates that review and must not be interpreted as its evidence.

Official vectors cover every official language and entropy size. The
interoperability tests compare every English entropy size, parsing, checksum
rejection, normalized Japanese seed input, and PBKDF2 output with an independent
mature Go implementation. Property tests cover policy satisfaction and round
trips.

The focused `TestSeedStrictFIPSRejectsShortBIP39Salt` check runs with
`GODEBUG=fips140=only` to verify typed, no-seed rejection rather than disabling
strict enforcement; ordinary BIP-39 vector tests run in the supported default
crypto mode.

Statistical checks use documented false-positive thresholds only to catch
obvious selection regressions. Race tests share immutable lists and default
generators. Fuzz targets cover alphabets, lists, mnemonic normalization and
malformed input, and passphrase parsing.

Mutation configuration targets checksum, policy, normalization, selection, and
error paths. Surviving relevant mutations block release; equivalent or
unreachable mutations require a reviewed exclusion with rationale.
