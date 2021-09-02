# External modules
from fastapi import FastAPI

# Internal models
from app.models.dto import PuzzleDTO


app = FastAPI()


@app.get("/v1/puzzles/{id}", response_model=PuzzleDTO)
async def get_tactic(id: str) -> PuzzleDTO:
    return PuzzleDTO(id=id, fen="FEN", url="https://domain.com")
