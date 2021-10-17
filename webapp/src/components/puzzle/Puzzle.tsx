import React from 'react';
import { Puzzle } from '../../types';
import { PuzzleDetails } from './PuzzleDetails';

interface Props {
  puzzle: Puzzle;
}

export function PuzzleView({ puzzle }: Props) {
  const { id, fen } = puzzle;
  return (
    <div>
      <h2>Puzzle ID: {id}</h2>
      <p>Position: {fen}</p>
      <PuzzleDetails puzzle={puzzle} />
    </div>
  )
};
