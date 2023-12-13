import Slider from 'react-slick';
import { Typography } from '@mui/material';
import Cory from '../../assets/Cory.jpg';
import Angela from '../../assets/Angela.jpg';
import Brian from '../../assets/Brian.jpg';
import Amanda from '../../assets/Amanda.jpg';
import Hector from '../../assets/Hector.jpg';
import Alex from '../../assets/Alex.jpg';

import './styles.css';

export default function TeamMemberCarousel(): JSX.Element {
  const settings = {
    dots: false,
    infinite: true,
    slidesToShow: 5,
    slidesToScroll: 1,
    autoplay: true,
    speed: 3000,
    autoplaySpeed: 1,
    cssEase: 'linear',
    slide: 'div',
    pauseOnHover: true,
  };

  return (
    <div>
      <Slider
        dots={settings.dots}
        infinite={settings.infinite}
        slidesToShow={settings.slidesToShow}
        slidesToScroll={settings.slidesToScroll}
        autoplay={settings.autoplay}
        speed={settings.speed}
        autoplaySpeed={settings.autoplaySpeed}
        cssEase={settings.cssEase}
        slide={settings.slide}
        pauseOnHover={settings.pauseOnHover}
      >
        <div className="image-container">
          <img src={Cory} alt="Cory" />
          <div className="overlay">
            <Typography className="overlay-name">Cory DeGuzman</Typography>
            <Typography className="overlay-title">
              Project Manager & Front-end Engineer
            </Typography>
          </div>
        </div>
        <div className="image-container">
          <img src={Angela} alt="Angela" />
          <div className="overlay">
            <Typography className="overlay-name">Angela Fisher</Typography>
            <Typography className="overlay-title">Back-end Engineer</Typography>
          </div>
        </div>
        <div className="image-container">
          <img src={Brian} alt="Brian" />
          <div className="overlay">
            <Typography className="overlay-name">Brian La</Typography>
            <Typography className="overlay-title">Back-end Engineer</Typography>
          </div>
        </div>
        <div className="image-container">
          <img src={Amanda} alt="Amanda" />
          <div className="overlay">
            <Typography className="overlay-name">Amanda Taing</Typography>
            <Typography className="overlay-title">Back-end Engineer</Typography>
          </div>
        </div>
        <div className="image-container">
          <img src={Hector} alt="Hector" />
          <div className="overlay">
            <Typography className="overlay-name">Hector Elias</Typography>
            <Typography className="overlay-title">
              Front-end Engineer
            </Typography>
          </div>
        </div>
        <div className="image-container">
          <img src={Alex} alt="Alex" />
          <div className="overlay">
            <Typography className="overlay-name">Alex Ho</Typography>
            <Typography className="overlay-title">
              Front-end Engineer
            </Typography>
          </div>
        </div>
      </Slider>
    </div>
  );
}
