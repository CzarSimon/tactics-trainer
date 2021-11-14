import React from 'react';
import { useHistory } from 'react-router';
import { Card, Tag } from 'antd';
import { Optional, Cycle } from '../../../../types';
import { EMTPY_DATE } from '../../../../constants';

import styles from './CycleCard.module.css';

interface Props {
  cycle: Cycle;
}

export function CycleCard({ cycle }: Props) {
  const history = useHistory();
  const { id, number, completedAt, createdAt } = cycle;
  const completed: boolean = isCompleted(completedAt);

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

function isCompleted(completedAt: Optional<string>): boolean {
  return completedAt !== undefined && completedAt !== EMTPY_DATE;
}
