import axios from 'axios';

const apiUrl = 'http://localhost:8080/penulis';  // Sesuaikan dengan URL backend Anda

// Fungsi untuk mendapatkan header otentikasi dengan token dari localStorage
const getAuthHeader = () => {
  const token = localStorage.getItem('token');  // Ambil token JWT yang disimpan di localStorage
  return token ? { Authorization: `Bearer ${token}` } : {};  // Mengembalikan header dengan token atau kosong jika tidak ada
};

// Mendapatkan penulis
export const getPenulis = async () => {
  try {
    const response = await axios.get(apiUrl, { headers: getAuthHeader() });  // Menambahkan header Authorization
    return response.data;
  } catch (error) {
    console.error('Error fetching penulis:', error);
    throw error;
  }
};

// Membuat penulis baru
export const createPenulis = async (penulis) => {
  try {
    const response = await axios.post(apiUrl, penulis, { headers: getAuthHeader() });  // Menambahkan header Authorization
    return response.data;
  } catch (error) {
    console.error('Error creating penulis:', error);
    throw error;
  }
};

// Memperbarui penulis
export const updatePenulis = async (id, penulis) => {
  try {
    const response = await axios.put(`${apiUrl}/${id}`, penulis, { headers: getAuthHeader() });  // Menambahkan header Authorization
    return response.data;
  } catch (error) {
    console.error('Error updating penulis:', error);
    throw error;
  }
};

// Menghapus penulis
export const deletePenulis = async (id) => {
  try {
    await axios.delete(`${apiUrl}/${id}`, { headers: getAuthHeader() });  // Menambahkan header Authorization
  } catch (error) {
    console.error('Error deleting penulis:', error);
    throw error;
  }
};
