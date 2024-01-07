import { Button, Typography } from '@mui/material';
import BusinessPartnerCards from '../LandingPage/Cards/BusinessPartnerCards';

import './styles/OtherEventsStyles.css';

export default function OtherEvents(): JSX.Element {
  return (
    <div className="otherEventsContainer">
      <div className="otherEvents-firstHalf">
        <div className="firstHalf-contentContainer">
          <h1>Explore Local Gems in your Cup of Tea</h1>
          <p>
            Embark on a journey to find local tea havens near you. CommuniTEA
            opens the door to a diverse tapestry of nearby businesses, each with
            its own unique blend of flavors and community charm. Whether
            you&apos;re a seasoned tea aficionado or a curious newcomer, let the
            essence of local tea culture unfold right in your neighborhood.
          </p>
          <Button
            variant="contained"
            sx={{
              backgroundColor: '#8B4513',
              borderRadius: 50,
              padding: 2.5,
              '&:hover': {
                backgroundColor: '#D2B48C',
                color: 'black',
              },
            }}
          >
            <Typography
              variant="body2"
              sx={{ fontFamily: 'Montserrat', fontSize: '1.5em' }}
            >
              Embark. Unite. Elevate your CommuniTEA.
            </Typography>
          </Button>
        </div>
      </div>
      <div className="secondHalf-contentContainer">
        <div>
          <BusinessPartnerCards />
        </div>
      </div>
    </div>
  );
}
