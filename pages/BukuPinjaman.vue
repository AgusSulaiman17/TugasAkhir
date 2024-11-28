<template>
  <div>
    <AppNavbar />
    <div class="container mt-8">
      <h2 class="mb-4 text-center text-primary">Buku yang Dipinjam</h2>

      <!-- Filter Pencarian dan Status -->
      <div class="mb-4 d-flex justify-content-between">
        <input v-model="searchQuery" type="text" class="form-control w-50" placeholder="Cari judul buku..." />
        <select v-model="filterStatus" class="form-control w-25">
          <option value="">Semua Status Pinjaman</option>
          <option value="Dipinjam">Dipinjam</option>
          <option value="Dikembalikan">Dikembalikan</option>
        </select>
        <select v-model="filterPengembalian" class="form-control w-25">
          <option value="">Semua Status Pengembalian</option>
          <option value="Pending">Pending</option>
          <option value="Approved">Approved</option>
          <option value="Rejected">Rejected</option>
        </select>
      </div>

      <!-- Pesan Error -->
      <div v-if="error" class="alert alert-danger text-center">{{ error }}</div>

      <!-- Spinner Loading -->
      <div v-if="loading" class="text-center my-5">
        <div class="spinner-border" role="status">
          <span class="sr-only">Loading...</span>
        </div>
      </div>

      <!-- Daftar Peminjaman -->
      <div v-else>
        <div v-if="filteredPeminjaman.length === 0" class="text-center">
          <p class="text-muted">Tidak ada buku yang ditemukan.</p>
        </div>

        <div v-else class="list-group">
          <div
            v-for="(item, index) in filteredPeminjaman"
            :key="index"
            class="list-group-item mt-3 shadow-sm p-3 mb-3 bg-white rounded"
          >
            <div class="d-flex justify-content-between align-items-center">
              <div>
                <h5 class="font-weight-bold text-primary">{{ item?.buku?.judul || 'Judul tidak tersedia' }}</h5>
                <p class="text-muted"><strong>Penulis:</strong> {{ item?.buku?.penulis?.nama || 'Penulis tidak tersedia' }}</p>
                <p><strong>Durasi:</strong> {{ item.durasi_hari || 'Durasi tidak tersedia' }} hari</p>
                <p>
                  <strong>Tanggal Kembali:</strong>
                  {{ item.tanggal_kembali ? formatTanggal(item.tanggal_kembali) : 'Tanggal tidak tersedia' }}
                </p>
                <p>
                  <strong>Status:</strong>
                  <span :class="item.status_kembali ? 'text-success' : 'text-danger'">
                    {{ item.status_kembali ? 'Dikembalikan' : 'Dipinjam' }}
                  </span>
                </p>
                <p>
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
              <div>
                <button v-if="!item.status_kembali" @click="returnBook(item.id_peminjaman)" class="btn btn-sm btn-danger">
                  Kembalikan Buku
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { getAllPeminjaman, returnBook } from '@/api/peminjaman';
import AppNavbar from '~/components/AppNavbar.vue';

export default {
  components: {
    AppNavbar,
  },
  data() {
    return {
      peminjaman: [],
      searchQuery: '',
      filterStatus: '',
      filterPengembalian: '',
      loading: false,
      error: null,
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
        this.loading = true;
        this.error = null;
        const data = await getAllPeminjaman();
        const userId = this.getUserId();
        this.peminjaman = data.filter((item) => item.id_user === userId);
      } catch (error) {
        console.error('Gagal memuat daftar peminjaman:', error);
        this.error = 'Gagal memuat daftar peminjaman. Coba lagi nanti.';
      } finally {
        this.loading = false;
      }
    },
    async returnBook(id) {
      try {
        this.loading = true;
        await returnBook(id);
        alert('Buku berhasil dikembalikan!');
        const updatedPeminjaman = this.peminjaman.map((item) => {
          if (item.id_peminjaman === id) {
            item.status_kembali = true;
            item.status_pengembalian = 'Pending';
          }
          return item;
        });
        this.peminjaman = updatedPeminjaman;
      } catch (error) {
        console.error('Gagal mengembalikan buku:', error);
        this.error = 'Gagal mengembalikan buku. Coba lagi nanti.';
      } finally {
        this.loading = false;
      }
    },
    formatTanggal(date) {
      return new Date(date).toLocaleDateString('id-ID');
    },
    getUserId() {
      const token = localStorage.getItem('token');
      if (!token) return null;
      try {
        const payload = JSON.parse(atob(token.split('.')[1]));
        return payload.id_user;
      } catch (error) {
        console.error('Gagal memproses token JWT:', error);
        return null;
      }
    },
  },
  mounted() {
    this.fetchPeminjaman();
  },
};
</script>
