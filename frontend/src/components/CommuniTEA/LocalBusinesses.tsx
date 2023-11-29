import { Button } from '@mui/material';
import useResponsiveHeight from '../../Hooks/useResponsiveHeight';
import BusinessPartnerCards from '../Landing Page/Cards/BusinessPartnerCards';

export default function BusinessShowcase(): JSX.Element {
  const responsiveHeight = useResponsiveHeight(); // custom hook
  return (
    <>
      <div>
        <Button
          variant="contained"
          sx={{
            backgroundColor: '#8B4513',
            fontFamily: 'Montserrat',
            borderRadius: 10,
            fontSize: '1vw',
            padding: 2.5,
            transform: 'translate(7vw, 2vw)',
            '&:hover': {
              backgroundColor: '#D2B48C',
              color: 'black',
            },
          }}
        >
          Embark. Unite. Elevate your CommuniTEA.
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
                fontSize: '2vw',
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
