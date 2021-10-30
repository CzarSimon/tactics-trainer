import React from 'react';
import { screen, render } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { ErrorContainer } from './ErrorContainer';
import { ErrorContext } from '../../state/error/ErrorContext';
import { ErrorInfo } from '../../types';

function setError(error: ErrorInfo) {
  console.log(error.title);
}

test('error container: not visible when no error', async () => {
  const error: ErrorInfo = {
    title: 'Test Error',
    details: 'Error Details',
  };

  render(
    <ErrorContext.Provider value={{ setError }}>
      <ErrorContainer />
    </ErrorContext.Provider>,
  );

  const errorTitle = screen.queryByText(error.title);
  expect(errorTitle).toBeFalsy();

  const errorDetails = screen.queryByText(error.details);
  expect(errorDetails).toBeFalsy();
});

test('error container: visible when error present', async () => {
  const error: ErrorInfo = {
    title: 'Test Error',
    details: 'Error Details',
  };

  render(
    <ErrorContext.Provider value={{ error, setError }}>
      <ErrorContainer />
    </ErrorContext.Provider>,
  );

  const errorTitle = screen.getByText(error.title);
  expect(errorTitle).toBeInTheDocument();

  const errorDetails = screen.getByText(error.details);
  expect(errorDetails).toBeInTheDocument();
});
