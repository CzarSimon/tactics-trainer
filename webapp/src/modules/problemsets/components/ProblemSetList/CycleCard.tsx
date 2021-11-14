import React from 'react';
import { useHistory } from 'react-router';
import { Card, Tag } from 'antd';
import { Cycle } from '../../../../types';
import { cycleIsCompleted } from '../../../../util';

import styles from './CycleCard.module.css';

interface Props {
  cycle: Cycle;
}

export function CycleCard({ cycle }: Props) {
  const history = useHistory();
  const { id, number, createdAt } = cycle;
  const completed: boolean = cycleIsCompleted(cycle);

  const onClick = () => {
    if (completed) {
      return;
    }

    history.push(`/cycles/${id}`);
  };

  return (
    <Card hoverable={!completed} onClick={onClick} className={styles.CycleCard}>
      <h4>Cycle {number}</h4>
      <div className={styles.CycleInfo}>
        <p>{createdAt}</p>
        {!completed && <Tag color="green">Active</Tag>}
        {completed && <Tag>Completed</Tag>}
      </div>
    </Card>
  );
}
