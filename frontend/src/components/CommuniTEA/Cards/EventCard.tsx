import {
  Card,
  CardActionArea,
  CardContent,
  CardMedia,
  Divider,
  Typography,
} from '@mui/material';
import { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import getEvents from '../../../Axios/getEvents';
import LikeIcon from '../../../assets/LikeIcon.png';

interface Event {
  id: number;
  image: string;
  name: string;
  title: string;
  Location: string;
  Date: string;
  Time: string;
  Price: string;
  Address: string;
  Attending: string;
}

export default function EventCards(): JSX.Element {
  // TODO: Add functionality to the card
  // TODO: Add Map functionality
  const [events, setEvents] = useState<Event[]>([]);

  useEffect(() => {
    getEvents()
      .then((response) => {
        setEvents(response);
      })
      .catch((error) => {
        console.error('Failed to fetch events', error);
      });
  }, []);

  return (
    <div style={{ width: '100%', marginTop: '2vh' }}>
      {events.map((event) => (
        <Card
          key={event.name}
          sx={{ display: 'flex', width: '100%', marginBottom: 2 }}
        >
          <CardActionArea
            component={Link}
            to={`/event/${event.id}`}
            sx={{ width: 'inherit', display: 'flex' }}
          >
            <CardMedia
              sx={{ height: 250, width: '60%' }}
              image={event.image}
              title={event.name}
            />
            <div
              style={{ display: 'flex', flexDirection: 'column', width: '40%' }}
            >
              <CardContent sx={{ marginBottom: '2vh' }}>
                <Typography gutterBottom={true} variant="h5" component="div">
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
