import React, { useEffect } from 'react';
import { initLoggerAndHttpclient, readUser } from './init';
import { Routes } from './routes/Routes';
import { ErrorContainer } from './modules/error/ErrorContainer';
import { useAuth } from './state/auth/hooks';

import 'antd/dist/antd.css';
import styles from './App.module.css';

function App() {
  const { authenticate } = useAuth();
  useEffect(() => {
    initLoggerAndHttpclient();
    const user = readUser();
    if (user) {
      authenticate(user);
    }
  }, []); // eslint-disable-line react-hooks/exhaustive-deps

  return (
    <div className={styles.App}>
      <ErrorContainer />
      <Routes />
    </div>
  );
}

export default App;
