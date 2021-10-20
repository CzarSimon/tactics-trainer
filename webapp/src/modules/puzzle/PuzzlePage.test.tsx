import React from 'react';
import { act, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { render } from '../../testutils';
import { PuzzlePage } from './PuzzlePage';
import { mockRequests } from '../../api/httpclient';
import { Puzzle } from '../../types';

// Puzzle id Should equal to the puzzleId value in the react-router-dom useParams mock below.
const puzzle: Puzzle = {
  id: '0bfb3be7-3a30-44e7-8db9-5134795aae84',
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

jest.mock('react-router-dom', () => ({
  ...jest.requireActual('react-router-dom'),
  useParams: () => ({
    puzzleId: '0bfb3be7-3a30-44e7-8db9-5134795aae84',
  }),
}));

test('renders puzzle page', async () => {
  mockRequests({
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

  await act(async () => {
    render(<PuzzlePage />);
  });

  const title = screen.getByRole('heading', { name: /^black to move$/i });
  expect(title).toBeInTheDocument();

  const puzzleDetailsCollapse = screen.getByText(/^puzzle details$/i);
  expect(puzzleDetailsCollapse).toBeInTheDocument();

  // Check that puzzle details are not visible
  const rating = screen.queryByText(/^rating: 1215$/i);
  expect(rating).toBeNull();
  const themesHeading = screen.queryByRole('heading', { name: /^themes$/i });
  expect(themesHeading).toBeNull();
  const aTheme = screen.queryByText(/^crushing$/i);
  expect(aTheme).toBeNull();

  // userEvent.click(puzzleDetailsCollapse);
});
