# Standard library
from typing import List

# Internal modules
from app.csvutil import get_reader
from app.models.dto import PuzzleDTO


def import_puzzles() -> None:
    with get_reader("../lichess_db_puzzle.csv") as reader:
        for row in reader:
            print(PuzzleDTO.from_dict(row))


import_puzzles()
