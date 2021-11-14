import { MoveResult } from '../types';

export const DEV_MODE = true;

export const CLIENT_ID_KEY = '@tactics-trainer:webapp:CLIENT_ID';
export const AUTH_TOKEN_KEY = '@tactics-trainer:webapp:AUTH_TOKEN';
export const CURRENT_USER_KEY = '@tactics-trainer:webapp:CURRENT_USER_KEY';

export const EMTPY_DATE = '0001-01-01T00:00:00Z';

export const NEW_PUZZLE_DELAY = 800;

// Move results
export const WRONG_MOVE: MoveResult = 'WRONG_MOVE';
export const CORRECT_MOVE: MoveResult = 'CORRECT_MOVE';
export const PUZZLE_SOLVED: MoveResult = 'PUZZLE_SOLVED';

export const APP_NAME = 'tactics-trainer/webapp';
export const APP_VERSION = 'x.y.z';
