import type { ReactElement } from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import LandingPage from './components/Landing_Page/LandingPage';
import CommuniTeaPage from './components/CommuniTEA/CommuniTeaPage';

function App(): ReactElement {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<LandingPage />} />
        <Route path="/communitea" element={<CommuniTeaPage />} />
      </Routes>
    </Router>
  );
}

export default App;
