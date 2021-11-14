import React from 'react';
import { Button, Spin } from 'antd';
import { PuzzleView } from '../../../components/puzzle/Puzzle';
import { usePuzzle } from '../../../hooks';
import { Cycle } from '../../../types';
import { cycleIsCompleted } from '../../../util';
import { CompletedMessage } from './CompletedMessage';

import styles from './CycleView.module.css';
import { useHistory } from 'react-router';

interface Props {
  cycle: Cycle;
  onSolvedPuzzle: () => void;
}

export function CycleView({ cycle, onSolvedPuzzle }: Props) {
  const { currentPuzzleId } = cycle;
  const history = useHistory();
  const puzzle = usePuzzle(currentPuzzleId);
  if (!puzzle) {
    return <Spin size="large" />;
  }

  const onClose = () => {
    history.goBack();
  };

  return (
    <div className={styles.CycleView}>
      <Button shape="circle" size="large" onClick={onClose} className={styles.CloseButton}>
        X
      </Button>
      {cycleIsCompleted(cycle) ? (
        <CompletedMessage cycle={cycle} />
      ) : (
        <PuzzleView puzzle={puzzle} onSolved={onSolvedPuzzle} />
      )}
    </div>
  );
}
