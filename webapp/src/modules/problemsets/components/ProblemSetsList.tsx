import React from 'react';
import { Optional, ProblemSet } from '../../../types';
import { ProblemSetCard } from './ProblemSetCard';

interface Props {
  problemSets: Optional<ProblemSet[]>;
  select: (id: string) => void;
}

export function ProblemSetsList({ problemSets, select }: Props) {
  return (
    <div>
      <h1>Problem Sets</h1>
      {!problemSets && <p>Loading</p>}
      {problemSets && problemSets.map((s) => <ProblemSetCard key={s.id} problemSet={s} select={select} />)}
    </div>
  );
}
