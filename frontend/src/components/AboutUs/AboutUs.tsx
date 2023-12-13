import Footer from '../LandingPage/Footer';
import NavBar from '../LandingPage/Navbar';
import HeroBanner from './HeroBanner';
import TeamMemberCarousel from './TeamMemberCarousel';

export default function AboutUs(): JSX.Element {
  return (
    <div style={{ overflowX: 'hidden' }}>
      <NavBar />
      <HeroBanner />
      <TeamMemberCarousel />
      <Footer />
    </div>
  );
}
