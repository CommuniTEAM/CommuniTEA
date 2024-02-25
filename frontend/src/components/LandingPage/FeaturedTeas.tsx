import '../../App.css';
import FeaturedTeaCards from './Cards/FeaturedTeaCards';
import './styles/FeaturedTeasStyles.css';

export default function FeaturedTeas(): JSX.Element {
  return (
    <div className="featuredTeasContainer">
      <div>
        <h1 className="featuredTeasHeading">Explore Our Featured Teas</h1>
      </div>
      <div className="featuredTeasCardsContainer">
        <FeaturedTeaCards />
      </div>
    </div>
  );
}
