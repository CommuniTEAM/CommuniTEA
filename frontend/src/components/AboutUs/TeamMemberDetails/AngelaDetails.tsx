import { Divider, Typography } from '@mui/material';
import AngelaCropped from '../../../assets/AngelaCropped.jpg';
import Github from '../../../assets/GitHubBlack.png';
import LinkedIn from '../../../assets/LinkedIn.png';

import '../styles/TeamMemberDetailsStyles.css';

export default function AngelaDetails(): JSX.Element {
  return (
    <div className="aboutUsBodyContainer">
      {/* Image */}
      <div className="aboutUsfirstHalf">
        <img src={AngelaCropped} alt="Cory" />
      </div>

      {/* Full Name */}
      <div className="aboutUsSecondHalf">
        <div>
          <Typography variant="h2" sx={{ fontFamily: 'Montserrat' }}>
            Angela Fisher
          </Typography>
        </div>

        {/* Divider */}
        <Divider sx={{ bgcolor: 'black' }} />

        {/* Title / Position */}
        <div className="titlePositionContainer">
          <div>
            <Typography
              variant="h5"
              sx={{ fontFamily: 'Montserrat', fontStyle: 'italic' }}
            >
              DevOps & Back-end Engineer
            </Typography>
          </div>

          {/* Social Links */}
          <div style={{ display: 'flex' }}>
            <a
              href="https://www.linkedin.com/in/angelajfisher/"
              target="_blank"
              rel="noreferrer noopener"
            >
              <img src={LinkedIn} alt="LinkedIn" />
            </a>
            <a
              href="https://github.com/angelajfisher"
              target="_blank"
              rel="noreferrer noopener"
            >
              <img src={Github} alt="Github" />
            </a>
          </div>
        </div>

        {/* Team Member Description */}
        <div className="aboutUsDescription">
          <Typography
            variant="body1"
            sx={{ fontFamily: 'Montserrat', fontSize: '1.25em' }}
          >
            Add description here.
          </Typography>
        </div>
      </div>
    </div>
  );
}
