import React from 'react';
import { Link } from 'react-router-dom';
import { Typography, IconButton } from '@mui/material';
import GitHubIcon from '../../assets/GitHubIcon.png';

export default function Footer(): JSX.Element {
  const footerMenuItems = [
    'About Us',
    'Contact Us',
    'Privacy Policy',
    'Terms of Service',
    'FAQ',
  ];

  interface MenuItemPaths {
    'About Us': string
    'Contact Us': string
    // Add other menu items
    [key: string]: string | undefined
  }

  const menuItemPaths: MenuItemPaths = {
    'About Us': '/about',
    'Contact Us': '#',
    'Privacy Policy': '#',
    'Terms of Service': '#',
    FAQ: '#',
  };

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
          maxWidth: '25%',
        }}
      >
        <Typography
          style={{ color: '#fff', fontFamily: 'Montserrat', fontSize: '.75vw' }}
        >
          © 2023 CommuniTeam. All Rights Reserved.
        </Typography>
      </div>
      <div
        style={{
          display: 'flex',
          flexGrow: 2,
          justifyContent: 'center',
          alignItems: 'center',
          maxWidth: '50%',
        }}
      >
        {footerMenuItems.map((item, index) => (
          <React.Fragment key={item}>
            {menuItemPaths[item] !== undefined ? (
              <Link
                to={menuItemPaths[item] ?? '/'}
                style={{ textDecoration: 'none', color: '#fff' }}
              >
                <Typography
                  style={{
                    fontFamily: 'Montserrat',
                    textTransform: 'none',
                    fontSize: '.75vw',
                  }}
                >
                  {item}
                </Typography>
              </Link>
            ) : (
              <Typography
                style={{
                  color: '#fff',
                  fontFamily: 'Montserrat',
                  textTransform: 'none',
                  fontSize: '.75vw',
                }}
              >
                {item}
              </Typography>
            )}
            {index < footerMenuItems.length - 1 && (
              <Typography
                style={{
                  color: '#fff',
                  fontFamily: 'Montserrat',
                  margin: '0 10px',
                }}
              >
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
          maxWidth: '25%',
        }}
      >
        <IconButton
          component="a"
          href="https://github.com/CommuniTEAM/CommuniTEA"
          target="_blank"
          rel="noopener noreferrer"
        >
          <img
            src={GitHubIcon}
            alt="GitHub"
            style={{ width: '1.5vw', color: 'white' }}
          />
        </IconButton>
      </div>
    </div>
  );
}
