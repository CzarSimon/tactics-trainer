import React, { useState } from 'react';
import { Spin, Button, Row } from 'antd';
import { PlusOutlined } from '@ant-design/icons';
import { Optional, ProblemSet } from '../../../../types';
import { ProblemSetCard } from './ProblemSetCard';
import { ProblemSetModal } from './ProblemSetModal';
import { ProfileAvatar } from '../../../profileAvatar';

import styles from './ProblemSetsList.module.css';

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
      <ProfileAvatar />
      <div className={styles.ListTitleRow}>
        <h1>Problem Sets</h1>
        <Button
          aria-label="new-problem-set-button"
          type="primary"
          shape="circle"
          size="large"
          onClick={onCreateNew}
          className={styles.NewButton}
          icon={<PlusOutlined />}
        />
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
