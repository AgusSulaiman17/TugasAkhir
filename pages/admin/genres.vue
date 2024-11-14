<template>
  <div>
    <h1>Genres</h1>
    <form @submit.prevent="createGenre">
      <input v-model="newGenre" placeholder="Add a new genre" />
      <button type="submit">Create Genre</button>
    </form>

    <div v-if="genres.length">
      <ul>
        <li v-for="genre in genres" :key="genre.id_genre">
          {{ genre.nama }}
          <button @click="editGenre(genre.id_genre)">Edit</button>
          <button @click="deleteGenre(genre.id_genre)">Delete</button>
        </li>
      </ul>
    </div>

    <div v-else>
      <p>No genres available</p>
    </div>

    <div v-if="isEditing">
      <h2>Edit Genre</h2>
      <input v-model="editedGenre" placeholder="Edit genre name" />
      <button @click="updateGenre">Update</button>
      <button @click="cancelEdit">Cancel</button>
    </div>
  </div>
</template>

<script>
import { getGenres, createGenre, updateGenre, deleteGenre } from '~/api/genres'

export default {
  data() {
    return {
      genres: [],
      newGenre: '',
      editedGenre: '',
      isEditing: false,
      genreToEdit: null
    }
  },
  created() {
    this.loadGenres(); // Memanggil fungsi async di sini
  },
  methods: {
    // Memuat genre dari API
    async loadGenres() {
      try {
        this.genres = await getGenres();
      } catch (error) {
        console.error('Error loading genres', error);
      }
    },
    // Menambah genre baru
    async createGenre() {
      try {
        const genre = { nama: this.newGenre };
        const newGenre = await createGenre(genre);
        this.genres.push(newGenre);
        this.newGenre = '';
      } catch (error) {
        console.error('Error creating genre', error);
      }
    },
    // Mengedit genre yang ada
    editGenre(id) {
      this.genreToEdit = this.genres.find(genre => genre.id_genre === id);
      this.editedGenre = this.genreToEdit.nama;
      this.isEditing = true;
    },
    // Memperbarui genre
    async updateGenre() {
      try {
        const updatedGenre = { nama: this.editedGenre };
        const genre = await updateGenre(this.genreToEdit.id_genre, updatedGenre);
        const index = this.genres.findIndex(g => g.id_genre === genre.id_genre);
        this.genres.splice(index, 1, genre);
        this.isEditing = false;
        this.editedGenre = '';
        this.genreToEdit = null;
      } catch (error) {
        console.error('Error updating genre', error);
      }
    },
    // Membatalkan edit genre
    cancelEdit() {
      this.isEditing = false;
      this.editedGenre = '';
      this.genreToEdit = null;
    },
    // Menghapus genre
    async deleteGenre(id) {
      try {
        await deleteGenre(id);
        this.genres = this.genres.filter(genre => genre.id_genre !== id);
      } catch (error) {
        console.error('Error deleting genre', error);
      }
    }
  }
}
</script>
