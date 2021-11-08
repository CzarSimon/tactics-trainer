import { Button } from 'antd';
import React from 'react';
import { User } from '../../types';

import styles from './ProfileInfo.module.css';

interface Props {
  user: User;
  logout: () => void;
}

export function ProfileInfo({ user, logout }: Props) {
  const { username } = user;
  return (
    <div className={styles.ProfileInfo}>
      <p>{username}</p>
      <Button type="ghost" shape="round" onClick={logout}>
        Log out
      </Button>
    </div>
  );
}
