import React from 'react';
import Chessboard from 'chessboardjsx';
import { Puzzle } from '../../types';
import { PuzzleDetails } from './PuzzleDetails';

interface Props {
  puzzle: Puzzle;
}

export function PuzzleView({ puzzle }: Props) {
  const { fen } = puzzle;

  return (
    <div>
      <h1>Black to move</h1>
      <Chessboard position={fen} orientation="black" />
      <PuzzleDetails puzzle={puzzle} />
    </div>
  )
};
