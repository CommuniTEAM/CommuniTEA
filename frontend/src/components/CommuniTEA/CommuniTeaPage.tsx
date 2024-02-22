import Footer from '../LandingPage/Footer';
import NavBar from '../LandingPage/Navbar';
import Filters from './Filters';
import GoogleMapsMultiplePins from './GoogleMapsMultiplePins';
import LocalBusinesses from './LocalBusinesses';
import LocationSearch from './LocationSearch';

import './styles/CommuniTeaPage.css';

export default function CommuniTeaPage(): JSX.Element {
  const apiKey = import.meta.env.VITE_GOOGLE_MAPS_API_KEY;

  // TO DO: Swap hardcoded locations with API when ready
  const locations = [
    { lat: 47.6205, lng: -122.3493, name: 'Seattle Center' },
    { lat: 47.6097, lng: -122.3331, name: 'Afternoon Tea' },
    { lat: 47.689253, lng: -122.354975, name: "Coyle's Bakeshop" },
    { lat: 47.597797, lng: -122.312026, name: 'East-West Chanoyu Center' },
    { lat: 47.621077, lng: -122.331414, name: 'Cafe Hagen Modern Cafe' },
    // ... more locations
  ];

  return (
    <>
      <NavBar />
      <div className="heroContainer">
        <LocationSearch />
        <div className="communiteaContainer">
          <Filters />
          <GoogleMapsMultiplePins apiKey={apiKey} locations={locations} />
        </div>
      </div>
      <LocalBusinesses />
      <Footer />
    </>
  );
}
