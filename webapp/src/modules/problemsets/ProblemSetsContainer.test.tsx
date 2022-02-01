import React from 'react';
import { screen, waitFor } from '@testing-library/react';
import { render } from '../../testutils';
import { mockRequests } from '../../api/httpclient';
import { Cycle, ProblemSet } from '../../types';
import { ProblemSetsContainer } from './PromblemSetsContainer';
import userEvent from '@testing-library/user-event';
import { EMTPY_DATE } from '../../constants';

const problemSet1: ProblemSet = {
  id: 'ps-1',
  name: 'name PS 1',
  description: 'descriptin of PS 1',
  themes: [],
  ratingInterval: '1500 - 1600',
  userId: 'user-1',
  archived: false,
  puzzleIds: [],
  createdAt: ' 2021-11-04T09:54:34Z',
  updatedAt: ' 2021-11-04T09:54:34Z',
};

const problemSet2: ProblemSet = {
  id: 'ps-2',
  name: 'name PS 2',
  description: 'descriptin of PS 2',
  themes: ['endgame', 'long'],
  ratingInterval: '1800 - 1900',
  userId: 'user-1',
  archived: false,
  puzzleIds: [],
  createdAt: '2021-11-04T09:54:34Z',
  updatedAt: '2021-11-04T09:54:34Z',
};

const cycles: Cycle[] = [
  {
    id: 'c-0',
    number: 1,
    problemSetId: 'ps-1',
    currentPuzzleId: 'p-4',
    completedAt: '2021-11-12T09:54:34Z',
    createdAt: '2021-11-04T09:54:34Z',
    updatedAt: '2021-11-04T09:54:34Z',
  },
  {
    id: 'c-1',
    number: 2,
    problemSetId: 'ps-1',
    currentPuzzleId: 'p-0',
    completedAt: EMTPY_DATE,
    createdAt: '2021-11-13T09:54:34Z',
    updatedAt: '2021-11-13T09:54:34Z',
  },
];

test('check that problem sets load and can be interacted with', async () => {
  mockRequests({
    '/api/puzzle-server/v1/problem-sets': {
      body: [problemSet1, problemSet2],
      metadata: {
        method: 'GET',
        requestId: 'list-problem-sets-request-id',
        status: 200,
        url: '/api/puzzle-server/v1/problem-sets',
      },
    },
    '/api/puzzle-server/v1/problem-sets/ps-1': {
      body: { ...problemSet1, themes: ['p-0', 'p-1', 'p-2', 'p-3', 'p-4'] },
      metadata: {
        method: 'GET',
        requestId: 'get-problem-set-request-id',
        status: 200,
        url: '/api/puzzle-server/v1/problem-sets/ps-1',
      },
    },
    '/api/puzzle-server/v1/problem-sets/ps-1/cycles': {
      body: cycles,
      metadata: {
        method: 'GET',
        requestId: 'get-problem-set-cycles-request-id',
        status: 200,
        url: '/api/puzzle-server/v1/problem-sets/ps-1/cycles',
      },
    },
  });

  render(<ProblemSetsContainer />);
  expect(screen.getByRole('heading', { name: /^problem sets$/i })).toBeInTheDocument();
  const newSetButton = screen.getByRole('button', { name: /new-problem-set-button/i });
  expect(newSetButton).toBeInTheDocument();
  for (const ps of [problemSet1, problemSet2]) {
    await waitFor(
      () => {
        expect(screen.getByText(ps.name)).toBeInTheDocument();
        expect(screen.getByText(ps.ratingInterval)).toBeInTheDocument();
        if (ps.themes.length > 0) {
          expect(screen.getByText(/themes/i)).toBeInTheDocument();
        }
        expect(screen.queryByText(ps.description!)).toBeFalsy();
      },
      { timeout: 100 },
    );
  }

  expect(screen.queryByRole('heading', { name: /^cycle 2$/i })).toBeFalsy();
  expect(screen.queryByRole('heading', { name: /^cycle 1$/i })).toBeFalsy();
  expect(screen.queryByText(/^active$/i)).toBeFalsy();
  expect(screen.queryByText(/^completed$/i)).toBeFalsy();

  const ps1Name = screen.getByText(problemSet1.name);
  userEvent.click(ps1Name);
  await waitFor(
    () => {
      expect(screen.getByText(problemSet1.description!)).toBeInTheDocument();
      expect(screen.getByRole('heading', { name: /^cycles$/i })).toBeInTheDocument();
    },
    { timeout: 100 },
  );

  const newCycleButton = screen.getByRole('button', { name: /^start new cycle$/i }) as HTMLButtonElement;
  expect(newCycleButton).toBeInTheDocument();
  expect(newCycleButton.disabled).toBeTruthy();

  expect(screen.getByRole('heading', { name: /^cycle 2$/i })).toBeInTheDocument();
  expect(screen.getByRole('heading', { name: /^cycle 1$/i })).toBeInTheDocument();

  expect(screen.getByText(/^active$/i)).toBeInTheDocument();
  expect(screen.getByText(/^completed$/i)).toBeInTheDocument();

  userEvent.click(newSetButton);
  expect(window.location.pathname).toBe('/problem-sets/new');
});

test('check that problem sets load and a cycle can be created', async () => {
  mockRequests({
    '/api/puzzle-server/v1/problem-sets': {
      body: [problemSet1],
      metadata: {
        method: 'GET',
        requestId: 'list-problem-sets-request-id',
        status: 200,
        url: '/api/puzzle-server/v1/problem-sets',
      },
    },
    '/api/puzzle-server/v1/problem-sets/ps-1': {
      body: { ...problemSet1, themes: ['p-0', 'p-1', 'p-2', 'p-3', 'p-4'] },
      metadata: {
        method: 'GET',
        requestId: 'get-problem-set-request-id',
        status: 200,
        url: '/api/puzzle-server/v1/problem-sets/ps-1',
      },
    },
    '/api/puzzle-server/v1/problem-sets/ps-1/cycles': {
      body: [],
      metadata: {
        method: 'GET',
        requestId: 'get-problem-set-cycles-request-id',
        status: 200,
        url: '/api/puzzle-server/v1/problem-sets/ps-1/cycles',
      },
    },
  });

  render(<ProblemSetsContainer />);
  expect(screen.getByRole('heading', { name: /^problem sets$/i })).toBeInTheDocument();
  await waitFor(
    () => {
      expect(screen.getByText(problemSet1.name)).toBeInTheDocument();
      expect(screen.getByText(problemSet1.ratingInterval)).toBeInTheDocument();
      if (problemSet1.themes.length > 0) {
        expect(screen.getByText(/themes/i)).toBeInTheDocument();
      }
      expect(screen.queryByText(problemSet1.description!)).toBeFalsy();
    },
    { timeout: 100 },
  );

  const ps1Name = screen.getByText(problemSet1.name);
  userEvent.click(ps1Name);
  await waitFor(
    () => {
      expect(screen.getByText(problemSet1.description!)).toBeInTheDocument();
      expect(screen.getByRole('heading', { name: /^cycles$/i })).toBeInTheDocument();
    },
    { timeout: 100 },
  );

  const newCycleButton = screen.getByRole('button', { name: /^start new cycle$/i }) as HTMLButtonElement;
  expect(newCycleButton).toBeInTheDocument();
  expect(newCycleButton.disabled).toBeFalsy();
});

test('check that you can goto a cycle', async () => {
  mockRequests({
    '/api/puzzle-server/v1/problem-sets': {
      body: [problemSet1],
      metadata: {
        method: 'GET',
        requestId: 'list-problem-sets-request-id',
        status: 200,
        url: '/api/puzzle-server/v1/problem-sets',
      },
    },
    '/api/puzzle-server/v1/problem-sets/ps-1': {
      body: { ...problemSet1, themes: ['p-0', 'p-1', 'p-2', 'p-3', 'p-4'] },
      metadata: {
        method: 'GET',
        requestId: 'get-problem-set-request-id',
        status: 200,
        url: '/api/puzzle-server/v1/problem-sets/ps-1',
      },
    },
    '/api/puzzle-server/v1/problem-sets/ps-1/cycles': {
      body: cycles,
      metadata: {
        method: 'GET',
        requestId: 'get-problem-set-cycles-request-id',
        status: 200,
        url: '/api/puzzle-server/v1/problem-sets/ps-1/cycles',
      },
    },
  });

  render(<ProblemSetsContainer />);
  expect(screen.getByRole('heading', { name: /^problem sets$/i })).toBeInTheDocument();
  await waitFor(
    () => {
      expect(screen.getByText(problemSet1.name)).toBeInTheDocument();
      expect(screen.getByText(problemSet1.ratingInterval)).toBeInTheDocument();
      if (problemSet1.themes.length > 0) {
        expect(screen.getByText(/themes/i)).toBeInTheDocument();
      }
      expect(screen.queryByText(problemSet1.description!)).toBeFalsy();
    },
    { timeout: 100 },
  );

  const ps1Name = screen.getByText(problemSet1.name);
  userEvent.click(ps1Name);
  await waitFor(
    () => {
      expect(screen.getByText(problemSet1.description!)).toBeInTheDocument();
      expect(screen.getByRole('heading', { name: /^cycles$/i })).toBeInTheDocument();
    },
    { timeout: 100 },
  );

  const cycle2Card = screen.getByRole('heading', { name: /^cycle 2$/i });
  expect(cycle2Card).toBeInTheDocument();
  expect(screen.getByRole('heading', { name: /^cycle 1$/i })).toBeInTheDocument();

  userEvent.click(cycle2Card);
  expect(window.location.pathname).toBe('/cycles/c-1');
});
