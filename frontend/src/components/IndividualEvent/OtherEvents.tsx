import { Button, Typography } from '@mui/material';
import useResponsiveHeight from '../../Hooks/useResponsiveHeight';
import BusinessPartnerCards from '../LandingPage/Cards/BusinessPartnerCards';

export default function OtherEvents(): JSX.Element {
  const responsiveHeight = useResponsiveHeight(); // custom hook
  return (
    <>
      <div>
        <Button
          variant="contained"
          sx={{
            backgroundColor: '#8B4513',
            borderRadius: 10,
            width: '30vw',
            padding: 2.5,
            transform: 'translate(7vw, 2vw)',
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
      <div
        style={{
          height: responsiveHeight,
          backgroundColor: 'rgba(85, 107, 47, 0.6)',
          display: 'flex',
        }}
      >
        <div style={{ flexGrow: 1, maxWidth: '50%' }}>
          <div style={{ width: '80%', transform: 'translate(4vw, 5vw)' }}>
            <h1
              style={{
                fontFamily: 'Roboto Slab',
                fontSize: '2em',
                fontWeight: 700,
              }}
            >
              Explore Local Gems in your Cup of Tea
            </h1>
            <p
              style={{
                fontFamily: 'Roboto',
                fontSize: '1vw',
                fontWeight: 300,
              }}
            >
              Embark on a journey to find local tea havens near you. CommuniTEA
              opens the door to a diverse tapestry of nearby businesses, each
              with its own unique blend of flavors and community charm. Whether
              you&apos;re a seasoned tea aficionado or a curious newcomer, let
              the essence of local tea culture unfold right in your
              neighborhood.
            </p>
          </div>
        </div>
        <div
          style={{
            flexGrow: 1,
            maxWidth: '60%',
            display: 'flex',
            flexDirection: 'column',
            justifyContent: 'center',
          }}
        >
          <div>
            <BusinessPartnerCards />
          </div>
        </div>
      </div>
    </>
  );
}
