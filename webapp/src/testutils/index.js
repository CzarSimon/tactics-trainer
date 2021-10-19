import React from 'react';
import PropTypes from 'prop-types';
import { render as rtlRender } from '@testing-library/react';
import { BrowserRouter as Router } from 'react-router-dom';
import log, { ConsoleHandler, level } from '@czarsimon/remotelogger';

const logHandlers = { console: new ConsoleHandler(level.DEBUG) };
log.configure(logHandlers);

export function render(ui, { ...renderOptions } = {}) {
  function Wrapper({ children }) {
    return <Router>{children}</Router>;
  }

  Wrapper.propTypes = {
    children: PropTypes.any,
  };
  return rtlRender(ui, { wrapper: Wrapper, ...renderOptions });
}
