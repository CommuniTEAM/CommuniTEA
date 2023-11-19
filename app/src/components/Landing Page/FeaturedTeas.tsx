import '../../App.css'
import useResponsiveHeight from '../../Hooks/useResponsiveHeight'
import OolongTeaCard from './Cards/OolongTeaCard'
import GreenTeaCard from './Cards/GreenTeaCard'
import ChamomileTeaCard from './Cards/ChamomileTeaCard'
import DarjeelingTeaCard from './Cards/DarjeelingTeaCard'

export default function FeaturedTeas (): JSX.Element {
  const responsiveHeight = useResponsiveHeight() // custom hook

  return (
    <div style={{ backgroundColor: '#D2B48C', height: responsiveHeight }}>
      <div>
        <h1
          style={{
            fontFamily: 'Roboto Slab',
            fontSize: '2vw',
            fontWeight: 300,
            textAlign: 'center',
            paddingTop: '5vh'
          }}
        >
          Explore Our Featured Teas
        </h1>
      </div>
      <div
        style={{
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'space-evenly'
        }}
      >
        <OolongTeaCard />
        <GreenTeaCard />
        <ChamomileTeaCard />
        <DarjeelingTeaCard />
      </div>
    </div>
  )
}
