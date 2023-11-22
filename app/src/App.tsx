import type { ReactElement } from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import LandingPage from './components/Landing Page/LandingPage';
import CommuniTeaPage from './components/CommuniTEA/CommuniTeaPage';
import EventDetailPage from './components/IndividualEvent/EventDeatailPage';

function App(): ReactElement {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<LandingPage />} />
        <Route path="/communitea" element={<CommuniTeaPage />} />
        <Route path="/event/:eventId" element={<EventDetailPage />} />
      </Routes>
    </Router>
  );
}

export default App;
