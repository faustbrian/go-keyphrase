# Migration and adoption

## From v1 of this module

Change module imports to `github.com/faustbrian/go-keyphrase/v2` and its
`/v2/bip39`, `/v2/password`, `/v2/passphrase`, `/v2/wordlist`,
`/v2/wordlist/eff`, and `/v2/keyphrasetest` packages. Do not add a `replace`
directive or a version-specific source directory. `bip39.Seed` retains BIP-39
seed bytes but uses the maintained Go PBKDF2 primitive, so cancellation is
observed before and after its fixed 2,048 rounds instead of during the rounds.
Under strict `GODEBUG=fips140=only`, BIP-39's short salt is rejected by that
primitive; see [BIP-39 behavior](bip39.md) before adopting in that environment.

## From ad-hoc random strings

Inventory the old alphabet, length semantics, required classes, exclusions,
encoding, and downstream normalization. Express them as a `password.Policy`,
compare `Analyze` with the old claim, and reject configuration that cannot meet
the intended floor. Replace modulo selection and post-generation repair with a
single `Generator` call. Keep the old format accepted only where stored values
must remain compatible.

## From another passphrase library

Pin the exact list bytes and metadata before migration. A list reorder changes
deterministic fixtures and may change compatibility even when words are equal.
Validate separators and casing, compare exact outcome counts, and test parsing
of existing phrases before switching generation.

## From another BIP-39 implementation

Run the repository interoperability vectors for every deployed language and
passphrase normalization case. Treat mnemonic text, passphrase, and seed as
secrets. Keep wallet, address, BIP-32, and custody logic in the existing wallet
component; replace only the BIP-39 seam.

## Rollout

Adopt behind an application-owned interface, inject deterministic sources only
in tests, observe error codes rather than values, establish output-lifetime
rules, run `make check`, and obtain security review before a future release.
