<template>
  <div>
    <AppNavbar />
    <div class="container mt-8">
      <h2 class="mb-4 text-center text-ijomuda">Buku yang Dipinjam</h2>
      <div class="d-flex flex-wrap gap-3 mb-4 align-items-center">
        <!-- Search Input -->
        <div class="input-group w-50 mr-4">
          <input v-model="searchQuery" type="text" class="form-control border-end-0 rounded-start shadow-sm"
            placeholder="Cari judul buku..." />
          <button class="btn bg-ijotua border-start-0 rounded-end shadow-sm">
            <i class="bi bi-search"></i>
          </button>
        </div>

        <!-- Dropdown Status Pinjaman -->
        <div class="flex-grow-1">
          <select v-model="filterStatus" class="form-select shadow-sm">
            <option value="">Semua Status Pinjaman</option>
            <option value="Dipinjam">Dipinjam</option>
            <option value="Dikembalikan">Dikembalikan</option>
          </select>
        </div>

        <!-- Dropdown Status Pengembalian -->
        <div class="flex-grow-1">
          <select v-model="filterPengembalian" class="form-select shadow-sm">
            <option value="">Semua Status Pengembalian</option>
            <option value="Pending">Pending</option>
            <option value="Approved">Approved</option>
            <option value="Rejected">Rejected</option>
          </select>
        </div>
      </div>
      <div v-if="error" class="alert alert-danger text-center">{{ error }}</div>
      <div v-if="loading" class="text-center my-5">
        <div class="spinner-border" role="status">
          <span class="sr-only">Loading...</span>
        </div>
      </div>
      <div v-else>
        <div v-if="filteredPeminjaman.length === 0" class="text-center">
          <p class="text-muted">Tidak ada buku yang ditemukan.</p>
        </div>
        <div v-else class="list-group">
          <div v-for="(item, index) in filteredPeminjaman" :key="index" class="card mb-4 shadow-lg">
            <div class="row g-0">
              <!-- Gambar Buku -->
              <div class="col-md-3 d-flex justify-content-center align-items-center">
                <img v-if="item.buku.gambar" :src="item.buku.gambar.startsWith('http')
                  ? item.buku.gambar
                  : `http://localhost:8080${item.buku.gambar}`" :alt="item.buku.judul" class="img-fluid rounded-start"
                  style="max-height: 150px; object-fit: cover;" />
                <img v-else src="/path/to/fallback-image.jpg" alt="Placeholder" class="img-fluid rounded-start"
                  style="max-height: 150px; object-fit: cover;" />
              </div>

              <!-- Detail Buku -->
              <div class="col-md-9">
                <div class="card-body d-flex justify-content-between align-items-center">
                  <div>
                    <h5 class="card-title text-ijomuda">{{ item?.buku?.judul || 'Judul tidak tersedia' }}</h5>
                    <p class="card-text">
                      <strong>Durasi:</strong> {{ item.durasi_hari || 'Durasi tidak tersedia' }} hari
                    </p>
                    <p class="card-text">
                      <strong>Tanggal Kembali:</strong>
                      {{ item.tanggal_kembali ? formatTanggal(item.tanggal_kembali) : 'Tanggal tidak tersedia' }}
                    </p>
                    <p class="card-text">
                      <strong>Status:</strong>
                      <span :class="item.status_kembali ? 'text-success' : 'text-danger'">
                        {{ item.status_kembali ? 'Dikembalikan' : 'Dipinjam' }}
                      </span>
                    </p>
                    <p class="card-text">
                      <strong>Status Pengembalian:</strong>
                      <span :class="{
                        'text-warning': item.status_pengembalian === 'Pending',
                        'text-success': item.status_pengembalian === 'Approved',
                        'text-danger': item.status_pengembalian === 'Rejected',
                      }">
                        {{ item.status_pengembalian || 'Belum Dikembalikan' }}
                      </span>
                    </p>
                  </div>
                  <!-- Tombol Aksi -->
                  <div class="d-flex flex-column gap-2">
                    <button v-if="!item.status_kembali" @click="showReturnConfirmation(item.id_peminjaman)"
                      class="btn bg-ijomuda btn-sm mb-1" :disabled="isLoadingForReturn">
                      <span v-if="isLoadingForReturn" class="spinner-border spinner-border-sm" role="status"
                        aria-hidden="true"></span>
                      <span v-else> Kembalikan Buku </span>
                    </button>
                    <button v-if="item.status_pengembalian === 'Approved'"
                      @click="showDeleteConfirmation(item.id_peminjaman)" class="btn bg-merah btn-sm">
                      Hapus Peminjaman
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Success Notification Modal -->
    <NotificationModal v-if="showSuccessModal" :isVisible="showSuccessModal" :messageTitle="successTitle"
      :messageBody="successMessage" @close="closeSuccessModal" />

    <!-- Confirmation Modal for Returning Book -->
    <NotificationModal v-if="showReturnConfirmationModal" :isVisible="showReturnConfirmationModal"
      :messageTitle="'Konfirmasi Pengembalian Buku'" :messageBody="'Apakah Anda yakin ingin mengembalikan buku ini?'"
      @close="cancelReturn">
      <template #footer>
        <button @click="confirmReturn" class="btn btn-abu" :disabled="isLoadingForReturn">
          <span v-if="isLoadingForReturn" class="spinner-border spinner-border-sm" role="status"
            aria-hidden="true"></span>
          <span v-else> Ya, Kembalikan </span>
        </button>
        <button @click="cancelReturn" class="btn btn-cancel">Batal</button>
      </template>
    </NotificationModal>

    <!-- Confirmation Modal for Deleting Peminjaman -->
    <NotificationModal v-if="showDeleteConfirmationModal" :isVisible="showDeleteConfirmationModal"
      :messageTitle="'Konfirmasi Hapus Peminjaman'" :messageBody="'Apakah Anda yakin ingin menghapus peminjaman ini?'"
      @close="cancelDelete">
      <template #footer>
        <button @click="confirmDelete" class="btn bg-merah">Ya, Hapus</button>
        <button @click="cancelDelete" class="btn bg-ijomuda">Batal</button>
      </template>
    </NotificationModal>
  </div>
</template>

<script>
import { getAllPeminjaman, returnBook, deletePeminjaman } from '@/api/peminjaman';
import AppNavbar from '~/components/AppNavbar.vue';
import NotificationModal from '~/components/NotificationModal.vue';

export default {
  components: {
    AppNavbar,
    NotificationModal
  },
  data() {
    return {
      peminjaman: [],
      searchQuery: '',
      filterStatus: '',
      filterPengembalian: '',
      loading: false,
      error: null,
      showReturnConfirmationModal: false,
      showDeleteConfirmationModal: false,
      showSuccessModal: false,
      successTitle: '',
      successMessage: '',
      returnBookId: null,
      deleteBookId: null,
      isLoadingForReturn: false, // State untuk kontrol loading pada konfirmasi pengembalian
    };
  },
  computed: {
    filteredPeminjaman() {
      return this.peminjaman.filter((item) => {
        const matchesSearch = item?.buku?.judul?.toLowerCase().includes(this.searchQuery.toLowerCase());
        const matchesStatus =
          this.filterStatus === '' ||
          (this.filterStatus === 'Dipinjam' && !item.status_kembali) ||
          (this.filterStatus === 'Dikembalikan' && item.status_kembali);
        const matchesPengembalian =
          this.filterPengembalian === '' || item.status_pengembalian === this.filterPengembalian;
        return matchesSearch && matchesStatus && matchesPengembalian;
      });
    },
  },
  methods: {
    async fetchPeminjaman() {
      try {
        this.error = null;
        const data = await getAllPeminjaman();
        const userId = this.getUserId();
        this.peminjaman = data.filter((item) => item.id_user === userId);
      } catch (error) {
        console.error('Gagal memuat daftar peminjaman:', error);
        this.error = 'Gagal memuat daftar peminjaman. Coba lagi nanti.';
      }
    },

    showReturnConfirmation(id) {
      this.returnBookId = id;
      this.showReturnConfirmationModal = true;
    },

    showDeleteConfirmation(id) {
      this.deleteBookId = id;
      this.showDeleteConfirmationModal = true;
    },

    cancelReturn() {
      this.showReturnConfirmationModal = false;
    },

    cancelDelete() {
      this.showDeleteConfirmationModal = false;
    },

    async confirmReturn() {
      this.isLoadingForReturn = true;
      try {
        await returnBook(this.returnBookId);
        this.successTitle = 'Pengembalian Buku Berhasil';
        this.successMessage = 'Tunggu Konfirmasi pengembalian dari admin';
        this.showSuccessModal = true;
        this.showReturnConfirmationModal = false;
        // Ambil data baru dari backend
        await this.fetchPeminjaman();
      } catch (error) {
        console.error('Gagal mengembalikan buku:', error);
        this.error = 'Gagal mengembalikan buku. Coba lagi nanti.';
      } finally {
        this.isLoadingForReturn = false;
      }
    },


    async confirmDelete() {
      try {
        this.loading = true;
        await deletePeminjaman(this.deleteBookId);
        this.successTitle = 'Hapus Peminjaman Berhasil';
        this.successMessage = 'Peminjaman berhasil dihapus!';
        this.showSuccessModal = true;
        this.showDeleteConfirmationModal = false;
        // Remove the deleted peminjaman from the list
        this.peminjaman = this.peminjaman.filter(item => item.id_peminjaman !== this.deleteBookId);
      } catch (error) {
        console.error('Gagal menghapus peminjaman:', error);
        this.error = 'Gagal menghapus peminjaman. Coba lagi nanti.';
      } finally {
        this.loading = false;
      }
    },

    closeSuccessModal() {
      this.showSuccessModal = false;
    },

    formatTanggal(date) {
      const options = { year: 'numeric', month: 'long', day: 'numeric' };
      return new Date(date).toLocaleDateString('id-ID', options);
    },

    getUserId() {
      const user = JSON.parse(localStorage.getItem('user'));
      return user ? user.id_user : null;
    }
  },
  mounted() {
    this.fetchPeminjaman();
  }
};
</script>

<style scoped>
.card {
  border-radius: 12px;
  overflow: hidden;
  transition: transform 0.2s ease-in-out, box-shadow 0.2s ease-in-out;
}

.card:hover {
  transform: translateY(-5px);
  box-shadow: 0 10px 20px rgba(0, 0, 0, 0.15);
}

.card-title {
  font-weight: bold;
  font-size: 1.25rem;
}

.card-text {
  margin-bottom: 0.5rem;
}

.img-fluid {
  border-radius: 12px;
}

.btn-outline-danger:hover,
.btn-outline-success:hover {
  background-color: #2e3b33;
  color: #fff;
}

/* Input Search */
.input-group input {
  border-radius: 8px 0 0 8px;
  transition: box-shadow 0.3s;
}

.input-group input:focus {
  box-shadow: 0 4px 8px #2e3b33;
  outline: none;
}

.input-group button {
  background-color: #2e3b33;
  border-radius: 0 8px 8px 0;
  transition: background-color 0.3s, box-shadow 0.3s;
}

.input-group button:hover {
  background-color: #2e3b33;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
}

.input-group .bi-search {
  font-size: 1.2rem;
}

/* Dropdown */
.form-select {
  border-radius: 8px;
  padding: 0.5rem;
  font-size: 0.95rem;
  transition: box-shadow 0.3s;
}

.form-select:focus {
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  outline: none;
}

/* Hover and active states */
.form-select:hover,
.input-group input:hover {
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

.form-select option {
  padding: 0.5rem;
}

.btn-outline-danger,
.btn-outline-success {
  background-color: #70a799;
  /* Warna dasar sama seperti tombol bg-ijomuda */
  color: #fff;
  padding: 0.75rem 1.5rem;
  border: none;
  border-radius: 8px;
  font-size: 0.875rem;
  cursor: pointer;
  transition: all 0.3s ease-in-out;
  position: relative;
  overflow: hidden;
}

.btn-outline-danger {
  background-color: #e74c3c;
  /* Warna dasar merah */
}



.btn-outline-danger::before,
.btn-outline-success::before {
  content: '';
  position: absolute;
  top: 50%;
  left: 50%;
  width: 300%;
  height: 300%;
  background-color: rgba(255, 255, 255, 0.2);
  transition: all 0.4s ease-in-out;
  border-radius: 50%;
  transform: translate(-50%, -50%) scale(0);
}

.btn-outline-danger:hover,
.btn-outline-success:hover {
  background-color: #5e9278;
  box-shadow: 0 12px 24px rgba(0, 0, 0, 0.2);
}

.btn-outline-danger:hover {
  background-color: #c0392b;
  /* Warna merah lebih gelap saat hover */
}

.btn-outline-danger:hover::before,
.btn-outline-success:hover::before {
  transform: translate(-50%, -50%) scale(1);
}

.btn-outline-danger:focus {
  outline: none;
  box-shadow: 0 0 0 0.2rem rgba(231, 76, 60, 0.5);
}

.btn-outline-success:focus {
  outline: none;
  box-shadow: 0 0 0 0.2rem rgba(112, 167, 153, 0.5);
}
</style>
