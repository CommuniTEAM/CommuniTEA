import '../../App.css';
import './styles/HeroBannerStyles.css';

function HeroBanner(): JSX.Element {
  return (
    <div className="landingHeroBannerContainer">
      <div className="heroBannerOverlay" />
      <div className="heroBannerTextContainer">
        <h1 className="heroBannerHeading">Contact Us</h1>
      </div>
    </div>
  );
}

export default HeroBanner;
