import React from 'react';
import { useHistory } from "react-router-dom";
import { Button, Typography } from 'antd';
import { getRandomPuzzleID } from '../../api/puzzleApi';

import styles from './Home.module.css';

const { Title } = Typography;

export function Home() {
  const location = useHistory();

  const gotoRandomPuzzle = () => {
    const puzzleId = getRandomPuzzleID();
    location.push(`puzzles/${puzzleId}`);
  }

  return (
    <div className={styles.Home}>
      <Title>Tactics Trainer</Title>
      <Button type="primary" size="large" onClick={gotoRandomPuzzle}>
        Get random puzzle
      </Button>
    </div>
  )
}
