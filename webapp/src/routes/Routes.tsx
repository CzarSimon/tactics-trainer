import React from 'react';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import { Home } from '../modules/home/Home';
import { LoginContainer } from '../modules/login/LoginContainer';
import { PuzzlePage } from '../modules/puzzle/PuzzlePage';
import { SignupContainer } from '../modules/signup/SignupContainer';

import styles from './Routes.module.css';

export function Routes() {
  return (
    <Router>
      <div className={styles.Content}>
        <Switch>
          <Route path="/puzzles/:puzzleId">
            <PuzzlePage />
          </Route>
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
      </div>
    </Router>
  );
}
