import React, { useEffect } from 'react';
import Chessboard from 'chessboardjsx';
import { ChessInstance } from 'chess.js';
import { Typography, Result } from 'antd';
import { Color, Puzzle } from '../../types';
import { PuzzleDetails } from './PuzzleDetails';
import { usePuzzleState } from '../../hooks';

import styles from './PuzzleView.module.css';

const Chess = require("chess.js");

const { Title } = Typography;

interface Props {
  puzzle: Puzzle;
}

interface Move {
  sourceSquare: string;
  targetSquare: string;
}

export function PuzzleView({ puzzle }: Props) {
  const color = getInitalTurn(puzzle.fen);
  const { fen, move, computerMove, done } = usePuzzleState(puzzle);

  useEffect(() => {
    setTimeout(() => {
      move(computerMove)
    }, 300);
  }, [computerMove]); // eslint-disable-line react-hooks/exhaustive-deps

  const handleMove = ({sourceSquare, targetSquare}: Move) => {
    const moveStr = `${sourceSquare}${targetSquare}`;
    move(moveStr);
  }

  return (
    <div className={styles.PuzzleView}>
      {!done && <Title className={styles.Title}>{color} to move</Title>}
      {done && <Result className={styles.Success} status='success' title='Puzzle solved!! 🎉' />}
      <Chessboard
        position={fen}
        orientation={color}
        onDrop={handleMove}
        draggable={!done}
      />
      <div className={styles.PuzzleDetails}>
        <PuzzleDetails puzzle={puzzle} />
      </div>
    </div>
  )
};

function getInitalTurn(fen: string): Color {
  const chess: ChessInstance = new Chess(fen);
  return (chess.turn() === chess.BLACK) ? "white" : "black";
}
