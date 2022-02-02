import React, { ReactNode } from 'react';
import { Button } from 'antd';
import { MenuOutlined } from '@ant-design/icons';

import styles from './PageHeader.module.css';
import { useMenuState } from '../../state/menu/hooks';

interface Props {
  title: string;
  extra?: ReactNode;
}

export function PageHeader({ title, extra }: Props) {
  const { setOpen } = useMenuState();

  return (
    <div className={styles.PageHeader}>
      <div className={styles.Title}>
        <Button type="text" icon={<MenuOutlined />} onClick={() => setOpen(true)} />
        <h1>{title}</h1>
      </div>
      {extra}
    </div>
  );
}
