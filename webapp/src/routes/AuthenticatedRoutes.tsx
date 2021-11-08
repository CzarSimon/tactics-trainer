import React from 'react';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import { NewProblemSetContainer, ProblemSetsContainer } from '../modules/problemsets';
import { ProfileAvatar } from '../modules/profileAvatar';
import { PuzzlePage } from '../modules/puzzle/PuzzlePage';

export function AuthenticatedRoutes() {
  return (
    <Router>
      <ProfileAvatar />
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
