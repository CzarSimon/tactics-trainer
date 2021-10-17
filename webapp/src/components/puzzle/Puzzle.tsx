import React, { useEffect } from 'react';
import Chessboard from 'chessboardjsx';
import { ChessInstance } from 'chess.js';
import { Color, Puzzle } from '../../types';
import { PuzzleDetails } from './PuzzleDetails';
import { usePuzzleState } from '../../hooks';

const Chess = require("chess.js");

interface Props {
  puzzle: Puzzle;
}

interface Move {
  sourceSquare: string;
  targetSquare: string;
}

export function PuzzleView({ puzzle }: Props) {
  const color = getInitalTurn(puzzle.fen);
  const { fen, move, computerMove, done } = usePuzzleState(puzzle);

  useEffect(() => {
    setTimeout(() => {
      move(computerMove)
    }, 300);
  }, [computerMove]); // eslint-disable-line react-hooks/exhaustive-deps

  const handleMove = ({sourceSquare, targetSquare}: Move) => {
    const moveStr = `${sourceSquare}${targetSquare}`;
    move(moveStr);
  }

  return (
    <div>
      <h1>{color} to move</h1>
      {done && <p>Solved!! 🎉</p>}
      <Chessboard
        position={fen}
        orientation={color}
        onDrop={handleMove}
        draggable={!done}
      />
      <PuzzleDetails puzzle={puzzle} />
    </div>
  )
};

function getInitalTurn(fen: string): Color {
  const chess: ChessInstance = new Chess(fen);
  return (chess.turn() === chess.BLACK) ? "white" : "black";
}
