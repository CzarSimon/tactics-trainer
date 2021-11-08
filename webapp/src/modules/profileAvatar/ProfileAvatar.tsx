import React from 'react';
import { Avatar, Popover } from 'antd';
import log from '@czarsimon/remotelogger';
import { useAuth } from '../../state/auth/hooks';
import { ProfileInfo } from './ProfileInfo';

import styles from './ProfileAvatar.module.css';

export function ProfileAvatar() {
  const { user, logout } = useAuth();
  if (!user) {
    return null;
  }

  const onLogout = () => {
    log.info(`user(id=${user.id}) logged out`);
    logout();
  };

  return (
    <div className={styles.ProfileAvatar}>
      <Popover placement="topRight" content={<ProfileInfo user={user} logout={onLogout} />}>
        <Avatar size="large" className={styles.Avatar}>
          {user.username[0].toUpperCase()}
        </Avatar>
      </Popover>
    </div>
  );
}
