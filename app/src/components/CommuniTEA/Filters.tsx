import {
  FormControl,
  FormControlLabel,
  FormGroup,
  FormLabel,
} from '@mui/material';
import EventCards from './EventCard';

export default function Filters(): JSX.Element {
  // TODO: Add functionality to the checkboxes

  return (
    <div style={{ display: 'flex', width: '50%' }}>
      <div
        style={{
          paddingLeft: '1vw',
          width: '30%',
        }}
      >
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
      <div
        style={{
          display: 'flex',
          flexDirection: 'column',
          width: '70%',
          justifyContent: 'center',
          alignItems: 'center',
        }}
      >
        <EventCards />
      </div>
    </div>
  );
}
