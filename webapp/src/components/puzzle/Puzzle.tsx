import React, { useEffect, useState } from 'react';
import Chessboard from 'chessboardjsx';
import { ChessInstance } from 'chess.js';
import { Typography, Result } from 'antd';
import { Chess, Color, Puzzle } from '../../types';
import { PuzzleDetails } from './PuzzleDetails';
import { usePuzzleState } from '../../hooks';

import styles from './PuzzleView.module.css';
import { PUZZLE_SOLVED } from '../../constants';

const { Title } = Typography;
const COMPUTER_MOVE_DELAY_MS = 200;

interface Props {
  puzzle: Puzzle;
  onSolved?: () => void;
}

interface Move {
  sourceSquare: string;
  targetSquare: string;
}

export function PuzzleView({ puzzle, onSolved }: Props) {
  const color = getInitalTurn(puzzle.fen);
  const [draggable, setDraggable] = useState<boolean>(true);
  const { fen, move, computerMove, done } = usePuzzleState(puzzle);

  useEffect(() => {
    const timeout = setTimeout(() => {
      move(computerMove);
    }, COMPUTER_MOVE_DELAY_MS);

    return () => {
      clearTimeout(timeout);
    };
  }, [computerMove]); // eslint-disable-line react-hooks/exhaustive-deps

  const handleMove = ({ sourceSquare, targetSquare }: Move) => {
    const moveStr = `${sourceSquare}${targetSquare}`;
    const result = move(moveStr);
    if (result === PUZZLE_SOLVED && onSolved) {
      onSolved();
    }
  };

  if (draggable && done) {
    setTimeout(() => setDraggable(false), COMPUTER_MOVE_DELAY_MS);
  }

  return (
    <div className={styles.PuzzleView}>
      <div className={styles.Chessboard}>
        <Chessboard position={fen} orientation={color} onDrop={handleMove} draggable={draggable} width={750} />
      </div>
      <div className={styles.PuzzleInfo}>
        {!done && <Title className={styles.Title}>{color} to move</Title>}
        {done && <Result className={styles.Success} status="success" title="Puzzle solved!! 🎉" />}
        <div className={styles.PuzzleDetails}>
          <PuzzleDetails puzzle={puzzle} />
        </div>
      </div>
    </div>
  );
}

function getInitalTurn(fen: string): Color {
  const chess: ChessInstance = new Chess(fen);
  return chess.turn() === chess.BLACK ? 'white' : 'black';
}
