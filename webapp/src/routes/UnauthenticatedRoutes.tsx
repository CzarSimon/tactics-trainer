import React from 'react';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import { AboutPageContainer } from '../modules/aboutpage';
import { LoginContainer } from '../modules/login';
import { SignupContainer } from '../modules/signup';

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
        <Route path="/about">
          <AboutPageContainer />
        </Route>
        <Route path="/">
          <SignupContainer />
        </Route>
      </Switch>
    </Router>
  );
}
