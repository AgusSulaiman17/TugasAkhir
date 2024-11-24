import axios from 'axios';

const apiUrl = 'http://localhost:8080/buku';

const getAuthHeader = () => {
  const token = localStorage.getItem('token');
  return token ? { Authorization: `Bearer ${token}` } : {};
};

export const getBukuList = async () => {
  try {
    const response = await axios.get(apiUrl, { headers: getAuthHeader() });
    return response.data;
  } catch (error) {
    console.error('Error fetching buku list:', error);
    throw error;
  }
};

export const getBukuById = async (id) => {
  console.log('ID yang diterima:', id);
  if (!id) {
    throw new Error('ID buku tidak diberikan');
  }

  try {
    const response = await axios.get(`${apiUrl}/${id}`, { headers: getAuthHeader() });
    return response.data;
  } catch (error) {
    console.error('Error fetching buku by ID:', error);
    throw error;
  }
};

export const createBuku = async (buku) => {
  const formData = new FormData();
  formData.append('judul', buku.judul);
  formData.append('id_penulis', buku.id_penulis);
  formData.append('id_genre', buku.id_genre);
  formData.append('deskripsi', buku.deskripsi);
  formData.append('jumlah', buku.jumlah);
  if (buku.gambar) {
    formData.append('gambar', buku.gambar);
  }

  try {
    const response = await axios.post(apiUrl, formData, { headers: getAuthHeader() });
    return response.data;
  } catch (error) {
    console.error('Error creating buku:', error);
    throw error;
  }
};

export const updateBuku = async (buku) => {
  const formData = new FormData();
  formData.append('judul', buku.judul);
  formData.append('id_penulis', buku.id_penulis);
  formData.append('id_genre', buku.id_genre);
  formData.append('deskripsi', buku.deskripsi);
  formData.append('jumlah', buku.jumlah);
  if (buku.gambar) {
    formData.append('gambar', buku.gambar);
  }

  try {
    const response = await axios.put(`${apiUrl}/${buku.id_buku}`, formData, { headers: getAuthHeader() });
    return response.data;
  } catch (error) {
    console.error('Error updating buku:', error);
    throw error;
  }
};

export const deleteBuku = async (id_buku) => {
  try {
    await axios.delete(`${apiUrl}/${id_buku}`, { headers: getAuthHeader() });
  } catch (error) {
    console.error('Error deleting buku:', error);
    throw error;
  }
};
