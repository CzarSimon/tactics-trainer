import * as ChessJS from 'chess.js';

export const Chess = typeof ChessJS === 'function' ? ChessJS : ChessJS.Chess;

// Utility types
export type Optional<T> = T | undefined;

export type TypedMap<T> = Record<string, T>;

export interface ApiResponse<T> {
  data?: T;
  error?: Error;
}

export interface ErrorInfo {
  title: string;
  details: string;
}

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

export interface ProblemSet {
  id: string;
  name: string;
  description?: string;
  themes: string[];
  ratingInterval: string;
  userId: string;
  puzzleIds: string[];
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
  signup: (req: AuthenticationRequest) => void;
  user?: User;
  authenticated: boolean;
  authenticate: (user: User) => void;
}

// Client
export interface Client {
  id: string;
  sessionId: string;
}
