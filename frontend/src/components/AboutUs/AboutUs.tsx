import NavBar from '../LandingPage/Navbar';
import HeroBanner from './HeroBanner';
import TeamMemberCarousel from './TeamMemberCarousel';

export default function AboutUs(): JSX.Element {
  return (
    <>
      <NavBar />
      <HeroBanner />
      <TeamMemberCarousel />
      <h1>ABOUT US PAGE</h1>
      <p>Add stuff here</p>
    </>
  );
}
