import { Fetch, HttpClient, HTTPResponse, MockTransport, Headers } from '@czarsimon/httpclient';
import { level, Handlers, ConsoleHandler } from '@czarsimon/remotelogger';
import { AUTH_TOKEN_KEY } from '../constants';
import { Client, TypedMap } from '../types';

export let httpclient = new HttpClient({
  baseHeaders: {
    'Content-Type': 'application/json',
  },
});

export function initHttpclient(client: Client, handlers: Handlers) {
  const baseHeaders: Headers = {
    ...httpclient.getHeaders(),
    'X-Client-ID': client.id,
    'X-Session-ID': client.sessionId,
  };

  const token = localStorage.getItem(AUTH_TOKEN_KEY);
  if (token) {
    baseHeaders['Authorization'] = `Bearer ${token}`;
  }

  httpclient = new HttpClient({
    logHandlers: handlers,
    transport: new Fetch(),
    baseHeaders,
  });
}

export function setHeader(name: string, value: string) {
  httpclient.setHeaders({
    ...httpclient.getHeaders(),
    [name]: value,
  });
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any
type MockResponses = TypedMap<HTTPResponse<any>>;

export function mockRequests(resonses: MockResponses) {
  httpclient = new HttpClient({
    logHandlers: { console: new ConsoleHandler(level.DEBUG) },
    transport: new MockTransport(resonses),
  });
}
