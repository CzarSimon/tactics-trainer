import React from 'react';
import { Button } from 'antd';
import { useProblemSetCycles, useCreateNewProblemSetCycle } from '../../../../hooks';
import { CycleCard } from './CycleCard';
import { Cycle, Optional } from '../../../../types';
import { EMTPY_DATE } from '../../../../constants';

import styles from './CycleList.module.css';

interface Props {
  problemSetId: string;
}

export function CycleList({ problemSetId }: Props) {
  const cycles = useProblemSetCycles(problemSetId);
  const createNew = useCreateNewProblemSetCycle();

  const startNewCycle = async () => {
    createNew(problemSetId);
  };

  const activeCyclesExist = checkForActiveCycles(cycles);
  return (
    <>
      <div className={styles.CycleList}>
        <div className={styles.TitleRow}>
          <h2>Cycles</h2>
          <Button type="primary" shape="round" onClick={startNewCycle} disabled={activeCyclesExist}>
            Start new cycle
          </Button>
        </div>
        {cycles && cycles.reverse().map((c) => <CycleCard cycle={c} key={c.id} />)}
      </div>
    </>
  );
}

function checkForActiveCycles(cycles: Optional<Cycle[]>): boolean {
  if (!cycles) {
    return false;
  }

  for (const c of cycles) {
    if (c.completedAt === EMTPY_DATE || c.completedAt === undefined) return true;
  }

  return false;
}
