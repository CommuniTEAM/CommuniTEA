import Footer from '../Landing_Page/Footer';
import NavBar from '../Landing_Page/Navbar';
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
