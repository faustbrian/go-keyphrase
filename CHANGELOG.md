# Changelog

All notable changes will be documented here. The project follows Semantic
Versioning after its first stable release.

## Unreleased

## 2.0.1 - 2026-10-02

### Changed

- Refresh the reusable CI workflow while retaining the verified v1.4.0
  tooling and the package-owned verification contract.

### Documentation

- Correct support and troubleshooting guidance for disabled public issue and
  discussion channels, linking the documentation and private security route.

## 2.0.0 - 2026-09-30

### Migration

- Move the public module and all package imports to
  `github.com/faustbrian/go-keyphrase/v2` for the changed BIP-39 seed
  cancellation contract. Replace v1 imports with the corresponding `/v2`
  paths; seed bytes remain specification-compatible, but cancellation is now
  observed before and after the fixed derivation rather than during its rounds.

### Changed

- Adopt the checksum-verified `go-library-tools` v1.4.0 CLI and immutable W14
  reusable workflow, and align local configuration, inventory, cohesion,
  repository, online specification, workflow, and implementation gates without
  changing the public API or runtime behavior.

- Adopt the checksum-verified `go-library-tools` v1.3.0 CLI, schema-v2 cohesion
  metadata, and repository-local cohesion gate without changing the public API
  or runtime behavior.
- Pin reusable CI to the immutable v1.3.0 cohesion-enforcement workflow and add
  versioned Golib ecosystem navigation.
- Adopt the versioned shared `golib` repository contract for local and hosted
  verification while retaining package-owned API and mutation evidence.
- Align isolated dependency checks with standalone package module paths.
- Delegate BIP-39 seed derivation to the standard-library `crypto/pbkdf2`
  primitive, failing closed if it rejects the specified parameters.
- Report strict FIPS-only rejection of BIP-39's short salt as `CodeDerivation`
  without a seed; this environment cannot be treated as generally supported
  for BIP-39 seed derivation.
- Update the production `x/text` dependency to v0.41.0 and pin the test
  oracle's transitive `x/crypto` dependency to v0.56.0.

### Documentation

- Document lifecycle, ownership, concurrency, shutdown, and bounded performance
  behavior, and point module metadata to the dedicated performance guide.

- Document stable-v1 maturity and portable-Go boundaries, link the executable
  example and complete support navigation, and correct the security reporting
  route and post-v1 compatibility wording.

- Link ecosystem and Domain utilities family guidance to the immutable v1.4.0
  documentation release.

- Add a repository-local documentation index, remove completed implementation
  plans, and clarify the pending independent cryptographic review.

## 1.0.0 - 2026-08-26

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
