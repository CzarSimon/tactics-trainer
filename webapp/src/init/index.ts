import { v4 as uuid } from 'uuid';
import log, { Handlers, ConsoleHandler, HttploggerHandler, level } from '@czarsimon/remotelogger';
import { CLIENT_ID_KEY, DEV_MODE, APP_NAME, APP_VERSION, CURRENT_USER_KEY } from '../constants';
import { Optional, Client, User } from '../types';
import { initHttpclient } from '../api/httpclient';

export function readUser(): Optional<User> {
  const serializedUser = localStorage.getItem(CURRENT_USER_KEY);
  if (!serializedUser) {
    return;
  }

  return JSON.parse(serializedUser);
}

export function initLoggerAndHttpclient(): void {
  const client = getClientInfo();
  const handlers = getLogHandlers(client);
  initHttpclient(client, handlers);

  log.configure(handlers);
  log.debug('initiated remotelogger');
  log.debug('initiated httpclient');
}

function getLogHandlers(client: Client): Handlers {
  const consoleLevel = DEV_MODE ? level.DEBUG : level.ERROR;
  const httpLevel = DEV_MODE ? level.DEBUG : level.INFO;
  return {
    console: new ConsoleHandler(consoleLevel),
    httplogger: new HttploggerHandler(httpLevel, {
      url: '/api/httplogger/v1/logs',
      app: APP_NAME,
      version: APP_VERSION,
      sessionId: client.sessionId,
      clientId: client.id,
    }),
  };
}

export function teardown(): void {
  log.info('closed application');
}

function getClientInfo(): Client {
  return {
    id: getOrCreateId(CLIENT_ID_KEY),
    sessionId: uuid(),
  };
}

function getOrCreateId(key: string): string {
  const id = localStorage.getItem(key);
  if (id) {
    return id;
  }

  const newId = uuid();
  localStorage.setItem(key, newId);
  return newId;
}
