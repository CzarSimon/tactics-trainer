import React from 'react';
import { Signup } from './components/Signup';
import { AuthenticationRequest } from '../../types';
import { useAuth } from '../../state/auth/hooks';
import { Redirect } from 'react-router';

export function SignupContainer() {
  const { authenticated, signup } = useAuth();
  const handleSignup = (req: AuthenticationRequest) => {
    signup(req);
  };

  return authenticated ? <Redirect to="/" /> : <Signup submit={handleSignup} />;
}
