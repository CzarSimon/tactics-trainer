import React from 'react';
import log from '@czarsimon/remotelogger';
import { useArchiveProblemSet } from '../../../../../hooks';
import { ProblemSet } from '../../../../../types';
import { DropdownMenu } from '../../common/DropdownMenu';

import styles from './Title.module.css';

interface Props {
  problemSet: ProblemSet;
}

export function Title({ problemSet }: Props) {
  const { id, name } = problemSet;

  const archiveProblemSet = useArchiveProblemSet();
  const archive = () => {
    log.info(`archived problemSet(id=${id})`);
    archiveProblemSet(id);
  };

  return (
    <div className={styles.Title}>
      <h3>{name}</h3>
      <DropdownMenu onArchive={archive} />
    </div>
  );
}
