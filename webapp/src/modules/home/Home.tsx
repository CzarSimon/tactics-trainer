import React from 'react';
import { useHistory } from "react-router-dom"
import { getRandomPuzzleID } from '../../api/puzzleApi';

export function Home() {
  const location = useHistory();

  const gotoRandomPuzzle = () => {
    const puzzleId = getRandomPuzzleID();
    location.push(`puzzles/${puzzleId}`);
  }

  return (
    <div>
      <h1>Tactics Trainer</h1>
      <button onClick={gotoRandomPuzzle}>Get random puzzle</button>
    </div>
  )
}
