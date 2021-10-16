import { Fetch, HttpClient } from "@czarsimon/httpclient";
import { Handlers } from "@czarsimon/remotelogger";
import { Client } from "../types";

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
  })
}

export function setHeader(name: string, value: string) {
  httpclient.setHeaders({
    ...httpclient.getHeaders(),
    [name]: value,
  });
}
