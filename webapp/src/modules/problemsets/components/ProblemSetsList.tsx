import React from 'react';
import { Spin } from 'antd';
import { Optional, ProblemSet } from '../../../types';
import { ProblemSetCard } from './ProblemSetCard';

import styles from './ProblemSetsList.module.css';

interface Props {
  problemSets: Optional<ProblemSet[]>;
  select: (id: string) => void;
}

export function ProblemSetsList({ problemSets, select }: Props) {
  return (
    <div className={styles.ProblemSetsList}>
      <h1>Problem Sets</h1>
      <div className={styles.ListContent}>
        {!problemSets && <Spin size="large" />}
        {problemSets && problemSets.map((s) => <ProblemSetCard key={s.id} problemSet={s} select={select} />)}
      </div>
    </div>
  );
}
