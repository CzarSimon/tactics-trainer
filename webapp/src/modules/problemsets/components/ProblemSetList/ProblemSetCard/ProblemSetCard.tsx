import React from 'react';
import { Card, Col } from 'antd';
import { ProblemSet } from '../../../../../types';
import { Title } from './Title';

interface Props {
  problemSet: ProblemSet;
  select: (id: string) => void;
}

export function ProblemSetCard({ problemSet, select }: Props) {
  const { id, ratingInterval, themes } = problemSet;

  return (
    <Col xs={{ span: 24 }} lg={{ span: 6 }}>
      <Card
        title={<Title problemSet={problemSet} />}
        hoverable
        onClick={() => select(id)}
        style={{ borderRadius: '8px', height: '180px' }}
      >
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
    </Col>
  );
}
