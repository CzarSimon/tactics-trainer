import React from 'react';
import { Card, Tag } from 'antd';
import { Cycle } from '../../../../types';

import styles from './CycleCard.module.css';

interface Props {
  cycle: Cycle;
}

export function CycleCard({ cycle }: Props) {
  const { number, compleatedAt, createdAt } = cycle;
  console.log(compleatedAt);
  const compleated: boolean = compleatedAt !== undefined;

  return (
    <Card hoverable className={styles.CycleCard}>
      <h4>Cycle {number}</h4>
      <div className={styles.CycleInfo}>
        <p>{createdAt}</p>
        {!compleated && <Tag color="green">Active</Tag>}
        {compleated && <Tag color="green">Compleated</Tag>}
      </div>
    </Card>
  );
}
