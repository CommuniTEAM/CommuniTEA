import { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import {
  Button, Divider, Paper, Typography,
} from '@mui/material';
import NavBar from '../LandingPage/Navbar';
import getEventById from '../../Axios/getEventById';
import LikeIcon from '../../assets/LikeIcon.png';
import shareIcon from '../../assets/ShareIcon.png';
import MapPinIcon from '../../assets/MapPinIcon.png';
import OtherEvents from './OtherEvents';
import Footer from '../LandingPage/Footer';
import RsvpForm from './Forms/RsvpForm';

import './styles/IndividualEventStyles.css';

interface EventData {
  id?: number
  image?: string
  tag?: string
  host?: string
  name?: string
  title?: string
  location?: string
  date?: string
  startTime?: string
  endTime?: string
  price?: string
  address?: string
  attending?: string
  headline?: string
  headline2?: string
  description?: string
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
          if (response !== null && response !== undefined) {
            setEventData(response);
          } else {
            setEventData(null);
          }
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

  const parallaxStyle = {
    backgroundImage: `url(${eventData.image})`,
  };

  return (
    <div>
      <NavBar />
      <div className="heroBannerContainer" style={parallaxStyle} />
      <div className="individualEventContainer">
        <div className="iconSection">
          <div className="eventTag">
            <Typography
              sx={{
                fontFamily: 'Montserrat',
                color: '#29C6CF',
                fontSize: '1.25em',
              }}
            >
              {eventData.tag}
            </Typography>
          </div>

          <div className="eventIcons">
            <div className="likeContainer">
              <img src={LikeIcon} alt="Like Icon" />
              <p>Like</p>
            </div>
            <div className="shareContainer">
              <img src={shareIcon} alt="Share Icon" />
              <p>Share</p>
            </div>
          </div>
        </div>

        <Divider sx={{ marginTop: '2vh', marginBottom: '2vh' }} />

        <div className="mainSection">
          <div className="mainSection-firstHalf">
            <Typography
              variant="h2"
              sx={{ fontFamily: 'Montserrat', fontSize: '3em' }}
            >
              {eventData.name}
            </Typography>
            <div className="dateAndTimeContainer">
              <div className="date">
                <Typography
                  variant="h5"
                  sx={{ fontFamily: 'Montserrat', fontSize: '1.5em' }}
                >
                  {eventData.date}
                </Typography>
              </div>
              <div className="time">
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
            <div className="headline">
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
                marginTop: '2vh',
                padding: '30px 10px',
                borderRadius: '10px',
              }}
            >
              <div className="hostDetails">
                <div className="hostDetails-firstHalf">
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
                <div className="hostDetails-secondHalf">
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

            <div className="locationHeaderContainer">
              <Typography
                variant="h3"
                sx={{ fontFamily: 'Montserrat', fontSize: '3em' }}
              >
                Location
              </Typography>
              <div className="locationDetailsContainer">
                <img src={MapPinIcon} alt="Location" />
                <div>
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
            <div className="mapsContainer">
              <iframe
                src={`https://www.google.com/maps/embed/v1/place?key=${apiKey}&q=${eventData.location}`}
                style={{ width: '100%', height: '50vh', border: 0 }}
                title="Google Maps"
              />
            </div>

            <Divider sx={{ marginTop: '50px', marginBottom: '50px' }} />

            <div className="aboutEventContainer">
              <Typography variant="h3" sx={{ fontFamily: 'Montserrat' }}>
                About this Event
              </Typography>
              <div style={{ marginTop: '20px' }}>
                <Typography variant="body2" sx={{ fontFamily: 'Montserrat' }}>
                  {eventData.description}
                </Typography>
              </div>
              <div style={{ marginTop: '20px' }}>
                <Typography
                  variant="h4"
                  sx={{ fontFamily: 'Montserrat', marginBottom: '20px' }}
                >
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

            <div className="whyAttend">
              <Typography variant="h4" sx={{ fontFamily: 'Montserrat' }}>
                Why Attend?
              </Typography>
              <div style={{ marginTop: '1vh' }}>
                <Typography variant="body2" sx={{ fontFamily: 'Montserrat' }}>
                  {eventData.whyAttend}
                </Typography>
              </div>
            </div>

            <Divider sx={{ marginTop: '50px', marginBottom: '50px' }} />
          </div>

          <div className="mainSection-secondHalf">
            <div className="formContainer">
              <RsvpForm />
            </div>
            <div className="upcomingEventsContainer">
              <Divider sx={{ marginTop: '50px', marginBottom: '50px' }} />
              <Typography
                variant="h4"
                sx={{ fontFamily: 'Montserrat', textAlign: 'center' }}
              >
                Upcoming events hosted by
                {' '}
                <span className="eventHost">{eventData.host}</span>
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
