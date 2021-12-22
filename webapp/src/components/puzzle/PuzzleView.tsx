import React, { useEffect, useState } from 'react';
import Chessboard from 'chessboardjsx';
import { ChessInstance } from 'chess.js';
import { Typography, Result } from 'antd';
import log from '@czarsimon/remotelogger';
import { Chess, Color, Puzzle, Move, Optional, PromotionPiece } from '../../types';
import { PuzzleDetails } from './PuzzleDetails';
import { PromotionDialog } from './PromotionDialog';
import { usePuzzleState } from '../../hooks';
import { PUZZLE_SOLVED } from '../../constants';
import { enablePromotion, encodeMove } from '../../util/chessutil';

import styles from './PuzzleView.module.css';

const { Title } = Typography;
const COMPUTER_MOVE_DELAY_MS = 200;

interface Props {
  puzzle: Puzzle;
  onSolved?: () => void;
}

export function PuzzleView({ puzzle, onSolved }: Props) {
  const color = getInitalTurn(puzzle.fen);
  const [draggable, setDraggable] = useState<boolean>(true);
  const [pendingMove, setPendingMove] = useState<Optional<Move>>(undefined);
  const { fen, move, computerMove, done } = usePuzzleState(puzzle);

  useEffect(() => {
    const timeout = setTimeout(() => {
      move(computerMove);
    }, COMPUTER_MOVE_DELAY_MS);

    return () => {
      clearTimeout(timeout);
    };
  }, [computerMove]); // eslint-disable-line react-hooks/exhaustive-deps

  const handleMove = (m: Move) => {
    if (enablePromotion(m)) {
      setPendingMove(m);
      return;
    }

    const moveStr = encodeMove(m);
    executeMove(moveStr);
  };

  const handlePromotion = (piece: PromotionPiece) => {
    if (!pendingMove) {
      log.error(`attempted promotion to piece=${piece} with undefined pending move`);
      return;
    }

    const moveStr = encodeMove(pendingMove, piece);
    executeMove(moveStr);
    setPendingMove(undefined);
  };

  const executeMove = (moveStr: string) => {
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
      {pendingMove && (
        <PromotionDialog orientation={color} onCancel={() => setPendingMove(undefined)} onSelect={handlePromotion} />
      )}
      <div className={styles.Chessboard}>
        <Chessboard
          position={fen}
          orientation={color}
          onDrop={handleMove}
          draggable={draggable}
          width={getBoardWidth()}
        />
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

function getBoardWidth(): number {
  const height = window.innerHeight;
  const width = window.innerWidth;
  const margin = 48;

  if (width > height) {
    return height - margin * 2;
  }

  return width - margin * 2;
}
