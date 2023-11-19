import * as React from 'react'
import { Card, CardContent, CardMedia, Typography, CardActionArea } from '@mui/material'
import useMediaQuery from '@mui/material/useMediaQuery'
import OolongTea from '../../../assets/OolongTea.png'

export default function OolongTeaCard (): JSX.Element {
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
        <CardMedia sx={{ height: '60%' }} image={OolongTea} title="Oolong Tea" />
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
            Oolong Tea
          </Typography>
          <Typography
            variant="body2"
            color="text.secondary"
            sx={{ fontFamily: 'Montserrat', fontWeight: 300 }}
          >
            Oolong tea is made from the Camellia sinensis plant. Its dried leaves and leaf buds are
            used to make several different teas, including black and green teas. Oolong tea is
            fermented for longer than green tea, but less than black tea. It contains caffeine which
            affects thinking and alertness.
          </Typography>
        </CardContent>
      </CardActionArea>
    </Card>
  )
}
