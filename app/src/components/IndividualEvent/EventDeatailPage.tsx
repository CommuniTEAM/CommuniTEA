import { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import { Typography } from '@mui/material';
import NavBar from '../Landing Page/Navbar';
import getEventById from '../../Axios/getEventById';

interface EventData {
  id: number
  image: string
  name: string
  title: string
  Location: string
  Date: string
  Time: string
  Price: string
  Address: string
  Attending: string
}

export default function EventDetailPage(): JSX.Element {
  const { eventId } = useParams<string>();
  const [eventData, setEventData] = useState<EventData | null>(null);
  const [loading, setLoading] = useState<boolean>(true);

  useEffect(() => {
    const fetchEventData = async (): Promise<void> => {
      try {
        if (eventId !== null && eventId !== undefined) {
          const response = await getEventById(Number(eventId));
          setEventData(response);
        }
      } catch (error) {
        // eslint-disable-next-line no-console
        console.error('Failed to fetch events', error);
      } finally {
        setLoading(false);
      }
    };

    void fetchEventData();
  }, [eventId]);

  if (loading) {
    return <div>loading...</div>;
  }

  if (eventData == null) {
    return <div>Event not found</div>;
  }

  // console.log(eventData);

  return (
    <div>
      <NavBar />
      <img src={eventData.image} alt={eventData.name} />
      <Typography variant="h1">{eventData.name}</Typography>
      {eventData.date}
      <Typography variant="body2">{eventData.headline}</Typography>
      <Typography variant="body2">{eventData.headline2}</Typography>
    </div>
  );
}
