#!/usr/bin/env python3
"""Import an exact authenticated upstream snapshot, not regenerate its vectors."""

import argparse
import hashlib
import os
from pathlib import Path
import tempfile
import urllib.request

ROOT = Path(__file__).resolve().parent.parent
REVISION = "b57a5ad77a981e743f4167ab2f7927a55c1e82a8"
SOURCE = f"https://raw.githubusercontent.com/trezor/python-mnemonic/{REVISION}/"
# These are reviewed source locks, not evidence of the historical local producer.
ASSETS = (
    ("vectors.json", "bip39/testdata/vectors.json", 152400,
     "fa3b937b7cff9c9b8ecd3aa011faeb8d6dd67993174b72326e83f4de8fdb30f8"),
    ("LICENSE", "bip39/testdata/python-mnemonic-LICENSE", 1084,
     "d5e3c7c62a84e80073201e2f6e5130e9e6804fa05f8ac4f8b26a13c7d3969697"),
)


def download(source, size, digest):
    """Reject changed bytes, oversized bodies and redirects away from the pin."""
    url = SOURCE + source
    with urllib.request.urlopen(url, timeout=20) as response:
        if response.geturl() != url:
            raise ValueError(f"unexpected redirect for {source}")
        data = response.read(size + 1)
    if len(data) != size or hashlib.sha256(data).hexdigest() != digest:
        raise ValueError(f"source lock mismatch for {source}")
    return data


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--write", action="store_true",
                        help="atomically replace each local asset after checking all inputs")
    args = parser.parse_args()
    # Authenticate every input before touching any local asset. Each replacement
    # is atomic; this is not a multi-file transaction or crash-durability promise.
    assets = [(target, download(source, size, digest))
              for source, target, size, digest in ASSETS]
    for target, data in assets:
        path = ROOT / target
        if args.write:
            temporary = None
            try:
                with tempfile.NamedTemporaryFile(dir=path.parent, delete=False) as output:
                    temporary = Path(output.name)
                    output.write(data)
                    output.flush()
                    os.fchmod(output.fileno(), 0o644)
                os.replace(temporary, path)
            finally:
                if temporary is not None:
                    temporary.unlink(missing_ok=True)
        if path.read_bytes() != data:
            raise ValueError(f"local asset differs from source lock: {target}")
    print(f"Verified exact python-mnemonic snapshot {REVISION}")


if __name__ == "__main__":
    main()
