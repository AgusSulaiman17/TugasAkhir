<template>
  <div>
    <AppNavbar />
    <div class="admin-peminjaman mt-8 mx-auto" style="max-width: 1200px;">
      <h1 class="text-center mb-4">Data Peminjaman</h1>

      <!-- Tombol untuk membuka modal filter -->
      <button class="btn btn-primary mb-4" @click="openFilterModal">Filter Data</button>

      <!-- Modal Filter -->
      <div v-if="isModalOpen" class="modal-overlay" @click="closeFilterModal">
        <div class="modal-content" @click.stop>
          <h4 class="text-center">Filter Peminjaman</h4>

          <div class="filter-section">
            <!-- Filter Berdasarkan Nama Buku -->
            <label for="namaBuku" class=" fw-bold">Filter berdasarkan nama buku:</label>
            <select id="namaBuku" class="form-select  mx-4" v-model="selectedBookName">
              <option value="">Pilih Nama Buku</option>
              <option v-for="book in books" :key="book.id_buku" :value="book.judul">{{ book.judul }}</option>
            </select>

            <!-- Filter Berdasarkan ID User -->
            <label for="idUser" class=" fw-bold ">Filter berdasarkan ID User:</label>
            <input type="number" id="idUser" class="form-control  mx-4" v-model="selectedUserId" placeholder="Masukkan ID User" />

            <!-- Filter Berdasarkan Tanggal Pinjam -->
            <label for="tanggal" class=" fw-bold ">Filter berdasarkan tanggal pinjam:</label>
            <input type="date" id="tanggal" class="form-control  mx-4" style="max-width: 200px;" v-model="selectedDate" />

            <!-- Filter Berdasarkan Tanggal Kembali -->
            <label for="tanggalKembali" class=" fw-bold ">Filter berdasarkan tanggal kembali:</label>
            <input type="date" id="tanggalKembali" class="form-control  mx-4" style="max-width: 200px;" v-model="selectedReturnDate" />

            <!-- Filter Berdasarkan Status -->
            <label for="status" class="fw-bold">Filter berdasarkan status:</label>
            <select id="status" class="form-select  mx-4" v-model="selectedStatus">
              <option value="">Semua Status Pengembalian</option>
              <option value="Pending">Pending</option>
              <option value="Approved">Approved</option>
              <option value="Rejected">Rejected</option>
            </select>

            <!-- Filter Berdasarkan Status Peminjaman -->
            <label for="statusPinjam" class=" fw-bold">Filter berdasarkan status peminjaman:</label>
            <select id="statusPinjam" class="form-select  mx-4" v-model="filterStatus">
              <option value="">Semua Status Pinjaman</option>
              <option value="Dipinjam">Dipinjam</option>
              <option value="Dikembalikan">Dikembalikan</option>
            </select>

            <div class="d-flex justify-content-between">
              <button class="btn bg-ijomuda mt-3 text-white" @click="applyFilters">Terapkan Filter</button>
              <button class="btn btn-secondary mt-3" @click="resetFilter">Reset Filter</button>
            </div>
          </div>

          <button class="btn btn-danger " @click="closeFilterModal">Tutup</button>
        </div>
      </div>

      <table class="table table-striped table-bordered table-hover">
        <thead class="bg-ijomuda text-white">
          <tr>
            <th>#</th>
            <th>ID User</th>
            <th>Buku</th>
            <th>Jumlah Buku</th>
            <th>Tanggal Pinjam</th>
            <th>Tanggal Kembali</th>
            <th>Status Peminjaman</th>
            <th>Status Pengembalian</th>
            <th>Denda</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(item, index) in filteredData" :key="item.id_peminjaman">
            <td>{{ index + 1 }}</td>
            <td>{{ item.id_user }}</td>
            <td>{{ item.buku.judul }}</td>
            <td>{{ item.jumlah_pinjam }}</td>
            <td>{{ new Date(item.tanggal_pinjam).toLocaleDateString() }}</td>
            <td>{{ new Date(item.tanggal_kembali).toLocaleDateString() }}</td>
            <td>
              <span :class="item.status_kembali ? 'text-success' : 'text-danger'">
                {{ item.status_kembali ? 'Dikembalikan' : 'Dipinjam' }}
              </span>
            </td>
            <td>
              <span :class="getStatusClass(item.status_pengembalian)">
                {{ item.status_pengembalian }}
              </span>
            </td>
            <td>Rp {{ item.denda.toLocaleString() }}</td>
          </tr>
          <tr v-if="filteredData.length === 0">
            <td colspan="9" class="text-center">Tidak ada data peminjaman untuk filter yang dipilih</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
import { getAllPeminjaman } from "@/api/peminjaman";
import AppNavbar from "~/components/AppNavbar.vue";

export default {
  components: {
    AppNavbar,
  },
  data() {
    return {
      dataPeminjaman: [],
      books: [], // Daftar buku untuk dropdown
      selectedDate: "",
      selectedStatus: "",
      selectedReturnDate: "",
      filterStatus: "",
      selectedBookName: "", // Filter berdasarkan nama buku
      selectedUserId: "", // Filter berdasarkan ID User
      filteredData: [],
      isModalOpen: false,
    };
  },
  async mounted() {
    try {
      this.dataPeminjaman = await getAllPeminjaman();
      this.books = this.dataPeminjaman.map(item => item.buku); // Mengambil daftar buku
      this.filteredData = this.dataPeminjaman;
    } catch (error) {
      console.error("Gagal memuat data peminjaman:", error);
    }
  },
  methods: {
    openFilterModal() {
      this.isModalOpen = true;
    },
    closeFilterModal() {
      this.isModalOpen = false;
    },
    applyFilters() {
      this.filterData(); // Panggil filterData setelah memilih filter
      this.closeFilterModal(); // Tutup modal setelah filter diterapkan
    },
    filterData() {
      let filtered = this.dataPeminjaman;

      // Filter berdasarkan nama buku
      if (this.selectedBookName) {
        filtered = filtered.filter((item) =>
          item.buku.judul.toLowerCase().includes(this.selectedBookName.toLowerCase())
        );
      }

      // Filter berdasarkan ID User
      if (this.selectedUserId) {
        filtered = filtered.filter((item) => item.id_user === parseInt(this.selectedUserId));
      }

      // Filter berdasarkan tanggal pinjam
      if (this.selectedDate) {
        filtered = filtered.filter(
          (item) =>
            new Date(item.tanggal_pinjam).toISOString().split("T")[0] ===
            this.selectedDate
        );
      }

      // Filter berdasarkan status pengembalian
      if (this.selectedStatus) {
        filtered = filtered.filter(
          (item) => item.status_pengembalian === this.selectedStatus
        );
      }

      // Filter berdasarkan status peminjaman (Dipinjam / Dikembalikan)
      if (this.filterStatus) {
        filtered = filtered.filter(
          (item) =>
            (this.filterStatus === "Dipinjam" && !item.status_kembali) ||
            (this.filterStatus === "Dikembalikan" && item.status_kembali)
        );
      }

      // Filter berdasarkan tanggal kembali
      if (this.selectedReturnDate) {
        filtered = filtered.filter(
          (item) =>
            new Date(item.tanggal_kembali).toISOString().split("T")[0] ===
            this.selectedReturnDate
        );
      }

      this.filteredData = filtered;
    },
    resetFilter() {
      this.selectedDate = "";
      this.selectedStatus = "";
      this.selectedReturnDate = "";
      this.filterStatus = "";
      this.selectedBookName = ""; // Reset nama buku
      this.selectedUserId = ""; // Reset ID User
      this.filteredData = this.dataPeminjaman;
    },
    getStatusClass(status) {
      return status === "Pending"
        ? "text-warning"
        : status === "Approved"
          ? "text-success"
          : "text-danger";
    },
  },
};
</script>

<style scoped>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.4); /* Lebih transparan */
  z-index: 1000;
  display: flex;
  justify-content: center;
  align-items: center;
  transition: opacity 0.3s ease; /* Efek transisi */
  z-index: 9999999;
}

.modal-content {
  background-color: #fff;
  padding: 30px;
  width: 500px;
  border-radius: 12px; /* Sudut lebih melengkung */
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2); /* Bayangan */
  animation: modalFadeIn 0.3s ease-out; /* Animasi muncul */
}

@keyframes modalFadeIn {
  from {
    opacity: 0;
    transform: scale(0.8);
  }
  to {
    opacity: 1;
    transform: scale(1);
  }
}

h4 {
  font-size: 1.5rem;
  margin-bottom: 20px;
  text-align: center;
}

label{
  margin: 0;
  margin-top: 5px;
  color: #bbff00;
}

.filter-section {
  margin-bottom: 20px;
}

.filter-section label {
  font-weight: bold;
  color: #333;
}

.filter-section input, .filter-section select {
  width: calc(100% - 40px);
  padding: 10px;
  margin-top: 8px;
  border: 1px solid #ccc;
  border-radius: 8px;
  box-sizing: border-box;
}

button {
  font-size: 1rem;
  padding: 10px 20px;
  border-radius: 8px;
}

button.btn-primary {
  background-color: #007bff;
  border: none;
  color: white;
  transition: background-color 0.3s;
}

button.btn-primary:hover {
  background-color: #0056b3;
}

button.btn-secondary {
  background-color: #6c757d;
  border: none;
  color: white;
  transition: background-color 0.3s;
}

button.btn-secondary:hover {
  background-color: #5a6268;
}



button.btn-danger {
  background-color: #e74c3c;
  border: none;
  color: white;
}

button.btn-danger:hover {
  background-color: #c0392b;
}

.d-flex {
  justify-content: flex-end;
}

.d-flex button {
  margin-left: 10px;
}
</style>
