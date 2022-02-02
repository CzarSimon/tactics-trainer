import React, { useState } from 'react';
import { MenuContext } from './MenuContext';

interface Props {
  children: JSX.Element;
}

export function MenuProvider({ children }: Props) {
  const [open, setOpen] = useState<boolean>(false);

  return <MenuContext.Provider value={{ open, setOpen }}>{children}</MenuContext.Provider>;
}
