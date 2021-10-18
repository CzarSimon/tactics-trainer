import React from 'react';
import { screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { render } from '../../testutils';
import { Home } from './Home';

test('renders home page', () => {
  render(<Home />);
  const title = screen.getByRole('heading', { name: /^tactics trainer$/i });
  expect(title).toBeInTheDocument();
  const button = screen.getByRole('button', { name: /^get random puzzle$/i });
  expect(button).toBeInTheDocument();

  expect(window.location.pathname).toBe('/');
  userEvent.click(button);
  expect(window.location.pathname).toMatch(/^\/puzzles\/[0-9a-f-]{36}/);
});
