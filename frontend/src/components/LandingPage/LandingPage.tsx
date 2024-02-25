import '../../App.css';
import BusinessShowcase from './BusinessShowcase';
import FeaturedTeas from './FeaturedTeas';
import Footer from './Footer';
import HeroBanner from './HeroBanner';
import HowItWorks from './HowItWorks';
import NavBar from './Navbar';
import ValueProp from './ValueProp';

export default function LandingPage(): JSX.Element {
  return (
    <div>
      <NavBar />
      <HeroBanner />
      <ValueProp />
      <FeaturedTeas />
      <HowItWorks />
      <BusinessShowcase />
      <Footer />
    </div>
  );
}
