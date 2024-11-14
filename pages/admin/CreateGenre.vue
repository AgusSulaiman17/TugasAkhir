<template>
  <div>
    <h1>Tambah Genre</h1>
    <form @submit.prevent="submitForm">
      <div>
        <label for="nama">Nama Genre:</label>
        <input type="text" v-model="genre.nama" required />
      </div>
      <button type="submit">Simpan</button>
    </form>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      genre: {
        nama: ''
      }
    };
  },
  methods: {
    async submitForm() {
      try {
        await axios.post('/api/genres', this.genre);
        this.$router.push('/admin/genres');
      } catch (error) {
        console.error('Error creating genre:', error);
      }
    }
  }
};
</script>
