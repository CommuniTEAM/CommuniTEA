import { useState, useEffect } from 'react';
import { fetchTeas } from '../../Axios/teaService';

interface Tea {
  id: number
  name: string
  img_url: string
  description: string
  brew_time: number
  brew_temp: number
  published: boolean
}
export default function TestFetch(): JSX.Element {
  const [teas, setTeas] = useState<Tea[]>([]);

  useEffect(() => {
    fetchTeas(true)
      .then((teasData) => {
        setTeas(teasData);
      })
      .catch((error) => {
        console.error(error);
      });
  }, []);

  return (
    <div>
      <h1>Test Fetch</h1>
      <ul>
        {teas.map((tea) => (
          <li key={tea.id}>
            <h2>{tea.name}</h2>
            <img src={tea.img_url} alt={tea.name} />
            <p>{tea.description}</p>
            <p>
              Brew Time:
              {tea.brew_time}
            </p>
            <p>
              Brew Temp:
              {tea.brew_temp}
              °F
            </p>
          </li>
        ))}
      </ul>
    </div>
  );
}
