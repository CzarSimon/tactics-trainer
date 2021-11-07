import React from 'react';
import { screen } from '@testing-library/react';
import { render } from '../../testutils';
import { NewProblemSetContainer } from './NewProblemSetContainer';
import userEvent from '@testing-library/user-event';

test('check that new problem set form loads and can be interacted with', async () => {
  render(<NewProblemSetContainer />);
  expect(screen.getByRole('heading', { name: /^create new problem set$/i })).toBeInTheDocument();

  const nameInput = screen.getByPlaceholderText('Name');
  expect(nameInput).toBeInTheDocument();

  const descriptionInput = screen.getByPlaceholderText('Description');
  expect(descriptionInput).toBeInTheDocument();

  const themesDropdown = screen.getByText('Themes');
  expect(themesDropdown).toBeInTheDocument();

  const sizeInput = screen.getByLabelText('Number of puzzles');
  expect(sizeInput).toBeInTheDocument();

  const createButton = screen.getByRole('button', { name: /^create$/i });
  expect(createButton).toBeInTheDocument();

  const cancelButton = screen.getByRole('button', { name: /^cancel$/i });
  expect(cancelButton).toBeInTheDocument();
});

test('check that new problem set form can be closed', () => {
  render(<NewProblemSetContainer />);
  expect(screen.getByRole('heading', { name: /^create new problem set$/i })).toBeInTheDocument();

  const cancelButton = screen.getByRole('button', { name: /^cancel$/i });
  expect(cancelButton).toBeInTheDocument();

  userEvent.click(cancelButton);
  expect(window.location.pathname).toBe('/');
});
