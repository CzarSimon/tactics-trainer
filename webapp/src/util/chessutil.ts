import { Move, PromotionPiece } from '../types';

type OptionalPromotionPiece = PromotionPiece | '';

const WHITE_PAWN = 'wP';
const BLACK_PAWN = 'bP';

export function encodeMove({ sourceSquare, targetSquare }: Move, promotion: OptionalPromotionPiece = ''): string {
  return `${sourceSquare}${targetSquare}${promotion}`;
}

export function enablePromotion({ targetSquare, piece }: Move): boolean {
  if (piece === WHITE_PAWN && targetSquare.endsWith('8')) {
    return true;
  }

  if (piece === BLACK_PAWN && targetSquare.endsWith('1')) {
    return true;
  }

  return false;
}
