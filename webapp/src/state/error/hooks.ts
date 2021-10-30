import { useContext } from 'react';
import { ErrorContext, ErrorState } from './ErrorContext';

export function useError(): ErrorState {
  return useContext(ErrorContext);
}
