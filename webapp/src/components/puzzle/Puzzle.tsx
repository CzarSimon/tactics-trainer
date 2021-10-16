import React from 'react';
import { Puzzle } from '../../types';

interface Props {
  puzzle: Puzzle;
}

export function PuzzleView({ puzzle }: Props) {
  const { id, rating, popularity, gameUrl } = puzzle;
  return (
    <div>
      <h2>Puzzle ID: {id}</h2>
      <p>Rating: {rating}</p>
      <p>popularity: {popularity}</p>
      <a href={gameUrl}>Link to game</a>
    </div>
  )
};
