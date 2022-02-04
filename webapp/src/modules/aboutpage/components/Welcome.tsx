import React from 'react';
import { portraitMode } from '../../../util';
import { StepContent } from '../types';

import styles from './Step.module.css';

function Welcome() {
  const className = portraitMode() ? styles.PortraitStep : styles.LandscapeStep;

  return (
    <div className={className}>
      <h2>Welcome to tactics trainer!</h2>
      <p>The goal of this tool is simple, to help you improve your chess tactics.</p>
      <p>How is it diffent from all other tactics trainers out there you ask?</p>
      <p>
        Tactics trainer is heavlity inspired by the fantastic book{' '}
        <a href="https://www.amazon.com/Woodpecker-Method-Axel-Smith-Tikkanen/dp/1784830550#:~:text=The%20quick%20explanation%20of%20the,re%2Dprogramming%20your%20unconscious%20mind.">
          The Woodpecker Method
        </a>{' '}
        by Axel Smith & Hans Tikkanen.
      </p>
      <p>
        It lets you build your own collections of puzzles with themes and difficullty levels that fits you perfectly and
        then you can use those in order to form a training regiment that will allow you to blunder less, calculate and
        spot tactics better.
      </p>
    </div>
  );
}

export const welcomeStep: StepContent = {
  title: 'Welcome',
  content: <Welcome />,
};
