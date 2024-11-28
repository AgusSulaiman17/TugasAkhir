<template>
  <div>
    <AppNavbar />
    <div class="container mt-8">
      <!-- Tampilkan Pesan Error -->
      <div v-if="error" class="alert alert-danger">{{ error }}</div>

      <!-- Tampilkan Detail Buku -->
      <div v-if="loading" class="text-center my-5">
        <div class="spinner-border" role="status">
          <span class="sr-only">Loading...</span>
        </div>
      </div>

      <div v-else-if="buku" class="card">
        <div class="row no-gutters">
          <div class="col-md-4">
            <img v-if="buku.gambar"
              :src="buku.gambar.startsWith('http') ? buku.gambar : `http://localhost:8080${buku.gambar}`"
              :alt="buku.judul" class="card-img" style="height: 100%; object-fit: cover;" />
            <div v-else class="text-center p-3">Tidak ada gambar</div>
          </div>
          <div class="col-md-8">
            <div class="card-body">
              <h3 class="card-title">{{ buku.judul }}</h3>
              <p class="card-text"><strong>Penulis:</strong> {{ buku.penulis.nama || 'Tidak diketahui' }}</p>
              <p class="card-text"><strong>Genre:</strong> {{ buku.genre.nama || 'Tidak diketahui' }}</p>
              <p class="card-text"><strong>Deskripsi:</strong> {{ buku.deskripsi || 'Tidak tersedia' }}</p>
              <p class="card-text"><strong>Jumlah:</strong> {{ buku.jumlah }}</p>
              <button @click="showModal = true" class="btn btn-success mt-3">Pinjam Buku</button>
            </div>
          </div>
        </div>
      </div>

      <!-- Modal Peminjaman -->
      <div v-if="showModal" class="modal fade show d-block " tabindex="-1" role="dialog"
        style="background: rgba(0, 0, 0, 0.5);">
        <div class="modal-dialog mt-8" role="document">
          <div class="modal-content">
            <div class="modal-header">
              <h5 class="modal-title">Formulir Peminjaman Buku</h5>
              <button type="button" class="close" @click="showModal = false" aria-label="Close">
                <span aria-hidden="true">&times;</span>
              </button>
            </div>
            <div class="modal-body">
              <!-- Loading Spinner -->
              <div v-if="loading" class="spinner-overlay">
                <div class="spinner-border text-primary" role="status">
                  <span class="sr-only">Loading...</span>
                </div>
              </div>

              <form @submit.prevent="submitPeminjaman">
                <div class="form-group">
                  <label for="durasi">Durasi (hari)</label>
                  <input type="number" v-model.number="form.durasiHari" placeholder="Masukkan durasi (hari)" required />
                </div>
                <div class="form-group">
                  <label for="jamKembali">Jam Pengembalian</label>
                  <input type="time" id="jamKembali" v-model="form.jamKembali" class="form-control" required />
                </div>
                <button type="submit" class="btn btn-primary">Kirim</button>
              </form>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { getBukuById } from '@/api/buku';
import { createPeminjaman } from '@/api/peminjaman';

import AppNavbar from '~/components/AppNavbar.vue';

export default {
  components: {
    AppNavbar,
  },
  data() {
    return {
      buku: null,
      loading: false,
      error: null,
      showModal: false,
      form: {
        durasiHari: '',
        jamKembali: '',
      },
    };
  },
  methods: {
    async fetchBuku() {
      try {
        this.loading = true;
        this.error = null;
        const id = this.$route.params.id_buku;
        if (!id) {
          throw new Error('ID buku tidak diberikan');
        }
        const data = await getBukuById(id);
        this.buku = data;
      } catch (error) {
        console.error(error);
        this.error = 'Gagal memuat detail buku. Coba lagi nanti.';
      } finally {
        this.loading = false;
      }
    },
    async submitPeminjaman() {
      try {
        if (!this.form.durasiHari || !this.form.jamKembali) {
          this.error = 'Semua kolom harus diisi';
          return;
        }

        const peminjamanData = {
          id_buku: this.buku.id_buku,
          durasi_hari: this.form.durasiHari,
          jam_kembali: this.form.jamKembali,
        };

        this.loading = true;

        const response = await createPeminjaman(peminjamanData);

        if (response && response.id_peminjaman) {
          this.showModal = false;
          this.$router.push('/BukuPinjaman');
          alert('Peminjaman berhasil dibuat!');
        } else {
          throw new Error('Respons dari server tidak valid.');
        }
      } catch (error) {
        console.error('Error saat membuat peminjaman:', error);
        this.error = error.response?.data?.message || error.message || 'Gagal mengajukan peminjaman. Coba lagi nanti.';
      } finally {
        this.loading = false;
      }
    },
  },
  mounted() {
    this.fetchBuku();
  },
};
</script>

<style scoped>
.card {
  border-radius: 10px;
  overflow: hidden;
}

.card-img {
  border-top-left-radius: 10px;
  border-bottom-left-radius: 10px;
}

.modal {
  display: flex;
  align-items: center;
  justify-content: center;
}

.spinner-overlay {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  display: flex;
  justify-content: center;
  align-items: center;
  background: rgba(255, 255, 255, 0.8);
  z-index: 10;
}

.modal-content {
  position: relative;
}
</style>
