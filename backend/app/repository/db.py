# External modules
import databases
from sqlmodel import create_engine


_DATABASE_URL: str = "sqlite:///./test.db"
database = databases.Database(_DATABASE_URL)
engine = create_engine(_DATABASE_URL)
