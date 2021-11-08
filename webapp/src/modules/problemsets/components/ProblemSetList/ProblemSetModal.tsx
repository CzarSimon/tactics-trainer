import React, { useState } from 'react';
import { Divider, Modal } from 'antd';
import { useProblemSet } from '../../../../hooks';
import { CycleList } from './CycleList';

interface Props {
  id: string;
  onClose: () => void;
}

export function ProblemSetModal({ id, onClose }: Props) {
  const problemSet = useProblemSet(id);
  const [open, setOpen] = useState(true);
  if (!problemSet) {
    return null;
  }

  const close = () => {
    setOpen(false);
    onClose();
  };

  const { name, description, ratingInterval, puzzleIds, createdAt } = problemSet;
  return (
    <Modal visible={open} onCancel={close} footer={null}>
      <h1>{name}</h1>
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
