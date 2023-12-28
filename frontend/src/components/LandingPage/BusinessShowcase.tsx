import { Button } from '@mui/material';
import BusinessPartnerCards from './Cards/BusinessPartnerCards';

export default function BusinessShowcase(): JSX.Element {
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
          // height: responsiveHeight,
          backgroundColor: 'rgba(85, 107, 47, 0.6)',
          display: 'flex',
        }}
      >
        <div style={{ flexGrow: 1, maxWidth: '50%' }}>
          <div style={{ width: '80%', transform: 'translate(4vw, 5vw)' }}>
            <h1
              style={{
                fontFamily: 'Roboto Slab',
                fontSize: '3vw',
                fontWeight: 700,
              }}
            >
              BUSINESSES
            </h1>
            <p
              style={{
                fontFamily: 'Roboto',
                fontSize: '1.75vw',
                fontWeight: 300,
              }}
            >
              Showcase your unique tea selections to a wider audience and boost
              community engagement. Join CommuniTEA and let tea lovers discover
              and enjoy your offerings!
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
