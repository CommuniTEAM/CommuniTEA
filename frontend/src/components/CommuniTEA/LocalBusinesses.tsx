import { Button } from '@mui/material';
import BusinessPartnerCards from '../LandingPage/Cards/BusinessPartnerCards';

import './styles/CommuniTeaPage.css';

export default function BusinessShowcase(): JSX.Element {
  return (
    <div className="ctaHeroContainer">
      <div className="firstHalf">
        <div>
          <h1 className="ctaHeroTitle">
            Explore Local Gems in your Cup of Tea
          </h1>
          <p className="ctaHeroText">
            Embark on a journey to find local tea havens near you. CommuniTEA
            opens the door to a diverse tapestry of nearby businesses, each with
            its own unique blend of flavors and community charm. Whether
            you&apos;re a seasoned tea aficionado or a curious newcomer, let the
            essence of local tea culture unfold right in your neighborhood.
          </p>

          <Button
            variant="contained"
            sx={{
              marginTop: 5,
              backgroundColor: '#8B4513',
              fontFamily: 'Montserrat',
              borderRadius: 10,
              fontSize: '1.2rem',
              padding: 2.5,
              '&:hover': {
                backgroundColor: '#D2B48C',
                color: 'black',
              },
            }}
          >
            Embark. Unite. Elevate your CommuniTEA.
          </Button>
        </div>
      </div>
      <div className="secondHalf">
        <BusinessPartnerCards />
      </div>
    </div>
  );
}
