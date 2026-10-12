# Source documentation audit

## Scope and reproduction

Documentation-only audit of base `b3e74310dd060bcb36a584b50c024e5616b800d2`.
One root module, `github.com/faustbrian/go-keyphrase/v2`, owns seven packages:
root keyphrase, bip39, keyphrasetest, passphrase, password, wordlist and
wordlist/eff. External test packages remain included. Reproduce the complete
tracked inventory with `git ls-files '*.go'` and
`golib source-docs inventory --json`; module/package declarations and native
gate selection are in modules.json, packages.json and .golib.yaml.

Before source edits, initial inventory reported 58 missing public field/method
comments. Existing APIs are already substantially documented; the semantic
work is entropy precision and source-quality assumptions, exact normalization
and lookup identities, cancellation checkpoints, retained/aliased secret state,
metadata authentication limits, operation-specific failure/buffer semantics,
and assurance-oracle/workload rationale. These are Tier A contract
clarifications, not changes to generation, validation or cryptography.

Direct consumers compose selectors with password/passphrase/BIP-39 generators,
validate and index immutable lists, and expose secrets at reviewed boundaries.
Named risks are misleading strength claims, untrusted provenance, hidden alias
assumptions, inability to interrupt supplied reads and incorrect Unicode
identity. Smallest checks: format, native AST/token/directive equivalence,
render all seven public packages, native docs/examples, exact native go-doc
baseline generation and stable regeneration, local links, changed-content
markers/secrets and ordinary complete final diff. No unchanged runtime,
race, fuzz, statistical, interoperability, strict-FIPS or benchmark campaign.

## Source-backed rationale and caveats

- Rejection sampling removes modulo bias only when source bytes are uniform
  and independent. Custom Source quality, cancellation and concurrent-use
  safety remain caller-owned. Selector stores the supplied interface; options
  run during construction and are not retained. Source errors fail even when
  returned with a full count; zero/invalid progress fails without retrying.
  The attempt cap counts candidate draws, not individual partial-read calls.
- Context checks precede each source read; default crypto/rand reads cannot
  be interrupted by this package. Cancellation can arrive after a final read
  and still yield success. Password preparation/unranking and passphrase
  preparation/assembly have no checkpoints. Contexts must be nonnil when
  reached; a valid Selector Fill with an empty destination does no read or
  context check. Seed checks before/after fixed PBKDF2, clearing canceled output.
- Constrained output counts use exact big.Int dynamic programming and unranking
  rather than forcing classes at particular positions. Bits uses the leading
  at most 53 count bits plus a shift and floating log2; it is an estimate,
  not an exact arbitrary-precision logarithm or cryptographic certification.
  Threshold checks use that estimate. Password Validate checks content/class
  membership, not whether the distribution meets MinimumEntropyBits; call
  Analyze for that admission check. Passphrase Parse does prepare its policy,
  including the entropy floor, but a parsed phrase is not proof of random origin.
- Unicode policies operate on code points, not graphemes or encoded byte length;
  alphabets reject per-symbol NFKD collisions but outputs are not normalized.
  Byte policies preserve arbitrary bytes including zero/non-UTF-8. Callers
  must not mutate policy slices, class slices or affix pointers during a call.
- Wordlist validation preserves raw entry order/identity and newline-terminated
  checksum bytes. NFKD collisions are always rejected; WithNFKD checks admission,
  not rewriting. Unique prefixes use NFC display code points (via NFKD/NFC),
  not raw byte prefixes. Index is exact, without normalization/case folding.
  SourceSHA256 is checksum-shaped provenance metadata, not a downloaded-source
  verification performed by New. A caller-created matching checksum does not
  establish dictionary authority, quality or cryptographic randomness.
- Immutable list accessors copy the word slice; string contents remain immutable.
  BIP-39 and EFF loaders memoize shared immutable lists and errors. Selector and
  generators add no synchronization; source concurrency safety controls sharing.
  Deterministic test sources have a mutex and finite/cyclic state, never security
  suitability; their context check does not preempt a lock wait.
- Secret is a slice, not an immutable value: assignment/conversion aliases its
  buffer, Clear affects its visible length and those aliases only, not spare
  capacity. Formatting/text/JSON/slog redaction
  does not prevent explicit extraction, binary encoders, or arbitrary wrappers
  exposing raw bytes. Mnemonic stores private entropy without a wipe method;
  Entropy returns a copy whose clearing does not clear the Mnemonic. Parsed
  word strings cannot be wiped. All erasure remains best effort.
- BIP-39 uses official entropy sizes and SHA-256 checksum packing; Parse checks
  all-list vocabulary ambiguity before checksum validity. Parsing is NFKD plus
  Unicode Fields, case-sensitive. Seed uses NFKD, the exact mnemonic+passphrase
  salt, SHA-512 PBKDF2/2048 rounds/64 bytes; a wrong passphrase still yields a
  seed without authentication. Strict FIPS-only short-salt rejection is a typed
  derivation error, not a reason to substitute parameters. No wallet/custody
  capability, modern password-storage KDF claim or secure-erasure claim exists.
- Count, symbol-byte, state-cell and grouped-mask operation limits bound named
  work, not total process memory. List and phrase ingress must also be bounded
  by the application. ChiSquared assumes its uint64 count sum is representable;
  it is a statistic, not a p-value, cryptographic-source test or overflow checker.

Evidence is the production sources and deterministic assertions listed below,
docs/entropy.md, docs/wordlists.md, docs/bip39.md, docs/errors-and-limits.md,
and the security review's stated source/residual-risk boundaries. Historical
review and release evidence are immutable; this pass does not enlarge them.

## Provenance, variants and native operations

There are no tracked generated Go files, generator commands, build tags or
platform-suffix variants. Four embed directives bind ten BIP-39 dictionaries
and three EFF dictionaries. Bitcoin BIPs revision
`8c369ac8e60629ac6c032ffe21bb5ec5b35213d7` and per-file hashes are in
bip39/lists.go; EFF source rows are transformed to their word columns with raw
and transformed pins in wordlist/eff/eff.go. No runtime network refresh occurs.
The repository's third-party notices and word-list guide retain licensing and
source/update instructions. No new corpus producer or regenerated asset is needed.
The committed BIP-39 vectors are attributed to python-mnemonic and compared
with the independent Go oracle. No separate upstream vector revision, checksum
or producer is recorded in tracked repository metadata. This is an unresolved
provenance question for a future corpus update, not an invented verified pin;
corpus bytes and historical evidence remain unchanged.

API mode is go-doc. Concatenate `go doc -all` stdout in this exact configured
order: `.`, `./bip39`, `./keyphrasetest`, `./passphrase`, `./password`,
`./wordlist`, `./wordlist/eff`; remove terminal newlines and append one newline.
Update api/baseline.txt with that native generation, not signature-only output.
Native docs runs `go test -run '^Example' ./...`. No additional package-owned
snapshot generator exists. Go is pinned to 1.27.2; retain the /v2 module identity.

Four fuzz targets cover bounded Unicode/class policies with generated-value
validation, lists with count/checksum preservation, canonical BIP-39 parse
round trips, and a fixed three-word passphrase parse count. Deterministic zero
sources do not constitute distribution/security oracles. Ten benchmarks cover
sampling, constrained/large/rejected/source-failure password policies, EFF
phrase generation/parsing and BIP-39 encode/normalize/seed; comments identify
setup and timed work. Three examples expose only public deterministic fixtures or
boolean checks of generated values. Brute-force small policy enumeration, independent
PBKDF2/mnemonic comparisons, literal normalization fixtures, error-redaction
assertions and checkpoint-controlled sources have distinct purposes. No saved
active-fuzz crash corpus or unfinished implementation marker is tracked.

## Tracked Go coverage

All tracked Go sources, API declarations/fields/interface methods, test
fixtures/assertions and variant directives are audited; no file is excluded
because it is private or test-only.

| File | Audit boundary |
| --- | --- |
| `a_random_guard_test.go` | Public/internal failure partitions, ownership and invariant assertions |
| `benchmark_test.go` | Workload, setup boundary and output lifetime |
| `bip39/benchmark_test.go` | Workload, setup boundary and output lifetime |
| `bip39/bip39.go` | Checksum/bit packing, normalization, ambiguity, seed and secret state |
| `bip39/bip39_internal_test.go` | Public/internal failure partitions, ownership and invariant assertions |
| `bip39/bip39_test.go` | Public/internal failure partitions, ownership and invariant assertions |
| `bip39/example_test.go` | Executable public workflow |
| `bip39/fuzz_test.go` | Bounded mutation inputs and stated oracle |
| `bip39/interoperability_test.go` | Public/internal failure partitions, ownership and invariant assertions |
| `bip39/lists.go` | Ten official embedded lists, identity/checksums and once-owned loading |
| `bip39/normalization_contract_test.go` | Public/internal failure partitions, ownership and invariant assertions |
| `concurrency_test.go` | Public/internal failure partitions, ownership and invariant assertions |
| `doc.go` | Canonical source-quality/ownership/cancellation contract |
| `errors.go` | Secret-safe category/cause formatting and wrapping |
| `example_test.go` | Executable public workflow |
| `keyphrasetest/source.go` | Mutex-owned finite/cyclic fixture sources and statistic constraints |
| `keyphrasetest/source_test.go` | Public/internal failure partitions, ownership and invariant assertions |
| `observation_test.go` | Public/internal failure partitions, ownership and invariant assertions |
| `passphrase/benchmark_test.go` | Workload, setup boundary and output lifetime |
| `passphrase/coverage_internal_test.go` | Public/internal failure partitions, ownership and invariant assertions |
| `passphrase/fuzz_test.go` | Bounded mutation inputs and stated oracle |
| `passphrase/passphrase.go` | List/casing/separators, independent affixes, parse and count semantics |
| `passphrase/passphrase_test.go` | Public/internal failure partitions, ownership and invariant assertions |
| `password/benchmark_test.go` | Workload, setup boundary and output lifetime |
| `password/byte.go` | Distinct byte policy, copied preparation and buffer semantics |
| `password/byte_test.go` | Public/internal failure partitions, ownership and invariant assertions |
| `password/coverage_internal_test.go` | Public/internal failure partitions, ownership and invariant assertions |
| `password/example_test.go` | Executable public workflow |
| `password/fuzz_test.go` | Bounded mutation inputs and stated oracle |
| `password/password.go` | Exact constrained counts/unranking and Unicode policy contracts |
| `password/password_test.go` | Public/internal failure partitions, ownership and invariant assertions |
| `password/property_test.go` | Public/internal failure partitions, ownership and invariant assertions |
| `password/resource_internal_test.go` | Public/internal failure partitions, ownership and invariant assertions |
| `random.go` | Source interface, rejection sampling, bounds and cancellation |
| `random_internal_test.go` | Public/internal failure partitions, ownership and invariant assertions |
| `random_test.go` | Public/internal failure partitions, ownership and invariant assertions |
| `secret.go` | Mutable buffer aliasing, redaction and best-effort clearing |
| `secret_test.go` | Public/internal failure partitions, ownership and invariant assertions |
| `wordlist/eff/eff.go` | Three pinned embedded dictionaries and shared lazy loading |
| `wordlist/eff/eff_test.go` | Public/internal failure partitions, ownership and invariant assertions |
| `wordlist/fuzz_test.go` | Bounded mutation inputs and stated oracle |
| `wordlist/list.go` | Raw/NFKD/NFC identity, validated metadata, copies and limits |
| `wordlist/list_internal_test.go` | Public/internal failure partitions, ownership and invariant assertions |
| `wordlist/list_test.go` | Public/internal failure partitions, ownership and invariant assertions |
| `wordlist/normalization_contract_test.go` | Public/internal failure partitions, ownership and invariant assertions |

## Disposition

Coverage: all 45 tracked Go files (12 implementation/test-support, 33 assurance/
example files), one module, seven package directories and 222 public declarations/
members, including Source.ReadContext, plus the private cryptoSource adapter
method. No executable token, parsed AST or compiler/security/embed directive
changed against the exact base. All tracked files are formatted by Go 1.27.2.
All seven packages render with native go doc; api/baseline.txt was generated in
configured order twice with byte-identical results. The native docs selector
ran the three examples; the other four packages have no executable examples.
Source documentation inventory/check has zero findings. All thirteen embedded
dictionary hashes/counts match existing pins. Local links, unfinished-marker
and changed-current-file secret checks, full final diff and ACK metadata pass.
Task-owned Go caches and disposable check artifacts are removed. No runtime fix or new
security guarantee is introduced. Random-source quality, ingress resources,
secret lifetime and downstream normalization remain application-owned. The vector-origin revision/checksum/producer question remains recorded above;
no other unresolved source-policy question has been identified. Hosted CI, source
publication, release qualification and broad runtime assurance are separate
boundaries not established by this documentation-only audit.
