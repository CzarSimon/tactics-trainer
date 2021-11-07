import React from 'react';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import { Home } from '../modules/home/Home';
import { LoginContainer } from '../modules/login/LoginContainer';
import { SignupContainer } from '../modules/signup/SignupContainer';

export function UnuthenticatedRoutes() {
  return (
    <Router>
      <Switch>
        <Route path="/signup">
          <SignupContainer />
        </Route>
        <Route path="/login">
          <LoginContainer />
        </Route>
        <Route path="/">
          <Home />
        </Route>
      </Switch>
    </Router>
  );
}
