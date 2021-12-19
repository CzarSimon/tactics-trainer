import React, { useState } from 'react';
import { Modal } from 'antd';
import { Color, PromotionPiece } from '../../types';
import { ReactComponent as BlackQueen } from '../../assets/svg/BlackQueen.svg';
import { ReactComponent as WhiteQueen } from '../../assets/svg/WhiteQueen.svg';
import { ReactComponent as BlackRook } from '../../assets/svg/BlackRook.svg';
import { ReactComponent as WhiteRook } from '../../assets/svg/WhiteRook.svg';
import { ReactComponent as BlackBishop } from '../../assets/svg/BlackBishop.svg';
import { ReactComponent as WhiteBishop } from '../../assets/svg/WhiteBishop.svg';
import { ReactComponent as BlackKnight } from '../../assets/svg/BlackKnight.svg';
import { ReactComponent as WhiteKnight } from '../../assets/svg/WhiteKnight.svg';

import styles from './PromotionDialog.module.css';

interface Props {
  onCancel: () => void;
  onSelect: (piece: PromotionPiece) => void;
  orientation: Color;
}

export function PromotionDialog({ onCancel, onSelect, orientation }: Props) {
  const [visible, setVisible] = useState<boolean>(true);

  const cancel = () => {
    onCancel();
    setVisible(false);
  };

  const queen = orientation === 'black' ? <BlackQueen /> : <WhiteQueen />;
  const rook = orientation === 'black' ? <BlackRook /> : <WhiteRook />;
  const bishop = orientation === 'black' ? <BlackBishop /> : <WhiteBishop />;
  const knight = orientation === 'black' ? <BlackKnight /> : <WhiteKnight />;

  return (
    <Modal visible={visible} onCancel={cancel} footer={null} closable={false}>
      <div className={styles.Content}>
        <div onClick={() => onSelect('q')} className={styles.Button}>
          {queen}
        </div>
        <div onClick={() => onSelect('r')} className={styles.Button}>
          {rook}
        </div>
        <div onClick={() => onSelect('b')} className={styles.Button}>
          {bishop}
        </div>
        <div onClick={() => onSelect('n')} className={styles.Button}>
          {knight}
        </div>
      </div>
    </Modal>
  );
}
