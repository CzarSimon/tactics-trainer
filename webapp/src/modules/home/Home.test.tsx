import React from 'react';
import { render, screen } from '@testing-library/react';
import { Home } from './Home';

test('renders home page', () => {
  render(<Home />);
  const title = screen.getByRole('heading', { name: /^tactics trainer$/i });
  expect(title).toBeInTheDocument();
  const button = screen.getByRole('button', { name: /^get random puzzle$/i });
  expect(button).toBeInTheDocument();
});
