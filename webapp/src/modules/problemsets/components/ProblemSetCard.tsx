import React from 'react';
import { Card } from 'antd';
import { ProblemSet } from '../../../types';

interface Props {
  problemSet: ProblemSet;
  select: (id: string) => void;
}

export function ProblemSetCard({ problemSet, select }: Props) {
  const { id, name, ratingInterval, themes } = problemSet;

  return (
    <div>
      <Card title={name} hoverable onClick={() => select(id)}>
        <p>{ratingInterval}</p>
        {themes && <p>{themes.join(' ')}</p>}
      </Card>
    </div>
  );
}
