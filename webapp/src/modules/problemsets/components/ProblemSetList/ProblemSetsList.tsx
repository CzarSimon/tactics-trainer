import React, { useState } from 'react';
import { Spin, Button, Row } from 'antd';
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
        <Row style={{ width: '100%', margin: '0' }} gutter={[32, 32]}>
          {problemSets &&
            problemSets.map((s) => <ProblemSetCard key={s.id} problemSet={s} select={selectProblemSet} />)}
        </Row>
      </div>

      {selectedId && <ProblemSetModal id={selectedId} onClose={unselectProblemSet} />}
    </div>
  );
}
