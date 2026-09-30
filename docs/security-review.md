# Independent cryptographic design review

Status: approved for the reviewed v2 source candidate

Reviewer: Codex, independent read-only reviewer

Organization: OpenAI

Review date: 2026-09-30

Reviewed commit: 4185b83263be3b13aa26ff0ec84c8dfd5d16c174

Scope: rejection sampling; constrained distribution counting and unranking;
entropy claims; BIP-39 bit packing, NFKD normalization, PBKDF2 parameters,
official vectors, and Japanese interoperability; word-list provenance; error
and secret disclosure; resource bounds; cancellation and v1-to-v2 compatibility;
concurrency; tests and release gates; and strict-FIPS short-salt behavior.

Findings: No unresolved findings. The prior v1 cancellation-contract
incompatibility and undocumented strict-FIPS short-salt behavior were resolved
through the v2 module transition, explicit migration guidance, fail-closed
derivation errors, and focused regression coverage.

Resolutions: The candidate uses Go 1.27 `crypto/pbkdf2` with the BIP-39 NFKD,
salt, SHA-512, iteration, and output parameters; checks cancellation before and
after fixed derivation; preserves specification-compatible seed bytes; and
consistently exposes the changed contract through `/v2`.

Residual risks: Go cannot guarantee erasure of strings or copied secret
buffers. Injected sources remain caller-trusted for cryptographic quality,
prompt cancellation, and concurrency safety. PBKDF2 cannot be interrupted
mid-operation. Strict FIPS-only mode rejects short BIP-39 salts. Entropy counts
do not establish guessing resistance after disclosure, reuse, user
modification, or downstream normalization.

Approval: Approved for the reviewed v2 source candidate; this is an
independent agent review, not a human third-party certification.

The existing `v1.0.0` tag predates this review and is not evidence of its
completion. The reviewed source is the commit above; later documentation-only
changes do not expand the reviewed cryptographic scope. A future release still
requires the applicable source, CI, and release gates.
