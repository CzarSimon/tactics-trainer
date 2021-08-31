# Standard library
from typing import List

# Internal modules
from app import csvutil
from app.models.dto import PuzzleDTO


def read_puzzles() -> List[PuzzleDTO]:
    rows = csvutil.read("../lichess_db_puzzle.csv")
    return [PuzzleDTO.from_dict(row) for row in rows]
