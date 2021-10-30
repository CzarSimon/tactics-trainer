import { createContext } from 'react';
import { ErrorInfo } from '../../types';

export interface ErrorState {
  error?: ErrorInfo;
  setError: (error: ErrorInfo) => void;
}

const initalState: ErrorState = {
  setError: (error: ErrorInfo) => {}, // eslint-disable-line
};

export const ErrorContext = createContext<ErrorState>(initalState);
