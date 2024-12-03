// api/user.js

import axios from 'axios';

const API_URL = 'http://localhost:8080'; // Ganti dengan URL API backend Anda

// Fungsi untuk mendapatkan header Authorization
const getAuthHeader = () => {
  const token = localStorage.getItem('token');
  return token ? { Authorization: `Bearer ${token}` } : {};
};

// Fungsi untuk membuat pengguna baru
export const createUser = async (userData) => {
  try {
    const response = await axios.post(`${API_URL}/user`, userData, {
      headers: getAuthHeader(),
    });
    return response.data;
  } catch (error) {
    throw new Error('Error creating user');
  }
};

// Fungsi untuk mengambil data pengguna berdasarkan ID
export const getUser = async (id) => {
  try {
    const response = await axios.get(`${API_URL}/user/${id}`, {
      headers: getAuthHeader(), // Menggunakan getAuthHeader untuk header Authorization
    });
    return response.data;
  } catch (error) {
    throw new Error('Error fetching user');
  }
};

// Fungsi untuk mengambil semua data pengguna
export const getUsers = async () => {
  try {
    const response = await axios.get(`${API_URL}/users`, {
      headers: getAuthHeader(), // Menggunakan getAuthHeader untuk header Authorization
    });
    return response.data;
  } catch (error) {
    throw new Error('Error fetching users');
  }
};

// Fungsi untuk memperbarui data pengguna
export const updateUser = async (id, userData) => {
  try {
    const response = await axios.put(`${API_URL}/user/${id}`, userData, {
      headers: getAuthHeader(), // Menggunakan getAuthHeader untuk header Authorization
    });
    return response.data;
  } catch (error) {
    throw new Error('Error updating user');
  }
};

// Fungsi untuk menghapus pengguna
export const deleteUser = async (id) => {
  try {
    const response = await axios.delete(`${API_URL}/user/${id}`, {
      headers: getAuthHeader(), // Menggunakan getAuthHeader untuk header Authorization
    });
    return response.data;
  } catch (error) {
    throw new Error('Error deleting user');
  }
};
