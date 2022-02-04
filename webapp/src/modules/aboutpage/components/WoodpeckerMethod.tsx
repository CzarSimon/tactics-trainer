import React from 'react';
import { portraitMode } from '../../../util';
import { StepContent } from '../types';

import styles from './Step.module.css';

function WoodpeckerMethod() {
  const className = portraitMode() ? styles.PortraitStep : styles.LandscapeStep;

  return (
    <div className={className}>
      <h2>The Woodpecker Method</h2>
      <p>
        The woodpecker method is developed and described by GMs Axel Smith and Hans Tikkanen and described in their
        wonderful book by the same name.
      </p>
      <p>
        In short the method aims to develop chess players tactical pattern recogintion by first selecting a large set of
        chess puzzles. After that the chess student should solve all the puzzles in the set. After all puzzles are
        solved the set of puzzles should be solved again, and then again and then again and so forth.
      </p>
      <p>
        Surely by the description above you have figured out that this takes alot of work. However what you get out is
        not only an improved calculation ability, you will also develop an intuitive pattern recognition that will make
        you much faster to spot tactics in you own games.
      </p>
      <p>
        Tactics trainer is heavily inspired by the woodpecker method, it lets you create your own sets of tactics
        puzzles, we call them problem sets, and then repeatedly solve them in what we call cycles.
      </p>
      <p>
        Before you get started, click the next button and you will learn how to create your own problem sets and how to
        get started solving the puzzles in the set.
      </p>
      <p>
        For a more detailed description of the woodpecker method i highly recommend you buy the{' '}
        <a href="https://www.amazon.com/Woodpecker-Method-Axel-Smith-Tikkanen/dp/1784830550#:~:text=The%20quick%20explanation%20of%20the,re%2Dprogramming%20your%20unconscious%20mind.">
          book
        </a>{' '}
        as it explains the method better than I ever could :)
      </p>
    </div>
  );
}

export const woodpeckerStep: StepContent = {
  title: 'The Woodpecker Method',
  content: <WoodpeckerMethod />,
};
