import type { ReactElement } from 'react';
import './App.css';
import TestButton from './components/TestButton';

function App(): ReactElement {
  return (
    <div>
      Hello CommuniTEAM Members! Welcome to TypeScript :)
      <br />
      <TestButton />
    </div>
  );
}

export default App;
