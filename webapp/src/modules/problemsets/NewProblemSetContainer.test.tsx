import React from 'react';
import { screen, waitFor } from '@testing-library/react';
import { render } from '../../testutils';
import { mockRequests } from '../../api/httpclient';
import { NewProblemSetContainer } from './NewProblemSetContainer';
import userEvent from '@testing-library/user-event';
import { ProblemSet } from '../../types';

const problemSet: ProblemSet = {
  id: 'ps-1',
  name: 'name PS 1',
  description: 'descriptin of PS 1',
  themes: [],
  ratingInterval: '1500 - 1600',
  userId: 'user-1',
  puzzleIds: ['p-0', 'p-1'],
  createdAt: ' 2021-11-04T09:54:34Z',
  updatedAt: ' 2021-11-04T09:54:34Z',
};

test('check that new problem set form loads and can be interacted with', async () => {
  mockRequests({
    '/api/puzzle-server/v1/problem-sets': {
      body: problemSet,
      metadata: {
        method: 'POST',
        requestId: 'create-problem-set-request-id',
        status: 200,
        url: '/api/puzzle-server/v1/problem-sets',
      },
    },
  });

  render(<NewProblemSetContainer />);
  expect(screen.getByRole('heading', { name: /^create new problem set$/i })).toBeInTheDocument();

  const nameInput = screen.getByPlaceholderText('Name') as HTMLInputElement;
  expect(nameInput).toBeInTheDocument();

  const descriptionInput = screen.getByPlaceholderText('Description') as HTMLInputElement;
  expect(descriptionInput).toBeInTheDocument();

  expect(screen.getByText(/rating interval/i)).toBeInTheDocument();

  const themesDropdown = screen.getByText('Themes');
  expect(themesDropdown).toBeInTheDocument();

  const sizeInput = screen.getByLabelText('Number of puzzles');
  expect(sizeInput).toBeInTheDocument();

  const cancelButton = screen.getByRole('button', { name: /^cancel$/i });
  expect(cancelButton).toBeInTheDocument();

  const createButton = screen.getByRole('button', { name: /^create$/i });
  expect(createButton).toBeInTheDocument();

  // Check that required warnings ARE NOT displayed.
  expect(screen.queryByText(/name is required/i)).toBeFalsy();

  userEvent.click(createButton);

  await waitFor(
    () => {
      // Check that required warnings ARE displayed.
      expect(screen.getByText(/name is required/i)).toBeInTheDocument();
    },
    { timeout: 1000 },
  );

  userEvent.type(nameInput, 'name PS 1');
  userEvent.type(descriptionInput, 'descriptin of PS 1');
  expect(nameInput.value).toBe('name PS 1');
  expect(descriptionInput.value).toBe('descriptin of PS 1');

  userEvent.click(createButton);
});

test('check that new problem set form can be closed', () => {
  render(<NewProblemSetContainer />);
  expect(screen.getByRole('heading', { name: /^create new problem set$/i })).toBeInTheDocument();

  const cancelButton = screen.getByRole('button', { name: /^cancel$/i });
  expect(cancelButton).toBeInTheDocument();

  userEvent.click(cancelButton);
  expect(window.location.pathname).toBe('/');
});
