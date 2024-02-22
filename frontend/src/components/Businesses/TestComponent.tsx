import { Button, Modal, TextField } from '@mui/material';
import axios from 'axios';
import { useEffect, useState } from 'react';

interface Business {
  id: string;
  name: string;
  street_address: string;
  city: string;
  state: string;
  zipcode: string;
  business_owner_id: string;
}

const BASE_URL = `${import.meta.env.VITE_API_HOST}`;

export default function BusinessTable(): JSX.Element {
  const [businesses, setBusinesses] = useState<Business[]>([]);
  const [openModal, setOpenModal] = useState<boolean>(false);
  const [newBusiness, setNewBusiness] = useState<Business>({
    id: '',
    name: '',
    street_address: '',
    city: '',
    state: '',
    zipcode: '',
    business_owner_id: '',
  });

  useEffect(() => {
    const fetchBusinessData = async () => {
      try {
        const response = await axios.get(`${BASE_URL}/businesses`);
        setBusinesses(response.data.businesses);
      } catch (error) {
        console.error('Error fetching businesses', error);
      }
    };

    fetchBusinessData();
  }, []);

  const handleInputChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>,
    field: keyof Business,
  ) => {
    const { value } = e.target;
    setNewBusiness((prevBusiness) => ({
      ...prevBusiness,
      [field]: value,
    }));
  };

  const handleSubmit = async () => {
    try {
      const headers = {
        'Content-Type': 'application/json',
      };

      if (newBusiness.id) {
        // Update existing business
        await axios.put(
          `${BASE_URL}/businesses/${newBusiness.id}`,
          newBusiness,
          { headers },
        );
        console.log('Business updated successfully');
      } else {
        // Add new business
        await axios.post(`${BASE_URL}/businesses`, newBusiness, { headers });
        console.log('Business added successfully');
      }

      setOpenModal(false); // Close the modal

      // Fetch updated business data
      const response = await axios.get(`${BASE_URL}/businesses`);
      setBusinesses(response.data.businesses);
    } catch (error) {
      console.error('Error updating or adding business', error);
    }
  };

  const handleUpdate = (business: Business) => {
    setNewBusiness(business); // Populate the modal fields with the selected business
    setOpenModal(true); // Open the modal
  };

  const handleDelete = async (id: string) => {
    try {
      await axios.delete(`${BASE_URL}/businesses/${id}`);

      // Remove the deleted business from the list
      setBusinesses((prevBusinesses) =>
        prevBusinesses.filter((business) => business.id !== id),
      );

      console.log('Business deleted successfully');
    } catch (error) {
      console.error('Error deleting business', error);
    }
  };

  return (
    <div style={{ padding: '100px 50px 0 100px' }}>
      <h2>Businesses</h2>
      <Button variant="contained" onClick={() => setOpenModal(true)}>
        Add Business
      </Button>
      <Modal open={openModal} onClose={() => setOpenModal(false)}>
        <div
          style={{
            display: 'flex',
            justifyContent: 'center',
            alignItems: 'center',
            height: '100vh',
          }}
        >
          <div
            style={{
              backgroundColor: 'white',
              padding: '20px',
              borderRadius: '10px',
              boxShadow: '0px 4px 16px rgba(0, 0, 0, 0.1)',
            }}
          >
            <h3>{newBusiness.id ? 'Update Business' : 'Add New Business'}</h3>
            <TextField
              label="Name"
              value={newBusiness.name}
              onChange={(e) => handleInputChange(e, 'name')}
              fullWidth={true}
              style={{ marginBottom: '10px' }}
            />
            <TextField
              label="Address"
              value={newBusiness.street_address}
              onChange={(e) => handleInputChange(e, 'street_address')}
              fullWidth={true}
              style={{ marginBottom: '10px' }}
            />
            <TextField
              label="City"
              value={newBusiness.city}
              onChange={(e) => handleInputChange(e, 'city')}
              fullWidth={true}
              style={{ marginBottom: '10px' }}
              InputProps={{
                readOnly: true,
              }}
            />
            <TextField
              label="State"
              value={newBusiness.state}
              onChange={(e) => handleInputChange(e, 'state')}
              fullWidth={true}
              style={{ marginBottom: '10px' }}
              InputProps={{
                readOnly: true,
              }}
            />
            <TextField
              label="Zipcode"
              value={newBusiness.zipcode}
              onChange={(e) => handleInputChange(e, 'zipcode')}
              fullWidth={true}
              style={{ marginBottom: '10px' }}
              InputProps={{
                readOnly: true,
              }}
            />
            <TextField
              label="Business Owner"
              value={newBusiness.business_owner_id}
              onChange={(e) => handleInputChange(e, 'business_owner_id')}
              fullWidth={true}
              InputProps={{
                readOnly: true,
              }}
              style={{ marginBottom: '10px' }}
            />
            <Button
              variant="contained"
              onClick={handleSubmit}
              style={{ marginTop: '10px' }}
            >
              {newBusiness.id ? 'Update' : 'Add'}
            </Button>
          </div>
        </div>
      </Modal>
      <table>
        <thead>
          <tr>
            <th>Name</th>
            <th>Address</th>
            <th>City</th>
            <th>State</th>
            <th>Zipcode</th>
          </tr>
        </thead>
        <tbody>
          {businesses.map((business) => (
            <tr key={business.id}>
              <td>{business.name}</td>
              <td>{business.street_address}</td>
              <td>{business.city}</td>
              <td>{business.state}</td>
              <td>{business.zipcode}</td>
              <td>
                <Button
                  variant="contained"
                  onClick={() => handleUpdate(business)}
                >
                  Update
                </Button>
                <Button
                  variant="contained"
                  onClick={() => handleDelete(business.id)}
                  style={{ marginLeft: '10px' }}
                >
                  Delete
                </Button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}
