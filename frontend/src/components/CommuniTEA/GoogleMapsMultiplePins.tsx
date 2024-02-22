import PropTypes from 'prop-types';
import { useEffect, useRef } from 'react';

import './styles/CommuniTeaPage.css';

interface Location {
  lat: number;
  lng: number;
  name: string;
}

interface GoogleMapsMultiplePinsProps {
  apiKey: string;
  locations: Location[];
}

interface MarkerAndInfoWindow {
  marker: google.maps.Marker;
  infowindow: google.maps.InfoWindow;
}

function GoogleMapsMultiplePins({
  apiKey,
  locations,
}: GoogleMapsMultiplePinsProps): JSX.Element {
  const mapRef = useRef<HTMLDivElement>(null);
  const markers = useRef<MarkerAndInfoWindow[]>([]); // Ref to store markers

  useEffect(() => {
    const loadGoogleMapScript = () => {
      const script = document.createElement('script');
      script.src = `https://maps.googleapis.com/maps/api/js?key=${apiKey}`;
      script.async = true;
      script.onload = initializeMap;
      document.body.appendChild(script);
      return () => {
        document.body.removeChild(script);
      };
    };

    const initializeMap = () => {
      if (mapRef.current !== null && locations.length > 0) {
        const map = new window.google.maps.Map(mapRef.current, {
          zoom: 10,
          center: { lat: locations[0].lat, lng: locations[0].lng },
        });
        addMarkers(map);
      }
    };

    const addMarkers = (map: google.maps.Map) => {
      for (const location of locations) {
        const marker = new window.google.maps.Marker({
          position: { lat: location.lat, lng: location.lng },
          map,
          title: location.name,
        });

        const infowindow = new window.google.maps.InfoWindow({
          content: `<div><strong>${location.name}</strong></div>`,
        });

        marker.addListener('click', () => {
          infowindow.open({
            anchor: marker,
            map,
            shouldFocus: false,
          });
        });

        markers.current.push({ marker, infowindow });
      }
    };

    return loadGoogleMapScript();
  }, [apiKey, locations, locations[0]?.lat, locations[0]?.lng]); // Include locations[0]?.lng here

  useEffect(() => {
    return () => {
      for (const { marker, infowindow } of markers.current) {
        infowindow.close();
        marker.setMap(null);
      }
    };
  }, []); // Adding locations as a dependency here

  return <div className="multiplePinsMap" ref={mapRef} />;
}

GoogleMapsMultiplePins.propTypes = {
  apiKey: PropTypes.string.isRequired,
  locations: PropTypes.arrayOf(
    PropTypes.shape({
      lat: PropTypes.number.isRequired,
      lng: PropTypes.number.isRequired,
      name: PropTypes.string.isRequired,
    }),
  ).isRequired,
};

export default GoogleMapsMultiplePins;
