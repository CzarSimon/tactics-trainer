import { httpclient, setHeader } from './httpclient';
import { AuthenticationRequest, AuthenticationResponse } from '../types';
import { wrapAndLogError } from './util';
import { AUTH_TOKEN_KEY } from '../constants';

const IAM_SERVER_URL = '/api/iam-server';

export async function login(req: AuthenticationRequest): Promise<void> {
  const { body, error, metadata } = await httpclient.post<AuthenticationResponse>({
    url: `${IAM_SERVER_URL}/v1/login`,
    body: req,
  });

  if (!body) {
    throw wrapAndLogError(`failed to login(username=${req.username})`, error, metadata);
  }

  setHeader('Authorization', `Bearer ${body.token}`);
  localStorage.setItem(AUTH_TOKEN_KEY, body.token);
}
