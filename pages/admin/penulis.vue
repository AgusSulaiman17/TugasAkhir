<template>
  <div class="penulis mt-8">
    <AppNavbar />
    <div class="container mx-auto" style="max-width: 1200px;">
      <h1 class="text-center mb-4">Manage Authors</h1>

      <!-- Button to open Add Author modal -->
      <button @click="openAddPenulisModal" class="btn bg-ijotua mb-4">Add Author<b-icon-plus></b-icon-plus></button>

      <!-- Author Table -->
      <table class="table table-striped table-bordered text-center">
        <thead class="bg-ijomuda text-white">
          <tr>
            <th>#</th>
            <th>Author Name</th>
            <th>Biography</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(penulisItem, index) in penulis" :key="penulisItem.id_penulis">
            <td>{{ index + 1 }}</td>
            <td>{{ penulisItem.nama }}</td>
            <td v-if="penulisItem.biografi" class="biography-text">{{ penulisItem.biografi }}</td>
            <td>
              <button @click="editPenulis(penulisItem)" class="btn bg-kuning btn-sm"><b-icon-pencil></b-icon-pencil></button>
              <button @click="openDeleteConfirmation(penulisItem)" class="btn bg-merah btn-sm"><b-icon-trash></b-icon-trash></button>
            </td>
          </tr>
        </tbody>
      </table>

      <p v-if="penulis.length === 0" class="text-center mt-4">No authors available. Add a new one!</p>

      <!-- Add Author Modal -->
      <div v-if="isAddPenulisModalOpen" class="modal-overlay" @click="closeAddPenulisModal">
        <div class="modal-content" @click.stop>
          <h4 class="text-center mb-4">Add New Author</h4>
          <input v-model="newPenulis" placeholder="Enter author name" class="form-control mb-4" />
          <textarea v-model="newBiografi" placeholder="Enter author biography" class="form-control mb-4" rows="4"></textarea>
          <div class="d-flex justify-content-between">
            <button @click="createPenulis" class="btn bg-ijomuda">Add Author</button>
            <button @click="closeAddPenulisModal" class="btn bg-ijotua">Cancel</button>
          </div>
        </div>
      </div>

      <!-- Edit Author Modal -->
      <div v-if="isEditing" class="modal-overlay" @click="cancelEdit">
        <div class="modal-content" @click.stop>
          <h4 class="text-center mb-4">Edit Author</h4>
          <input v-model="editedPenulis" placeholder="Edit author name" class="form-control mb-4" />
          <textarea v-model="editedBiografi" placeholder="Edit author biography" class="form-control mb-4" rows="4"></textarea>
          <div class="d-flex justify-content-between">
            <button @click="updatePenulis" class="btn bg-ijomuda">Save Changes</button>
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
      v-if="showDeleteConfirmation"
      :isVisible="showDeleteConfirmation"
      :messageTitle="'Delete Author'"
      :messageBody="'Are you sure you want to delete this author?'"
      @close="cancelDelete"
    >
      <template #footer>
        <button @click="confirmDelete" class="btn btn-abu">Yes, Delete</button>
        <button @click="cancelDelete" class="btn btn-cancel">Cancel</button>
      </template>
    </NotificationModal>
  </div>
</template>

<script>
import { getPenulis, createPenulis, updatePenulis, deletePenulis } from '~/api/penulis';
import AppNavbar from '~/components/AppNavbar.vue';


export default {
  middleware: 'admin',
  name: 'Penulis',
  components: {
    AppNavbar,
    NotificationModal: () => import('~/components/NotificationModal.vue')
  },
  data() {
    return {
      penulis: [],
      newPenulis: '',
      newBiografi: '',
      editedPenulis: '',
      editedBiografi: '',
      isEditing: false,
      penulisToEdit: null,
      isAddPenulisModalOpen: false,
      showSuccessModal: false,
      successTitle: '',
      successMessage: '',
      showDeleteConfirmation: false,
      penulisToDelete: null,
    };
  },
  created() {
    this.loadPenulis();
  },
  methods: {
    async loadPenulis() {
      try {
        this.penulis = await getPenulis();
      } catch (error) {
        console.error('Error loading authors', error);
      }
    },
    async createPenulis() {
      if (!this.newPenulis.trim()) {
        alert('Author name cannot be empty!');
        return;
      }
      try {
        const penulis = { nama: this.newPenulis, biografi: this.newBiografi };
        const newPenulis = await createPenulis(penulis);
        this.penulis.push(newPenulis);
        this.resetAddPenulisForm();
        this.showSuccessNotification('Success', 'Author successfully added!');
      } catch (error) {
        console.error('Error creating author', error);
      }
    },
    async updatePenulis() {
      if (!this.editedPenulis.trim()) {
        alert('Author name cannot be empty!');
        return;
      }
      try {
        const updatedPenulis = { nama: this.editedPenulis, biografi: this.editedBiografi };
        const penulis = await updatePenulis(this.penulisToEdit.id_penulis, updatedPenulis);
        const index = this.penulis.findIndex((p) => p.id_penulis === penulis.id_penulis);
        this.penulis.splice(index, 1, penulis);
        this.cancelEdit();
        this.showSuccessNotification('Success', 'Author successfully updated!');
      } catch (error) {
        console.error('Error updating author', error);
      }
    },
    async deletePenulis(id) {
      try {
        await deletePenulis(id);
        this.penulis = this.penulis.filter((penulisItem) => penulisItem.id_penulis !== id);
        this.showSuccessNotification('Success', 'Author successfully deleted!');
      } catch (error) {
        console.error('Error deleting author', error);
      }
    },

    // Open confirmation modal
    openDeleteConfirmation(penulisItem) {
      this.penulisToDelete = penulisItem;
      this.showDeleteConfirmation = true;
    },

    // Cancel delete operation
    cancelDelete() {
      this.penulisToDelete = null;
      this.showDeleteConfirmation = false;
    },

    // Confirm delete operation
    confirmDelete() {
      if (this.penulisToDelete) {
        this.deletePenulis(this.penulisToDelete.id_penulis);
      }
      this.cancelDelete();
    },
    openAddPenulisModal() {
      this.isAddPenulisModalOpen = true;
    },
    closeAddPenulisModal() {
      this.isAddPenulisModalOpen = false;
      this.resetAddPenulisForm();
    },
    editPenulis(penulisItem) {
      this.penulisToEdit = penulisItem;
      this.editedPenulis = penulisItem.nama;
      this.editedBiografi = penulisItem.biografi;
      this.isEditing = true;
    },
    cancelEdit() {
      this.isEditing = false;
      this.editedPenulis = '';
      this.editedBiografi = '';
      this.penulisToEdit = null;
    },
    resetAddPenulisForm() {
      this.newPenulis = '';
      this.newBiografi = '';
      this.isAddPenulisModalOpen = false;
    },
    showSuccessNotification(title, message) {
      this.successTitle = title;
      this.successMessage = message;
      this.showSuccessModal = true;
    },
    closeSuccessModal() {
      this.showSuccessModal = false;
    }
  }
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

.btn-primary {
  background-color: #34c38f;
  color: white;
  border: none;
  padding: 8px 12px;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.btn-warning {
  background-color: #f1c40f;
  color: white;
  border: none;
  padding: 8px 12px;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.btn-danger {
  background-color: #e74c3c;
  color: white;
  border: none;
  padding: 8px 12px;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.btn-success {
  background-color: #28a745;
  color: white;
  border: none;
  padding: 8px 12px;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.btn-secondary {
  background-color: #6c757d;
  color: white;
  border: none;
  padding: 8px 12px;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.text-center {
  text-align: center;
}

.text-white {
  color: white;
}

.biography-text {
  text-align: justify;
}

.d-flex {
  display: flex;
}

.justify-content-between {
  justify-content: space-between;
}

.modal-overlay .modal-content .text-center {
  margin-bottom: 20px;
}

.modal-overlay .modal-content .form-control {
  margin-bottom: 12px;
}
</style>
