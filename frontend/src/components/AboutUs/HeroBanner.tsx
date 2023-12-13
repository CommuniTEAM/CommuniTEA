import React from 'react';
import Group from '../../assets/Group.png';
import '../../App.css';

function HeroBanner(): JSX.Element {
  return (
    <div
      style={{
        position: 'relative',
        width: '100vw',
        height: '85vh',
        textAlign: 'center',
        color: 'white',
      }}
    >
      <img
        src={Group}
        alt="Group"
        style={{ width: '100vw', height: '85vh', objectFit: 'cover' }}
      />
      <div
        style={{
          position: 'absolute',
          top: 0,
          left: 0,
          width: '100%',
          height: '85vh',
          backgroundColor: 'rgba(0, 0, 0, 0.7)',
        }}
      />
      <div
        style={{
          position: 'absolute',
          top: '50%',
          left: '50%',
          transform: 'translate(-50%, -50%)',
          padding: '20px',
          boxSizing: 'border-box',
        }}
      >
        <h1
          style={{
            fontFamily: 'Montserrat',
            fontSize: '5vw',
            fontWeight: 300,
          }}
        >
          MEET OUR TEAM
        </h1>
      </div>
    </div>
  );
}

export default HeroBanner;
