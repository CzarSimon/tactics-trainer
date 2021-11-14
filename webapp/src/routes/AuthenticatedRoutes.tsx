import React from 'react';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import { CycleContainer } from '../modules/cycles';
import { NewProblemSetContainer, ProblemSetsContainer } from '../modules/problemsets';
import { ProfileAvatar } from '../modules/profileAvatar';

export function AuthenticatedRoutes() {
  return (
    <Router>
      <ProfileAvatar />
      <Switch>
        <Route path="/cycles/:cycleId">
          <CycleContainer />
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
