import { httpclient } from './httpclient';
import { Puzzle } from '../types';
import { wrapAndLogError } from './util';

const PUZZLE_URL = '/api/puzzle-server/v1/puzzles'

export async function getPuzzle(id: string): Promise<Puzzle> {
  const { body, error, metadata } = await httpclient.get<Puzzle>({ url: `${PUZZLE_URL}/${id}` });

  if (!body) {
    throw wrapAndLogError(`failed to fetch puzzle(id=${id})`, error, metadata);
  }

  return body;
}

export function getRandomPuzzleID(): string {
  const ids = [
    '0deddb83-f8e7-4cb3-8a1c-94dcd6f4bc94',
    '73205470-e281-4da8-85f1-2756969a4bc8',
    'b6d35bd6-72b0-4ca8-9134-88eda6a45b47',
    'd58ca880-2685-4512-994a-426e731cc784',
    '8cdc2ea3-e810-4917-bfc4-75cd59f5ad72',
    '0eaf0e3c-0255-482b-aabf-7d3e60b15b7d',
    '3d32485b-3a02-424c-bd47-6123583c4e6f',
    'a663a3f5-3011-46e5-bb4b-1cd6ec5be5f1',
    '9e0795c9-1842-4ef0-9032-79acd2fa222a',
    '4632e2e3-20aa-42db-a950-dd21f719e867',
  ]

  const randomIdx = Math.floor(Math.random() * ids.length);
  return ids[randomIdx];
}


