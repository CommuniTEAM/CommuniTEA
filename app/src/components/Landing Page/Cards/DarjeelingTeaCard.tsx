import * as React from 'react'
import { Card, CardContent, CardMedia, Typography, CardActionArea } from '@mui/material'
import useMediaQuery from '@mui/material/useMediaQuery'
import DarjeelingTea from '../../../assets/DarjeelingTea.jpg'

export default function DarjeelingTeaCard (): JSX.Element {
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
        <CardMedia sx={{ height: '60%' }} image={DarjeelingTea} title="Darjeeling Tea" />
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
            Darjeeling
          </Typography>
          <Typography
            variant="body2"
            color="text.secondary"
            sx={{ fontFamily: 'Montserrat', fontWeight: 300 }}
          >
            Darjeeling tea is a type of black tea produced in India. Darjeeling tea has a fruity
            aroma and a golden or bronze color, depending on the way it&apos;s brewed. Tea experts
            say it has notes (flavors) of citrus fruit, flowers, and even a vegetal quality.
            Darjeeling tastes sweeter and less bitter than other forms of black tea.
          </Typography>
        </CardContent>
      </CardActionArea>
    </Card>
  )
}
