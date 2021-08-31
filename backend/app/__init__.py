# External modules
from fastapi import FastAPI

# Internal models
from app.models.dto import TacticDTO


app = FastAPI()


@app.get("/v1/tactics/{id}", response_model=TacticDTO)
def get_tactic(id: str) -> TacticDTO:
    return TacticDTO(id=id, fen="FEN", url="https://domain.com")
