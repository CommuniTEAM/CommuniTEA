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

const AddBusinessModal = ({
  open,
  handleClose,
  handleSubmit,
}: { open: boolean; handleClose: () => void; handleSubmit: () => void }) => {
  const [newBusiness, setNewBusiness] = useState<Business>({
    id: '',
    name: '',
    street_address: '',
    city: '',
    state: '',
    zipcode: '',
    business_owner_id: '',
  });

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

  return (
    <Modal open={open} onClose={handleClose}>
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
          <h3>Add New Business</h3>
          <TextField
            label="Name"
            value={newBusiness.name}
            onChange={(e) => handleInputChange(e, 'name')}
            fullWidth={true}
            style={{ marginBottom: '10px' }}
          />
          <TextField
            label="Street Address"
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
          />
          <TextField
            label="State"
            value={newBusiness.state}
            onChange={(e) => handleInputChange(e, 'state')}
            fullWidth={true}
            style={{ marginBottom: '10px' }}
          />
          <TextField
            label="Zipcode"
            value={newBusiness.zipcode}
            onChange={(e) => handleInputChange(e, 'zipcode')}
            fullWidth={true}
            style={{ marginBottom: '10px' }}
          />
          <TextField
            label="Business Owner ID"
            value={newBusiness.business_owner_id}
            onChange={(e) => handleInputChange(e, 'business_owner_id')}
            fullWidth={true}
            style={{ marginBottom: '10px' }}
          />
          <Button
            variant="contained"
            onClick={handleSubmit}
            style={{ marginTop: '10px' }}
          >
            Add
          </Button>
        </div>
      </div>
    </Modal>
  );
};

const UpdateBusinessModal = ({
  open,
  handleClose,
  handleSubmit,
  business,
}: {
  open: boolean;
  handleClose: () => void;
  handleSubmit: () => void;
  business: Business;
}) => {
  const [updatedBusiness, setUpdatedBusiness] = useState<Business>({
    ...business,
  });

  const handleInputChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>,
    field: keyof Business,
  ) => {
    const { value } = e.target;
    setUpdatedBusiness((prevBusiness) => ({
      ...prevBusiness,
      [field]: value,
    }));
  };

  useEffect(() => {
    setUpdatedBusiness({ ...business });
  }, [business]);

  return (
    <Modal open={open} onClose={handleClose}>
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
          <h3>Update Business</h3>
          <TextField
            label="Name"
            value={updatedBusiness.name}
            onChange={(e) => handleInputChange(e, 'name')}
            fullWidth={true}
            style={{ marginBottom: '10px' }}
          />
          <TextField
            label="Street Address"
            value={updatedBusiness.street_address}
            onChange={(e) => handleInputChange(e, 'street_address')}
            fullWidth={true}
            style={{ marginBottom: '10px' }}
          />
          <TextField
            label="City"
            value={updatedBusiness.city}
            onChange={(e) => handleInputChange(e, 'city')}
            fullWidth={true}
            style={{ marginBottom: '10px' }}
          />
          <TextField
            label="State"
            value={updatedBusiness.state}
            onChange={(e) => handleInputChange(e, 'state')}
            fullWidth={true}
            style={{ marginBottom: '10px' }}
          />
          <TextField
            label="Zipcode"
            value={updatedBusiness.zipcode}
            onChange={(e) => handleInputChange(e, 'zipcode')}
            fullWidth={true}
            style={{ marginBottom: '10px' }}
          />
          <TextField
            label="Business Owner ID"
            value={updatedBusiness.business_owner_id}
            onChange={(e) => handleInputChange(e, 'business_owner_id')}
            fullWidth={true}
            style={{ marginBottom: '10px' }}
          />
          <Button
            variant="contained"
            onClick={handleSubmit}
            style={{ marginTop: '10px' }}
          >
            Update
          </Button>
        </div>
      </div>
    </Modal>
  );
};

export default function BusinessTable(): JSX.Element {
  const [businesses, setBusinesses] = useState<Business[]>([]);
  const [openModal, setOpenModal] = useState<boolean>(false);
  const [modalMode, setModalMode] = useState<'add' | 'update'>('add');
  const [selectedBusiness, setSelectedBusiness] = useState<Business | null>(
    null,
  );

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

  const handleOpenModal = (mode: 'add' | 'update', business?: Business) => {
    setModalMode(mode);
    setSelectedBusiness(business ? { ...business } : null);
    setOpenModal(true);
  };

  const handleCloseModal = () => {
    setOpenModal(false);
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

  const handleSubmit = async () => {
    try {
      const headers = {
        'Content-Type': 'application/json',
      };

      if (modalMode === 'add') {
        // Add new business
        await axios.post(`${BASE_URL}/businesses`, selectedBusiness, {
          headers,
        });
        console.log('Business added successfully');
      } else if (modalMode === 'update' && selectedBusiness) {
        // Update existing business
        await axios.put(
          `${BASE_URL}/businesses/${selectedBusiness.id}`,
          selectedBusiness,
          { headers },
        );
        console.log('Business updated successfully');
      }

      setOpenModal(false); // Close the modal

      // Fetch updated business data
      const response = await axios.get(`${BASE_URL}/businesses`);
      setBusinesses(response.data.businesses);
    } catch (error) {
      console.error('Error updating or adding business', error);
    }
  };

  return (
    <div style={{ padding: '100px 50px 0 100px' }}>
      <h2>Businesses</h2>
      <Button variant="contained" onClick={() => handleOpenModal('add')}>
        Add Business
      </Button>
      {modalMode === 'add' && (
        <AddBusinessModal
          open={openModal}
          handleClose={handleCloseModal}
          handleSubmit={handleSubmit}
        />
      )}
      {modalMode === 'update' && selectedBusiness && (
        <UpdateBusinessModal
          open={openModal}
          handleClose={handleCloseModal}
          handleSubmit={handleSubmit}
          business={selectedBusiness}
        />
      )}
      <table>
        <thead>
          <tr>
            <th>Name</th>
            <th>Street Address</th>
            <th>City</th>
            <th>State</th>
            <th>Zipcode</th>
            <th>Business Owner ID</th>
            <th>Actions</th>
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
              <td>{business.business_owner_id}</td>
              <td>
                <Button
                  variant="contained"
                  onClick={() => handleOpenModal('update', business)}
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
