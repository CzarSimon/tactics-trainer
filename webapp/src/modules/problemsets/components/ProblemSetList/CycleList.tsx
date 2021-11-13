import React from 'react';
import { Button } from 'antd';
import { useProblemSetCycles, useCreateNewProblemSetCycle } from '../../../../hooks';

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

  return (
    <>
      <div className={styles.TitleRow}>
        <h2>Cycles</h2>
        <Button type="primary" shape="round" onClick={startNewCycle}>
          Start new cycle
        </Button>
        {cycles &&
          cycles.map((c) => (
            <div key={c.id}>
              <p>
                <b>Number:</b> {c.number}
              </p>
              <p>
                <b>Created at:</b> {c.createdAt}
              </p>
            </div>
          ))}
      </div>
    </>
  );
}
