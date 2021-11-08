import React from 'react';
import { useHistory } from 'react-router';
import log from '@czarsimon/remotelogger';
import { useProblemSets } from '../../hooks';
import { ProblemSetsList } from './components/ProblemSetList';

export function ProblemSetsContainer() {
  const problemSets = useProblemSets();
  const history = useHistory();
  const onCreateNew = () => {
    log.info('creating new problem set');
    history.push('/problem-sets/new');
  };

  return <ProblemSetsList problemSets={problemSets} onCreateNew={onCreateNew} />;
}
