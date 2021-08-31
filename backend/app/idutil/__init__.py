# Standard library
from uuid import uuid4


def new() -> str:
    return str(uuid4()).lower()
