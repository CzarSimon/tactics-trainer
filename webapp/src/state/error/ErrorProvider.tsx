import React, { useState } from 'react';
import { ErrorContext } from './ErrorContext';
import { ErrorInfo, Optional } from '../../types';

interface Props {
  children: JSX.Element;
}

export function ErrorProvider({ children }: Props) {
  const [error, setError] = useState<Optional<ErrorInfo>>(undefined);

  return <ErrorContext.Provider value={{ error, setError }}>{children}</ErrorContext.Provider>;
}
