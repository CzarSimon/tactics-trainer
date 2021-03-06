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

export interface ApiStatus {
  status: string;
}

export type EmptyFn<T = void> = () => T;

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
  archived: boolean;
  puzzleIds: string[];
  createdAt: string;
  updatedAt: string;
}

export interface CreateProblemSetRequest {
  name: string;
  description: string;
  filter: PuzzleFilter;
}

export interface PuzzleFilter {
  themes: string[];
  minRating: number;
  maxRating: number;
  minPopularity: number;
  size: number;
}

export interface Cycle {
  id: string;
  number: number;
  problemSetId: string;
  currentPuzzleId: string;
  completedAt?: string;
  createdAt: string;
  updatedAt: string;
}

// Chess types
export type Color = 'black' | 'white';

export interface Move {
  sourceSquare: string;
  targetSquare: string;
  piece: string;
}

export type PromotionPiece = 'q' | 'r' | 'b' | 'n';

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
export type MoveResult = 'WRONG_MOVE' | 'CORRECT_MOVE' | 'PUZZLE_SOLVED';

export interface UsePuzzleStateResult {
  fen: string;
  move: (move: string) => MoveResult;
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
  logout: () => void;
}

// Client
export interface Client {
  id: string;
  sessionId: string;
}
