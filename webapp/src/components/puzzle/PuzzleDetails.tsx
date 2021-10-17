import React from 'react';
import { Puzzle } from '../../types';

interface Props {
  puzzle: Puzzle;
}

export function PuzzleDetails({ puzzle }: Props) {
  const { id, rating, popularity, themes, gameUrl } = puzzle;
  return (
    <div>
      <h3>Puzzle details</h3>
      <p>Rating: {rating}</p>
      <p>popularity: {popularity}</p>
      <h4>Themes</h4>
      {themes.map(theme => (
        <p key={`${id}:${theme}`}>{theme}</p>
      ))}
      <a href={gameUrl}>Link to game</a>
    </div>
  )
};
