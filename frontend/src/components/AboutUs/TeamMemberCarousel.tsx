import Slider from 'react-slick';
import { Typography } from '@mui/material';

import Cory from '../../assets/Cory.jpg';
import Angela from '../../assets/Angela.jpg';
import Brian from '../../assets/Brian.jpg';

import './styles/TeamMemberDetailsStyles.css';

interface TeamMemberCarouselProps {
  onSelectTeamMember: (memberName: string) => void;
}

export default function TeamMemberCarousel({
  onSelectTeamMember,
}: TeamMemberCarouselProps): JSX.Element {
  const handleSelectTeamMember = (name: string): void => {
    onSelectTeamMember(name);
  };

  const handleKeyPress = (event: any, memberName: any): void => {
    if (event.key === 'Enter' || event.key === ' ') {
      onSelectTeamMember(memberName);
    }
  };

  const settings = {
    dots: false,
    infinite: true,
    slidesToShow: 3,
    slidesToScroll: 3,
    autoplay: true,
    speed: 3000,
    autoplaySpeed: 1,
    cssEase: 'linear',
    slide: 'div',
    pauseOnHover: true,
    responsive: [
      {
        breakpoint: 1024,
        settings: {
          dots: false,
          infinite: true,
          slidesToShow: 3,
          slidesToScroll: 1,
          autoplay: true,
          speed: 3000,
          autoplaySpeed: 1,
          cssEase: 'linear',
          slide: 'div',
          pauseOnHover: true,
        },
      },
      {
        breakpoint: 800,
        settings: {
          dots: false,
          infinite: true,
          slidesToShow: 3,
          slidesToScroll: 1,
          autoplay: true,
          speed: 3000,
          autoplaySpeed: 1,
          cssEase: 'linear',
          slide: 'div',
          pauseOnHover: true,
        },
      },
      {
        breakpoint: 600,
        settings: {
          dots: false,
          infinite: true,
          slidesToShow: 2,
          slidesToScroll: 1,
          autoplay: true,
          speed: 3000,
          autoplaySpeed: 1,
          cssEase: 'linear',
          slide: 'div',
          pauseOnHover: true,
        },
      },
      {
        breakpoint: 480,
        settings: {
          dots: false,
          infinite: true,
          slidesToShow: 1,
          slidesToScroll: 1,
          autoplay: true,
          speed: 3000,
          autoplaySpeed: 1,
          cssEase: 'linear',
          slide: 'div',
          pauseOnHover: true,
        },
      },
    ],
  };

  return (
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
      responsive={settings.responsive}
    >
      <div
        className="image-container"
        onClick={() => {
          handleSelectTeamMember('Cory');
        }}
        onKeyDown={(e) => {
          handleKeyPress(e, 'Cory');
        }}
        role="button"
        tabIndex={0}
        aria-label="Cory's details"
      >
        <img src={Cory} alt="Cory" />
        <div className="overlay">
          <Typography id="overlay-name">Cory DeGuzman</Typography>
          <Typography id="overlay-title">
            Project Manager & Front-end Engineer
          </Typography>
        </div>
      </div>

      <div
        className="image-container"
        onClick={() => {
          handleSelectTeamMember('Angela');
        }}
        onKeyDown={(e) => {
          handleKeyPress(e, 'Angela');
        }}
        role="button"
        tabIndex={0}
        aria-label="Angela's details"
      >
        <img src={Angela} alt="Angela" />
        <div className="overlay">
          <Typography id="overlay-name">Angela Fisher</Typography>
          <Typography id="overlay-title">DevOps & Back-end Engineer</Typography>
        </div>
      </div>

      <div
        className="image-container"
        onClick={() => {
          handleSelectTeamMember('Brian');
        }}
        onKeyDown={(e) => {
          handleKeyPress(e, 'Brian');
        }}
        role="button"
        tabIndex={0}
        aria-label="Brian's details"
      >
        <img src={Brian} alt="Brian" />
        <div className="overlay">
          <Typography id="overlay-name">Brian La</Typography>
          <Typography id="overlay-title">Back-end Engineer</Typography>
        </div>
      </div>
    </Slider>
  );
}
