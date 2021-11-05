import React from 'react';
import { Card } from 'antd';
import { ProblemSet } from '../../../types';

import styles from './ProblemSetCard.module.css';

interface Props {
  problemSet: ProblemSet;
  select: (id: string) => void;
}

export function ProblemSetCard({ problemSet, select }: Props) {
  const { id, name, ratingInterval, themes } = problemSet;

  return (
    <div className={styles.ProblemSetCard}>
      <Card title={name} hoverable onClick={() => select(id)} style={{ height: '200px' }}>
        <p>
          <b>Rating interval: </b>
          {ratingInterval}
        </p>
        {themes.length > 0 && (
          <p>
            <b>Themes: </b>
            {themes.join(' ')}
          </p>
        )}
      </Card>
    </div>
  );
}
