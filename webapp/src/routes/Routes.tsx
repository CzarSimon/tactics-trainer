import React from 'react';
import { useIsAuthenticated } from '../state/auth/hooks';
import { AuthenticatedRoutes } from './AuthenticatedRoutes';
import { UnuthenticatedRoutes } from './UnauthenticatedRoutes';

export function Routes() {
  const authenticated = useIsAuthenticated();
  return authenticated ? <AuthenticatedRoutes /> : <UnuthenticatedRoutes />;
}
