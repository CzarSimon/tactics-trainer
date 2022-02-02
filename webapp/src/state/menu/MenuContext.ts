import { createContext } from 'react';

export interface MenuState {
  open: boolean;
  setOpen: (value: boolean) => void;
}

const initalState: MenuState = {
  open: false,
  setOpen: (value: boolean) => {}, // eslint-disable-line
};

export const MenuContext = createContext<MenuState>(initalState);
