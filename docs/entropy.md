# Entropy interpretation

[Documentation index](README.md)

`Outcomes` is the exact number of distinct byte or code-point sequences the
configured generator can emit. `Bits` estimates `log2(Outcomes)` from the
leading at most 53 count bits plus a binary shift. Low-bit truncation and
floating-point rounding mean it is not an exact logarithm or certified floor.

For an unconstrained alphabet of size `A` and length `L`, outcomes are `A^L`.
For required classes, `password` counts the full valid language with dynamic
programming, then samples a rank uniformly from that language. It does not use
the misleading `L*log2(A)` value when constraints exclude outputs.

For passphrases, outcomes are `listSize^wordCount` multiplied by the exact
outcomes of independent affixes. A casing transform is accepted only when it
keeps list entries unique. Separator collisions are rejected because parsing
would otherwise be ambiguous.

Entropy is not a password-strength score. Reuse, exposure, user edits,
predictable injected randomness, biased hardware, downstream normalization,
rate limits, and attacker knowledge change practical guessing resistance.
`MinimumEntropyBits` rejects policies below a caller-selected mathematical
floor using that floating-point estimate; it does not choose an appropriate
floor for the caller's threat model. `password.Validate` checks password
content but does not compare the output count against that floor; use
`password.Analyze` for policy admission.

Continue with [API reference](api.md).
