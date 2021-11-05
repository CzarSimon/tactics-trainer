import React from 'react';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import { LoginContainer } from '../modules/login/LoginContainer';
import { ProblemSetsContainer } from '../modules/problemsets/PromblemSetsContainer';
import { PuzzlePage } from '../modules/puzzle/PuzzlePage';
import { SignupContainer } from '../modules/signup/SignupContainer';

export function Routes() {
  return (
    <Router>
      <Switch>
        <Route path="/signup">
          <SignupContainer />
        </Route>
        <Route path="/login">
          <LoginContainer />
        </Route>
        <Route path="/puzzles/:puzzleId">
          <PuzzlePage />
        </Route>
        <Route path="/">
          <ProblemSetsContainer />
        </Route>
      </Switch>
    </Router>
  );
}
