import axios from 'axios';

const apiUrl = 'http://localhost:8080'; // Sesuaikan dengan URL backend Anda

const getAuthHeader = () => {
  const token = localStorage.getItem('token');
  return token ? { Authorization: `Bearer ${token}` } : {};
};

export const getAllPeminjaman = async () => {
  const response = await axios.get(`${apiUrl}/peminjaman`, { headers: getAuthHeader() });
  return response.data;
};

export const createPeminjaman = async (data) => {
  try {
    console.log('Mengirim data peminjaman:', data); // Log data yang dikirim
    const response = await axios.post(`${apiUrl}/peminjaman`, data, { headers: getAuthHeader() });
    console.log('Respons dari server:', response.data); // Log respons server
    return response.data;
  } catch (error) {
    if (error.response) {
      // Server memberikan respons error
      console.error('Error dari server:', error.response.data);
    } else if (error.request) {
      // Tidak ada respons dari server
      console.error('Tidak ada respons dari server:', error.request);
    } else {
      // Error lain
      console.error('Terjadi kesalahan:', error.message);
    }
    throw error; // Lempar error agar bisa ditangani di frontend
  }
};

export const returnBook = async (id) => {
  const response = await axios.post(`${apiUrl}/peminjaman/${id}/kembalikan`, {}, { headers: getAuthHeader() });
  return response.data;
};

export const deletePeminjaman = async (id) => {
  const response = await axios.delete(`${apiUrl}/peminjaman/${id}`, { headers: getAuthHeader() });
  return response.data;
};
