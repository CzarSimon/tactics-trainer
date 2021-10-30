import React from 'react';
import PropTypes from 'prop-types';
import { render as rtlRender } from '@testing-library/react';
import { BrowserRouter as Router } from 'react-router-dom';
import log, { ConsoleHandler, level } from '@czarsimon/remotelogger';
import { QueryClient, QueryClientProvider } from 'react-query';
import { AuthProvider } from '../state/auth/AuthProvider';

const logHandlers = { console: new ConsoleHandler(level.DEBUG) };
log.configure(logHandlers);

// eslint-disable-next-line no-undef
global.matchMedia =
  global.matchMedia || // eslint-disable-line no-undef
  function () {
    return {
      addListener: jest.fn(),
      removeListener: jest.fn(),
    };
  };

export function render(ui, { ...renderOptions } = {}) {
  function Wrapper({ children }) {
    const queryClient = new QueryClient();
    return (
      <QueryClientProvider client={queryClient}>
        <AuthProvider>
          <Router>{children}</Router>
        </AuthProvider>
      </QueryClientProvider>
    );
  }

  Wrapper.propTypes = {
    children: PropTypes.any,
  };
  return rtlRender(ui, { wrapper: Wrapper, ...renderOptions });
}
