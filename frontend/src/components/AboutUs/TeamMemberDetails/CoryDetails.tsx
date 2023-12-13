import { Divider, Typography } from '@mui/material';
import CoryCropped from '../../../assets/CoryCropped.jpg';
import Github from '../../../assets/GitHubBlack.png';
import LinkedIn from '../../../assets/LinkedIn.png';

export default function CoryDetails(): JSX.Element {
  const responsiveFontSize = 'calc(0.5vw + 0.5em)';

  return (
    <div style={{ display: 'flex' }}>
      <div style={{ flexGrow: 1, maxWidth: '100vw', width: '50%' }}>
        <img src={CoryCropped} alt="Cory" style={{ maxWidth: '100%' }} />
      </div>

      <div
        style={{
          flexGrow: 1,
          paddingTop: '1vh',
          paddingLeft: '2vw',
          paddingRight: '2vw',
          width: '50%',
        }}
      >
        <div>
          <Typography variant="h2" sx={{ fontFamily: 'Montserrat' }}>
            Cory DeGuzman
          </Typography>
        </div>

        <Divider sx={{ bgcolor: 'black' }} />

        <div
          style={{
            display: 'flex',
            justifyContent: 'space-between',
            paddingBottom: 25,
          }}
        >
          <div>
            <Typography
              variant="h5"
              sx={{ fontFamily: 'Montserrat', fontStyle: 'italic' }}
            >
              Project Manager & Front-end Developer
            </Typography>
          </div>

          <div>
            <div style={{ display: 'flex', alignItems: 'center' }}>
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
        </div>

        <div>
          <Typography
            variant="body1"
            sx={{ fontFamily: 'Montserrat', fontSize: responsiveFontSize }}
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
