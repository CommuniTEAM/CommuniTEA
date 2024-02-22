import { Button, TextField, Typography } from '@mui/material';
import { Form, Formik } from 'formik';
import * as Yup from 'yup';

import '../styles/RsvpFormStyles.css';

// Define the validation schema
const validationSchema = Yup.object().shape({
  first_name: Yup.string().required('First Name is Required'),
  last_name: Yup.string().required('Last Name is Required'),
  email: Yup.string().email('Invalid email').required('Required'),
  phone: Yup.string().required('Phone Number is Required'),
  num_attendees: Yup.number().required('Number of Attendees is Required'),
});

export default function RsvpForm(): JSX.Element {
  return (
    <div className="formFieldsContainer">
      <Typography
        variant="h4"
        sx={{
          color: 'white',
          fontFamily: 'Montserrat',
          textAlign: 'center',
          marginBottom: '20px',
        }}
      >
        RSVP FOR THIS EVENT!
      </Typography>
      <Formik
        initialValues={{
          first_name: '',
          last_name: '',
          email: '',
          phone: '',
          num_attendees: '',
        }}
        validationSchema={validationSchema}
        onSubmit={(values, { setSubmitting }) => {
          console.log(values);
          setSubmitting(false);
        }}
      >
        {({ isSubmitting, errors, touched, handleChange, handleBlur }) => (
          <Form>
            <div className="inputFieldsContainer">
              <TextField
                type="text"
                name="first_name"
                label="First Name"
                placeholder="First Name"
                variant="outlined"
                fullWidth={true}
                margin="dense"
                error={!!(touched.first_name && errors.first_name)}
                onChange={handleChange}
                onBlur={handleBlur}
                sx={{
                  backgroundColor: 'white',
                  borderRadius: '10px',
                }}
              />

              <TextField
                type="text"
                name="last_name"
                label="Last Name"
                placeholder="Last Name"
                variant="outlined"
                fullWidth={true}
                margin="dense"
                error={Boolean(touched.last_name && errors.last_name)}
                onChange={handleChange}
                onBlur={handleBlur}
                sx={{
                  backgroundColor: 'white',
                  borderRadius: '10px',
                }}
              />

              <TextField
                type="email"
                name="email"
                label="Email"
                placeholder="Email"
                variant="outlined"
                fullWidth={true}
                margin="dense"
                error={Boolean(touched.email && errors.email)}
                onChange={handleChange}
                onBlur={handleBlur}
                sx={{
                  backgroundColor: 'white',
                  borderRadius: '10px',
                }}
              />

              <TextField
                type="phone"
                name="phone"
                label="Phone Number"
                placeholder="Phone Number"
                variant="outlined"
                fullWidth={true}
                margin="dense"
                error={Boolean(touched.phone && errors.phone)}
                onChange={handleChange}
                onBlur={handleBlur}
                sx={{
                  backgroundColor: 'white',
                  borderRadius: '10px',
                }}
              />

              <TextField
                type="number"
                name="num_attendees"
                label="Number of Attendees"
                placeholder="Number of Attendees"
                variant="outlined"
                fullWidth={true}
                margin="dense"
                error={Boolean(touched.num_attendees && errors.num_attendees)}
                onChange={handleChange}
                onBlur={handleBlur}
                sx={{
                  backgroundColor: 'white',
                  marginBottom: '30px',
                  borderRadius: '10px',
                }}
              />

              <Button
                type="submit"
                disabled={isSubmitting}
                variant="contained"
                sx={{
                  width: '100%',
                  borderRadius: '10px',
                  backgroundColor: '#87CEEB',
                  color: 'black',
                  fontFamily: 'Montserrat',
                  marginBottom: '5%',
                  padding: '20px 0',
                }}
              >
                JOIN
              </Button>
            </div>
          </Form>
        )}
      </Formik>
    </div>
  );
}
