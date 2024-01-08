import '../../App.css';
import './styles/HeroBannerStyles.css';

function HeroBanner(): JSX.Element {
  return (
    <div className="landingHeroBannerContainer">
      <div className="heroBannerOverlay" />
      <div className="heroBannerTextContainer">
        <h1 className="heroBannerHeading">
          Discover the World of Tea with CommuniTEA!
        </h1>
        <p className="heroBannerParagraph">
          Connecting Tea Enthusiasts with Local Businesses for a Shared Tea
          Experience
        </p>
      </div>
    </div>
  );
}

export default HeroBanner;
