import { Divider, Typography } from '@mui/material';
import CoryCropped from '../../../assets/CoryCropped.jpg';
import Github from '../../../assets/GitHubBlack.png';
import LinkedIn from '../../../assets/LinkedIn.png';

import '../styles/TeamMemberDetailsStyles.css';

export default function CoryDetails(): JSX.Element {
  return (
    <div className="aboutUsBodyContainer">
      {/* Image */}
      <div className="aboutUsfirstHalf">
        <img src={CoryCropped} alt="Cory" />
      </div>

      {/* Full Name */}
      <div className="aboutUsSecondHalf">
        <div>
          <Typography variant="h2" sx={{ fontFamily: 'Montserrat' }}>
            Cory DeGuzman
          </Typography>
        </div>

        {/* Divider */}
        <Divider sx={{ bgcolor: 'black' }} />

        {/* Title / Position */}
        <div className="titlePositionContainer">
          <div>
            <Typography
              variant="h5"
              sx={{ fontFamily: 'Montserrat', fontStyle: 'italic' }}
            >
              Project Manager & Front-end Engineer
            </Typography>
          </div>

          {/* Social Links */}
          <div style={{ display: 'flex' }}>
            <a
              href="https://www.linkedin.com/in/cory-deguzman/"
              target="_blank"
              rel="noreferrer noopener"
            >
              <img src={LinkedIn} alt="LinkedIn" />
            </a>
            <a
              href="https://github.com/deguzmancory"
              target="_blank"
              rel="noreferrer noopener"
            >
              <img src={Github} alt="Github" />
            </a>
          </div>
        </div>

        {/* Team Member Description */}
        <div className="aboutUsDescription">
          <Typography
            variant="body1"
            sx={{ fontFamily: 'Montserrat', fontSize: '1.25em' }}
          >
            Cory DeGuzman is an adept Front End Developer and a proactive
            Software Engineer at Neon Blvd, whose proficiency in creating
            user-centric web applications is evident from his impactful work.
            <br />
            <br />
            With a solid foundation laid at DataHouse and JNK Enterprises, Cory
            has a track record of excelling in project management and
            development, particularly with React.js, JavaScript, and Material
            UI.
            <br />
            <br />
            He has effectively led teams and managed projects as the Project
            Manager and Lead Front-end Engineer for platforms like CommuniTEA
            and Pawsitive Health, utilizing tools like Jira and Figma to deliver
            engaging and innovative solutions.
            <br />
            <br />
            His commitment to lifelong learning is further demonstrated by his
            certifications from Hack Reactor in Software Engineering and AWS
            Cloud Practitioner, and he is poised to further his expertise by
            pursuing a Bachelor of Science in Software Engineering from Western
            Governor&apos;s University come February 2024.
            <br />
            <br />
            Based in Las Vegas, NV, Cory&apos;s enthusiasm for technology and
            innovation is the driving force behind his successful career, making
            him a valuable asset to his team and a visionary in the field of web
            development.
          </Typography>
        </div>
      </div>
    </div>
  );
}
