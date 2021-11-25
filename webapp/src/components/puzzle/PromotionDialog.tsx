import React, { useState } from 'react';
import { Button, Modal } from 'antd';
import { PromotionPiece } from '../../types';

import styles from './PromotionDialog.module.css';

interface Props {
  onCancel: () => void;
  onSelect: (piece: PromotionPiece) => void;
}

export function PromotionDialog({ onCancel, onSelect }: Props) {
  const [visible, setVisible] = useState<boolean>(true);

  const cancel = () => {
    onCancel();
    setVisible(false);
  };

  return (
    <Modal visible={visible} onCancel={cancel} footer={null} closable={false}>
      <div className={styles.Content}>
        <Button shape="circle" size="large" onClick={() => onSelect('q')} className={styles.Button}>
          Q
        </Button>
        <Button shape="circle" size="large" onClick={() => onSelect('r')} className={styles.Button}>
          R
        </Button>
        <Button shape="circle" size="large" onClick={() => onSelect('b')} className={styles.Button}>
          B
        </Button>
        <Button shape="circle" size="large" onClick={() => onSelect('n')} className={styles.Button}>
          N
        </Button>
      </div>
    </Modal>
  );
}
