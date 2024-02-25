import type { ReactElement } from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import AboutUs from './components/AboutUs/AboutUs';
import CommuniTeaPage from './components/CommuniTEA/CommuniTeaPage';
import ContactUsPage from './components/ContactUs/ContactUs';
import EventDetailPage from './components/IndividualEvent/EventDetailPage';
import JoditTest from './components/Jodit/JoditTest';
import LandingPage from './components/LandingPage/LandingPage';
import NotFound from './components/NotFound';
import TestFetch from './components/WikiTEAdia/TestFetch';

function App(): ReactElement {
  return (
    <Router>
      <Routes>
        <Route path="*" element={<NotFound />} />
        <Route path="/" element={<LandingPage />} />
        <Route path="/communitea" element={<CommuniTeaPage />} />
        <Route path="/event/:eventId" element={<EventDetailPage />} />
        <Route path="/about" element={<AboutUs />} />
        <Route path="/contact-us" element={<ContactUsPage />} />
        <Route path="/test" element={<TestFetch />} />
        <Route path="/jodit" element={<JoditTest />} />
      </Routes>
    </Router>
  );
}

export default App;
