// Util
export type Optional<T> = T | undefined;

// Puzzle
export interface Puzzle {
  id: string;
  externalId: string;
  fen: string;
  moves: string[];
  rating: number;
  ratingDeviation: number;
  popularity: number;
  themes: string[];
  gameUrl: string;
  createdAt: string;
  updatedAt: string;
}

// Chess types
export type Color = 'black' | 'white';

// Hook types
export interface UsePuzzleStateResult {
  fen: string;
  move: (move: string) => void;
  computerMove: string;
  correctMove: string;
  done: boolean;
}

// Client
export interface Client {
  id: string;
  sessionId: string;
}
