import type { ReactElement } from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import LandingPage from './components/LandingPage/LandingPage';
import CommuniTeaPage from './components/CommuniTEA/CommuniTeaPage';
import EventDetailPage from './components/IndividualEvent/EventDeatailPage';
import AboutUs from './components/AboutUs/AboutUs';

function App(): ReactElement {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<LandingPage />} />
        <Route path="/communitea" element={<CommuniTeaPage />} />
        <Route path="/event/:eventId" element={<EventDetailPage />} />
        <Route path="/about" element={<AboutUs />} />
      </Routes>
    </Router>
  );
}

export default App;
