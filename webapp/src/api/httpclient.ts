import { Fetch, HttpClient, HTTPResponse, MockTransport } from '@czarsimon/httpclient';
import { level, Handlers, ConsoleHandler } from '@czarsimon/remotelogger';
import { Client, TypedMap } from '../types';

export let httpclient = new HttpClient({
  baseHeaders: {
    'Content-Type': 'application/json',
  },
});

export function initHttpclient(client: Client, handlers: Handlers) {
  httpclient = new HttpClient({
    logHandlers: handlers,
    baseHeaders: {
      ...httpclient.getHeaders(),
      'X-Client-ID': client.id,
      'X-Session-ID': client.sessionId,
    },
    transport: new Fetch(),
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
