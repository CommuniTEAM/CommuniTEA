import Footer from '../LandingPage/Footer';
import NavBar from '../LandingPage/Navbar';
import HeroBanner from './HeroBanner';
import TeamMemberCarousel from './TeamMemberCarousel';
import CoryDetails from './TeamMemberDetails/CoryDetails';

export default function AboutUs(): JSX.Element {
  return (
    <div style={{ overflowX: 'hidden' }}>
      <NavBar />
      <HeroBanner />
      <CoryDetails />
      <TeamMemberCarousel />
      <Footer />
    </div>
  );
}
