import { useEffect, useRef } from 'react';
import PropTypes from 'prop-types';

import './styles/CommuniTeaPage.css';

interface Location {
  lat: number
  lng: number
  name: string
}

interface GoogleMapsMultiplePinsProps {
  apiKey: string
  locations: Location[]
}

interface MarkerAndInfoWindow {
  marker: google.maps.Marker
  infowindow: google.maps.InfoWindow
}

function GoogleMapsMultiplePins({
  apiKey,
  locations,
}: GoogleMapsMultiplePinsProps): JSX.Element {
  const mapRef = useRef<HTMLDivElement>(null);
  const markers = useRef<MarkerAndInfoWindow[]>([]); // Ref to store markers

  useEffect(() => {
    // Load the Google Maps script
    const script = document.createElement('script');
    script.src = `https://maps.googleapis.com/maps/api/js?key=${apiKey}`;
    script.async = true;
    script.onload = () => {
      if (mapRef.current !== null) {
        // Initialize the map
        const map = new window.google.maps.Map(mapRef.current, {
          zoom: 10,
          center: { lat: locations[0].lat, lng: locations[0].lng },
        });

        // Add markers and InfoWindows to the map
        locations.forEach((location) => {
          const marker = new window.google.maps.Marker({
            position: { lat: location.lat, lng: location.lng },
            map,
            title: location.name, // This sets the tooltip text
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

          markers.current.push({ marker, infowindow }); // Store the marker and infowindow
        });
      }
    };

    document.body.appendChild(script);

    // Cleanup
    return () => {
      document.body.removeChild(script);
      // Clean up markers and InfoWindows
      markers.current.forEach((item) => {
        item.infowindow.close();
        item.marker.setMap(null);
      });
      markers.current = [];
    };
  }, [apiKey, locations]);

  return <div className="multiplePinsMap" ref={mapRef} />;
}

export default GoogleMapsMultiplePins;

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
