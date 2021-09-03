# External modules
from fastapi import FastAPI

# Internal models
from app.models.dto import PuzzleDTO
from app.repository import database


app = FastAPI()


@app.get("/v1/puzzles/{id}", response_model=PuzzleDTO)
async def get_tactic(id: str) -> PuzzleDTO:
    return PuzzleDTO(id=id, fen="FEN", url="https://domain.com")


@app.on_event("startup")
async def startup() -> None:
    await database.connect()


@app.on_event("shutdown")
async def teardown() -> None:
    await database.disconnect()
