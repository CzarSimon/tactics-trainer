import { useContext } from 'react';
import { useHistory } from 'react-router-dom';
import * as api from '../../api';
import { setHeader } from '../../api/httpclient';
import { AuthenticationRequest, UseAuthResult } from '../../types';
import { AUTH_TOKEN_KEY } from '../../constants';
import { AuthContext } from './AuthContext';

export function useAuth(): UseAuthResult {
  const { user, authenticated, authenticate } = useContext(AuthContext);
  const history = useHistory();

  const login = async (req: AuthenticationRequest) => {
    const { token, user } = await api.login(req);
    authenticate(user);
    storeToken(token);
    history.push('/');
  };

  return {
    login,
    user,
    authenticated,
  };
}

function storeToken(token: string) {
  setHeader('Authorization', `Bearer ${token}`);
  localStorage.setItem(AUTH_TOKEN_KEY, token);
}
