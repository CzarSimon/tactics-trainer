import React from 'react';
import { PuzzleView } from '../../components/puzzle/Puzzle';
import { usePuzzle } from '../../hooks';

interface Props{
  id: string;
}

export function PuzzleSolver({ id }: Props) {
  const puzzle = usePuzzle(id);
  if (!puzzle) {
    return <p>Loading...</p>
  }

  return <PuzzleView puzzle={puzzle} />
}
