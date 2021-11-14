import React from 'react';
import { Spin } from 'antd';
import { useParams } from 'react-router';
import log from '@czarsimon/remotelogger';
import { useCycle, useUpdateCycle } from '../../hooks';
import { CycleView } from './components/CycleView';
import { NEW_PUZZLE_DELAY } from '../../constants';

interface ParamTypes {
  cycleId: string;
}

export function CycleContainer() {
  const { cycleId } = useParams<ParamTypes>();
  const cycle = useCycle(cycleId);
  const updateCycle = useUpdateCycle();
  if (!cycle) {
    return <Spin size="large" />;
  }

  const onSolvedPuzzle = () => {
    const { id, currentPuzzleId } = cycle;
    log.info(`Solved puzzle(id=${currentPuzzleId}) in Cycle(id=${id})`);
    setTimeout(() => updateCycle(id), NEW_PUZZLE_DELAY);
  };

  return <CycleView cycle={cycle} onSolvedPuzzle={onSolvedPuzzle} />;
}
