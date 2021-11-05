import React from 'react';
import { useHistory } from 'react-router';
import log from '@czarsimon/remotelogger';
import { useProblemSets } from '../../hooks';
import { ProblemSetsList } from './components/ProblemSetsList';

export function ProblemSetsContainer() {
  const problemSets = useProblemSets();
  const history = useHistory();

  const selectProblemSet = (id: string) => {
    log.info(`selected problem set id=${id}`);
    history.push(`/problem-sets/${id}`);
  };

  return <ProblemSetsList problemSets={problemSets} select={selectProblemSet} />;
}
