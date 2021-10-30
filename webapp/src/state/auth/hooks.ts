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

  const login = async (req: AuthenticationRequest) => {
    const { data, error } = await api.login(req);
    if (error) {
      setError({
        title: 'Login failed',
        details: error.message,
      });
      return;
    }
    if (data) {
      authenticate(data.user);
      storeAuthInfo(data);
      history.push('/');
    }
  };

  return {
    login,
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
