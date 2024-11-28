import axios from 'axios';

const apiUrl = 'http://localhost:8080/genres';  // Sesuaikan dengan URL backend Anda

// Fungsi untuk mendapatkan header otentikasi dengan token dari localStorage
const getAuthHeader = () => {
  const token = localStorage.getItem('token');  // Ambil token JWT yang disimpan di localStorage
  return token ? { Authorization: `Bearer ${token}` } : {};  // Mengembalikan header dengan token atau kosong jika tidak ada
};

// Mendapatkan genres
export const getGenres = async () => {
  try {
    const response = await axios.get(apiUrl, { headers: getAuthHeader() });  // Menambahkan header Authorization
    return response.data;
  } catch (error) {
    console.error('Error fetching genres:', error);
    throw error;
  }
};

// Membuat genre baru
export const createGenre = async (genre) => {
  try {
    const response = await axios.post(apiUrl, genre, { headers: getAuthHeader() });  // Menambahkan header Authorization
    return response.data;
  } catch (error) {
    console.error('Error creating genre:', error);
    throw error;
  }
};``

// Memperbarui genre
export const updateGenre = async (id, genre) => {
  try {
    const response = await axios.put(`${apiUrl}/${id}`, genre, { headers: getAuthHeader() });  // Menambahkan header Authorization
    return response.data;
  } catch (error) {
    console.error('Error updating genre:', error);
    throw error;
  }
};

// Menghapus genre
export const deleteGenre = async (id) => {
  try {
    await axios.delete(`${apiUrl}/${id}`, { headers: getAuthHeader() });  // Menambahkan header Authorization
  } catch (error) {
    console.error('Error deleting genre:', error);
    throw error;
  }
};
