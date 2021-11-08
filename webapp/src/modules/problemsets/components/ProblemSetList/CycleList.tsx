import React from 'react';
import { Button } from 'antd';
import log from '@czarsimon/remotelogger';

import styles from './CycleList.module.css';

interface Props {
  problemSetId: string;
}

export function CycleList({ problemSetId }: Props) {
  const startNewCycle = () => {
    log.info(`Should create new cycle for ProblemSet(id=${problemSetId})`);
  };

  return (
    <>
      <div className={styles.TitleRow}>
        <h2>Cycles</h2>
        <Button type="primary" shape="round" onClick={startNewCycle}>
          Start new cycle
        </Button>
      </div>
    </>
  );
}
