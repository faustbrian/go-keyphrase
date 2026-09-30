# BIP-39 behavior

`bip39` supports 128, 160, 192, 224, and 256 bits of entropy and the resulting
12, 15, 18, 21, and 24-word mnemonics. It creates and validates checksum bits,
normalizes mnemonic and passphrase text with NFKD, uses all ten official lists,
and derives a 64-byte seed with PBKDF2-HMAC-SHA512, 2,048 rounds, and salt
`"mnemonic" + passphrase` exactly as specified.

Seed derivation uses Go's maintained `crypto/pbkdf2` primitive. It checks
context cancellation before and after the fixed derivation, not between rounds.
In strict `GODEBUG=fips140=only` mode, the primitive rejects the BIP-39 salt
when an empty or short passphrase leaves it below its minimum length. `Seed`
returns a typed `CodeDerivation` error and no seed; it does not bypass FIPS
enforcement or substitute a non-BIP-39 salt. This mode is not a supported
general BIP-39 seed-derivation environment. Applications needing both BIP-39
and strict FIPS-only operation must resolve that policy conflict outside this
library rather than retrying with altered parameters.

`Parse` checks vocabulary across every official list before checksum
validation. If all words occur in multiple lists, it returns
`CodeAmbiguousLanguage` and safe language candidates instead of guessing.
`ParseLanguage` is appropriate when the surrounding protocol already supplies
the language. Whitespace is normalized for parsing; `Mnemonic.String` renders
Japanese vectors with ideographic spaces and other languages with ASCII spaces.

The tests consume the complete vector set pinned from Trezor's
`python-mnemonic` repository and compare English mnemonic and seed output with
the independently maintained `tyler-smith/go-bip39` implementation.

BIP-39 is included only for mnemonic and seed interoperability. This module
does not implement BIP-32, BIP-44, wallets, addresses, private keys,
transactions, custody, account discovery, or chain-specific behavior.
