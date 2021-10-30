import * as ChessJS from 'chess.js';

export const Chess = typeof ChessJS === 'function' ? ChessJS : ChessJS.Chess;

// Utility types
export type Optional<T> = T | undefined;

export type TypedMap<T> = Record<string, T>;

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

// IAM types
export interface User {
  id: string;
  username: string;
  role: string;
  createdAt: string;
  updatedAt: string;
}

export interface AuthenticationRequest {
  username: string;
  password: string;
}

export interface AuthenticationResponse {
  token: string;
  user: User;
}

// Hook types
export interface UsePuzzleStateResult {
  fen: string;
  move: (move: string) => void;
  computerMove: string;
  correctMove: string;
  done: boolean;
}

export interface UseAuthResult {
  login: (req: AuthenticationRequest) => void;
  user?: User;
  authenticated: boolean;
}

// Client
export interface Client {
  id: string;
  sessionId: string;
}
