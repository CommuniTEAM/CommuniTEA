import { Card, CardContent, CardMedia, Typography, CardActionArea } from '@mui/material'
import { useTheme } from '@mui/material/styles'
import useMediaQuery from '@mui/material/useMediaQuery'
import ChamomileTea from '../../../assets/ChamomileTea.jpg'
import DarjeelingTea from '../../../assets/DarjeelingTea.jpg'
import OolongTea from '../../../assets/OolongTea.png'
import GreenTea from '../../../assets/GreenTea.jpg'

export default function FeaturedTeaCards (): JSX.Element {
  const theme = useTheme()
  const largeScreen = useMediaQuery(theme.breakpoints.up('lg'))
  const mediumScreen = useMediaQuery(theme.breakpoints.between('md', 'lg'))
  const smallScreen = useMediaQuery(theme.breakpoints.down('md'))

  const getImageHeight = (): string => {
    if (largeScreen) return '10vw'
    if (mediumScreen) return '10vw'
    if (smallScreen) return '10vw'
    return '60vw'
  }

  const getCardHeight = (): string => {
    if (largeScreen) return '18vw'
    if (mediumScreen) return '20vw'
    if (smallScreen) return '20vw'
    return '60vw'
  }

  // This is test data. Replace when API is ready.
  const teaData = [
    {
      name: 'Chamomile Tea',
      description:
        'True to the origins of its name, Chamomile has gentle notes of apple, and there is a mellow, honey-like sweetness in the cup. It has a silky mouthfeel and yet remains a clean, delicately floral herbal tea, and even from the very first sip it feels wonderfully soothing.',
      image: ChamomileTea,
      title: 'Chamomile Tea'
    },
    {
      name: 'Darjeeling',
      description:
        "Darjeeling tea is a type of black tea produced in India. Darjeeling tea has a fruity aroma and a golden or bronze color, depending on the way it's brewed. Tea experts say it has notes (flavors) of citrus fruit, flowers, and even a vegetal quality. Darjeeling tastes sweeter and less bitter than other forms of black tea.",
      image: DarjeelingTea,
      title: 'Darjeeling Tea'
    },
    {
      name: 'Oolong Tea',
      description:
        'Oolong tea is made from the Camellia sinensis plant. Its dried leaves and leaf buds are used to make several different teas, including black and green teas. Oolong tea is fermented for longer than green tea, but less than black tea. It contains caffeine which affects thinking and alertness.',
      image: OolongTea,
      title: 'Oolong Tea'
    },
    {
      name: 'Green Tea',
      description:
        'Green tea, also known as unoxidized tea, is made solely from the leaves of the camellia sinensis plant. The leaves are plucked, slightly withered, then immediately cooked to preserve the green quality and prevent oxidization.',
      image: GreenTea,
      title: 'Green Tea'
    }
  ]

  return (
    <>
      {teaData.map((tea) => (
        <Card key={tea.name} sx={{ width: '20vw', height: getCardHeight(), borderRadius: 10 }}>
          <CardActionArea sx={{ height: '100%' }}>
            <CardMedia sx={{ height: getImageHeight() }} image={tea.image} title={tea.title} />
            <CardContent sx={{ height: '40%' }}>
              <Typography
                gutterBottom
                variant="body1"
                component="div"
                sx={{ fontFamily: 'Montserrat', color: '#29C6CF', fontWeight: 300 }}
              >
                Featured
              </Typography>
              <Typography
                gutterBottom
                variant="h5"
                component="div"
                sx={{ fontFamily: 'Montserrat', fontWeight: 700 }}
              >
                {tea.name}
              </Typography>
              <Typography variant="body2" color="text.secondary" sx={{ fontFamily: 'Montserrat' }}>
                {tea.description}
              </Typography>
            </CardContent>
          </CardActionArea>
        </Card>
      ))}
    </>
  )
}
