import './styles/TeamMemberDetailsStyles.css';

function HeroBanner(): JSX.Element {
  return (
    <div className="heroBannerContainer">
      <div className="heroBannerOverlay" />
      <div className="heroBannerTextContainer">
        <h1>MEET OUR TEAM</h1>
      </div>
    </div>
  );
}

export default HeroBanner;
