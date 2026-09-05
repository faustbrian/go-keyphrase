# Changelog

All notable changes will be documented here. The project follows Semantic
Versioning after its first stable release.

## Unreleased

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

### Documentation

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
