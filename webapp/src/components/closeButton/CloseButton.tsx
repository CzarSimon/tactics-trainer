import React from 'react';
import { Button } from 'antd';
import { CloseOutlined } from '@ant-design/icons';

import styles from './CloseButton.module.css';

interface Props {
  onClose: () => void;
}

export function CloseButton({ onClose }: Props) {
  return (
    <div className={styles.CloseButton}>
      <Button aria-label="close-button" shape="circle" size="large" onClick={onClose} icon={<CloseOutlined />} />
    </div>
  );
}
