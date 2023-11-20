import { TextField } from '@mui/material'
import MapPinIcon from '../../assets/MapPinIcon.png'

export default function LocationSearch (): JSX.Element {
  // TODO: Add functionality to the searchbox
  return (
    <div style={{ maxWidth: '50%', backgroundColor: '#FFFFF0', marginTop: '1vw' }}>
      <div style={{ display: 'flex', alignItems: 'center' }}>
        <div style={{ display: 'flex', alignItems: 'center', flexGrow: 2, marginLeft: '1vw' }}>
          <img
            src={MapPinIcon}
            alt="Map Pin Icon"
            style={{ width: '25px', marginRight: '1vw', marginLeft: '1vw' }}
          />
          <TextField
            id="location-search"
            label="Enter your location"
            variant="standard"
            sx={{ width: '100%' }}
          />
        </div>
        {/* <div style={{ flexGrow: 1, textAlign: 'center' }}>
          <Button
            variant="contained"
            sx={{ backgroundColor: '#D2B48C', color: '#000000', width: '80%' }}
          >
            Search
          </Button>
        </div> */}
      </div>
    </div>
  )
}
