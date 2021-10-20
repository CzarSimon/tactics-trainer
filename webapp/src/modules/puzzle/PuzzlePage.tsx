import React from 'react';
import { useParams } from 'react-router-dom';
import { PuzzleView } from '../../components/puzzle/Puzzle';
import { usePuzzle } from '../../hooks';

interface ParamTypes {
  puzzleId: string;
}

export function PuzzlePage() {
  const { puzzleId } = useParams<ParamTypes>();

  const puzzle = usePuzzle(puzzleId);
  if (!puzzle) {
    return <p>Loading...</p>;
  }

  return <PuzzleView puzzle={puzzle} />;
}
