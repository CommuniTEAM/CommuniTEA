import { Grid } from '@mui/material';
import One from '../../assets/One.svg';
import Two from '../../assets/Two.svg';
import Three from '../../assets/Three.svg';
import './styles/HowItWorksStyles.css';

export default function HowItWorks(): JSX.Element {
  return (
    <div className="container">
      <h1 className="howItWorksTitle">How does it work?</h1>
      <Grid container spacing={2} justifyContent="center">
        {/* Grid item for Step One */}
        <Grid item xs={12} sm={6} md={4}>
          <div className="howItWorksGridItem">
            <img src={One} alt="Step one" className="howItWorksImage" />
            <h1 className="howItWorksStepTitle">Discover</h1>
            <p className="howItWorksDescription">
              Find your favorite teas from curated local businesses
            </p>
          </div>
        </Grid>

        {/* Grid item for Step Two */}
        <Grid item xs={12} sm={6} md={4}>
          <div className="howItWorksGridItem">
            <img src={Two} alt="Step two" className="howItWorksImage" />
            <h1 className="howItWorksStepTitle">Connect</h1>
            <p className="howItWorksDescription">
              Explore tea-related events and connect enthusiasts.
            </p>
          </div>
        </Grid>

        {/* Grid item for Step Three */}
        <Grid item xs={12} sm={6} md={4}>
          <div className="howItWorksGridItem">
            <img src={Three} alt="Step three" className="howItWorksImage" />
            <h1 className="howItWorksStepTitle">Experience</h1>
            <p className="howItWorksDescription">
              Transform your tea-ppreciation into a collective experience with
              CommuniTEA.
            </p>
          </div>
        </Grid>
      </Grid>
    </div>
  );
}
