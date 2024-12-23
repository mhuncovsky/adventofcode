import os
from pathlib import Path

def readlines(path: os.PathLike | str) -> list[str]:
    return Path(path).read_text().splitlines()

