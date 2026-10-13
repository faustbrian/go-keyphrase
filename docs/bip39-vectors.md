# BIP-39 vector provenance and import

[Documentation index](README.md)

## Current snapshot authority

The maintained corpus is the exact `vectors.json` snapshot in the official
[Trezor python-mnemonic repository](https://github.com/trezor/python-mnemonic/tree/b57a5ad77a981e743f4167ab2f7927a55c1e82a8).
The import authority is immutable revision
`b57a5ad77a981e743f4167ab2f7927a55c1e82a8`, not a branch, release label or
installed runtime dependency. HTTPS verifies the source server; checked
SHA-256 locks verify the exact downloaded bytes. No signed-commit claim is made.

| Upstream path | Local path | Bytes | SHA-256 |
| --- | --- | --- | --- |
| `vectors.json` | `bip39/testdata/vectors.json` | 152400 | `fa3b937b7cff9c9b8ecd3aa011faeb8d6dd67993174b72326e83f4de8fdb30f8` |
| `LICENSE` | `bip39/testdata/python-mnemonic-LICENSE` | 1084 | `d5e3c7c62a84e80073201e2f6e5130e9e6804fa05f8ac4f8b26a13c7d3969697` |

The upstream MIT notice is retained unchanged. The corpus contains twelve
language arrays of twenty-four vectors; this library consumes the ten official
BIP-39 languages it supports. Russian and Turkish entries and the fourth-column
BIP-32 extended private keys remain in the snapshot but do not expand this
library's language or BIP-32 guarantees.

## Historical boundary

The existing local corpus, including the parentless first commit
`ab30bf09b537265aed18a00ac47beaaa7364faa4`, has exactly these vector bytes.
That identity establishes correspondence with the pinned upstream snapshot,
not who produced or imported the old local file. Repository history records
no historical generator revision or execution environment; that attribution
remains unknown. The archived source audit is preserved as historical evidence.

This migration makes the current and future **snapshot import** explicit without
changing any vector. It does not pretend to regenerate the multilingual corpus:
the pinned upstream `generate_vectors.py` generates English vectors only, and
its dependency ranges do not establish the original multilingual environment.

## Reproduce or update

Run from the repository root with Python 3 and HTTPS access:

```bash
python3 scripts/import-bip39-vectors.py
python3 scripts/import-bip39-vectors.py --write
python3 scripts/import-bip39-vectors.py
```

The default command downloads and verifies both pinned assets and checks local
byte identity without writing. `--write` verifies every input first, then uses
an atomic replacement for each asset; it is not a multi-file transaction or a
crash-durability promise. Downloads have a twenty-second socket timeout, exact
size/digest checks and refusal of redirects away from the pinned URL. Failed
input verification cannot replace either local asset; temporary files are
removed on normal failure paths. No upstream Python generator or third-party
Python dependency is executed.

For a different upstream revision, review the official commit, corpus changes
and license first; change the fixed revision and both asset locks in the import
script and this guide together. Run the importer twice and compare outputs.
Preserve the old audit and record semantic changes before adopting new bytes.
Changed vectors require the affected BIP-39 interoperability tests; an unchanged
snapshot/provenance update does not imply a new runtime conformance campaign or
module release.
