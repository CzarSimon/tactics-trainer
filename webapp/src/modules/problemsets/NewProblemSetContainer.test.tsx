import React from 'react';
import { screen } from '@testing-library/react';
import { render } from '../../testutils';
import { NewProblemSetContainer } from './NewProblemSetContainer';

test('check that new problem set form loads and can be interacted with', async () => {
  render(<NewProblemSetContainer />);
  expect(screen.getByRole('heading', { name: /^create new problem set$/i })).toBeInTheDocument();
});
