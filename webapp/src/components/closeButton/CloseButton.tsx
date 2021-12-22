import React from 'react';
import { Button } from 'antd';

import styles from './CloseButton.module.css';

interface Props {
  onClose: () => void;
}

export function CloseButton({ onClose }: Props) {
  return (
    <div className={styles.CloseButton}>
      <Button shape="circle" size="large" onClick={onClose}>
        X
      </Button>
    </div>
  );
}
