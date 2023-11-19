import '../../App.css'
import useMediaQuery from '@mui/material/useMediaQuery'
import OolongTeaCard from './Cards/OolongTeaCard'
import GreenTeaCard from './Cards/GreenTeaCard'
import ChamomileTeaCard from './Cards/ChamomileTeaCard'
import DarjeelingTeaCard from './Cards/DarjeelingTeaCard'

export default function FeaturedTeas (): JSX.Element {
  const largeScreen = useMediaQuery('(min-width:1500px)')
  const mediumScreen = useMediaQuery('(min-width:1000px)')
  const smallScreen = useMediaQuery('(min-width:400px)')

  const calculateHeight = (): string => {
    if (largeScreen) return '60vh'
    if (mediumScreen) return '50vh'
    if (smallScreen) return '30vw'
    return 'auto'
  }

  return (
    <div style={{ backgroundColor: '#D2B48C', height: calculateHeight() }}>
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
