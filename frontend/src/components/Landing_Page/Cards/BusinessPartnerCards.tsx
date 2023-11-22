import {
  Card,
  CardContent,
  CardMedia,
  Typography,
  CardActionArea,
} from '@mui/material';
import Slider from 'react-slick';
import { useTheme } from '@mui/material/styles';
import useMediaQuery from '@mui/material/useMediaQuery';
import CoylesBakeshop from '../../../assets/CoylesBakeshop.jpg';
import MiroTea from '../../../assets/MiroTea.png';
import QueenMaryTea from '../../../assets/QueenMaryTea.png';
import 'slick-carousel/slick/slick.css';
import 'slick-carousel/slick/slick-theme.css';

export default function BusinessPartnerCards(): JSX.Element {
  const theme = useTheme();
  const largeScreen = useMediaQuery(theme.breakpoints.up('lg'));
  const mediumScreen = useMediaQuery(theme.breakpoints.between('md', 'lg'));
  const smallScreen = useMediaQuery(theme.breakpoints.down('md'));

  const getImageHeight = (): string => {
    if (largeScreen) return '10vw';
    if (mediumScreen) return '10vw';
    if (smallScreen) return '10vw';
    return '60vw';
  };

  const getCardHeight = (): string => {
    if (largeScreen) return '18vw';
    if (mediumScreen) return '20vw';
    if (smallScreen) return '20vw';
    return '60vw';
  };

  // This is test data. Replace when API is ready with rotation of featured businesses.
  const businessesData = [
    {
      image: CoylesBakeshop,
      name: "Coyle's Bakeshop",
      title: "Coyle's Bakeshop",
      address: ' 8300 GREENWOOD AVE N',
      contact: '206-257-4636',
      website: 'https://www.coylesbakeshop.com/',
    },
    {
      image: MiroTea,
      name: 'Miro Tea',
      title: 'Miro Tea',
      address: ' 5405 BALLARD AVE NW',
      contact: '206-782-6832',
      website: 'https://mirotea.blogspot.com/',
    },
    {
      image: QueenMaryTea,
      name: 'Queen Mary Tea',
      title: 'Queen Mary Tea',
      address: ' 2912 NE 55TH ST',
      contact: '206-257-4636',
      website: 'https://queenmarytea.com/',
    },
    {
      image: CoylesBakeshop,
      name: "Coyle's Bakeshop",
      title: "Coyle's Bakeshop",
      address: ' 8300 GREENWOOD AVE N',
      contact: '206-257-4636',
      website: 'https://www.coylesbakeshop.com/',
    },
    {
      image: MiroTea,
      name: 'Miro Tea',
      title: 'Miro Tea',
      address: ' 5405 BALLARD AVE NW',
      contact: '206-782-6832',
      website: 'https://mirotea.blogspot.com/',
    },
    {
      image: QueenMaryTea,
      name: 'Queen Mary Tea',
      title: 'Queen Mary Tea',
      address: ' 2912 NE 55TH ST',
      contact: '206-257-4636',
      website: 'https://queenmarytea.com/',
    },
    {
      image: CoylesBakeshop,
      name: "Coyle's Bakeshop",
      title: "Coyle's Bakeshop",
      address: ' 8300 GREENWOOD AVE N',
      contact: '206-257-4636',
      website: 'https://www.coylesbakeshop.com/',
    },
    {
      image: MiroTea,
      name: 'Miro Tea',
      title: 'Miro Tea',
      address: ' 5405 BALLARD AVE NW',
      contact: '206-782-6832',
      website: 'https://mirotea.blogspot.com/',
    },
    {
      image: QueenMaryTea,
      name: 'Queen Mary Tea',
      title: 'Queen Mary Tea',
      address: ' 2912 NE 55TH ST',
      contact: '206-257-4636',
      website: 'https://queenmarytea.com/',
    },
    {
      image: CoylesBakeshop,
      name: "Coyle's Bakeshop",
      title: "Coyle's Bakeshop",
      address: ' 8300 GREENWOOD AVE N',
      contact: '206-257-4636',
      website: 'https://www.coylesbakeshop.com/',
    },
    {
      image: MiroTea,
      name: 'Miro Tea',
      title: 'Miro Tea',
      address: ' 5405 BALLARD AVE NW',
      contact: '206-782-6832',
      website: 'https://mirotea.blogspot.com/',
    },
    {
      image: QueenMaryTea,
      name: 'Queen Mary Tea',
      title: 'Queen Mary Tea',
      address: ' 2912 NE 55TH ST',
      contact: '206-257-4636',
      website: 'https://queenmarytea.com/',
    },
  ];

  // Slider settings
  const settings = {
    dots: true,
    infinite: true,
    speed: 500,
    slidesToShow: 3,
    slidesToScroll: 3,
    responsive: [
      {
        breakpoint: 1024,
        settings: {
          slidesToShow: 2,
          slidesToScroll: 2,
        },
      },
      {
        breakpoint: 600,
        settings: {
          slidesToShow: 1,
          slidesToScroll: 1,
        },
      },
    ],
  };

  return (
    <div style={{ padding: '0 40px' }}>
      {/* eslint-disable-next-line react/jsx-props-no-spreading */}
      <Slider {...settings}>
        {businessesData.map((business) => (
          <Card
            key={business.name}
            sx={{ width: '100%', height: getCardHeight(), borderRadius: 10 }}
          >
            <CardActionArea sx={{ height: '100%' }}>
              <CardMedia
                sx={{ height: getImageHeight() }}
                image={business.image}
                title={business.title}
              />
              <CardContent sx={{ height: '40%' }}>
                <Typography
                  gutterBottom
                  variant="body1"
                  component="div"
                  sx={{
                    fontFamily: 'Montserrat',
                    fontWeight: 300,
                    color: '#29C6CF',
                  }}
                >
                  Partner
                </Typography>
                <Typography
                  gutterBottom
                  variant="h5"
                  component="div"
                  sx={{ fontFamily: 'Montserrat', fontWeight: 700 }}
                >
                  {business.title}
                </Typography>
                <Typography
                  variant="body2"
                  color="text.secondary"
                  sx={{ fontFamily: 'Montserrat', fontWeight: 300 }}
                >
                  Address:
                  {business.address}
                </Typography>
                <Typography
                  variant="body2"
                  color="text.secondary"
                  sx={{ fontFamily: 'Montserrat', fontWeight: 300 }}
                >
                  Contact:
                  {business.contact}
                </Typography>
                <Typography
                  variant="body2"
                  color="text.secondary"
                  sx={{ fontFamily: 'Montserrat', fontWeight: 300 }}
                >
                  Website:
                  {business.website}
                </Typography>
              </CardContent>
            </CardActionArea>
          </Card>
        ))}
      </Slider>
    </div>
  );
}
