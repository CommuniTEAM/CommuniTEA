import { useState } from 'react';
import Footer from '../LandingPage/Footer';
import NavBar from '../LandingPage/Navbar';
import HeroBanner from './HeroBanner';
import TeamMemberCarousel from './TeamMemberCarousel';
import CoryDetails from './TeamMemberDetails/CoryDetails';
import AngelaDetails from './TeamMemberDetails/AngelaDetails';
import BrianDetails from './TeamMemberDetails/BrianDetails';

import './styles/TeamMemberDetailsStyles.css';

export default function AboutUs(): JSX.Element {
  const [selectedTeamMember, setSelectedTeamMember] = useState<null | string>(
    null,
  );

  return (
    <div className="aboutUsContainer">
      <NavBar />
      <HeroBanner />
      {selectedTeamMember === 'Cory' && <CoryDetails />}
      {selectedTeamMember === 'Angela' && <AngelaDetails />}
      {selectedTeamMember === 'Brian' && <BrianDetails />}
      <TeamMemberCarousel onSelectTeamMember={setSelectedTeamMember} />
      <Footer />
    </div>
  );
}
