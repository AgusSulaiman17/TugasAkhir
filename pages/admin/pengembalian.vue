<template>
  <div>
    <AppNavbar />
    <div class="container mt-8">
      <h2 class="mb-4 text-center text-primary">Konfirmasi Pengembalian Buku</h2>

      <!-- Tampilkan Pesan Error -->
      <div v-if="error" class="alert alert-danger text-center">{{ error }}</div>

      <!-- Tampilkan Spinner Jika Data Sedang Dimuat -->
      <div v-if="loading" class="text-center my-5">
        <div class="spinner-border" role="status">
          <span class="sr-only">Loading...</span>
        </div>
      </div>

      <!-- Tampilkan Daftar Pengembalian Pending -->
      <div v-else>
        <div v-if="pengembalianPending.length === 0" class="text-center">
          <p class="text-muted">Tidak ada pengembalian yang perlu dikonfirmasi.</p>
        </div>

        <div v-else class="list-group">
          <div v-for="(item, index) in pengembalianPending" :key="index"
            class="list-group-item mt-3 shadow-sm p-3 mb-3 bg-white rounded">
            <div class="d-flex justify-content-between align-items-center">
              <div>
                <h5 class="font-weight-bold text-primary">ID Peminjaman: {{ item.id_peminjaman }}</h5>
                <p><strong>ID User:</strong> {{ item.id_user }}</p>
                <p><strong>Tanggal Peminjaman:</strong> {{ formatTanggal(item.tanggal_pinjam) }}</p>
                <p><strong>Tanggal Pengembalian:</strong> {{ formatTanggal(item.tanggal_kembali) }}</p>
                <p><strong>Denda:</strong> Rp {{ formatCurrency(item.denda) }}</p>
              </div>
              <div>
                <!-- Tombol Konfirmasi -->
                <button @click="verifyReturn(item.id_peminjaman)" class="btn btn-sm btn-success">
                  Konfirmasi
                </button>

                <!-- Tombol Tolak -->
                <button @click="showRejectModal(item.id_peminjaman)" class="btn btn-sm btn-danger">
                  Tolak
                </button>
              </div>
            </div>

            <!-- Modal untuk Menulis Alasan Penolakan -->
            <div v-if="isRejecting === item.id_peminjaman" class="mt-3">
              <textarea v-model="rejectionReason" class="form-control" rows="4" placeholder="Masukkan alasan penolakan..."></textarea>
              <button @click="rejectReturn(item.id_peminjaman)" class="btn btn-warning mt-2">Kirim Alasan</button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>


<script>
import { getPengembalianPending, verifyReturn, rejectReturn } from '@/api/peminjaman';
import AppNavbar from '~/components/AppNavbar.vue';

export default {
  components: {
    AppNavbar,
  },
  data() {
    return {
      pengembalianPending: [],
      loading: false,
      error: null,
      isRejecting: null, // ID peminjaman yang sedang dalam proses penolakan
      rejectionReason: '', // Alasan penolakan yang dimasukkan
    };
  },
  methods: {
    async fetchPengembalianPending() {
      try {
        this.loading = true;
        this.error = null;
        this.pengembalianPending = await getPengembalianPending();
      } catch (error) {
        console.error('Gagal memuat pengembalian:', error);
        this.error = 'Gagal memuat pengembalian. Coba lagi nanti.';
      } finally {
        this.loading = false;
      }
    },
    async verifyReturn(id) {
      try {
        await verifyReturn(id);
        alert('Pengembalian berhasil dikonfirmasi!');
        this.fetchPengembalianPending();
      } catch (error) {
        console.error('Gagal mengonfirmasi pengembalian:', error);
        this.error = 'Gagal mengonfirmasi pengembalian. Coba lagi nanti.';
      }
    },
    async rejectReturn(id) {
      if (!this.rejectionReason) {
        alert('Alasan penolakan harus diisi.');
        return;
      }
      try {
        await rejectReturn(id, { alasan: this.rejectionReason }); // Kirim alasan ke backend
        alert('Pengembalian berhasil ditolak!');
        this.fetchPengembalianPending();
        this.isRejecting = null;
        this.rejectionReason = '';
      } catch (error) {
        console.error('Gagal menolak pengembalian:', error);
        this.error = 'Gagal menolak pengembalian. Coba lagi nanti.';
      }
    },

    showRejectModal(id) {
      this.isRejecting = id;
      this.rejectionReason = '';
    },
    formatTanggal(date) {
      return new Date(date).toLocaleDateString('id-ID');
    },
    formatCurrency(amount) {
      return amount?.toLocaleString('id-ID') || '0';
    },
  },
  mounted() {
    this.fetchPengembalianPending();
  },
};
</script>
