import React from 'react';
import { useHistory } from 'react-router';
import { Spin } from 'antd';
import { PuzzleView } from '../../../components/puzzle/PuzzleView';
import { usePuzzle } from '../../../hooks';
import { Cycle } from '../../../types';
import { cycleIsCompleted, portraitMode } from '../../../util';
import { CompletedMessage } from './CompletedMessage';
import { CloseButton } from '../../../components/closeButton/CloseButton';

import styles from './CycleView.module.css';

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

  const mobileView = portraitMode();
  const style = mobileView ? {} : { padding: '48px' };

  return (
    <div className={styles.CycleView} style={style}>
      {!mobileView && <CloseButton onClose={onClose} />}
      {cycleIsCompleted(cycle) ? (
        <CompletedMessage cycle={cycle} />
      ) : (
        <PuzzleView puzzle={puzzle} onSolved={onSolvedPuzzle} />
      )}
    </div>
  );
}
