import useResponsiveHeight from '../../Hooks/useResponsiveHeight'
import One from '../../assets/One.svg'
import Two from '../../assets/Two.svg'
import Three from '../../assets/Three.svg'
import '../../App.css'

export default function HowItWorks (): JSX.Element {
  const responsiveHeight = useResponsiveHeight() // custom hook

  return (
    <div style={{ height: responsiveHeight }}>
      <h1
        style={{
          fontFamily: 'Roboto Slab',
          fontSize: '2vw',
          fontWeight: 300,
          color: '#416543',
          textAlign: 'center'
        }}
      >
        How does it work?
      </h1>
      <div style={{ display: 'flex', justifyContent: 'space-evenly' }}>
        <div
          style={{
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center',
            flexGrow: 1,
            maxWidth: '30%'
          }}
        >
          <img src={One} alt="Step one" style={{ width: '25vh' }} />
          <h1
            style={{
              fontFamily: 'Roboto Slab',
              fontSize: '1.5vw',
              fontWeight: 300,
              color: '#416543',
              textAlign: 'center'
            }}
          >
            Discover
          </h1>
          <p
            style={{
              fontFamily: 'Roboto',
              fontSize: '1vw',
              fontWeight: 300,
              color: '#416543'
            }}
          >
            Find your favorite teas from curated local businesses
          </p>
        </div>
        <div
          style={{
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center',
            flexGrow: 1,
            maxWidth: '30%'
          }}
        >
          <img src={Two} alt="Step two" style={{ width: '25vh' }} />
          <h1
            style={{
              fontFamily: 'Roboto Slab',
              fontSize: '1.5vw',
              fontWeight: 300,
              color: '#416543',
              textAlign: 'center'
            }}
          >
            Connect
          </h1>
          <p
            style={{
              fontFamily: 'Roboto',
              fontSize: '1vw',
              fontWeight: 300,
              color: '#416543'
            }}
          >
            Explore tea-related events and connect enthusiasts.
          </p>
        </div>
        <div
          style={{
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center',
            flexGrow: 1,
            maxWidth: '30%'
          }}
        >
          <img src={Three} alt="Step three" style={{ width: '25vh' }} />
          <h1
            style={{
              fontFamily: 'Roboto Slab',
              fontSize: '1.5vw',
              fontWeight: 300,
              color: '#416543',
              textAlign: 'center'
            }}
          >
            Experience
          </h1>
          <p
            style={{
              fontFamily: 'Roboto',
              fontSize: '1vw',
              fontWeight: 300,
              color: '#416543'
            }}
          >
            Transform your tea-ppreciation into a collective experience with CommuniTEA
          </p>
        </div>
      </div>
    </div>
  )
}
