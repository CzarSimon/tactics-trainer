import { useQuery } from 'react-query';
import { getPuzzle } from '../api/puzzleApi';
import { Puzzle, Optional } from '../types';

export function usePuzzle(id: string): Optional<Puzzle> {
  const { data } = useQuery<Puzzle, Error>(['puzzle', id], () => getPuzzle(id));
  return data;
}

