import { Button } from '@mui/material';
import { useTheme } from '@mui/material/styles';
import BusinessPartnerCards from './Cards/BusinessPartnerCards';
import './styles/BusinessShowcaseStyles.css';

export default function BusinessShowcase(): JSX.Element {
  const theme = useTheme();

  return (
    <div className="businessShowcaseContainer">
      <div className="businessShowcaseTextContainer">
        <h1 className="businessShowcaseTitle">BUSINESSES</h1>
        <p className="businessShowcaseDescription">
          Showcase your unique tea selections to a wider audience and boost
          community engagement. Join CommuniTEA and let tea lovers discover and
          enjoy your offerings!
        </p>
        <Button
          className="businessShowcaseButton"
          variant="contained"
          sx={{
            backgroundColor: '#8B4513',
            fontFamily: 'Montserrat',
            borderRadius: '50px',
            fontSize: '1.5rem',
            padding: '20px 40px',
            '&:hover': {
              backgroundColor: '#D2B48C',
              color: 'black',
            },
            [theme.breakpoints.down('md')]: {
              fontSize: '1rem',
            },
          }}
        >
          Embark. Unite. Elevate your CommuniTEA.
        </Button>
      </div>
      <div className="businessPartnerCardsContainer">
        <div>
          <BusinessPartnerCards />
        </div>
      </div>
    </div>
  );
}
