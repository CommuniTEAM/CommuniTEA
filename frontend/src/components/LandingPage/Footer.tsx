import { IconButton, Typography } from '@mui/material';
import React from 'react';
import { Link } from 'react-router-dom';
import GitHubIcon from '../../assets/GitHubIcon.png';

import './styles/FooterStyles.css';

export default function Footer(): JSX.Element {
  const footerMenuItems = [
    'About Us',
    'Contact Us',
    'Privacy Policy',
    'Terms of Service',
    'FAQ',
  ];

  interface MenuItemPaths {
    'About Us': string;
    'Contact Us': string;
    // Add other menu items
    [key: string]: string | undefined;
  }

  const menuItemPaths: MenuItemPaths = {
    'About Us': '/about',
    'Contact Us': '/contact-us',
    'Privacy Policy': '#',
    'Terms of Service': '#',
    FAQ: '#',
  };

  return (
    <div className="footerContainer">
      <div className="firstThird">
        <Typography
          style={{ color: '#fff', fontFamily: 'Montserrat', fontSize: '1em' }}
        >
          © 2023 CommuniTeam. All Rights Reserved.
        </Typography>
      </div>
      <div className="secondThird">
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
                    fontSize: '1em',
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
                  fontSize: '1em',
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
      <div className="thirdThird">
        <IconButton
          component="a"
          href="https://github.com/CommuniTEAM/CommuniTEA"
          target="_blank"
          rel="noopener noreferrer"
        >
          <img
            src={GitHubIcon}
            alt="GitHub"
            style={{ width: '1em', color: 'white' }}
          />
        </IconButton>
      </div>
    </div>
  );
}
