# Changelog

All notable changes will be documented here. The project follows Semantic
Versioning after its first stable release.

## Unreleased

### Changed

- Align isolated dependency checks with standalone package module paths.

### Documentation

- Add a repository-local documentation index, remove completed implementation
  plans, and clarify the pending independent cryptographic review.

## 1.0.0 - 2026-08-25

### Changed

- Exclude intentional nested modules from root local-proxy archives so local,
  bootstrap, CI, and public module checksums describe the same source
  boundary.

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

### Changed

- Publish the module from its standalone `github.com/faustbrian/go-keyphrase` identity while preserving its documented API and behavior.

### Documentation

- Replace obsolete repository links and workflow claims with standalone
  package targets and current release guidance.

- Add package discovery documentation.

- Delegate local mutation checks to the canonical exact-100 repository runner
  and remove the superseded Gremlins installation and configuration.
- Add unbiased password, passphrase, and BIP-39 generation.
- Add pinned EFF and official BIP-39 word lists.
- Add typed failures, exact entropy, resource limits, vectors, fuzzing,
  benchmarks, documentation, and local release gates.
