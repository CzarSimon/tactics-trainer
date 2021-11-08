import React, { useState } from 'react';
import { Spin, Button } from 'antd';
import { Optional, ProblemSet } from '../../../../types';
import { ProblemSetCard } from './ProblemSetCard';

import styles from './ProblemSetsList.module.css';
import { ProblemSetModal } from './ProblemSetModal';

interface Props {
  problemSets: Optional<ProblemSet[]>;
  onCreateNew: () => void;
}

export function ProblemSetsList({ problemSets, onCreateNew }: Props) {
  const [selectedId, setSelectedId] = useState<Optional<string>>(undefined);
  const selectProblemSet = (id: string) => {
    setSelectedId(id);
  };
  const unselectProblemSet = () => {
    setSelectedId(undefined);
  };

  return (
    <div className={styles.ProblemSetsList}>
      <div className={styles.ListTitleRow}>
        <h1>Problem Sets</h1>
        <Button type="primary" shape="round" size="large" onClick={onCreateNew}>
          Create new problem set
        </Button>
      </div>
      <div className={styles.ListContent}>
        {!problemSets && <Spin size="large" />}
        {problemSets && problemSets.map((s) => <ProblemSetCard key={s.id} problemSet={s} select={selectProblemSet} />)}
      </div>
      {selectedId && <ProblemSetModal id={selectedId} onClose={unselectProblemSet} />}
    </div>
  );
}
