import { EMTPY_DATE } from '../constants';
import { Cycle } from '../types';

export function cycleIsCompleted({ completedAt }: Cycle): boolean {
  return completedAt !== undefined && completedAt !== EMTPY_DATE;
}

export function portraitMode(): boolean {
  return window.innerHeight > window.innerWidth;
}
