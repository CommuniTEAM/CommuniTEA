import '../../App.css';
import FeaturedTeaCards from './Cards/FeaturedTeaCards';

export default function FeaturedTeas(): JSX.Element {
  return (
    <div style={{ backgroundColor: '#D2B48C' }}>
      <div>
        <h1
          style={{
            fontFamily: 'Roboto Slab',
            fontSize: '2vw',
            fontWeight: 300,
            textAlign: 'center',
            paddingTop: '5vh',
          }}
        >
          Explore Our Featured Teas
        </h1>
      </div>
      <div
        style={{
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'space-evenly',
        }}
      >
        <FeaturedTeaCards />
      </div>
    </div>
  );
}
