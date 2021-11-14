import React from 'react';
import { Button } from 'antd';
import { useProblemSetCycles, useCreateNewProblemSetCycle } from '../../../../hooks';

import styles from './CycleList.module.css';
import { CycleCard } from './CycleCard';

interface Props {
  problemSetId: string;
}

export function CycleList({ problemSetId }: Props) {
  const cycles = useProblemSetCycles(problemSetId);
  const createNew = useCreateNewProblemSetCycle();

  const startNewCycle = async () => {
    createNew(problemSetId);
  };

  return (
    <>
      <div className={styles.CycleList}>
        <div className={styles.TitleRow}>
          <h2>Cycles</h2>
          <Button type="primary" shape="round" onClick={startNewCycle}>
            Start new cycle
          </Button>
        </div>
        {cycles && cycles.map((c) => <CycleCard cycle={c} key={c.id} />)}
      </div>
    </>
  );
}
