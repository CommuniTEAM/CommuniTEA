import NavBar from './Navbar'
import HeroBanner from './HeroBanner'
import ValueProp from './ValueProp'
import FeaturedTeas from './FeaturedTeas'
import '../../App.css'
import HowItWorks from './HowItWorks'
import BusinessShowcase from './BusinessShowcase'
import Footer from './Footer'

export default function LandingPage (): JSX.Element {
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
  )
}
