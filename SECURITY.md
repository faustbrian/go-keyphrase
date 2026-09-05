# Security policy

Do not open a public issue for a suspected vulnerability. Report it privately
through [GitHub Security Advisories for
`faustbrian/go-keyphrase`](https://github.com/faustbrian/go-keyphrase/security/advisories/new)
and include the affected version or commit, reproduction, impact, and proposed
mitigation when available.

Generated values, mnemonics, entropy, passphrases, seeds, private keys, and
production randomness traces must not be included in reports. Replace them with
deterministic public fixtures.

Published releases follow [`COMPATIBILITY.md`](COMPATIBILITY.md) and
[`DEPRECATION.md`](DEPRECATION.md). At and after v1, incompatible exported API
or documented behavior changes require a new major version. The independent
cryptographic design review remains pending; its status is tracked in
[`docs/security-review.md`](docs/security-review.md).
