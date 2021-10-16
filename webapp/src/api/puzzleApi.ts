import log from '@czarsimon/remotelogger';
import { httpclient } from './httpclient';
import { Optional, Puzzle } from '../types';

const PUZZLE_URL = '/api/puzzle-server/v1/puzzles'

export async function getPuzzle(id: string): Promise<Puzzle> {
  const { body, error } = await httpclient.get<Puzzle>({ url: `${PUZZLE_URL}/${id}` });

  if (!body) {
    throw handleGetPuzzleError(id, error);
  }

  return body;
}

function handleGetPuzzleError(id: string, error: Optional<Error>): Error {
  const message = `failed to fetch puzzle(id=${id}), error=${error?.message}`;
  log.error(message);
  return new Error(message);
}



