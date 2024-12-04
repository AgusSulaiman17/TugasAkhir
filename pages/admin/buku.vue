<template>
  <div>
    <AppNavbar />
    <div class="container mt-8">
      <h2 class="mb-4 text-center text-primary">Manajemen Buku</h2>
      <div>
        <button @click="openCreateModal" class="btn btn-primary mb-3">Tambah Buku</button>
      </div>

      <!-- Tampilkan Pesan Error -->
      <div v-if="error" class="alert alert-danger text-center">{{ error }}</div>

      <!-- Tampilkan Spinner Jika Data Sedang Dimuat -->
      <div v-if="loading" class="text-center my-5">
        <div class="spinner-border" role="status">
          <span class="sr-only">Loading...</span>
        </div>
      </div>

      <!-- Tampilkan Daftar Buku -->
      <div v-else>
        <div v-if="bukuList.length === 0" class="text-center">
          <p class="text-muted">Tidak ada buku yang terdaftar.</p>
        </div>

        <div v-else>
          <table class="table table-striped table-bordered">
            <thead class=" bg-ijomuda">
              <tr>
                <th>ID Buku</th>
                <th>Judul</th>
                <th>Penulis</th>
                <th>Genre</th>
                <th>Deskripsi</th>
                <th>Jumlah</th>
                <th>Gambar</th>
                <th>Aksi</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="buku in bukuList" :key="buku.id_buku">
                <td>{{ buku.id_buku }}</td>
                <td>{{ buku.judul }}</td>
                <td>{{ buku.penulis?.nama || 'Tidak diketahui' }}</td>
                <td>{{ buku.genre?.nama || 'Tidak diketahui' }}</td>
                <td>{{ buku.deskripsi }}</td>
                <td>{{ buku.jumlah }}</td>
                <td>
                  <img v-if="buku.gambar" :src="`http://localhost:8080${buku.gambar}`" :alt="buku.judul" class="card-img-top" style="width: 100px" />
                  <span v-else class="text-center">Tidak ada gambar</span>
                </td>
                <td>
                  <button @click="openEditModal(buku)" class="btn btn-sm btn-primary">Edit</button>
                  <button @click="deleteBuku(buku.id_buku)" class="btn btn-sm btn-danger">Hapus</button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Modal untuk menambah dan mengedit buku -->
      <div v-if="isModalOpen" class="modal-overlay">
        <div class="modal-dialog">
          <div class="modal-content">
            <div class="modal-header">
              <h5 class="modal-title">{{ isEditMode ? 'Edit Buku' : 'Tambah Buku' }}</h5>
              <button @click="closeModal" class="btn-close"></button>
            </div>
            <div class="modal-body">
              <form @submit.prevent="submitForm">
                <div class="mb-3">
                  <label for="judul" class="form-label">Judul</label>
                  <input type="text" id="judul" class="form-control" v-model="buku.judul" required />
                </div>
                <div class="mb-3">
                  <label for="id_penulis" class="form-label">Penulis</label>
                  <select id="id_penulis" class="form-select" v-model="buku.id_penulis" required>
                    <option v-for="penulis in penulisList" :key="penulis.id_penulis" :value="penulis.id_penulis">
                      {{ penulis.nama }}
                    </option>
                  </select>
                </div>
                <div class="mb-3">
                  <label for="id_genre" class="form-label">Genre</label>
                  <select id="id_genre" class="form-select" v-model="buku.id_genre" required>
                    <option v-for="genre in genreList" :key="genre.id_genre" :value="genre.id_genre">
                      {{ genre.nama }}
                    </option>
                  </select>
                </div>
                <div class="mb-3">
                  <label for="deskripsi" class="form-label">Deskripsi</label>
                  <textarea id="deskripsi" class="form-control" v-model="buku.deskripsi"></textarea>
                </div>
                <div class="mb-3">
                  <label for="jumlah" class="form-label">Jumlah</label>
                  <input type="number" id="jumlah" class="form-control" v-model="buku.jumlah" required />
                </div>
                <div class="mb-3">
                  <label for="gambar" class="form-label">Gambar</label>
                  <input type="file" id="gambar" class="form-control" @change="onFileChange" />
                </div>
                <div class="d-flex justify-content-end">
                  <button type="submit" class="btn btn-primary">
                    {{ isEditMode ? 'Update' : 'Tambah' }}
                  </button>
                  <button type="button" class="btn btn-secondary ms-2" @click="closeModal">Batal</button>
                </div>
              </form>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { getBukuList, createBuku, updateBuku, deleteBuku } from '~/api/buku.js';
import { getPenulis } from '@/api/penulis';
import { getGenres } from '@/api/genres';
import AppNavbar from '~/components/AppNavbar.vue';

export default {
  components: {
    AppNavbar
  },
  data() {
    return {
      bukuList: [],
      penulisList: [],
      genreList: [],
      buku: {
        id_buku: null,
        judul: '',
        id_penulis: '',
        id_genre: '',
        deskripsi: '',
        jumlah: 0,
        gambar: null,
      },
      isModalOpen: false,
      isEditMode: false,
      loading: false,
      error: null,
    };
  },
  async created() {
    await this.getBukuList();
    await this.fetchPenulisList();
    await this.fetchGenreList();
  },
  methods: {
    async getBukuList() {
      try {
        this.loading = true;
        this.error = null;
        this.bukuList = await getBukuList();
      } catch (error) {
        this.error = 'Gagal memuat daftar buku.';
        console.error('Error fetching buku list:', error);
      } finally {
        this.loading = false;
      }
    },
    async fetchPenulisList() {
      try {
        const data = await getPenulis();
        this.penulisList = data || [];
      } catch (error) {
        console.error('Error fetching penulis:', error);
      }
    },
    async fetchGenreList() {
      try {
        const data = await getGenres();
        this.genreList = data || [];
      } catch (error) {
        console.error('Error fetching genre:', error);
      }
    },
    openCreateModal() {
      this.isModalOpen = true;
      this.isEditMode = false;
      this.buku = {
        id_buku: null,
        judul: '',
        id_penulis: '',
        id_genre: '',
        deskripsi: '',
        jumlah: 0,
        gambar: null,
      };
    },
    openEditModal(buku) {
      this.isModalOpen = true;
      this.isEditMode = true;
      this.buku = { ...buku };
    },
    closeModal() {
      this.isModalOpen = false;
    },
    onFileChange(event) {
      const file = event.target.files[0];
      if (file) {
        this.buku.gambar = file;
      }
    },
    async submitForm() {
      try {
        if (this.isEditMode) {
          await updateBuku(this.buku);
        } else {
          await createBuku(this.buku);
        }
        this.getBukuList();
        this.closeModal();
      } catch (error) {
        this.error = 'Gagal menyimpan data.';
        console.error('Error while saving data:', error);
      }
    },
    async deleteBuku(id_buku) {
      if (confirm('Apakah Anda yakin ingin menghapus buku ini?')) {
        try {
          await deleteBuku(id_buku);
          this.getBukuList();
        } catch (error) {
          this.error = 'Gagal menghapus buku.';
          console.error('Error deleting buku:', error);
        }
      }
    }
  }
};
</script>


<style scoped>
/* Modal Overlay */
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 99999;
}

/* Modal Content */
.modal-dialog {
  background-color: #fff;
  width: 100%;
  max-width: 500px;
  border-radius: 8px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  padding: 20px;
  animation: fadeIn 0.3s ease-in-out;
}

/* Modal Header */
.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-bottom: 1px solid #ddd;
  padding-bottom: 10px;
}

.modal-title {
  font-size: 1.25rem;
  font-weight: bold;
  color: #333;
}

.btn-close {
  background: none;
  border: none;
  font-size: 1.5rem;
  color: #aaa;
  cursor: pointer;
}

/* Modal Body */
.modal-body {
  padding: 20px 0;
}

.form-label {
  font-weight: bold;
  color: #555;
}

.form-control, .form-select {
  border-radius: 4px;
  box-shadow: none;
  border: 1px solid #ddd;
  padding: 10px;
  font-size: 1rem;
}

button[type="submit"], .btn-secondary {
  padding: 10px 20px;
  font-size: 1rem;
}

button[type="submit"] {
  background-color: #28a745;
  color: #fff;
  border: none;
}

button[type="button"] {
  background-color: #6c757d;
  color: #fff;
  border: none;
}

/* Button Hover */
button[type="submit"]:hover {
  background-color: #218838;
}

button[type="button"]:hover {
  background-color: #5a6268;
}

/* Modal Fade In Animation */
@keyframes fadeIn {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}
</style>
