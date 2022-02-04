import React, { useState } from 'react';
import { Button, Steps } from 'antd';
import { PageHeader } from '../../../components/PageHeader';
import { welcomeStep } from './Welcome';
import { StepContent } from '../types';
import { LeftOutlined, RightOutlined } from '@ant-design/icons';
import { woodpeckerStep } from './WoodpeckerMethod';
import { problemSetsStep } from './ProblemSets';

import styles from './AboutPage.module.css';

const { Step } = Steps;

const steps: StepContent[] = [welcomeStep, woodpeckerStep, problemSetsStep];

export function AboutPage() {
  const [idx, setIdx] = useState<number>(0);
  const next = () => {
    setIdx(idx + 1);
  };

  const back = () => {
    setIdx(idx - 1);
  };

  return (
    <div className={styles.AboutPage}>
      <PageHeader title="About" />
      <Steps current={idx} labelPlacement="vertical">
        {steps.map((s) => (
          <Step key={`about-page-step-${s.title}`} />
        ))}
      </Steps>
      <div className={styles.Content}>
        <div>{steps[idx].content}</div>
        <div className={styles.ButtonGroup}>
          <Button
            icon={<LeftOutlined />}
            type="primary"
            disabled={idx === 0}
            onClick={back}
            shape="circle"
            className={styles.Button}
          />
          <Button
            icon={<RightOutlined />}
            type="primary"
            disabled={idx === steps.length - 1}
            onClick={next}
            shape="circle"
            className={styles.Button}
          />
        </div>
      </div>
    </div>
  );
}
