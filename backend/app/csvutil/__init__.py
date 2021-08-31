# Standard library
from csv import DictReader, DictWriter
from typing import Any, Dict, List


def read(filename: str, delimiter: str = ",") -> List[Dict[str, Any]]:
    with open(filename, "r", encoding="utf-8") as f:
        reader = DictReader(f, delimiter=delimiter)
        return [row for row in reader]
