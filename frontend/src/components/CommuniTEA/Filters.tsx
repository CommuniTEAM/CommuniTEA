import {
  FormControl,
  FormControlLabel,
  FormGroup,
  FormLabel,
} from '@mui/material';
import EventCards from './Cards/EventCard';

import './styles/CommuniTeaPage.css';

export default function Filters(): JSX.Element {
  // TODO: Add functionality to the checkboxes

  return (
    <div className="filtersContainer">
      <div className="filtersChild">
        <FormControl sx={{ m: 3 }} component="fieldset" variant="standard">
          <FormLabel component="legend" sx={{ fontSize: '1vw' }}>
            Date
          </FormLabel>
          <FormGroup sx={{ paddingBottom: '1vw' }}>
            <FormControlLabel
              label="Today"
              control={<input type="checkbox" />}
              sx={{ paddingLeft: '1vw' }}
            />
            <FormControlLabel
              label="Tomorrow"
              control={<input type="checkbox" />}
              sx={{ paddingLeft: '1vw' }}
            />
            <FormControlLabel
              label="This Weekend"
              control={<input type="checkbox" />}
              sx={{ paddingLeft: '1vw' }}
            />
            <FormControlLabel
              label="Pick a Date"
              control={<input type="checkbox" />}
              sx={{ paddingLeft: '1vw' }}
            />
          </FormGroup>

          <FormLabel component="legend" sx={{ fontSize: '1vw' }}>
            Price
          </FormLabel>
          <FormGroup sx={{ paddingBottom: '1vw' }}>
            <FormControlLabel
              label="Free"
              control={<input type="checkbox" />}
              sx={{ paddingLeft: '1vw' }}
            />
            <FormControlLabel
              label="Paid"
              control={<input type="checkbox" />}
              sx={{ paddingLeft: '1vw' }}
            />
          </FormGroup>

          <FormLabel component="legend" sx={{ fontSize: '1vw' }}>
            Category
          </FormLabel>
          <FormGroup sx={{ paddingBottom: '1vw' }}>
            <FormControlLabel
              label="Seminar"
              control={<input type="checkbox" />}
              sx={{ paddingLeft: '1vw' }}
            />
            <FormControlLabel
              label="Class"
              control={<input type="checkbox" />}
              sx={{ paddingLeft: '1vw' }}
            />
            <FormControlLabel
              label="Social"
              control={<input type="checkbox" />}
              sx={{ paddingLeft: '1vw' }}
            />
            <FormControlLabel
              label="Tasting"
              control={<input type="checkbox" />}
              sx={{ paddingLeft: '1vw' }}
            />
            <FormControlLabel
              label="Pairing"
              control={<input type="checkbox" />}
              sx={{ paddingLeft: '1vw' }}
            />
          </FormGroup>
        </FormControl>
      </div>
      <div className="businessCards">
        <EventCards />
      </div>
    </div>
  );
}
