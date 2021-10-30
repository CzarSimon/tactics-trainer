import React, { useEffect } from 'react';
import { initLoggerAndHttpclient } from './init';
import { Routes } from './routes/Routes';
import { ErrorContainer } from './modules/error/ErrorContainer';

import 'antd/dist/antd.css';
import styles from './App.module.css';

function App() {
  useEffect(() => {
    initLoggerAndHttpclient();
  }, []);

  return (
    <div className={styles.App}>
      <ErrorContainer />
      <Routes />
    </div>
  );
}

export default App;
