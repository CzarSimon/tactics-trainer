import React from 'react';
import { Button } from 'antd';
import { StepContent } from '../types';
import { useHistory } from 'react-router-dom';

import styles from './Step.module.css';
import { portraitMode } from '../../../util';

function ProblemSets() {
  const history = useHistory();
  const className = portraitMode() ? styles.PortraitStep : styles.LandscapeStep;

  return (
    <div className={className}>
      <h2>Problem sets</h2>
      <p>Problem sets are what we call a collection of puzzles that you use to repeatedly solve over and over again.</p>
      <p>
        You create your own problem sets by giving them a name, selecting the rating level of the puzzles and optionally
        choosing the themes you want the puzzles to have. You can also choose how large you want your sets to be. After
        that, click create and you will have your very first problem set. From there you can start solving the puzzles
        in the set, and if you keep at it your tacticall skill will improve.
      </p>
      <p>
        Each iteration of solving the puzzles in a set is called a cycle. For each cycle you should aim to solve the
        puzzles faster and faster. You will find that this comes naturally as your brain will unconsiously recognize the
        patterns in the puzzles, even if you dont remember them by heart.
      </p>
      <Button type="primary" onClick={() => history.push('/')} shape="round" className={styles.CTAButton}>
        Go to problem sets
      </Button>
    </div>
  );
}

export const problemSetsStep: StepContent = {
  title: 'Problem sets',
  content: <ProblemSets />,
};
