<template>
  <div class="genre mt-8">
    <AppNavbar />
    <div class="container">
      <h1 class="title">Manage Genres</h1>

      <!-- Tombol untuk membuka modal tambah genre -->
      <button @click="openAddGenreModal" class="btn btn-primary">Add Genre</button>

      <!-- Modal untuk menambah genre -->
      <div v-if="isAddGenreModalOpen" class="modal-overlay" @click="closeAddGenreModal">
        <div class="modal" @click.stop>
          <h2 class="mt-5">Add New Genre</h2>
          <input v-model="newGenre" placeholder="Enter genre name" class="input" required />
          <div class="modal-actions">
            <button @click="createGenre" class="btn btn-primary mt-2">Add Genre</button>
            <button @click="closeAddGenreModal" class="btn btn-cancel">Cancel</button>
          </div>
        </div>
      </div>

      <div v-if="genres.length" class="genre-list">
        <ul>
          <li v-for="genre in genres" :key="genre.id_genre" class="genre-item">
            <span>{{ genre.nama }}</span>
            <div class="genre-actions">
              <button @click="editGenre(genre.id_genre)" class="btn btn-edit">Edit</button>
              <button @click="deleteGenre(genre.id_genre)" class="btn btn-delete">Delete</button>
            </div>
          </li>
        </ul>
      </div>

      <div v-else class="empty-message">
        <p>No genres available. Add a new one!</p>
      </div>

      <div v-if="isEditing" class="edit-section">
        <h2>Edit Genre</h2>
        <input v-model="editedGenre" placeholder="Edit genre name" class="input" />
        <div class="edit-actions">
          <button @click="updateGenre" class="btn btn-update">Save Changes</button>
          <button @click="cancelEdit" class="btn btn-cancel">Cancel</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { getGenres, createGenre, updateGenre, deleteGenre } from '~/api/genres';
import AppNavbar from '~/components/AppNavbar.vue';

export default {
  middleware: 'admin',
  name: 'Genres',
  components: {
    AppNavbar
  },
  data() {
    return {
      genres: [],
      newGenre: '',
      editedGenre: '',
      isEditing: false,
      genreToEdit: null,
      isAddGenreModalOpen: false, // State to manage modal visibility
    };
  },
  created() {
    this.loadGenres();
  },
  methods: {
    async loadGenres() {
      try {
        this.genres = await getGenres();
      } catch (error) {
        console.error('Error loading genres', error);
      }
    },
    async createGenre() {
      if (!this.newGenre.trim()) {
        alert('Genre name cannot be empty!');
        return; // Berhenti jika input kosong
      }

      try {
        const genre = { nama: this.newGenre };
        const newGenre = await createGenre(genre);
        this.genres.push(newGenre);
        this.newGenre = '';
        this.closeAddGenreModal();
      } catch (error) {
        console.error('Error creating genre', error);
      }
    },
    openAddGenreModal() {
      this.isAddGenreModalOpen = true;
    },
    closeAddGenreModal() {
      this.isAddGenreModalOpen = false;
      this.newGenre = ''; // Reset the input when closing the modal
    },
    editGenre(id) {
      this.genreToEdit = this.genres.find((genre) => genre.id_genre === id);
      this.editedGenre = this.genreToEdit.nama;
      this.isEditing = true;
    },
    async updateGenre() {
      try {
        const updatedGenre = { nama: this.editedGenre };
        const genre = await updateGenre(this.genreToEdit.id_genre, updatedGenre);
        const index = this.genres.findIndex((g) => g.id_genre === genre.id_genre);
        this.genres.splice(index, 1, genre);
        this.isEditing = false;
        this.editedGenre = '';
        this.genreToEdit = null;
      } catch (error) {
        console.error('Error updating genre', error);
      }
    },
    cancelEdit() {
      this.isEditing = false;
      this.editedGenre = '';
      this.genreToEdit = null;
    },
    async deleteGenre(id) {
      try {
        await deleteGenre(id);
        this.genres = this.genres.filter((genre) => genre.id_genre !== id);
      } catch (error) {
        console.error('Error deleting genre', error);
      }
    },
  },
};
</script>

<style scoped>
.container {
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
  background-color: #f9f9f9;
  border-radius: 8px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

.title {
  font-size: 2rem;
  font-weight: bold;
  text-align: center;
  margin-bottom: 20px;
  color: #333;
}

.btn {
  padding: 10px 20px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-weight: bold;
  font-size: 1rem;
  transition: background-color 0.3s, transform 0.2s;
}

.btn-primary {
  background-color: #4CAF50;
  color: white;
}

.btn-primary:hover {
  background-color: #45a049;
  transform: scale(1.05);
}

.genre-list {
  margin-top: 20px;
}

.genre-item {
  background-color: white;
  padding: 15px;
  border-radius: 6px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 10px;
}

.genre-actions {
  display: flex;
  gap: 15px;
}

.btn-edit {
  background-color: #2196F3;
  color: white;
}

.btn-edit:hover {
  background-color: #0b7dda;
}

.btn-delete {
  background-color: #f44336;
  color: white;
}

.btn-delete:hover {
  background-color: #d32f2f;
}

.empty-message {
  text-align: center;
  color: #666;
}

.edit-section {
  margin-top: 30px;
  padding: 20px;
  background-color: white;
  border-radius: 6px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.modal {
  display: block;
  background-color: white;
  padding: 20px;
  border-radius: 6px;
  width: 400px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

.modal-actions {
  display: flex;
  justify-content: space-between;
  gap: 10px;
}

.btn-cancel {
  background-color: #9e9e9e;
  color: white;
}

.btn-cancel:hover {
  background-color: #757575;
}
</style>
