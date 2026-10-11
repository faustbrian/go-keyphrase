# Contributing

## Before Editing

1. Read [`AGENTS.md`](AGENTS.md) and the affected package's docs.
2. Run `go test ./...` or the affected package's narrower tests.
3. Identify owned dependencies and reverse dependants in `modules.json`.
4. Preserve unrelated work and generated/corpus provenance.

## Changes

Keep commits focused and conventional. Update every affected changelog with
the behavior and migration impact. Public API changes require compatibility
evidence and documentation. Specification behavior requires a decision record,
fixture coverage, and interoperability evidence.

New direct dependencies and dependency updates must follow the
[dependency governance policy](AGENTS.md#dependencies-and-supply-chain). Package-local
update bots are forbidden; the root policy owns every module and action update.

Specification-backed changes must follow the
[specification governance contract](AGENTS.md#design), update
the affected stable decision entries, and complete the Specification Decisions
section of the pull request template. An unresolved interpretation or stale
source pin is release-blocking; peer behavior cannot silently select policy.

Mutation is optional local or scheduled assurance, not a required PR status.
Confirmed behavioral or security defects still require a focused regression
and repair; retain reviewed narrow equivalent/invalid-mutant dispositions.

Do not add package-local workflows, permanent replacements, machine-specific
paths, bypass flags, broad mutation exclusions, or aggregate quality metrics
that hide a failing package.

## Verification

Standalone `go test ./...` runs deterministic tests, fuzz seeds and saved corpus
regressions without sibling checkouts or ecosystem ledgers. The manifest declares
no required services, build tags or test tags.

With the checksum-pinned public verifier selected by `.golib.yaml`, use:

```bash
make inventory
make check
make ci
```

`check` and `ci` use bounded local checks. Ordinary CI consumes the public
verifier without source bootstrapping. `ci` validates offline specification
state; `make specification-check` retains online authority refresh and is also
selected by `make check-full`.

Select one target for short active fuzz exploration. Optional local mutation
selects one package; neither command is an all-change prerequisite:

```bash
GOMAXPROCS=2 go test -run '^$' -fuzz '^FuzzAnalyzeAlphabet$' -fuzztime=2s -parallel=1 -timeout=2m ./password
golib mutation --module . --package ./password
```

Selection, normalization, secret cleanup, injected-source or concurrency changes
retain affected behavioral, race, official-vector and interoperability checks.
Generators are stateless; unrelated changes do not acquire a mandatory race
campaign. Preserve all pinned list bytes, provenance and secret-safe assertions.

`make check-full` preserves explicit full assurance, all four native 10s fuzz
budgets, embedded-list integrity, examples and benchmarks. See
[testing and release gates](docs/testing.md) for strict-FIPS and conformance
boundaries. Relevant release qualification remains
`golib release dry-run --module .`; it does not require a sibling catalog release.
Report unavailable, skipped and failed boundaries honestly; a local pass is not
hosted-CI or public-release qualification.

## Adding A Module

Follow [repository structure policy](AGENTS.md#repository-structure). New modules
require an explicit purpose, ownership boundary, dependency review, package
catalog entry, full quality gates, documentation, changelog, license, security
policy, compatibility plan, and release dry-run.
