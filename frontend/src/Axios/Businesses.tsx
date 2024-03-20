import axios from 'axios';

interface BusinessFormData {
  id: string;
  name: string;
  street_address: string;
  city: string;
  state: string;
  zipcode: string;
}

const BASE_URL = `${import.meta.env.VITE_API_HOST}`;

// Fetch all businesses
export const fetchBusinesses = async () => {
  try {
    const response = await axios.get(`${BASE_URL}/businesses`);
    return response.data;
  } catch (error) {
    console.error('Error fetching businesses', error);
    throw error;
  }
};

// Create new business
export const createBusiness = async (formData: BusinessFormData) => {
  try {
    const response = await axios.post(`${BASE_URL}/businesses`, formData);
    return response.data;
  } catch (error) {
    console.error('Error creating business', error);
    throw error;
  }
};

// Update business
export const updateBusiness = async (
  id: string,
  formData: BusinessFormData,
) => {
  try {
    const response = await axios.put(`${BASE_URL}/businesses/${id}`, formData);
    return response.data;
  } catch (error) {
    console.error('Error updating business', error);
    throw error;
  }
};
