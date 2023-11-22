import React from 'react';
import { Button, Typography, IconButton } from '@mui/material';
import GitHubIcon from '../../assets/GitHubIcon.png';

export default function Footer(): JSX.Element {
  const footerMenuItems = ['About Us', 'Contact Us', 'Privacy Policy', 'Terms of Service', 'FAQ'];

  return (
    <div
      style={{
        display: 'flex',
        backgroundColor: '#333',
        height: 64,
        paddingLeft: '1vw',
      }}
    >
      <div
        style={{
          display: 'flex',
          flexDirection: 'column',
          justifyContent: 'center',
          flexGrow: 1,
          maxWidth: '33%',
        }}
      >
        <Typography style={{ color: '#fff', fontFamily: 'Montserrat' }}>
          © 2023 CommuniTeam. All Rights Reserved.
        </Typography>
      </div>
      <div
        style={{
          display: 'flex',
          justifyContent: 'center',
          alignItems: 'center',
          flexGrow: 1,
          maxWidth: '33%',
        }}
      >
        {footerMenuItems.map((item, index) => (
          <React.Fragment key={item}>
            <Button
              style={{
                color: '#fff',
                fontFamily: 'Montserrat',
                textTransform: 'none',
                fontSize: '1rem',
              }}
            >
              {item}
            </Button>
            {index < footerMenuItems.length - 1 && (
              <Typography style={{ color: '#fff', fontFamily: 'Montserrat', margin: '0 10px' }}>
                |
              </Typography>
            )}
          </React.Fragment>
        ))}
      </div>
      <div
        style={{
          display: 'flex',
          flexDirection: 'column',
          justifyContent: 'center',
          alignItems: 'flex-end',
          flexGrow: 1,
          maxWidth: '33%',
        }}
      >
        <IconButton
          component="a"
          href="https://github.com/CommuniTEAM/CommuniTEA"
          target="_blank"
          rel="noopener noreferrer"
        >
          <img src={GitHubIcon} alt="GitHub" style={{ width: '1.5vw', color: 'white' }} />
        </IconButton>
      </div>
    </div>
  );
}
