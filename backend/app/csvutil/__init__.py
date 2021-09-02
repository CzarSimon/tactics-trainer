# Standard library
from contextlib import contextmanager
from csv import DictReader, DictWriter
from typing import Any, Dict, List, Iterator, Generator, Optional


class Reader:
    def __init__(self, filename: str, delimiter: str) -> None:
        self._file = open(filename, "r", encoding="utf-8")
        self._csv_reader = DictReader(self._file, delimiter=delimiter)

    def __iter__(self) -> "Reader":
        return self

    def __next__(self) -> Dict[str, str]:
        return self._csv_reader.__next__()

    def close(self) -> None:
        self._file.close


@contextmanager
def get_reader(filename: str, delimiter: str = ",") -> Generator[Reader, None, None]:
    reader = Reader(filename, delimiter)
    yield reader
    reader.close()
