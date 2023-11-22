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
  location: string
  date: string
  startTime: string
  endTime: string
  price: string
  address: string
  attending: string
  headline?: string
  headline2?: string
  description: string
  eventHighlight1?: string
  eventHighlight1Bullet1?: string
  eventHighlight1Bullet2?: string
  eventHighlight1Bullet3?: string
  eventHighlight2?: string
  eventHighlight2Bullet1?: string
  eventHighlight2Bullet2?: string
  eventHighlight2Bullet3?: string
  eventHighlight3?: string
  eventHighlight3Bullet1?: string
  eventHighlight3Bullet2?: string
  eventHighlight3Bullet3?: string
  eventHighlight4?: string
  eventHighlight4Bullet1?: string
  eventHighlight4Bullet2?: string
  eventHighlight4Bullet3?: string
  whyAttend?: string
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

  return (
    <div>
      <NavBar />
      <img src={eventData.image} alt={eventData.name} />
      <Typography variant="h1">{eventData.name}</Typography>
      <Typography variant="h5">{eventData.date}</Typography>
      <Typography variant="h5">{eventData.startTime}</Typography>
      <Typography variant="h5">{eventData.endTime}</Typography>
      <Typography variant="body2">{eventData.headline}</Typography>
      <Typography variant="body2">{eventData.headline2}</Typography>
      <Typography variant="h3">Location</Typography>
      <Typography variant="body2">{eventData.location}</Typography>
      <Typography variant="body2">{eventData.address}</Typography>
      <Typography variant="h3">About this Event</Typography>
      <Typography variant="body2">{eventData.description}</Typography>
      <Typography variant="h4">Event Highlights:</Typography>
      <ul>
        <li>
          <Typography variant="h5">{eventData.eventHighlight1}</Typography>
          <ul>
            <li>
              <Typography variant="body2">
                {eventData.eventHighlight1Bullet1}
              </Typography>
            </li>
            <li>
              <Typography variant="body2">
                {eventData.eventHighlight1Bullet2}
              </Typography>
            </li>
          </ul>
        </li>
        <li>
          <Typography variant="h5">{eventData.eventHighlight2}</Typography>
          <ul>
            <li>
              <Typography variant="body2">
                {eventData.eventHighlight2Bullet1}
              </Typography>
            </li>
            <li>
              <Typography variant="body2">
                {eventData.eventHighlight2Bullet2}
              </Typography>
            </li>
          </ul>
        </li>
        <li>
          <Typography variant="h5">{eventData.eventHighlight3}</Typography>
          <ul>
            <li>
              <Typography variant="body2">
                {eventData.eventHighlight3Bullet1}
              </Typography>
            </li>
            <li>
              <Typography variant="body2">
                {eventData.eventHighlight3Bullet2}
              </Typography>
            </li>
          </ul>
        </li>
        <li>
          <Typography variant="h5">{eventData.eventHighlight4}</Typography>
          <ul>
            <li>
              <Typography variant="body2">
                {eventData.eventHighlight4Bullet1}
              </Typography>
            </li>
            <li>
              <Typography variant="body2">
                {eventData.eventHighlight4Bullet2}
              </Typography>
            </li>
          </ul>
        </li>
      </ul>
    </div>
  );
}
