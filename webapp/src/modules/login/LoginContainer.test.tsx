import React from 'react';
import { screen, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { render } from '../../testutils';
import { mockRequests, httpclient } from '../../api/httpclient';
import { LoginContainer } from './LoginContainer';
import { AuthenticationResponse } from '../../types';
import { AUTH_TOKEN_KEY } from '../../constants';

beforeEach(() => {
  localStorage.clear();
});

const authResponse: AuthenticationResponse = {
  token: 'header.body.signature',
  user: {
    id: 'some-user-id',
    username: 'valid-username',
    role: 'USER',
    createdAt: '2021-10-28T09:11:44.668685Z',
    updatedAt: '2021-10-28T09:11:44.668685Z',
  },
};

test('login screen renders and login works', async () => {
  mockRequests({
    '/api/iam-server/v1/login': {
      body: authResponse,
      metadata: {
        method: 'POST',
        requestId: 'login-request-id',
        status: 200,
        url: '/api/iam-server/v1/login',
      },
    },
  });

  render(<LoginContainer />);
  const title = screen.getByRole('heading', { name: /^tactics trainer$/i });
  expect(title).toBeInTheDocument();

  const usernameInput = screen.getByPlaceholderText(/^username$/i) as HTMLInputElement;
  expect(usernameInput).toBeInTheDocument();
  expect(usernameInput.value).toBe('');

  const passwordInput = screen.getByPlaceholderText(/^password$/i) as HTMLInputElement;
  expect(passwordInput).toBeInTheDocument();
  expect(passwordInput.value).toBe('');

  const loginButton = screen.getByRole('button', { name: /^log in$/i });
  expect(loginButton).toBeInTheDocument();

  const singupLink = screen.getByRole('link', { name: /^sign up$/i });
  expect(singupLink).toBeInTheDocument();

  // Check that required warnings ARE NOT displayed.
  expect(screen.queryByText(/username is required/i)).toBeFalsy();
  expect(screen.queryByText(/password is required/i)).toBeFalsy();

  userEvent.click(loginButton);
  await waitFor(
    () => {
      // Check that required warnings ARE displayed.
      expect(screen.getByText(/username is required/i)).toBeInTheDocument();
      expect(screen.getByText(/password is required/i)).toBeInTheDocument();
    },
    { timeout: 1000 },
  );

  userEvent.type(usernameInput, 'valid-username');
  userEvent.type(passwordInput, 'some-valid-password');
  expect(usernameInput.value).toBe('valid-username');
  expect(passwordInput.value).toBe('some-valid-password');

  userEvent.click(loginButton);
  await waitFor(
    () => {
      expect(httpclient.getHeaders()['Authorization']).toBe('Bearer header.body.signature');
      expect(localStorage.getItem(AUTH_TOKEN_KEY)).toBe('header.body.signature');
    },
    { timeout: 1000 },
  );
  expect(window.location.pathname).toBe('/');
});

test('login: redirect to signup works', async () => {
  render(<LoginContainer />);

  const loginButton = screen.getByRole('button', { name: /^log in$/i });
  expect(loginButton).toBeInTheDocument();

  const singupLink = screen.getByRole('link', { name: /^sign up$/i });
  expect(singupLink).toBeInTheDocument();

  userEvent.click(singupLink);
  expect(window.location.pathname).toBe('/signup');
});
