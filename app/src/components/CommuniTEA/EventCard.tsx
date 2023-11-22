import {
  Card,
  CardActionArea,
  CardContent,
  CardMedia,
  Divider,
  Typography,
} from '@mui/material';
import TeaExpo from '../../assets/TeaExpo.jpg';
import LikeIcon from '../../assets/LikeIcon.png';
import TeaTasting from '../../assets/TeaTasting.png';
import FourSeasonsTea from '../../assets/FourSeasonsTea.png';
import DataCoffee from '../../assets/DataCoffee.png';
import TeaAndPainting from '../../assets/TeaAndPainting.jpg';

export default function EventCards(): JSX.Element {
  // TODO: Add functionality to the card
  // TODO: Add Map functionality

  // Test data, replace with API call
  const eventsData = [
    {
      image: TeaExpo,
      name: 'World Tea Conference and Expo',
      title: 'World Tea Conference and Expo',
      Location: 'Seattle Center',
      Date: '11/26/2023',
      Time: '5:00 PM',
      Price: 'from $50',
      Address: '305 Harrison St',
      Attending: '2000',
    },
    {
      image: TeaTasting,
      name: 'Tea Tasting',
      title: 'Tea Tasting',
      Location: "Coyle's Bakeshop",
      Date: '11/25/2023',
      Time: '7:00 PM',
      Price: 'Free',
      Address: '8300 Greenwood Ave N',
      Attending: '20',
    },
    {
      image: FourSeasonsTea,
      name: 'Four Seasons Tea Gatherings',
      title: 'Four Seasons Tea Gatherings',
      Location: 'East-West Chanoyu Center',
      Date: '11/24/2023',
      Time: '2:00 PM',
      Price: 'Free',
      Address: '1414 S Weller St. East Bldg. fl 2',
      Attending: '10',
    },
    {
      image: DataCoffee,
      name: 'Seattle Data Professionals',
      title: 'Seattle Data Professionals',
      Location: 'Cafe Hagen Modern Cafe',
      Date: '11/24/2023',
      Time: '10:00 AM',
      Price: 'Free',
      Address: '1252 Thomas St.',
      Attending: '22',
    },
    {
      image: TeaAndPainting,
      name: 'Paint & Sip: Night Owl',
      title: 'Paint & Sip: Night Owl',
      Location: 'Friday Afternoon Tea',
      Date: '11/18/2023',
      Time: '6:00 PM - 8:00 PM',
      Price: '$50',
      Address: '4228 Stone Way N',
      Attending: '30',
    },
  ];

  return (
    <div style={{ width: '100%', marginTop: '2vh' }}>
      {eventsData.map((event) => (
        <Card
          key={event.name}
          sx={{ display: 'flex', width: '100%', marginBottom: 2 }}
        >
          <CardActionArea sx={{ width: 'inherit', display: 'flex' }}>
            <CardMedia
              sx={{ height: 250, width: '60%' }}
              image={event.image}
              title={event.name}
            />
            <div
              style={{ display: 'flex', flexDirection: 'column', width: '40%' }}
            >
              <CardContent sx={{ marginBottom: '2vh' }}>
                <Typography gutterBottom variant="h5" component="div">
                  {event.name}
                </Typography>
                <Typography variant="body2" color="text.secondary">
                  Location:
                  {event.Location}
                </Typography>
                <Typography variant="body2" color="text.secondary">
                  Date:
                  {event.Date}
                </Typography>
                <Typography variant="body2" color="text.secondary">
                  Time:
                  {event.Time}
                </Typography>
                <Typography variant="body2" color="text.secondary">
                  Price:
                  {event.Price}
                </Typography>
              </CardContent>

              <Divider sx={{ width: '90%', alignSelf: 'center' }} />

              <div
                style={{
                  display: 'flex',
                  alignItems: 'center',
                  height: '100%',
                }}
              >
                <div style={{ flexGrow: 1, marginLeft: '1vw' }}>
                  <Typography variant="body2" color="text.secondary">
                    {event.Address}
                  </Typography>
                </div>
                <div style={{ display: 'flex', alignItems: 'center' }}>
                  <img
                    src={LikeIcon}
                    alt="Like Icon"
                    style={{ width: '.75vw' }}
                  />
                  <Typography variant="body2" color="text.secondary">
                    {' '}
                    {event.Attending}
                  </Typography>
                </div>
              </div>
            </div>
          </CardActionArea>
        </Card>
      ))}
    </div>
  );
}
