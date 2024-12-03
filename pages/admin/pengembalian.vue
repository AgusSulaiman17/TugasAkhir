<template>
  <div>
    <AppNavbar />
    <div class="container mt-8">
      <h2 class="mb-4 text-center text-ijomuda">Konfirmasi Pengembalian Buku</h2>

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
          <div
            v-for="(item, index) in pengembalianPending"
            :key="index"
            class="card mb-4 shadow-lg"
          >
            <div class="row g-0">
              <!-- Detail Buku -->
              <div class="col-md-12">
                <div class="card-body d-flex justify-content-between align-items-center">
                  <div>
                    <h5 class="card-title text-primary">ID Peminjaman: {{ item.id_peminjaman }}</h5>
                    <p class="card-text">
                      <strong>ID User:</strong> {{ item.id_user }}
                    </p>
                    <p class="card-text">
                      <strong>Tanggal Peminjaman:</strong> {{ formatTanggal(item.tanggal_pinjam) }}
                    </p>
                    <p class="card-text">
                      <strong>Tanggal Pengembalian:</strong> {{ formatTanggal(item.tanggal_kembali) }}
                    </p>
                    <p class="card-text">
                      <strong>Denda:</strong> Rp {{ formatCurrency(item.denda) }}
                    </p>
                  </div>

                  <!-- Tombol Aksi -->
                  <div class="d-flex flex-column gap-2">
                    <button
                      @click="verifyReturn(item.id_peminjaman)"
                      class="btn btn-outline-success btn-sm"
                    >
                      Konfirmasi
                    </button>
                    <button
                      @click="showRejectModal(item.id_peminjaman)"
                      class="btn btn-outline-danger btn-sm"
                    >
                      Tolak
                    </button>
                  </div>
                </div>

                <!-- Modal untuk Menulis Alasan Penolakan -->
                <div v-if="isRejecting === item.id_peminjaman" class="mt-3 px-3">
                  <textarea
                    v-model="rejectionReason"
                    class="form-control shadow-sm"
                    rows="4"
                    placeholder="Masukkan alasan penolakan..."
                  ></textarea>
                  <div class="d-flex gap-2 mt-2 mb-2">
                    <button
                      @click="rejectReturn(item.id_peminjaman)"
                      class="btn btn-warning shadow-sm mr-2"
                    >
                      Kirim Alasan
                    </button>
                    <button
                      @click="cancelRejectModal"
                      class="btn btn-secondary shadow-sm"
                    >
                      Batal
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>


<script>
import { getPengembalianPending, verifyReturn, rejectReturn } from "@/api/peminjaman";
import AppNavbar from "~/components/AppNavbar.vue";

export default {
  components: {
    AppNavbar,
  },
  data() {
    return {
      pengembalianPending: [],
      loading: false,
      error: null,
      isRejecting: null,
      rejectionReason: "",
    };
  },
  methods: {
    async fetchPengembalianPending() {
      try {
        this.loading = true;
        this.error = null;
        this.pengembalianPending = await getPengembalianPending();
      } catch (error) {
        this.error = "Gagal memuat pengembalian. Coba lagi nanti.";
      } finally {
        this.loading = false;
      }
    },
    async verifyReturn(id) {
      try {
        await verifyReturn(id);
        alert("Pengembalian berhasil dikonfirmasi!");
        this.fetchPengembalianPending();
      } catch (error) {
        this.error = "Gagal mengonfirmasi pengembalian. Coba lagi nanti.";
      }
    },
    async rejectReturn(id) {
      if (!this.rejectionReason) {
        alert("Alasan penolakan harus diisi.");
        return;
      }
      try {
        await rejectReturn(id, { alasan: this.rejectionReason });
        alert("Pengembalian berhasil ditolak!");
        this.fetchPengembalianPending();
        this.isRejecting = null;
        this.rejectionReason = "";
      } catch (error) {
        this.error = "Gagal menolak pengembalian. Coba lagi nanti.";
      }
    },
    showRejectModal(id) {
      this.isRejecting = id;
      this.rejectionReason = "";
    },
    formatTanggal(date) {
      return new Date(date).toLocaleDateString("id-ID");
    },
    formatCurrency(amount) {
      return amount?.toLocaleString("id-ID") || "0";
    },
    cancelRejectModal() {
    this.isRejecting = null;
    this.rejectionReason = "";
  },
  },
  mounted() {
    this.fetchPengembalianPending();
  },
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

.btn-outline-danger,
.btn-outline-success {
  background-color: #70a799;
  color: #fff;
  padding: 0.50rem 1rem;
  border: none;
  border-radius: 8px;
  font-size: 0.875rem;
  cursor: pointer;
  transition: all 0.3s ease-in-out;
  position: relative;
  overflow: hidden;
}
.btn-outline-danger {
  margin-top: 5px;
  background-color: #e74c3c;
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
  background-color: #c0392b; /* Warna merah lebih gelap saat hover */
}

.btn-outline-danger:hover::before,
.btn-outline-success:hover::before {
  transform: translate(-50%, -50%) scale(1);
}

.btn-outline-danger:focus{
  outline: none;
  box-shadow: 0 0 0 0.2rem rgba(231, 76, 60, 0.5);
}
.btn-outline-success:focus {
  outline: none;
  box-shadow: 0 0 0 0.2rem rgba(112, 167, 153, 0.5);
}
</style>
