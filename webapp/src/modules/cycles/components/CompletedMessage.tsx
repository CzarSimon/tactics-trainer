import React from 'react';
import { Result } from 'antd';
import { Cycle } from '../../../types';

interface Props {
  cycle: Cycle;
}

export function CompletedMessage({ cycle }: Props) {
  const title = `Cycle ${cycle.number} completed, good job! 🎉`;
  return <Result status="success" title={title} />;
}
