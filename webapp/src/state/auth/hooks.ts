import { useContext } from 'react';
import { useHistory } from 'react-router-dom';
import * as api from '../../api';
import { setHeader, removeHeader } from '../../api/httpclient';
import { AuthenticationRequest, AuthenticationResponse, UseAuthResult } from '../../types';
import { AUTH_TOKEN_KEY, CURRENT_USER_KEY } from '../../constants';
import { AuthContext } from './AuthContext';
import { useError } from '../error/hooks';

export function useAuth(): UseAuthResult {
  const { user, authenticated, authenticate, logout } = useContext(AuthContext);
  const history = useHistory();
  const { setError } = useError();

  const handleAuthResponse = (res?: AuthenticationResponse) => {
    if (res) {
      storeAuthInfo(res);
      setTimeout(() => {
        authenticate(res.user);
        history.push('/');
      }, 100);
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

  const onLogout = () => {
    logout();
    removeAuthInfo();
    history.push('/login');
  };

  return {
    login,
    signup,
    user,
    authenticated,
    authenticate,
    logout: onLogout,
  };
}

export function useIsAuthenticated(): boolean {
  const { authenticated } = useContext(AuthContext);
  return authenticated;
}

function storeAuthInfo({ user, token }: AuthenticationResponse) {
  setHeader('Authorization', `Bearer ${token}`);
  localStorage.setItem(AUTH_TOKEN_KEY, token);
  localStorage.setItem(CURRENT_USER_KEY, JSON.stringify(user));
}

function removeAuthInfo() {
  removeHeader('Authorization');
  localStorage.removeItem(AUTH_TOKEN_KEY);
  localStorage.removeItem(CURRENT_USER_KEY);
}
