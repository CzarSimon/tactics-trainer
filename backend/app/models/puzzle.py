# Standard libraries
from typing import Any, Dict, List

# External modules
from sqlmodel import SQLModel

# Internal modules
from app.models.dto import PuzzleDTO


class Puzzle(SQLModel):
    id: str
    external_id: str
    fen: str
    moves: str
    rating: int
    rating_deviation: int
    popularity: int
    number_of_plays: int
    themes: str
    game_url: str

    @classmethod
    def from_dto(cls, dto: PuzzleDTO) -> "Puzzle":
        return cls(
            id=dto.id,
            external_id=dto.external_id,
            fen=dto.fen,
            moves=_encode_list(dto.moves),
            rating=dto.rating,
            rating_deviation=dto.rating_deviation,
            popularity=dto.popularity,
            number_of_plays=dto.number_of_plays,
            themes=_encode_list(dto.themes),
            game_url=dto.game_url,
        )


def _encode_list(l: List[str]) -> str:
    return " ".join(l)
