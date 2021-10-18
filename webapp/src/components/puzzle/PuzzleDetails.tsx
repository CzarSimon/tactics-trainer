import React from 'react';
import { Table, Collapse } from 'antd';
import { Puzzle } from '../../types';

import styles from './PuzzleDetails.module.css';

const { Panel } = Collapse;

interface Props {
  puzzle: Puzzle;
}

const themeColums = [
  {
    title: 'Themes',
    dataIndex: 'theme',
    key: 'theme',
  },
];

export function PuzzleDetails({ puzzle }: Props) {
  const { id, rating, popularity, themes, gameUrl } = puzzle;
  const themeTableData = themes.map((theme) => ({ theme, key: `${id}:${theme}` }));

  return (
    <Collapse className={styles.Collapse}>
      <Panel header="Puzzle details" key="0">
        <div className={styles.PuzzleDetails}>
          <div className={styles.BasicInfo}>
            <p>Rating: {rating}</p>
            <p>Popularity: {popularity}</p>
            <a href={gameUrl}>Link to game</a>
          </div>
          <Table dataSource={themeTableData} columns={themeColums} pagination={false} />
        </div>
      </Panel>
    </Collapse>
  );
}
