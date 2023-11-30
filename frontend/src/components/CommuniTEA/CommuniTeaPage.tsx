import Footer from '../LandingPage/Footer';
import NavBar from '../LandingPage/Navbar';
import Filters from './Filters';
import LocationSearch from './LocationSearch';
import LocalBusinesses from './LocalBusinesses';

export default function CommuniTeaPage(): JSX.Element {
  // TODO: Add Map functionality
  return (
    <div>
      <NavBar />
      <LocationSearch />
      <Filters />
      <LocalBusinesses />
      <Footer />
    </div>
  );
}
