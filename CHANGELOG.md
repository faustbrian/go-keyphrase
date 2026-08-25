# Changelog

All notable changes will be documented here. The project follows Semantic
Versioning after its first stable release.

## Unreleased

### Changed

- Track the pinned documentation-tool lockfile so clean CI checkouts install
  the exact validated cspell dependency.

- Reconcile standalone dependency checksums against deterministic current
  module archives so CI, local verification, and release consumers resolve
  identical content.

- Harden standalone documentation validation with deterministic spelling and
  link checks, package-specific documentation gates, and repository-local
  contributor guidance.

### Documentation

- Correct stale package, standalone, and authoritative-source links in public
  documentation.

## 1.0.0 - 2026-08-25

### Changed

- Publish the module from its standalone `github.com/faustbrian/go-keyphrase` identity while preserving its documented API and behavior.

### Documentation

- Replace obsolete standalone-repository links and workflow claims with
  monorepo-canonical targets and current release guidance.

- Link the package README to the repository-wide Golib documentation portal.

- Delegate local mutation checks to the canonical exact-100 repository runner
  and remove the superseded Gremlins installation and configuration.
- Add unbiased password, passphrase, and BIP-39 generation.
- Add pinned EFF and official BIP-39 word lists.
- Add typed failures, exact entropy, resource limits, vectors, fuzzing,
  benchmarks, documentation, and local release gates.
