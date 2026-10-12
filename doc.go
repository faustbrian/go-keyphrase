// Package keyphrase provides shared cryptographic randomness and entropy
// contracts for password, passphrase, and BIP-39 mnemonic generation.
//
// It does not hash passwords, manage wallets, derive wallet keys, store
// secrets, or distribute secret material. Exact output counts describe the
// configured distribution, not resistance after disclosure, edits or downstream
// normalization. Uniform generation relies on a uniform, independent Source;
// custom source quality, cancellation and concurrency safety are caller-owned.
//
// Contexts must be nonnil when checked. Cancellation is cooperative around
// source reads; password preparation/unranking and phrase assembly are not
// interruptible. BIP-39 seed derivation checks before and after fixed PBKDF2.
// Secret buffers are caller-owned but slice aliases share storage; clearing
// is best effort and does not erase immutable strings or a Mnemonic's retained
// entropy. The package creates no background goroutines or shutdown owner.
package keyphrase
