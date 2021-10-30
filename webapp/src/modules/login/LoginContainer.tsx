import React from 'react';
import { Login } from './components/Login';
import { AuthenticationRequest } from '../../types';
import { useAuth } from '../../state/auth/hooks';
import { Redirect } from 'react-router';

export function LoginContainer() {
  const { authenticated, login } = useAuth();
  const handleLogin = (req: AuthenticationRequest) => {
    login(req);
  };

  return authenticated ? <Redirect to="/" /> : <Login submit={handleLogin} />;
}
