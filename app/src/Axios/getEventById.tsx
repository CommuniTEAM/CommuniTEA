import TeaExpo from '../assets/TeaExpo.jpg';
import TeaTasting from '../assets/TeaTasting.png';
import FourSeasonsTea from '../assets/FourSeasonsTea.png';
import DataCoffee from '../assets/DataCoffee.png';
import TeaAndPainting from '../assets/TeaAndPainting.jpg';

interface Event {
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

const eventData = [
  {
    id: 1,
    image: TeaExpo,
    name: 'World Tea Conference and Expo',
    title: 'World Tea Conference and Expo',
    Location: 'Seattle Center',
    Date: '11/26/2023',
    Time: '5:00 PM',
    Price: 'from $50',
    Address: '305 Harrison St',
    Attending: '2000',
    headline:
      'Join us on this extraordinary journey through the world of tea. Elevate your tea experience, broaden our horizons, and become a part of the vibrant tapestry that is the World Tea Conference and Expo! ',
    headline2: 'Let’s sip, learn, and celebrate tea together!',
  },
  {
    id: 2,
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
    id: 3,
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
    id: 4,
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
    id: 5,
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

const getEventById = async (eventId: number): Promise<Event> => new Promise((resolve, reject) => {
  setTimeout(() => {
    const foundEvent = eventData.find(
      (eventItem) => eventItem.id === eventId,
    );
    if (foundEvent !== undefined) {
      resolve(foundEvent);
    } else {
      reject(new Error('Event not found'));
    }
  }, 1000);
});

export default getEventById;
