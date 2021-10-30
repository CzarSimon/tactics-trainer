import React from 'react';
import { screen, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { render } from '../../testutils';
import { mockRequests, httpclient } from '../../api/httpclient';
import { SignupContainer } from './SignupContainer';
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

test('signup screen renders and signup works', async () => {
  mockRequests({
    '/api/iam-server/v1/signup': {
      body: authResponse,
      metadata: {
        method: 'POST',
        requestId: 'signup-request-id',
        status: 200,
        url: '/api/iam-server/v1/signup',
      },
    },
  });

  render(<SignupContainer />);

  const title = screen.getByRole('heading', { name: /^tactics trainer$/i });
  expect(title).toBeInTheDocument();

  const usernameInput = screen.getByPlaceholderText(/^username$/i) as HTMLInputElement;
  expect(usernameInput).toBeInTheDocument();
  expect(usernameInput.value).toBe('');

  const passwordInput = screen.getByPlaceholderText(/^password$/i) as HTMLInputElement;
  expect(passwordInput).toBeInTheDocument();
  expect(passwordInput.value).toBe('');

  const signupButton = screen.getByRole('button', { name: /^sign up$/i });
  expect(signupButton).toBeInTheDocument();

  // Check that required warnings ARE NOT displayed.
  expect(screen.queryByText(/username is required/i)).toBeFalsy();
  expect(screen.queryByText(/password is required/i)).toBeFalsy();

  userEvent.click(signupButton);
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

  userEvent.click(signupButton);
  await waitFor(
    () => {
      expect(httpclient.getHeaders()['Authorization']).toBe('Bearer header.body.signature');
      expect(localStorage.getItem(AUTH_TOKEN_KEY)).toBe('header.body.signature');
    },
    { timeout: 1000 },
  );
  expect(window.location.pathname).toBe('/');
});

test('signup: redirect to login works', () => {
  render(<SignupContainer />);

  const signupButton = screen.getByRole('button', { name: /^sign up$/i });
  expect(signupButton).toBeInTheDocument();

  const loginLink = screen.getByRole('link', { name: /^log in$/i });
  expect(loginLink).toBeInTheDocument();

  userEvent.click(loginLink);
  expect(window.location.pathname).toBe('/login');
});
