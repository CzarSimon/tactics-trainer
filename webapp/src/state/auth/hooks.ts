import { useContext } from 'react';
import { useHistory } from 'react-router-dom';
import * as api from '../../api';
import { setHeader } from '../../api/httpclient';
import { AuthenticationRequest, AuthenticationResponse, UseAuthResult } from '../../types';
import { AUTH_TOKEN_KEY, CURRENT_USER_KEY } from '../../constants';
import { AuthContext } from './AuthContext';
import { useError } from '../error/hooks';

export function useAuth(): UseAuthResult {
  const { user, authenticated, authenticate } = useContext(AuthContext);
  const history = useHistory();
  const { setError } = useError();

  const handleAuthResponse = (res?: AuthenticationResponse) => {
    if (res) {
      authenticate(res.user);
      storeAuthInfo(res);
      history.push('/');
    }
  };

  const login = async (req: AuthenticationRequest) => {
    const { data, error } = await api.login(req);
    if (error) {
      setError({
        title: 'Login failed',
        details: error.message,
      });
      return;
    }
    handleAuthResponse(data);
  };

  const signup = async (req: AuthenticationRequest) => {
    const { data, error } = await api.signup(req);
    if (error) {
      setError({
        title: 'Signup failed',
        details: error.message,
      });
      return;
    }
    handleAuthResponse(data);
  };

  return {
    login,
    signup,
    user,
    authenticated,
    authenticate,
  };
}

function storeAuthInfo({ user, token }: AuthenticationResponse) {
  setHeader('Authorization', `Bearer ${token}`);
  localStorage.setItem(AUTH_TOKEN_KEY, token);
  localStorage.setItem(CURRENT_USER_KEY, JSON.stringify(user));
}
