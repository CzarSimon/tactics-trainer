import React from 'react';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import { Home } from '../modules/home/Home';
import { PuzzlePage } from '../modules/puzzle/PuzzleSolver';

export function Routes() {
  return (
    <Router>
      <Switch>
        <Route path="/puzzles/:puzzleId">
          <PuzzlePage />
        </Route>
        <Route path="/">
          <Home />
        </Route>
      </Switch>
    </Router>
  );
}
