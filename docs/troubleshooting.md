# Troubleshooting

## Installation or compilation fails

Confirm that `go version` reports Go 1.26.6 or later and that the imported
module path is `github.com/faustbrian/go-keyphrase`. The public packages have
no platform-specific source files or external runtime service requirement.

## Generation returns an error

Inspect the typed error code rather than matching its message. The
[errors and resource limits guide](errors-and-limits.md) maps the stable error
categories and bounded-input rules. Do not include generated values, entropy,
mnemonics, passphrases, seeds, or production randomness traces in logs or
support reports.

## BIP-39 language detection is ambiguous

Shared vocabulary can make automatic detection ambiguous. Select the expected
language explicitly with `bip39.ParseLanguage`; see the [BIP-39 guide](bip39.md)
and [FAQ](faq.md).

## More help

Use [GitHub Issues](https://github.com/faustbrian/go-keyphrase/issues) for a
reproducible defect and [GitHub
Discussions](https://github.com/faustbrian/go-keyphrase/discussions) for
adoption questions. Report suspected vulnerabilities only through the private
process in [SECURITY.md](../SECURITY.md).
