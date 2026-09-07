# Performance and operational limits

Keyphrase generation is local and bounded, but it is not constant-cost. This
repository publishes benchmark entry points rather than a latency or throughput
guarantee. Measure the relevant policy, randomness source, platform, and Go
version in the consuming environment.

## Cost model

- Random selection uses unbiased rejection sampling. Each selection samples at
  most the selector's configured attempt limit, but completing one sample may
  require multiple randomness-source calls when a custom source returns partial
  progress. Source latency and contention therefore directly affect generation
  latency.
- Password policy analysis builds a bounded dynamic-programming table. Work and
  allocation grow with password length, alphabet size, and required classes;
  the limits in [errors and resource limits](errors-and-limits.md) reject work
  beyond the supported cell and operation budgets.
- Passphrase generation performs one bounded selection per word plus any prefix
  and suffix password work. Output allocation grows with the selected word and
  affix sizes and is capped before generation.
- The first BIP-39 list lookup loads and validates the embedded official lists
  once for the process. Later lookups reuse immutable validated lists. Seed
  derivation always performs the specification's 2,048 PBKDF2-HMAC-SHA512
  rounds and returns a 64-byte caller-owned result.

## Cancellation and resources

Operations run synchronously and create no goroutines or background work.
Random-source loops check contexts between reads, and BIP-39 seed derivation
checks periodically between rounds. An in-flight randomness read can return
promptly on cancellation only when the selected `Source` honors its
`ReadContext` contract; callers own timeout policy and any custom source
lifecycle.

Generated secrets and caller destinations consume memory proportional to their
bounded output size. Temporary buffers are cleared on failure and after use
where the API permits, but Go and the operating system may retain copies. Do
not use allocation counts or clearing as proof that secret bytes cannot remain
in memory; follow the [secret lifetime guidance](secret-lifetime.md).

Default selectors and generators may be shared concurrently. Custom sources
must provide their own concurrency safety, and callers must use separate
mutable destination buffers for overlapping operations. The package owns no
external resource and requires no close or shutdown step.

## Reproducing benchmarks

Run package benchmarks on the target environment:

```sh
go test ./... -run '^$' -bench . -benchmem
```

Record the Go version, platform, policy, randomness source, and whether the
result includes first-use BIP-39 list initialization when comparing results.
