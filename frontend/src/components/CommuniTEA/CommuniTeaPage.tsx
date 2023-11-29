import Footer from '../Landing Page/Footer';
import NavBar from '../Landing Page/Navbar';
import Filters from './Filters';
import LocationSearch from './LocationSearch';
import LocalBusinesses from './LocalBusinesses';
import GoogleMapsMultiplePins from './GoogleMapsMultiplePins';

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
    <div>
      <NavBar />
      <LocationSearch />
      <div style={{ display: 'flex' }}>
        <Filters />
        <GoogleMapsMultiplePins apiKey={apiKey} locations={locations} />
      </div>
      <LocalBusinesses />
      <Footer />
    </div>
  );
}
