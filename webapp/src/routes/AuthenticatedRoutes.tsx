import React from 'react';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import { NewProblemSetContainer } from '../modules/problemsets/NewProblemSetContainer';
import { ProblemSetsContainer } from '../modules/problemsets/PromblemSetsContainer';
import { PuzzlePage } from '../modules/puzzle/PuzzlePage';

export function AuthenticatedRoutes() {
  return (
    <Router>
      <Switch>
        <Route path="/puzzles/:puzzleId">
          <PuzzlePage />
        </Route>
        <Route path="/problem-sets/new">
          <NewProblemSetContainer />
        </Route>
        <Route path="/">
          <ProblemSetsContainer />
        </Route>
      </Switch>
    </Router>
  );
}
