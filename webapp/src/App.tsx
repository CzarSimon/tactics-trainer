import React, { useEffect } from 'react';
import { initLoggerAndHttpclient } from './init';
import { Routes } from './routes/Routes';

import 'antd/dist/antd.css';

function App() {
  useEffect(() => {
    initLoggerAndHttpclient();
  }, []);

  return (
    <div className="App">
      <Routes />
    </div>
  );
}

export default App;
