<template>
  <div>
    <AppNavbar />
    <div class="admin-peminjaman mt-8 mx-auto" style="max-width: 1200px;">
      <h1 class="text-center mb-4">Data Peminjaman</h1>
      <div class="filter-section d-flex justify-content-center align-items-center mb-4">
        <label for="tanggal" class="me-3 fw-bold">Filter berdasarkan tanggal:</label>
        <input type="date" id="tanggal" class="form-control me-3" style="max-width: 200px;" v-model="selectedDate" @change="filterByDate" />
        <button class="btn btn-secondary" @click="resetFilter">Semua Tanggal</button>
      </div>
      <table class="table table-striped table-bordered table-hover">
        <thead class="table-dark">
          <tr>
            <th>#</th>
            <th>ID User</th>
            <th>ID Buku</th>
            <th>Tanggal Pinjam</th>
            <th>Tanggal Kembali</th>
            <th>Status Pengembalian</th>
            <th>Denda</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(item, index) in filteredData" :key="item.id_peminjaman">
            <td>{{ index + 1 }}</td>
            <td>{{ item.id_user }}</td>
            <td>{{ item.id_buku }}</td>
            <td>{{ new Date(item.tanggal_pinjam).toLocaleDateString() }}</td>
            <td>{{ new Date(item.tanggal_kembali).toLocaleDateString() }}</td>
            <td>
              <span :class="getStatusClass(item.status_pengembalian)">
                {{ item.status_pengembalian }}
              </span>
            </td>
            <td>Rp {{ item.denda.toLocaleString() }}</td>
          </tr>
          <tr v-if="filteredData.length === 0">
            <td colspan="7" class="text-center">Tidak ada data peminjaman untuk tanggal ini</td>
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
      selectedDate: "",
      filteredData: [],
    };
  },
  async mounted() {
    try {
      this.dataPeminjaman = await getAllPeminjaman();
      this.filteredData = this.dataPeminjaman; // Tampilkan semua data awalnya
    } catch (error) {
      console.error("Gagal memuat data peminjaman:", error);
    }
  },
  methods: {
    filterByDate() {
      if (this.selectedDate) {
        this.filteredData = this.dataPeminjaman.filter(
          (item) =>
            new Date(item.tanggal_pinjam).toISOString().split("T")[0] ===
            this.selectedDate
        );
      } else {
        this.filteredData = this.dataPeminjaman;
      }
    },
    resetFilter() {
      this.selectedDate = "";
      this.filteredData = this.dataPeminjaman;
    },
    getStatusClass(status) {
      return status === "Pending"
        ? "text-warning"
        : status === "Selesai"
        ? "text-success"
        : "text-danger";
    },
  },
};
</script>

<style>
.admin-peminjaman {
  padding: 20px;
  background-color: #f9f9f9;
  border-radius: 8px;
  box-shadow: 0px 4px 6px rgba(0, 0, 0, 0.1);
}

.table th,
.table td {
  text-align: center;
  vertical-align: middle;
}

.filter-section {
  margin-bottom: 20px;
}

h1 {
  font-size: 1.8rem;
  font-weight: bold;
  color: #333;
}
</style>
