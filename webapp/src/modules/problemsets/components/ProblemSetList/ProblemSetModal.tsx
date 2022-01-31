import React, { useState } from 'react';
import { Divider, Modal } from 'antd';
import { useArchiveProblemSet, useProblemSet } from '../../../../hooks';
import { CycleList } from './CycleList';
import { DropdownMenu } from '../common/DropdownMenu';

import styles from './ProblemSetModal.module.css';

interface Props {
  id: string;
  onClose: () => void;
}

export function ProblemSetModal({ id, onClose }: Props) {
  const problemSet = useProblemSet(id);
  const archiveProblemSet = useArchiveProblemSet();
  const [open, setOpen] = useState(true);
  if (!problemSet) {
    return null;
  }

  const close = () => {
    setOpen(false);
    onClose();
  };

  const archive = () => {
    archiveProblemSet(id);
    close();
  };

  const { name, description, ratingInterval, puzzleIds, createdAt } = problemSet;
  return (
    <Modal visible={open} onCancel={close} footer={null} closable={false}>
      <div className={styles.Title}>
        <h1>{name}</h1>
        <DropdownMenu onClose={close} onArchive={archive} />
      </div>
      <p>{description}</p>
      <p>
        <b>Rating interval: </b>
        {ratingInterval}
      </p>
      <p>
        <b>Problems: </b>
        {puzzleIds.length}
      </p>
      <p>
        <b>Created: </b>
        {createdAt}
      </p>
      <Divider />
      <CycleList problemSetId={id} />
    </Modal>
  );
}
