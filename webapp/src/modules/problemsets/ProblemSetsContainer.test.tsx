import React from 'react';
import { screen, waitFor } from '@testing-library/react';
import { render } from '../../testutils';
import { mockRequests } from '../../api/httpclient';
import { ProblemSet } from '../../types';
import { ProblemSetsContainer } from './PromblemSetsContainer';

const problemSet1: ProblemSet = {
  id: 'ps-1',
  name: 'name PS 1',
  description: 'descriptin of PS 1',
  themes: [],
  ratingInterval: '1500 - 1600',
  userId: 'user-1',
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
  puzzleIds: [],
  createdAt: ' 2021-11-04T09:54:34Z',
  updatedAt: ' 2021-11-04T09:54:34Z',
};

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
  });

  render(<ProblemSetsContainer />);
  expect(screen.getByRole('heading', { name: /^Problem sets$/i })).toBeInTheDocument();
  for (const ps of [problemSet1, problemSet2]) {
    await waitFor(
      () => {
        expect(screen.getByText(ps.name)).toBeInTheDocument();
        expect(screen.getByText(ps.ratingInterval)).toBeInTheDocument();
        if (ps.themes.length > 0) {
          expect(screen.getByText(/themes/i)).toBeInTheDocument();
        }
      },
      { timeout: 100 },
    );
  }
});
