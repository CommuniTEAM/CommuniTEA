import HeroBannerImage from '../../assets/HeroBannerImage.jpg';
import '../../App.css';
import './styles/HeroBannerStyles.css';

function HeroBanner(): JSX.Element {
  return (
    <div className="heroBannerContainer">
      <div className="heroBannerOverlay" />
      <img
        src={HeroBannerImage}
        alt="Assorted teas and herbs"
        className="heroBannerImage"
      />
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
