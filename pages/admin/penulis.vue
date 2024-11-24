<template>
  <div class="penulis mt-8">
    <AppNavbar />
    <div class="container">
      <h1 class="title">Manage Authors</h1>

      <!-- Tombol untuk membuka modal tambah penulis -->
      <button @click="openAddPenulisModal" class="btn btn-primary">Add Author</button>

      <!-- Modal untuk menambah penulis -->
      <div v-if="isAddPenulisModalOpen" class="modal-overlay" @click="closeAddPenulisModal">
        <div class="modal" @click.stop>
          <h2 class="mt-5">Add New Author</h2>
          <input v-model="newPenulis" placeholder="Enter author name" class="input" />
          <textarea v-model="newBiografi" placeholder="Enter author biography" class="input" rows="4"></textarea>
          <div class="modal-actions">
            <button @click="createPenulis" class="btn btn-abu mt-2">Add Author</button>
            <button @click="closeAddPenulisModal" class="btn btn-cancel">Cancel</button>
          </div>
        </div>
      </div>

      <!-- Modal untuk mengedit penulis -->
      <div v-if="isEditing" class="modal-overlay" @click="cancelEdit">
        <div class="modal" @click.stop>
          <h2 class="mt-5">Edit Author</h2>
          <input v-model="editedPenulis" placeholder="Edit author name" class="input" />
          <textarea v-model="editedBiografi" placeholder="Edit author biography" class="input" rows="4"></textarea>
          <div class="modal-actions">
            <button @click="updatePenulis" class="btn btn-abu mt-2">Save Changes</button>
            <button @click="cancelEdit" class="btn btn-cancel">Cancel</button>
          </div>
        </div>
      </div>

      <div v-if="penulis.length" class="penulis-list">
        <ul>
          <li v-for="penulisItem in penulis" :key="penulisItem.id_penulis" class="penulis-item">
            <span>{{ penulisItem.nama }}</span>
            <span v-if="penulisItem.biografi" class="biography-text">{{ penulisItem.biografi }}</span>
            <div class="penulis-actions">
              <button @click="editPenulis(penulisItem.id_penulis)" class="btn btn-edit">Edit</button>
              <button @click="deletePenulis(penulisItem.id_penulis)" class="btn btn-delete">Delete</button>
            </div>
          </li>
        </ul>
      </div>

      <div v-else class="empty-message">
        <p>No authors available. Add a new one!</p>
      </div>

      <!-- Edit Section for managing penulis -->
      <div v-if="isEditing" class="edit-section">
        <h2>Edit Author</h2>
        <input v-model="editedPenulis" placeholder="Edit author name" class="input" />
        <textarea v-model="editedBiografi" placeholder="Edit author biography" class="input" rows="4"></textarea>
        <div class="edit-actions">
          <button @click="updatePenulis" class="btn btn-update">Save Changes</button>
          <button @click="cancelEdit" class="btn btn-cancel">Cancel</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { getPenulis, createPenulis, updatePenulis, deletePenulis } from '~/api/penulis';
import AppNavbar from '~/components/AppNavbar.vue';

export default {
  middleware: 'admin',
  name: 'Penulis',
  components: {
    AppNavbar
  },
  data() {
    return {
      penulis: [],
      newPenulis: '',
      newBiografi: '', // Menambahkan state untuk biografi baru
      editedPenulis: '',
      editedBiografi: '', // Menambahkan state untuk biografi yang diedit
      isEditing: false,
      penulisToEdit: null,
      isAddPenulisModalOpen: false,
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
        alert('Nama penulis tidak boleh kosong!');
        return; // Berhenti jika input kosong
      }

      try {
        const penulis = { nama: this.newPenulis, biografi: this.newBiografi };
        const newPenulis = await createPenulis(penulis);
        this.penulis.push(newPenulis);
        this.newPenulis = '';
        this.newBiografi = ''; // Reset biografi setelah penulis baru ditambahkan
        this.closeAddPenulisModal();
      } catch (error) {
        console.error('Error creating author', error);
      }
    },
    openAddPenulisModal() {
      this.isAddPenulisModalOpen = true;
    },
    closeAddPenulisModal() {
      this.isAddPenulisModalOpen = false;
      this.newPenulis = ''; // Reset input saat modal ditutup
      this.newBiografi = ''; // Reset biografi saat modal ditutup
    },
    editPenulis(id) {
      this.penulisToEdit = this.penulis.find((penulisItem) => penulisItem.id_penulis === id);
      this.editedPenulis = this.penulisToEdit.nama;
      this.editedBiografi = this.penulisToEdit.biografi;
      this.isEditing = true;
    },
    async updatePenulis() {
      try {
        const updatedPenulis = { nama: this.editedPenulis, biografi: this.editedBiografi };
        const penulis = await updatePenulis(this.penulisToEdit.id_penulis, updatedPenulis);
        const index = this.penulis.findIndex((p) => p.id_penulis === penulis.id_penulis);
        this.penulis.splice(index, 1, penulis);
        this.isEditing = false;
        this.editedPenulis = '';
        this.editedBiografi = '';
        this.penulisToEdit = null;
      } catch (error) {
        console.error('Error updating author', error);
      }
    },
    cancelEdit() {
      this.isEditing = false;
      this.editedPenulis = '';
      this.editedBiografi = '';
      this.penulisToEdit = null;
    },
    async deletePenulis(id) {
      try {
        await deletePenulis(id);
        this.penulis = this.penulis.filter((penulisItem) => penulisItem.id_penulis !== id);
      } catch (error) {
        console.error('Error deleting author', error);
      }
    }
  }
};
</script>

<style scoped>
.container{
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

.penulis-list {
  margin-top: 20px;
}

.penulis-item {
  background-color: white;
  padding: 15px;
  border-radius: 6px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 10px;
}

.penulis-actions {
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
