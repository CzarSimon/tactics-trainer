import React from 'react';
import { screen, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { render } from '../../testutils';
import { mockRequests } from '../../api/httpclient';
import { Cycle, Puzzle } from '../../types';
import { CycleContainer } from '.';

// Cycle id Should equal to the cycleId value in the react-router-dom useParams mock below.
const cycle: Cycle = {
  id: 'db01b275-38c4-497f-b526-736c9d682d3d',
  number: 1,
  problemSetId: 'ps-0',
  currentPuzzleId: '0bfb3be7-3a30-44e7-8db9-5134795aae84',
  createdAt: '2021-11-17T16:31:32.084102Z',
  updatedAt: '2021-11-17T16:31:32.084102Z',
};

const puzzle: Puzzle = {
  id: cycle.currentPuzzleId,
  externalId: '02F4S',
  fen: '8/4n1k1/4P3/3p2PP/rp1P4/3K4/P3R3/8 w - - 0 47',
  gameUrl: 'https://lichess.org/XkKfgmwN#93',
  moves: ['e2f2', 'a4a3', 'd3c2', 'a3a2', 'c2b3', 'a2f2'],
  popularity: 100,
  rating: 1215,
  ratingDeviation: 75,
  themes: ['crushing', 'endgame', 'interference', 'long', 'skewer'],
  updatedAt: '2021-09-17T16:31:32.084102Z',
  createdAt: '2021-09-17T16:31:32.084102Z',
};

test('fetches cycle, puzzle and renders puzzle', async () => {
  mockRequests({
    // using undefined in the route as a hack to get around having to mock useParams from react-router-dom
    ['/api/puzzle-server/v1/cycles/undefined']: {
      body: cycle,
      metadata: {
        method: 'GET',
        requestId: 'get-cycle-by-id',
        status: 200,
        url: `/api/puzzle-server/v1/cycles/${cycle.id}`,
      },
    },
    [`/api/puzzle-server/v1/puzzles/${puzzle.id}`]: {
      body: puzzle,
      metadata: {
        method: 'GET',
        requestId: 'get-puzzle-by-id',
        status: 200,
        url: `/api/puzzle-server/v1/puzzles/${puzzle.id}`,
      },
    },
  });

  render(<CycleContainer />);
  await waitFor(
    () => {
      expect(screen.getByRole('heading', { name: /^black to move$/i })).toBeInTheDocument();
    },
    { timeout: 100 },
  );

  const puzzleDetailsCollapse = screen.getByText(/^puzzle details$/i);
  expect(puzzleDetailsCollapse).toBeInTheDocument();

  // Check that puzzle details are not visible
  expect(screen.queryByText(/^rating: 1215$/i)).toBeNull();
  expect(screen.queryByText(/^themes$/i)).toBeNull();
  for (const theme of puzzle.themes) {
    expect(screen.queryByText(theme)).toBeNull();
  }

  userEvent.click(puzzleDetailsCollapse);

  // Check that puzzle details are now visible
  expect(screen.getByText(/^rating: 1215$/i)).toBeInTheDocument();
  expect(screen.getByText(/^themes$/i)).toBeInTheDocument();
  for (const theme of puzzle.themes) {
    expect(screen.getByText(theme)).toBeInTheDocument();
  }

  const closeButton = screen.getByRole('button', { name: /close-button/i });
  expect(closeButton).toBeInTheDocument();
  userEvent.click(closeButton);
});

test('fetches cycle, puzzle and renders puzzle', async () => {
  mockRequests({
    // using undefined in the route as a hack to get around having to mock useParams from react-router-dom
    ['/api/puzzle-server/v1/cycles/undefined']: {
      body: { ...cycle, completedAt: '2021-11-18T16:31:32.084102Z' },
      metadata: {
        method: 'GET',
        requestId: 'get-cycle-by-id',
        status: 200,
        url: `/api/puzzle-server/v1/cycles/${cycle.id}`,
      },
    },
    [`/api/puzzle-server/v1/puzzles/${puzzle.id}`]: {
      body: puzzle,
      metadata: {
        method: 'GET',
        requestId: 'get-puzzle-by-id',
        status: 200,
        url: `/api/puzzle-server/v1/puzzles/${puzzle.id}`,
      },
    },
  });

  render(<CycleContainer />);
  await waitFor(
    () => {
      expect(screen.getByText(/cycle 1 completed/i)).toBeInTheDocument();
    },
    { timeout: 100 },
  );

  expect(screen.queryByRole('heading', { name: /^black to move$/i })).toBeFalsy();
});
