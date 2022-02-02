import React from 'react';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import { AboutPage } from '../modules/aboutpage';
import { CycleContainer } from '../modules/cycles';
import { SideMenu } from '../modules/menu';
import { NewProblemSetContainer, ProblemSetsContainer } from '../modules/problemsets';

export function AuthenticatedRoutes() {
  return (
    <Router>
      <SideMenu />
      <Switch>
        <Route path="/cycles/:cycleId">
          <CycleContainer />
        </Route>
        <Route path="/problem-sets/new">
          <NewProblemSetContainer />
        </Route>
        <Route path="/about">
          <AboutPage />
        </Route>
        <Route path="/">
          <ProblemSetsContainer />
        </Route>
      </Switch>
    </Router>
  );
}
