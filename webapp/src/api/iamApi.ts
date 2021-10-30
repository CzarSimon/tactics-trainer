import { httpclient } from './httpclient';
import { AuthenticationRequest, AuthenticationResponse } from '../types';
import { wrapAndLogError } from './util';

const IAM_SERVER_URL = '/api/iam-server';

export async function login(req: AuthenticationRequest): Promise<AuthenticationResponse> {
  const { body, error, metadata } = await httpclient.post<AuthenticationResponse>({
    url: `${IAM_SERVER_URL}/v1/login`,
    body: req,
  });

  if (!body) {
    throw wrapAndLogError(`failed to login(username=${req.username})`, error, metadata);
  }

  return body;
}
