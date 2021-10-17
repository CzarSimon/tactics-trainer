import { useQuery } from 'react-query';
import { getPuzzle } from '../api/puzzleApi';
import { Puzzle, Optional } from '../types';

const IMMUTABLE_QUERY_OPTIONS = {
  retry: 0,
  refetchOnWindowFocus: false,
  refetchIntervalInBackground: false,
  refetchOnMount: false,
  refetchOnReconnect: false
};

export function usePuzzle(id: string): Optional<Puzzle> {
  const { data } = useQuery<Puzzle, Error>(['puzzle', id], () => getPuzzle(id), IMMUTABLE_QUERY_OPTIONS);
  return data;
};
