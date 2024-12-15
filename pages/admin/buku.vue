<template>
  <div>
    <AppNavbar />
    <div class="container mt-8">
      <h2 class="mb-4 text-center text-ijotua">Manajemen Buku</h2>
      <div>
        <button @click="openCreateModal" class="btn bg-ijotua mb-3">Tambah Buku <b-icon-plus></b-icon-plus></button>
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
            <thead class=" bg-ijomuda text-white">
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
                  <img v-if="buku.gambar" :src="`http://localhost:8080${buku.gambar}`" :alt="buku.judul"
                    class="card-img-top" style="width: 100px; border-radius: 4px;" />
                  <span v-else class="text-center">Tidak ada gambar</span>
                </td>
                <td>
                  <button @click="openEditModal(buku)" class="btn bg-kuning"><b-icon-pencil></b-icon-pencil></button>
                  <button @click="confirmDeleteBuku(buku.id_buku)" class="btn bg-merah"><b-icon-trash></b-icon-trash></button>
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
              <button @click="closeModal" class="btn-close"><i class="fas fa-times"></i></button>
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
                  <button type="submit" class="btn bg-ijomuda mr-2">
                    {{ isEditMode ? 'Update' : 'Tambah' }}
                  </button>
                  <button type="button" class="btn bg-merah ms-2" @click="closeModal">Batal</button>
                </div>
              </form>
            </div>
          </div>
        </div>
      </div>

      <!-- Notification Modal -->
      <NotificationModal :isVisible="showNotification" :messageTitle="notificationTitle" :messageBody="notificationBody"
        @close="showNotification = false" />

      <!-- Delete Confirmation Modal -->
      <NotificationModal v-if="showDeleteConfirmModal" :isVisible="showDeleteConfirmModal"
        :messageTitle="'Konfirmasi Hapus'" :messageBody="'Apakah Anda yakin ingin menghapus item ini?'"
        @close="closeDeleteConfirmModal">
        <template #footer>
          <button @click="deleteBuku()" class="btn bg-merah">Ya, Hapus</button>
          <button @click="closeDeleteConfirmModal()" class="btn bg-ijomuda">Batal</button>
        </template>
      </NotificationModal>
    </div>
  </div>
</template>


<script>
import { getBukuList, createBuku, updateBuku, deleteBuku } from '~/api/buku.js';
import { getPenulis } from '@/api/penulis';
import { getGenres } from '@/api/genres';
import AppNavbar from '~/components/AppNavbar.vue';
import NotificationModal from '~/components/NotificationModal.vue';

export default {
  components: {
    AppNavbar,
    NotificationModal
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
      confirmIdBuku: null,
      showNotification: false,
      showDeleteConfirmModal: false,
      notificationTitle: '',
      notificationBody: '',
      bookToDelete: null,
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
          this.notificationTitle = 'Update Sukses';
          this.notificationBody = 'Buku berhasil diperbarui.';
        } else {
          await createBuku(this.buku);
          this.notificationTitle = 'Tambah Sukses';
          this.notificationBody = 'Buku berhasil ditambahkan.';
        }
        this.showNotification = true;
        this.getBukuList();
        this.closeModal();
      } catch (error) {
        this.notificationTitle = 'Terjadi Kesalahan';
        this.notificationBody = 'Gagal menyimpan data buku.';
        this.showNotification = true;
      }
    },
    confirmDeleteBuku(id) {
      this.bookToDelete = id; // Simpan ID buku yang ingin dihapus
      this.showDeleteConfirmModal = true; // Tampilkan modal konfirmasi
    },
    closeDeleteConfirmModal() {
      this.showDeleteConfirmModal = false;
      this.bookToDelete = null; // Hapus ID yang tersimpan
    },
    async deleteBuku() {
      try {
        this.loading = true; // Tambahkan loading indicator saat proses hapus
        await deleteBuku(this.bookToDelete);
        this.notificationTitle = 'Hapus Sukses';
        this.notificationBody = 'Buku berhasil dihapus.';
        this.showNotification = true;
        this.getBukuList(); // Refresh daftar buku
      } catch (error) {
        console.error('Gagal menghapus buku:', error);
        this.notificationTitle = 'Terjadi Kesalahan';
        this.notificationBody = 'Gagal menghapus buku.';
        this.showNotification = true;
      } finally {
        this.loading = false;
        this.closeDeleteConfirmModal(); // Tutup modal konfirmasi
      }
    }
  }
};
</script>


<style scoped>
.container {
  max-width: 1200px;
  padding: 20px;
  background-color: #f9f9f9;
  border-radius: 8px;
}

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

.modal-dialog {
  background-color: #fff;
  max-width: 500px;
  width: 100%;
  border-radius: 8px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  padding: 20px;
  animation: fadeIn 0.3s ease-in-out;
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-bottom: 1px solid #ddd;
  padding-bottom: 10px;
  background: green;
}

.modal-title {
  font-size: 1.25rem;
  font-weight: bold;
  color: #ffff;

}

.btn-close {
  background: none;
  border: none;
  font-size: 1.5rem;
  color: #888;
  cursor: pointer;
}

.btn-close i {
  font-size: 1.2rem;
}

.table {
  width: 100%;
  margin-bottom: 1rem;
  color: #212529;
  border-radius: 8px;
}

.table-bordered {
  border: 1px solid #dee2e6;
}

.table th, .table td {
  border-top: 1px solid #dee2e6;
  padding: 8px 12px;
}

.text-center {
  text-align: center;
}

.alert {
  padding: 10px;
  margin: 10px 0;
  border-radius: 4px;
  color: white;
  background-color: #e74c3c;
}

.spinner-border {
  width: 2rem;
  height: 2rem;
  border-width: 0.3em;
}

@keyframes fadeIn {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}

</style>
