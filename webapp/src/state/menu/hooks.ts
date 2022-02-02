import { useContext } from 'react';
import { MenuContext, MenuState } from './MenuContext';

export function useMenuState(): MenuState {
  return useContext(MenuContext);
}
