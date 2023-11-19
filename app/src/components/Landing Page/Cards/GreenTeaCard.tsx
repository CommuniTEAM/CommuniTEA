import { Card, CardContent, CardMedia, Typography, CardActionArea } from '@mui/material'
import useMediaQuery from '@mui/material/useMediaQuery'
import GreenTea from '../../../assets/GreenTea.jpg'

export default function GreenTeaCard (): JSX.Element {
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
        <CardMedia sx={{ height: '60%' }} image={GreenTea} title="Oolong Tea" />
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
            Green Tea
          </Typography>
          <Typography
            variant="body2"
            color="text.secondary"
            sx={{ fontFamily: 'Montserrat', fontWeight: 300 }}
          >
            Green tea, also known as unoxidized tea, is made solely from the leaves of the camellia
            sinensis plant. The leaves are plucked, slightly withered, then immediately cooked to
            preserve the green quality and prevent oxidization.
          </Typography>
        </CardContent>
      </CardActionArea>
    </Card>
  )
}
