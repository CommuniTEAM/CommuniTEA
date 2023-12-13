import { Divider, Typography } from '@mui/material';
import AngelaCropped from '../../../assets/AngelaCropped.jpg';
import Github from '../../../assets/GitHubBlack.png';
import LinkedIn from '../../../assets/LinkedIn.png';

export default function AngelaDetails(): JSX.Element {
  const responsiveFontSize = 'calc(0.5vw + 0.5em)';

  return (
    <div style={{ display: 'flex' }}>
      {/* Image */}
      <div style={{ flexGrow: 1, maxWidth: '100vw', width: '50%' }}>
        <img src={AngelaCropped} alt="Angela" style={{ maxWidth: '100%' }} />
      </div>

      {/* Full Name */}
      <div
        style={{
          flexGrow: 1,
          paddingTop: '1vh',
          paddingLeft: '2vw',
          paddingRight: '2vw',
          width: '50%',
        }}
      >
        <div>
          <Typography variant="h2" sx={{ fontFamily: 'Montserrat' }}>
            Angela Fisher
          </Typography>
        </div>

        {/* Divider */}
        <Divider sx={{ bgcolor: 'black' }} />

        {/* Title / Position */}
        <div
          style={{
            display: 'flex',
            justifyContent: 'space-between',
            paddingBottom: 25,
          }}
        >
          <div>
            <Typography
              variant="h5"
              sx={{ fontFamily: 'Montserrat', fontStyle: 'italic' }}
            >
              Back-end Engineer
            </Typography>
          </div>

          {/* Social Links */}
          <div>
            <div style={{ display: 'flex', alignItems: 'center' }}>
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
        </div>

        {/* Team Member Description */}
        <div>
          <Typography
            variant="body1"
            sx={{ fontFamily: 'Montserrat', fontSize: responsiveFontSize }}
          >
            Add Details Here
          </Typography>
        </div>
      </div>
    </div>
  );
}
