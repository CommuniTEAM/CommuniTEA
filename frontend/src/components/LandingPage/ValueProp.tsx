import './styles/ValuePropStyles.css';
import PouringTea from '../../assets/PouringTea.png';
import PeopleDrinkingTea from '../../assets/PeopleDrinkingTea.jpg';
import FriendsDrinkingTea from '../../assets/FriendsDrinkingTea.jpg';

export default function ValueProp(): JSX.Element {
  return (
    <div className="valuePropContainer">
      <div className="imageSection">
        <div className="imageContainerFirst">
          <img src={PouringTea} alt="Pouring Tea" className="imageStyle" />
        </div>
        <div className="imageContainerSecond">
          <img
            src={PeopleDrinkingTea}
            alt="People Drinking Tea"
            className="imageStyle"
          />
        </div>
        <div className="imageContainerThird">
          <img
            src={FriendsDrinkingTea}
            alt="Friends Drinking Tea"
            className="imageStyle"
          />
        </div>
      </div>
      <div className="textSection">
        <div className="textContainer">
          <h1 className="headingStyle">
            CommuniTEA is your gateway to a vibrant world of tea exploration.
          </h1>
          <p className="paragraphStyle">
            Whether you are a seasoned connoisseur or just starting your tea
            journey, our platform brings together the finest tea varieties and
            the best local businesses, creating a tea community like no other.
          </p>
        </div>
      </div>
    </div>
  );
}
