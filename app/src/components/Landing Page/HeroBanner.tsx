import React from 'react';
import HeroBannerImage from '../../assets/HeroBannerImage.jpg';
import '../../App.css';

function HeroBanner(): JSX.Element {
  return (
    <div
      style={{
        position: 'relative',
        width: '100vw',
        textAlign: 'center',
        color: 'white',
      }}
    >
      <img
        src={HeroBannerImage}
        alt="Assorted teas and herbs"
        style={{ width: '100vw', height: 'auto', display: 'block' }}
      />
      <div
        style={{
          position: 'absolute',
          top: '25vw',
          left: '20vw',
          transform: 'translate(-50%, -50%)',
          width: '40vw',
          padding: '20px',
          boxSizing: 'border-box',
        }}
      >
        <h1
          style={{
            fontFamily: 'Roboto Slab',
            fontSize: '4vw',
            fontWeight: 300,
          }}
        >
          Discover the World of Tea with CommuniTEA!
        </h1>
        <p style={{ fontFamily: 'Roboto', fontSize: '2vw', fontWeight: 300 }}>
          Connecting Tea Enthusiasts with Local Businesses for a Shared Tea
          Experience
        </p>
      </div>
    </div>
  );
}

export default HeroBanner;
