import React from 'react';
import PropTypes from 'prop-types';
import { render as rtlRender } from '@testing-library/react';
import { BrowserRouter as Router } from 'react-router-dom';
import log, { ConsoleHandler, level } from '@czarsimon/remotelogger';
import { QueryClient, QueryClientProvider } from 'react-query';

const logHandlers = { console: new ConsoleHandler(level.DEBUG) };
log.configure(logHandlers);

Object.defineProperty(window, 'matchMedia', {
  writable: true,
  value: jest.fn().mockImplementation((query) => ({
    matches: false,
    media: query,
    onchange: null,
    addListener: jest.fn(), // deprecated
    removeListener: jest.fn(), // deprecated
    addEventListener: jest.fn(),
    removeEventListener: jest.fn(),
    dispatchEvent: jest.fn(),
  })),
});

export function render(ui, { ...renderOptions } = {}) {
  function Wrapper({ children }) {
    const queryClient = new QueryClient();
    return (
      <QueryClientProvider client={queryClient}>
        <Router>{children}</Router>
      </QueryClientProvider>
    );
  }

  Wrapper.propTypes = {
    children: PropTypes.any,
  };
  return rtlRender(ui, { wrapper: Wrapper, ...renderOptions });
}
