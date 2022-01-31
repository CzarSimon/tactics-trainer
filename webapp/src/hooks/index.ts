import { useState } from 'react';
import Form, { FormInstance } from 'antd/lib/form';
import { useQuery, useQueryClient } from 'react-query';
import { ChessInstance } from 'chess.js';
import log from '@czarsimon/remotelogger';
import {
  getPuzzle,
  getProblemSet,
  getProblemSets,
  createProblemSet,
  listProblemSetCycles,
  getCycle,
  updateCycle,
  createProblemSetCycle,
  archiveProblemSet,
} from '../api/puzzleApi';
import {
  Chess,
  Puzzle,
  Optional,
  UsePuzzleStateResult,
  ProblemSet,
  CreateProblemSetRequest,
  Cycle,
  MoveResult,
} from '../types';
import { CORRECT_MOVE, PUZZLE_SOLVED, WRONG_MOVE } from '../constants';

const DEFAULT_QUERY_OPTIONS = {
  retry: 0,
};

const IMMUTABLE_QUERY_OPTIONS = {
  ...DEFAULT_QUERY_OPTIONS,
  refetchOnWindowFocus: false,
  refetchIntervalInBackground: false,
  refetchOnMount: false,
  refetchOnReconnect: false,
};

export function usePuzzle(id: string): Optional<Puzzle> {
  const { data } = useQuery<Puzzle, Error>(['puzzle', id], () => getPuzzle(id), IMMUTABLE_QUERY_OPTIONS);
  return data;
}

export function useProblemSet(id: string): Optional<ProblemSet> {
  const { data } = useQuery<ProblemSet, Error>(['problem-sets', id], () => getProblemSet(id), IMMUTABLE_QUERY_OPTIONS);
  return data;
}

export function useProblemSets(): Optional<ProblemSet[]> {
  const { data } = useQuery<ProblemSet[], Error>('problem-sets', getProblemSets);
  return data;
}

export function useProblemSetCycles(id: string): Optional<Cycle[]> {
  const { data } = useQuery<Cycle[], Error>(
    ['problem-sets', id, 'cycles'],
    () => listProblemSetCycles(id),
    DEFAULT_QUERY_OPTIONS,
  );

  return data;
}

type MutateCycleFn = (id: string) => Promise<Cycle>;

export function useCreateNewProblemSetCycle(): MutateCycleFn {
  const queryClient = useQueryClient();

  const createNewCycle = async (id: string): Promise<Cycle> => {
    const cycle = await createProblemSetCycle(id);
    queryClient.invalidateQueries(['problem-sets', id, 'cycles']);
    return cycle;
  };

  return createNewCycle;
}

export function useUpdateCycle(): MutateCycleFn {
  const queryClient = useQueryClient();

  const updateCycleWrapper = async (id: string): Promise<Cycle> => {
    const cycle = await updateCycle(id);
    queryClient.invalidateQueries(['cycles', id]);
    return cycle;
  };

  return updateCycleWrapper;
}

export function useCycle(id: string): Optional<Cycle> {
  const { data } = useQuery<Cycle, Error>(['cycles', id], () => getCycle(id), DEFAULT_QUERY_OPTIONS);
  return data;
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any
type UseFormSelectHook = [FormInstance, (key: string, value: any) => void];

export function useFormSelect(): UseFormSelectHook {
  const [form] = Form.useForm();
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const onSelect = (key: string, value: any) => {
    form.setFieldsValue({ [key]: value });
  };

  return [form, onSelect];
}

type CreateProblemSetFn = (req: CreateProblemSetRequest) => Promise<ProblemSet>;

export function useCreateNewProblemSet(): CreateProblemSetFn {
  const queryClient = useQueryClient();

  const createNewProblemSet = async (req: CreateProblemSetRequest): Promise<ProblemSet> => {
    const ps = await createProblemSet(req);
    queryClient.invalidateQueries('problem-sets');
    return ps;
  };

  return createNewProblemSet;
}

type ArchiveProblemSetFn = (id: string) => Promise<void>;

export function useArchiveProblemSet(): ArchiveProblemSetFn {
  const queryClient = useQueryClient();

  const archive = async (id: string): Promise<void> => {
    await archiveProblemSet(id);
    queryClient.invalidateQueries('problem-sets');
  };

  return archive;
}

export function usePuzzleState({ fen, moves }: Puzzle): UsePuzzleStateResult {
  const [done, setDone] = useState<boolean>(false);
  const [moveIdx, setMoveIdx] = useState<number>(0);
  const [computerMove, setComputerMove] = useState<string>(moves[0]);
  const [correctMove, setCorrectMove] = useState<string>(moves[1]);

  const [position, setPosition] = useState<string>(fen);

  const updatePosition = (move: string): MoveResult => {
    log.debug(`Move: ${move}`);
    if (move !== correctMove && move !== computerMove) {
      log.info('Wrong move!');
      return WRONG_MOVE;
    }

    const chess: ChessInstance = new Chess(position);
    const validMove = chess.move(move, { sloppy: true });

    if (!validMove) {
      log.error('Invalid move!');
      return WRONG_MOVE;
    }

    setPosition(chess.fen);

    if (move === correctMove) {
      const nextIndex = moveIdx + 2;
      setMoveIdx(nextIndex);
      const [nextComputerMove, nextCorrectMove] = nextMoves(nextIndex, moves);
      if (!nextComputerMove) {
        setDone(true);
        return PUZZLE_SOLVED;
      }

      setComputerMove(nextComputerMove);
      setCorrectMove(nextCorrectMove);
    }

    return CORRECT_MOVE;
  };

  return {
    fen: position,
    move: updatePosition,
    computerMove,
    correctMove,
    done,
  };
}

function nextMoves(idx: number, moves: string[]): string[] {
  if (idx >= moves.length) {
    return [];
  }

  const nextComputerMove = moves[idx];
  const nextCorrectMove = moves[idx + 1];
  return [nextComputerMove, nextCorrectMove];
}
