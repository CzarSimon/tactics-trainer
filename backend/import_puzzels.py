# Standard library
from typing import List

# Internal modules
from app.csvutil import get_reader
from app.models import Puzzle
from app.models.dto import PuzzleDTO


def import_puzzles() -> None:
    with get_reader("../lichess_db_puzzle.csv") as reader:
        for row in reader:
            dto = PuzzleDTO.from_dict(row)
            puzzle = Puzzle.from_dto(dto)
            print(puzzle)


import_puzzles()
