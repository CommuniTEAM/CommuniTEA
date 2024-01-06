import type { ReactElement } from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import LandingPage from './components/LandingPage/LandingPage';
import CommuniTeaPage from './components/CommuniTEA/CommuniTeaPage';
import EventDetailPage from './components/IndividualEvent/EventDetailPage';
import AboutUs from './components/AboutUs/AboutUs';
import TestFetch from './components/WikiTEAdia/TestFetch';
import NotFound from './components/NotFound';

function App(): ReactElement {
  return (
    <Router>
      <Routes>
        <Route path="*" element={<NotFound />} status={404} />
        <Route path="/" element={<LandingPage />} />
        <Route path="/communitea" element={<CommuniTeaPage />} />
        <Route path="/event/:eventId" element={<EventDetailPage />} />
        <Route path="/about" element={<AboutUs />} />
        <Route path="/test" element={<TestFetch />} />
      </Routes>
    </Router>
  );
}

export default App;
