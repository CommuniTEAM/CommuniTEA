import Box from '@mui/material/Box';
import Rating from '@mui/material/Rating';
import Typography from '@mui/material/Typography';
import * as React from 'react';

export default function TestRating() {
  const [value, setValue] = React.useState<number | null>(2);

  return (
    <Box
      sx={{
        '& > legend': { mt: 2 },
      }}
    >
      <Typography component="legend">Controlled</Typography>
      <Rating
        name="simple-controlled"
        value={value}
        onChange={(event, newValue) => {
          setValue(newValue);
          console.log(event);
        }}
      />
      <Typography component="legend" sx={{ color: 'red' }}>
        Red Text
      </Typography>
      <Rating name="read-only" value={value} readOnly={true} />
      <Typography component="legend">Dont work</Typography>
      <Rating name="disabled" value={value} disabled={true} />
      <Typography component="legend">No rating given</Typography>
      <Rating name="no-value" value={null} />
    </Box>
  );
}
