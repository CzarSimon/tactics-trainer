import React, { useEffect } from 'react';
import { initLoggerAndHttpclient } from './init';
import { PuzzleSolver } from './modules/puzzle/PuzzleSolver';

function App() {
  useEffect(() => {
    initLoggerAndHttpclient();
  }, []);


  return (
    <div className="App">
      <h1>Tactics trainer</h1>
      <PuzzleSolver id='9e0795c9-1842-4ef0-9032-79acd2fa222a' />
    </div>
  );
}

export default App;
