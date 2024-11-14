<template>
  <div>
    <h1>Edit Genre</h1>
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
  created() {
    this.fetchGenre();
  },
  methods: {
    async fetchGenre() {
      const { id } = this.$route.params;
      try {
        const response = await axios.get(`/api/genres/${id}`);
        this.genre = response.data;
      } catch (error) {
        console.error('Error fetching genre:', error);
      }
    },
    async submitForm() {
      const { id } = this.$route.params;
      try {
        await axios.put(`/api/genres/${id}`, this.genre);
        this.$router.push('/admin/genres');
      } catch (error) {
        console.error('Error updating genre:', error);
      }
    }
  }
};
</script>
