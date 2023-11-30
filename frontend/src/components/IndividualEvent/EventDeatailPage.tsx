import { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import {
  Button, Divider, Paper, Typography,
} from '@mui/material';
import NavBar from '../LandingPage/Navbar';
import getEventById from '../../Axios/getEventById';
import LikeIcon from '../../assets/LikeIcon.png';
import shareIcon from '../../assets/shareIcon.png';
import MapPinIcon from '../../assets/MapPinIcon.png';
import OtherEvents from './OtherEvents';
import Footer from '../LandingPage/Footer';
import RsvpForm from './Forms/RsvpForm';

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
  const apiKey = import.meta.env.VITE_GOOGLE_MAPS_API_KEY;

  // Concat Street Address, City, State, and Zip Code
  // Create a variable to store the address
  // Call the variable in the Google Maps query

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
      <div
        style={{
          marginLeft: '5vw',
          marginRight: '5vw',
        }}
      >
        <div style={{ textAlign: 'center', marginTop: '2vh' }}>
          <img
            src={eventData.image}
            alt={eventData.name}
            style={{ borderRadius: '50px', width: '100%', height: '80vh' }}
          />
        </div>
        <div style={{ display: 'flex', marginTop: '2vh' }}>
          <div style={{ width: '90%' }}>
            <Typography
              sx={{
                fontFamily: 'Montserrat',
                color: '#29C6CF',
                fontSize: '.80vw',
              }}
            >
              {eventData.tag}
            </Typography>
          </div>
          <div
            style={{
              display: 'flex',
              width: '10%',
              justifyContent: 'space-between',
              alignItems: 'center',
            }}
          >
            <div
              style={{
                fontFamily: 'Montserrat',
                fontSize: '.80vw',
                display: 'flex',
                alignItems: 'center',
              }}
            >
              <img
                src={LikeIcon}
                alt="Like Icon"
                style={{ width: '1vw', marginRight: '.5vw' }}
              />
              Like
            </div>
            <div
              style={{
                fontFamily: 'Montserrat',
                fontSize: '.80vw',
              }}
            >
              <img
                src={shareIcon}
                alt="Share Icon"
                style={{ width: '1vw', marginRight: '.5vw' }}
              />
              Share
            </div>
          </div>
        </div>

        <Divider sx={{ marginTop: '2vh', marginBottom: '2vh' }} />

        <div style={{ display: 'flex' }}>
          <div style={{ width: '70%' }}>
            <Typography
              variant="h2"
              sx={{ fontFamily: 'Montserrat', fontSize: '3em' }}
            >
              {eventData.name}
            </Typography>
            <div style={{ display: 'flex', marginTop: '1vh' }}>
              <div style={{ width: '80%' }}>
                <Typography
                  variant="h5"
                  sx={{ fontFamily: 'Montserrat', fontSize: '1.5em' }}
                >
                  {eventData.date}
                </Typography>
              </div>
              <div style={{ display: 'flex', justifyContent: 'flex-end' }}>
                <Typography
                  variant="h5"
                  sx={{ fontFamily: 'Montserrat', fontSize: '1.5em' }}
                >
                  {eventData.startTime}
                  -
                  {eventData.endTime}
                  PST
                </Typography>
              </div>
            </div>
            <div style={{ paddingTop: '2vh' }}>
              <Typography
                variant="body2"
                sx={{ fontFamily: 'Montserrat', fontSize: '1em' }}
              >
                {eventData.headline}
              </Typography>
              <Typography
                variant="body2"
                sx={{
                  fontFamily: 'Montserrat',
                  fontStyle: 'italic',
                  marginTop: '1vh',
                  fontSize: '1em',
                }}
              >
                {eventData.headline2}
              </Typography>
            </div>

            <Paper
              elevation={0}
              sx={{
                backgroundColor: '#333',
                height: '5em',
                marginTop: '2vh',
              }}
            >
              <div style={{ display: 'flex', height: '100%' }}>
                <div
                  style={{
                    display: 'flex',
                    flexDirection: 'column',
                    width: '70%',
                    justifyContent: 'center',
                    marginLeft: '2vw',
                  }}
                >
                  <div>
                    <Typography
                      variant="h4"
                      sx={{
                        fontFamily: 'Montserrat',
                        color: 'white',
                        fontSize: '1.5em',
                      }}
                    >
                      Seattle Events
                    </Typography>
                  </div>
                  <div>
                    <Typography
                      variant="body2"
                      sx={{
                        fontFamily: 'Montserrat',
                        color: 'white',
                        fontSize: '1em',
                      }}
                    >
                      10k+ Followers
                    </Typography>
                  </div>
                </div>
                <div
                  style={{
                    display: 'flex',
                    flexDirection: 'column',
                    justifyContent: 'center',
                    width: '30%',
                    marginRight: '2vw',
                  }}
                >
                  <Button
                    variant="contained"
                    sx={{
                      borderRadius: '10px',
                      backgroundColor: '#87CEEB',
                      color: 'black',
                      fontFamily: 'Montserrat',
                    }}
                  >
                    Follow
                  </Button>
                </div>
              </div>
            </Paper>

            <Divider sx={{ marginTop: '2vh', marginBottom: '2vh' }} />

            <div style={{ marginTop: '2vh' }}>
              <Typography
                variant="h3"
                sx={{ fontFamily: 'Montserrat', fontSize: '3em' }}
              >
                Location
              </Typography>
              <div style={{ display: 'flex', marginTop: '1vh' }}>
                <img
                  src={MapPinIcon}
                  alt="Location"
                  style={{ width: '1.5em', height: '1.5em' }}
                />
                <div style={{ marginLeft: '3vh' }}>
                  <Typography
                    variant="body2"
                    sx={{ fontFamily: 'Montserrat', fontWeight: 'bold' }}
                  >
                    {eventData.location}
                  </Typography>
                  <Typography variant="body2" sx={{ fontFamily: 'Montserrat' }}>
                    {eventData.address}
                  </Typography>
                </div>
              </div>
            </div>
            <div>
              <iframe
                src={`https://www.google.com/maps/embed/v1/place?key=${apiKey}&q=${eventData.location}`}
                style={{ width: '100%', height: '50vh', border: 0 }}
                title="Google Maps"
              />
            </div>

            <Divider sx={{ marginTop: '2vh', marginBottom: '2vh' }} />

            <div>
              <Typography variant="h3" sx={{ fontFamily: 'Montserrat' }}>
                About this Event
              </Typography>
              <div style={{ marginTop: '1vh' }}>
                <Typography variant="body2" sx={{ fontFamily: 'Montserrat' }}>
                  {eventData.description}
                </Typography>
              </div>
              <div style={{ marginTop: '1vh' }}>
                <Typography variant="h4" sx={{ fontFamily: 'Montserrat' }}>
                  Event Highlights:
                </Typography>
              </div>
              <ul>
                <li>
                  <Typography variant="h5" sx={{ fontFamily: 'Montserrat' }}>
                    {eventData.eventHighlight1}
                  </Typography>
                  <ul>
                    <li>
                      <Typography
                        variant="body2"
                        sx={{ fontFamily: 'Montserrat' }}
                      >
                        {eventData.eventHighlight1Bullet1}
                      </Typography>
                    </li>
                    <li>
                      <Typography
                        variant="body2"
                        sx={{ fontFamily: 'Montserrat' }}
                      >
                        {eventData.eventHighlight1Bullet2}
                      </Typography>
                    </li>
                  </ul>
                </li>
                <li>
                  <Typography variant="h5" sx={{ fontFamily: 'Montserrat' }}>
                    {eventData.eventHighlight2}
                  </Typography>
                  <ul>
                    <li>
                      <Typography
                        variant="body2"
                        sx={{ fontFamily: 'Montserrat' }}
                      >
                        {eventData.eventHighlight2Bullet1}
                      </Typography>
                    </li>
                    <li>
                      <Typography
                        variant="body2"
                        sx={{ fontFamily: 'Montserrat' }}
                      >
                        {eventData.eventHighlight2Bullet2}
                      </Typography>
                    </li>
                  </ul>
                </li>
                <li>
                  <Typography variant="h5" sx={{ fontFamily: 'Montserrat' }}>
                    {eventData.eventHighlight3}
                  </Typography>
                  <ul>
                    <li>
                      <Typography
                        variant="body2"
                        sx={{ fontFamily: 'Montserrat' }}
                      >
                        {eventData.eventHighlight3Bullet1}
                      </Typography>
                    </li>
                    <li>
                      <Typography
                        variant="body2"
                        sx={{ fontFamily: 'Montserrat' }}
                      >
                        {eventData.eventHighlight3Bullet2}
                      </Typography>
                    </li>
                  </ul>
                </li>
                <li>
                  <Typography variant="h5" sx={{ fontFamily: 'Montserrat' }}>
                    {eventData.eventHighlight4}
                  </Typography>
                  <ul>
                    <li>
                      <Typography
                        variant="body2"
                        sx={{ fontFamily: 'Montserrat' }}
                      >
                        {eventData.eventHighlight4Bullet1}
                      </Typography>
                    </li>
                    <li>
                      <Typography
                        variant="body2"
                        sx={{ fontFamily: 'Montserrat' }}
                      >
                        {eventData.eventHighlight4Bullet2}
                      </Typography>
                    </li>
                  </ul>
                </li>
              </ul>
            </div>

            <div>
              <Typography variant="h4" sx={{ fontFamily: 'Montserrat' }}>
                Why Attend?
              </Typography>
              <div style={{ marginTop: '1vh' }}>
                <Typography variant="body2" sx={{ fontFamily: 'Montserrat' }}>
                  {eventData.whyAttend}
                </Typography>
              </div>
            </div>

            <Divider sx={{ marginTop: '2vh', marginBottom: '2vh' }} />
          </div>

          <div
            style={{
              display: 'flex',
              flexDirection: 'column',
              width: '30%',
              marginLeft: '2vw',
            }}
          >
            <div
              style={{
                backgroundColor: '#333',
              }}
            >
              <RsvpForm />
            </div>
            <div style={{ flexGrow: 1 }}>
              <Divider sx={{ marginTop: '2vh', marginBottom: '2vh' }} />
              <Typography
                variant="h4"
                sx={{ fontFamily: 'Montserrat', textAlign: 'center' }}
              >
                Upcoming events hosted by
                {' '}
                {eventData.host}
              </Typography>
            </div>
          </div>
        </div>
      </div>
      <OtherEvents />
      <Footer />
    </div>
  );
}
