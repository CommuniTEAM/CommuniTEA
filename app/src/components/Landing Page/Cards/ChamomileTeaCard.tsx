import { Card, CardContent, CardMedia, Typography, CardActionArea } from '@mui/material'
import useMediaQuery from '@mui/material/useMediaQuery'
import ChamomileTea from '../../../assets/ChamomileTea.jpg'

export default function ChamomileTeaCard (): JSX.Element {
  const largeScreen = useMediaQuery('(min-width:1500px)')
  const mediumScreen = useMediaQuery('(min-width:1000px)')
  const smallScreen = useMediaQuery('(min-width:400px)')

  const calculateHeight = (): string => {
    if (largeScreen) return '40vh'
    if (mediumScreen) return '30vh'
    if (smallScreen) return '20vw'
    return 'auto'
  }

  return (
    <Card sx={{ width: '20vw', height: calculateHeight(), borderRadius: 10 }}>
      <CardActionArea sx={{ height: '100%' }}>
        <CardMedia sx={{ height: '60%' }} image={ChamomileTea} title="Chamomile Tea" />
        <CardContent sx={{ height: '40%' }}>
          <Typography
            gutterBottom
            variant="body1"
            component="div"
            sx={{ fontFamily: 'Montserrat', color: '#29C6CF' }}
          >
            Featured
          </Typography>
          <Typography
            gutterBottom
            variant="h5"
            component="div"
            sx={{ fontFamily: 'Montserrat', fontWeight: 700 }}
          >
            Chamomile Tea
          </Typography>
          <Typography
            variant="body2"
            color="text.secondary"
            sx={{ fontFamily: 'Montserrat', fontWeight: 300 }}
          >
            True to the origins of its name, Chamomile has gentle notes of apple, and there is a
            mellow, honey-like sweetness in the cup. It has a silky mouthfeel and yet remains a
            clean, delicately floral herbal tea, and even from the very first sip it feels
            wonderfully soothing.
          </Typography>
        </CardContent>
      </CardActionArea>
    </Card>
  )
}
