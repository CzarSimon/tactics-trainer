import React from 'react';
import { Typography, Result } from 'antd';
import { Color, Puzzle } from '../../types';
import { PuzzleDetails } from './PuzzleDetails';
import { portraitMode } from '../../util';

import styles from './PuzzleInfo.module.css';

const { Title } = Typography;

interface Props {
  puzzle: Puzzle;
  color: Color;
  done: boolean;
}

export function PuzzleInfo({ puzzle, color, done }: Props) {
  return (
    <div className={styles.PuzzleInfo}>
      {!done && <Title className={styles.Title}>{color} to move</Title>}
      {done && <Result className={styles.Success} status="success" title="Puzzle solved!! 🎉" />}
      {!portraitMode() && (
        <div className={styles.PuzzleDetails}>
          <PuzzleDetails puzzle={puzzle} />
        </div>
      )}
    </div>
  );
}
