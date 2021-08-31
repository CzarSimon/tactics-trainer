# External modules
from pydantic import BaseModel


class TacticDTO(BaseModel):
    id: str
    fen: str
    url: str
