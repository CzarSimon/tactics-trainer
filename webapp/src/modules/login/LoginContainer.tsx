import React from 'react';
import { Login } from './components/Login';
import * as api from '../../api';
import { AuthenticationRequest } from '../../types';

export function LoginContainer() {
  const handleLogin = (req: AuthenticationRequest) => {
    api.login(req);
  };

  return <Login submit={handleLogin} />;
}
