import React from 'react';
import Chessboard from 'chessboardjsx';
import { Color, Move } from '../../types';

import styles from './Board.module.css';
import { portraitMode } from '../../util';

interface Props {
  fen: string;
  color: Color;
  draggable: boolean;
  wrongMove: boolean;
  handleMove: (move: Move) => void;
}

export function Board({ fen, color, draggable, wrongMove, handleMove }: Props) {
  return (
    <div className={wrongMove ? styles.ChessboardWithErrorHighlighting : styles.Chessboard}>
      <Chessboard
        position={fen}
        orientation={color}
        onDrop={handleMove}
        draggable={draggable && !wrongMove}
        width={getBoardWidth()}
      />
    </div>
  );
}

function getBoardWidth(): number {
  if (portraitMode()) {
    return window.innerWidth;
  }

  const margin = 48;
  return window.innerHeight - margin * 2;
}
