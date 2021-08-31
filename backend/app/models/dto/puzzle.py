# Standard libraries
from typing import Any, Dict, List

# External modules
from pydantic import BaseModel

# Internal modules
from app import idutil


class StrModel(BaseModel):
    def __str__(self) -> str:
        return super().__repr__()


class PuzzleDTO(StrModel):
    id: str
    external_id: str
    fen: str
    moves: List[str]
    rating: int
    rating_deviation: int
    popularity: int
    number_of_plays: int
    themes: List[str]
    game_url: str

    @classmethod
    def from_dict(cls, raw: Dict[str, Any]) -> "PuzzleDTO":
        external_id: str = raw["PuzzleId"]
        fen: str = raw["FEN"]
        rating: int = int(raw["Rating"])
        rating_deviation: int = int(raw["RatingDeviation"])
        popularity: int = int(raw["Popularity"])
        number_of_plays: int = int(raw["NbPlays"])
        moves: List[str] = raw["Moves"].split(" ")
        themes: List[str] = raw["Themes"].split(" ")
        game_url: str = raw["GameUrl"]
        return cls(
            id=idutil.new(),
            external_id=external_id,
            fen=fen,
            rating=rating,
            rating_deviation=rating_deviation,
            popularity=popularity,
            number_of_plays=number_of_plays,
            moves=moves,
            themes=themes,
            game_url=game_url,
        )
