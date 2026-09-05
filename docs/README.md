# Documentation

`keyphrase` generates passwords, EFF-list passphrases, and BIP-39 mnemonics
from caller-controlled policies and cryptographic randomness.

The module is stable at v1, requires Go 1.26.6 or later, and has no
platform-specific source files or external runtime backend. See the
[compatibility policy](../COMPATIBILITY.md) for the supported public contract.

## Getting started

- [Adoption](adoption.md)
- [Executable passphrase example](../example_test.go)
- [Policy examples](policies.md)
- [API reference](api.md)
- [FAQ](faq.md)
- [Troubleshooting](troubleshooting.md)

## Security model

- [Threat model](threat-model.md)
- [Entropy interpretation](entropy.md)
- [Secret lifetime](secret-lifetime.md)
- [Errors and resource limits](errors-and-limits.md)
- [Independent review status](security-review.md)

## Formats and data

- [BIP-39 behavior](bip39.md)
- [Word-list provenance](wordlists.md)

## Maintainers

- [Testing and release gates](testing.md)

## Project and release information

- [Support](../SUPPORT.md)
- [Security reporting](../SECURITY.md)
- [Compatibility](../COMPATIBILITY.md)
- [Deprecation policy](../DEPRECATION.md)
- [Changelog](../CHANGELOG.md)
- [License](../LICENSE)
- [Third-party notices](../THIRD_PARTY_NOTICES.md)
