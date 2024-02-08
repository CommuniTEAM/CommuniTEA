import axios from 'axios';

const BASE_URL = `${import.meta.env.VITE_API_HOST}/teas/`;

interface Tea {
  id: number
  name: string
  img_url: string
  description: string
  brew_time: number
  brew_temp: number
  published: boolean
}

// GET request to fetch all teas
export const fetchTeas = async (published: boolean): Promise<Tea[]> => {
  try {
    const response = await axios.get(`${BASE_URL}${published}`);
    return response.data;
  } catch (error) {
    console.error('Error fetching teas: ', error);
    throw error;
  }
};

// POST request to create a new tea
export const createTea = async (tea: any): Promise<Tea[]> => {
  try {
    const response = await axios.post(BASE_URL, tea);
    return response.data;
  } catch (error) {
    console.error('Error creating tea: ', error);
    throw error;
  }
};

// PUT request to update a tea
export const updateTea = async (tea: any): Promise<Tea[]> => {
  try {
    const response = await axios.put(`${BASE_URL}${tea.id}`, tea);
    return response.data;
  } catch (error) {
    console.error('Error updating tea: ', error);
    throw error;
  }
};

// DELETE request to delete a tea
export const deleteTea = async (id: number): Promise<Tea[]> => {
  try {
    const response = await axios.delete(`${BASE_URL}${id}`);
    return response.data;
  } catch (error) {
    console.error('Error deleting tea: ', error);
    throw error;
  }
};
