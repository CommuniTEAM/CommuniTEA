import TeaExpo from '../assets/TeaExpo.jpg';
import TeaTasting from '../assets/TeaTasting.png';
import FourSeasonsTea from '../assets/FourSeasonsTea.png';
import DataCoffee from '../assets/DataCoffee.png';
import TeaAndPainting from '../assets/TeaAndPainting.jpg';

interface Event {
  id?: number;
  image?: string;
  tag?: string;
  host?: string;
  name?: string;
  title?: string;
  location?: string;
  date?: string;
  startTime?: string;
  endTime?: string;
  price?: string;
  address?: string;
  attending?: string;
  headline?: string;
  headline2?: string;
  description: string;
  eventHighlight1?: string;
  eventHighlight1Bullet1?: string;
  eventHighlight1Bullet2?: string;
  eventHighlight1Bullet3?: string;
  eventHighlight2?: string;
  eventHighlight2Bullet1?: string;
  eventHighlight2Bullet2?: string;
  eventHighlight2Bullet3?: string;
  eventHighlight3?: string;
  eventHighlight3Bullet1?: string;
  eventHighlight3Bullet2?: string;
  eventHighlight3Bullet3?: string;
  eventHighlight4?: string;
  eventHighlight4Bullet1?: string;
  eventHighlight4Bullet2?: string;
  eventHighlight4Bullet3?: string;
  whyAttend?: string;
}

const eventData = [
  {
    id: 1,
    image: TeaExpo,
    tag: 'Seminar',
    host: 'Seattle Events',
    name: 'World Tea Conference and Expo',
    title: 'World Tea Conference and Expo',
    location: 'Seattle Center',
    date: 'November 26,2023',
    startTime: '4:00 PM',
    endTime: '8:00 PM',
    price: 'from $50',
    address: '305 Harrison St Seattle, WA 98109',
    attending: '2000',
    headline:
      'Join us on this extraordinary journey through the world of tea. Elevate your tea experience, broaden our horizons, and become a part of the vibrant tapestry that is the World Tea Conference and Expo! ',
    headline2: 'Let’s sip, learn, and celebrate tea together!',
    description:
      'Welcome to the premier global gathering for the tea community - the World Tea Conference and Expo! Join us for an immersive experience where tea enthusiasts, industry experts, and businesses converge to celebrate, learn, and explore the fascinating world of tea.',
    eventHighlight1: 'Explore the Global Tea Landscape',
    eventHighlight1Bullet1:
      'Immerse yourself in a diverse showcase of teas from around the world.',
    eventHighlight1Bullet2:
      'Engage with leading tea producers, distributors, and brands.',
    eventHighlight2: 'Educational Sessions and Workshops',
    eventHighlight2Bullet1:
      'Gain insights from industry leaders through a series of educational sessions.',
    eventHighlight2Bullet2:
      'Participate in hands-on workshops to redefine your tea knowledge and skills.',
    eventHighlight3: 'Networking Opportunities',
    eventHighlight3Bullet1:
      'Connect with fellow tea enthusiasts, professionals, and potential collaborators.',
    eventHighlight3Bullet2:
      'Forge valuable relationships that extends beyond the conference. ',
    eventHighlight4: 'Innovation and Trends:',
    eventHighlight4Bullet1:
      'Stay ahead of the curve by discovering the latest innovations and trends in the tea industry',
    eventHighlight4Bullet2:
      'Access exclusive previews of upcoming tea products and technologies',
    whyAttend:
      'The World Tea Conference and Expo is more than an event; it’s a celebration of the global tea community. Whether you’re a seasoned professional, a budding entrepreneur, or simply passionate about tea, this conference offers a platform to expand your knowledge, connect with like-minded individuals, and contribute to the ongoing evolution of the tea industry.',
  },
  {
    id: 2,
    image: TeaTasting,
    tag: '',
    host: '',
    name: 'Tea Tasting',
    title: 'Tea Tasting',
    location: "Coyle's Bakeshop",
    date: '11/25/2023',
    startTime: '7:00 PM',
    endTime: '8:30 PM',
    price: 'Free',
    address: '8300 Greenwood Ave N',
    attending: '20',
    headline:
      'Join us on this extraordinary journey through the world of tea. Elevate your tea experience, broaden our horizons, and become a part of the vibrant tapestry that is the World Tea Conference and Expo! ',
    headline2: 'Let’s sip, learn, and celebrate tea together!',
    description:
      'Welcome to the premier global gathering for the tea community - the World Tea Conference and Expo! Join us for an immersive experience where tea enthusiasts, industry experts, and businesses converge to celebrate, learn, and explore the fascinating world of tea.',
    eventHighlight1: 'Explore the Global Tea Landscape',
    eventHighlight1Bullet1:
      'Immerse yourself in a diverse showcase of teas from around the world.',
    eventHighlight1Bullet2:
      'Engage with leading tea producers, distributors, and brands.',
    eventHighlight2: 'Educational Sessions and Workshops',
    eventHighlight2Bullet1:
      'Gain insights from industry leaders through a series of educational sessions.',
    eventHighlight2Bullet2:
      'Participate in hands-on workshops to redefine your tea knowledge and skills.',
    eventHighlight3: 'Networking Opportunities',
    eventHighlight3Bullet1:
      'Connect with fellow tea enthusiasts, professionals, and potential collaborators.',
    eventHighlight3Bullet2:
      'Forge valuable relationships that extends beyond the conference. ',
    eventHighlight4: 'Innovation and Trends:',
    eventHighlight4Bullet1:
      'Stay ahead of the curve by discovering the latest innovations and trends in the tea industry',
    eventHighlight4Bullet2:
      'Access exclusive previews of upcoming tea products and technologies',
    whyAttend:
      'The World Tea Conference and Expo is more than an event; it’s a celebration of the global tea community. Whether you’re a seasoned professional, a budding entrepreneur, or simply passionate about tea, this conference offers a platform to expand your knowledge, connect with like-minded individuals, and contribute to the ongoing evolution of the tea industry.',
  },
  {
    id: 3,
    image: FourSeasonsTea,
    tag: '',
    host: '',
    name: 'Four Seasons Tea Gatherings',
    title: 'Four Seasons Tea Gatherings',
    location: 'East-West Chanoyu Center',
    date: '11/24/2023',
    startTime: '2:00 PM',
    endTime: '8:00 PM',
    price: 'Free',
    address: '1414 S Weller St.',
    attending: '10',
    headline:
      'Join us on this extraordinary journey through the world of tea. Elevate your tea experience, broaden our horizons, and become a part of the vibrant tapestry that is the World Tea Conference and Expo! ',
    headline2: 'Let’s sip, learn, and celebrate tea together!',
    description:
      'Welcome to the premier global gathering for the tea community - the World Tea Conference and Expo! Join us for an immersive experience where tea enthusiasts, industry experts, and businesses converge to celebrate, learn, and explore the fascinating world of tea.',
    eventHighlight1: 'Explore the Global Tea Landscape',
    eventHighlight1Bullet1:
      'Immerse yourself in a diverse showcase of teas from around the world.',
    eventHighlight1Bullet2:
      'Engage with leading tea producers, distributors, and brands.',
    eventHighlight2: 'Educational Sessions and Workshops',
    eventHighlight2Bullet1:
      'Gain insights from industry leaders through a series of educational sessions.',
    eventHighlight2Bullet2:
      'Participate in hands-on workshops to redefine your tea knowledge and skills.',
    eventHighlight3: 'Networking Opportunities',
    eventHighlight3Bullet1:
      'Connect with fellow tea enthusiasts, professionals, and potential collaborators.',
    eventHighlight3Bullet2:
      'Forge valuable relationships that extends beyond the conference. ',
    eventHighlight4: 'Innovation and Trends:',
    eventHighlight4Bullet1:
      'Stay ahead of the curve by discovering the latest innovations and trends in the tea industry',
    eventHighlight4Bullet2:
      'Access exclusive previews of upcoming tea products and technologies',
    whyAttend:
      'The World Tea Conference and Expo is more than an event; it’s a celebration of the global tea community. Whether you’re a seasoned professional, a budding entrepreneur, or simply passionate about tea, this conference offers a platform to expand your knowledge, connect with like-minded individuals, and contribute to the ongoing evolution of the tea industry.',
  },
  {
    id: 4,
    image: DataCoffee,
    tag: '',
    host: '',
    name: 'Seattle Data Professionals',
    title: 'Seattle Data Professionals',
    location: 'Cafe Hagen Modern Cafe',
    date: '11/24/2023',
    startTime: '10:00 AM',
    endTime: '9:00 PM',
    price: 'Free',
    address: '1252 Thomas St. Seattle, WA 98109',
    attending: '22',
    headline:
      'Join us on this extraordinary journey through the world of tea. Elevate your tea experience, broaden our horizons, and become a part of the vibrant tapestry that is the World Tea Conference and Expo! ',
    headline2: 'Let’s sip, learn, and celebrate tea together!',
    description:
      'Welcome to the premier global gathering for the tea community - the World Tea Conference and Expo! Join us for an immersive experience where tea enthusiasts, industry experts, and businesses converge to celebrate, learn, and explore the fascinating world of tea.',
    eventHighlight1: 'Explore the Global Tea Landscape',
    eventHighlight1Bullet1:
      'Immerse yourself in a diverse showcase of teas from around the world.',
    eventHighlight1Bullet2:
      'Engage with leading tea producers, distributors, and brands.',
    eventHighlight2: 'Educational Sessions and Workshops',
    eventHighlight2Bullet1:
      'Gain insights from industry leaders through a series of educational sessions.',
    eventHighlight2Bullet2:
      'Participate in hands-on workshops to redefine your tea knowledge and skills.',
    eventHighlight3: 'Networking Opportunities',
    eventHighlight3Bullet1:
      'Connect with fellow tea enthusiasts, professionals, and potential collaborators.',
    eventHighlight3Bullet2:
      'Forge valuable relationships that extends beyond the conference. ',
    eventHighlight4: 'Innovation and Trends:',
    eventHighlight4Bullet1:
      'Stay ahead of the curve by discovering the latest innovations and trends in the tea industry',
    eventHighlight4Bullet2:
      'Access exclusive previews of upcoming tea products and technologies',
    whyAttend:
      'The World Tea Conference and Expo is more than an event; it’s a celebration of the global tea community. Whether you’re a seasoned professional, a budding entrepreneur, or simply passionate about tea, this conference offers a platform to expand your knowledge, connect with like-minded individuals, and contribute to the ongoing evolution of the tea industry.',
  },
  {
    id: 5,
    image: TeaAndPainting,
    tag: '',
    host: '',
    name: 'Paint & Sip: Night Owl',
    title: 'Paint & Sip: Night Owl',
    location: 'Friday Afternoon Tea',
    date: '11/18/2023',
    startTime: '6:00 PM',
    endTime: '8:00 PM',
    price: '$50',
    address: '4228 Stone Way N',
    attending: '30',
    headline:
      'Join us on this extraordinary journey through the world of tea. Elevate your tea experience, broaden our horizons, and become a part of the vibrant tapestry that is the World Tea Conference and Expo! ',
    headline2: 'Let’s sip, learn, and celebrate tea together!',
    description:
      'Welcome to the premier global gathering for the tea community - the World Tea Conference and Expo! Join us for an immersive experience where tea enthusiasts, industry experts, and businesses converge to celebrate, learn, and explore the fascinating world of tea.',
    eventHighlight1: 'Explore the Global Tea Landscape',
    eventHighlight1Bullet1:
      'Immerse yourself in a diverse showcase of teas from around the world.',
    eventHighlight1Bullet2:
      'Engage with leading tea producers, distributors, and brands.',
    eventHighlight2: 'Educational Sessions and Workshops',
    eventHighlight2Bullet1:
      'Gain insights from industry leaders through a series of educational sessions.',
    eventHighlight2Bullet2:
      'Participate in hands-on workshops to redefine your tea knowledge and skills.',
    eventHighlight3: 'Networking Opportunities',
    eventHighlight3Bullet1:
      'Connect with fellow tea enthusiasts, professionals, and potential collaborators.',
    eventHighlight3Bullet2:
      'Forge valuable relationships that extends beyond the conference. ',
    eventHighlight4: 'Innovation and Trends:',
    eventHighlight4Bullet1:
      'Stay ahead of the curve by discovering the latest innovations and trends in the tea industry',
    eventHighlight4Bullet2:
      'Access exclusive previews of upcoming tea products and technologies',
    whyAttend:
      'The World Tea Conference and Expo is more than an event; it’s a celebration of the global tea community. Whether you’re a seasoned professional, a budding entrepreneur, or simply passionate about tea, this conference offers a platform to expand your knowledge, connect with like-minded individuals, and contribute to the ongoing evolution of the tea industry.',
  },
];

const getEventById = async (eventId: number): Promise<Event | undefined> =>
  new Promise((resolve, reject) => {
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
