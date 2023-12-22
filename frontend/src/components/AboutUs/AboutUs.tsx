import { useState } from 'react';
import Footer from '../LandingPage/Footer';
import NavBar from '../LandingPage/Navbar';
import HeroBanner from './HeroBanner';
import TeamMemberCarousel from './TeamMemberCarousel';
import CoryDetails from './TeamMemberDetails/CoryDetails';
import AngelaDetails from './TeamMemberDetails/AngelaDetails';
import BrianDetails from './TeamMemberDetails/BrianDetails';
// import AmandaDetails from './TeamMemberDetails/AmandaDetails';
// import HectorDetails from './TeamMemberDetails/HectorDetails';
// import AlexDetails from './TeamMemberDetails/AlexDetails';

export default function AboutUs(): JSX.Element {
  const [selectedTeamMember, setSelectedTeamMember] = useState<null | string>(
    null,
  );

  return (
    <div style={{ overflowX: 'hidden' }}>
      <NavBar />
      <HeroBanner />
      {selectedTeamMember === 'Cory' && <CoryDetails />}
      {selectedTeamMember === 'Angela' && <AngelaDetails />}
      {selectedTeamMember === 'Brian' && <BrianDetails />}
      {/* {selectedTeamMember === 'Amanda' && <AmandaDetails />}
      {selectedTeamMember === 'Hector' && <HectorDetails />}
      {selectedTeamMember === 'Alex' && <AlexDetails />} */}
      <TeamMemberCarousel onSelectTeamMember={setSelectedTeamMember} />
      <Footer />
    </div>
  );
}
