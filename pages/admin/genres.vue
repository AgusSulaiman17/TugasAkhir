<template>
  <div class="genre mt-8">
    <AppNavbar />
    <div class="container mx-auto" style="max-width: 1200px;">
      <h1 class="text-center mb-4">Manage Genres</h1>

      <!-- Button to open Add Genre modal -->
      <button @click="openAddGenreModal" class="btn bg-ijotua mb-4">Add Genre <b-icon-plus></b-icon-plus></button>

      <!-- Genre Table -->
      <table class="table table-striped table-bordered text-center">
        <thead class="bg-ijomuda text-white">
          <tr>
            <th>#</th>
            <th>Genre Name</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(genre, index) in genres" :key="genre.id_genre">
            <td>{{ index + 1 }}</td>
            <td>{{ genre.nama }}</td>
            <td>
              <button @click="editGenre(genre.id_genre)" class="btn bg-kuning btn-sm text-white"><b-icon-pencil></b-icon-pencil></button>
              <button @click="confirmDeleteGenre(genre.id_genre)" class="btn bg-merah btn-sm"><b-icon-trash></b-icon-trash></button>
            </td>
          </tr>
        </tbody>
      </table>

      <p v-if="genres.length === 0" class="text-center mt-4">No genres available. Add a new one!</p>

      <!-- Add Genre Modal -->
      <div v-if="isAddGenreModalOpen" class="modal-overlay" @click="closeAddGenreModal">
        <div class="modal-content" @click.stop>
          <h4 class="text-center mb-4">Add New Genre</h4>
          <input v-model="newGenre" placeholder="Enter genre name" class="form-control mb-4" />
          <div class="d-flex justify-content-between">
            <button @click="createGenre" class="btn bg-ijomuda">Add Genre</button>
            <button @click="closeAddGenreModal" class="btn bg-ijotua">Cancel</button>
          </div>
        </div>
      </div>

      <!-- Edit Genre Modal -->
      <div v-if="isEditing" class="modal-overlay" @click="cancelEdit">
        <div class="modal-content" @click.stop>
          <h4 class="text-center mb-4">Edit Genre</h4>
          <input v-model="editedGenre" placeholder="Edit genre name" class="form-control mb-4" />
          <div class="d-flex justify-content-between">
            <button @click="updateGenre" class="btn bg-ijomuda">Save Changes</button>
            <button @click="cancelEdit" class="btn bg-ijotua">Cancel</button>
          </div>
        </div>
      </div>

    </div>
        <!-- Success Notification Modal -->
        <NotificationModal
      v-if="showSuccessModal"
      :isVisible="showSuccessModal"
      :messageTitle="successTitle"
      :messageBody="successMessage"
      @close="closeSuccessModal"
    />

    <!-- Delete Confirmation Modal -->
    <NotificationModal
      v-if="showDeleteConfirmModal"
      :isVisible="showDeleteConfirmModal"
      :messageTitle="'Delete Confirmation'"
      :messageBody="'Are you sure you want to delete this genre?'"
      @close="closeDeleteConfirmModal"
    >
      <template #footer>
        <button @click="deleteGenre(confirmGenreId)" class="btn btn-delete">Yes, Delete</button>
        <button @click="closeDeleteConfirmModal" class="btn btn-cancel">Cancel</button>
      </template>
    </NotificationModal>
  </div>
</template>


<script>
import { getGenres, createGenre, updateGenre, deleteGenre } from '~/api/genres';
import AppNavbar from '~/components/AppNavbar.vue';
import NotificationModal from '~/components/NotificationModal.vue';

export default {
  middleware: 'admin',
  name: 'Genres',
  components: {
    AppNavbar,
    NotificationModal
  },
  data() {
    return {
      genres: [],
      newGenre: '',
      editedGenre: '',
      isEditing: false,
      genreToEdit: null,
      isAddGenreModalOpen: false,
      showSuccessModal: false,
      showDeleteConfirmModal: false,
      confirmGenreId: null,
      successTitle: '',
      successMessage: ''
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
        return; // Stop if input is empty
      }

      try {
        const genre = { nama: this.newGenre };
        const newGenre = await createGenre(genre);
        this.genres.push(newGenre);
        this.newGenre = '';
        this.closeAddGenreModal();
        this.showSuccessNotification('Success', 'Genre successfully added!');
      } catch (error) {
        console.error('Error creating genre', error);
      }
    },
    openAddGenreModal() {
      this.isAddGenreModalOpen = true;
    },
    closeAddGenreModal() {
      this.isAddGenreModalOpen = false;
      this.newGenre = ''; // Reset input when closing modal
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
        this.showSuccessNotification('Success', 'Genre successfully updated!');
      } catch (error) {
        console.error('Error updating genre', error);
      }
    },
    cancelEdit() {
      this.isEditing = false;
      this.editedGenre = '';
      this.genreToEdit = null;
    },
    confirmDeleteGenre(id) {
      this.confirmGenreId = id;
      this.showDeleteConfirmModal = true;
    },
    closeDeleteConfirmModal() {
      this.showDeleteConfirmModal = false;
      this.confirmGenreId = null;
    },
    async deleteGenre(id) {
      try {
        await deleteGenre(id);
        this.genres = this.genres.filter((genre) => genre.id_genre !== id);
        this.closeDeleteConfirmModal();
        this.showSuccessNotification('Success', 'Genre successfully deleted!');
      } catch (error) {
        console.error('Error deleting genre', error);
      }
    },
    showSuccessNotification(title, message) {
      this.successTitle = title;
      this.successMessage = message;
      this.showSuccessModal = true;
    },
    closeSuccessModal() {
      this.showSuccessModal = false;
      this.successTitle = '';
      this.successMessage = '';
    }
  },
};
</script>

<style scoped>
.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
  background-color: #f9f9f9;
  border-radius: 8px;
}

.table {
  margin-top: 20px;
}

.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.6);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 9999;
}

.modal-content {
  background: white;
  border-radius: 8px;
  padding: 20px;
  width: 400px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
}

.modal-content h4 {
  font-size: 1.5rem;
  font-weight: bold;
  text-align: center;
  margin-bottom: 20px;
}

</style>
