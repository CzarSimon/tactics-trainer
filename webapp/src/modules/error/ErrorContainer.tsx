import React from 'react';
import { useError } from '../../state/error/hooks';
import { ErrorDisplay } from './components/ErrorDisplay';

export function ErrorContainer() {
  const { error } = useError();

  return error ? <ErrorDisplay error={error} /> : null;
}
