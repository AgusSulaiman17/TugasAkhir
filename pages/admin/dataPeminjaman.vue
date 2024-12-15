<template>
  <div>
    <AppNavbar />
    <div class="mt-8 admin-peminjaman mx-auto" style="max-width: 1200px;">
      <h1 class="text-center mb-4">Data Peminjaman</h1>
      <div class="d-flex justify-content-between align-items-center mb-4">
        <button class="btn bg-ijotua text-white" @click="openFilterModal">
          Filter Data <b-icon-plus></b-icon-plus>
        </button>
        <div class="d-flex">
          <button class="btn bg-kuning text-white me-2" @click="exportToExcel">
            Ekspor ke Excel <b-icon icon="file-earmark-spreadsheet"></b-icon>
          </button>
        </div>
      </div>

      <!-- Modal Filter -->
      <div v-if="isModalOpen" class="modal-overlay" @click="closeFilterModal">
        <div class="modal-content" @click.stop>
          <h4 class="text-center">Filter Peminjaman </h4>
          <div class="filter-section">
            <label for="namaBuku" class="fw-bold">Filter berdasarkan nama buku:</label>
            <select id="namaBuku" class="form-select mx-4" v-model="selectedBookName">
              <option value="">Semua buku</option>
              <option v-for="book in books" :key="book.id_buku" :value="book.judul">{{ book.judul }}</option>
            </select>

            <label for="emailUser" class="fw-bold">Filter berdasarkan email pengguna:</label>
            <div class="d-flex">
              <input
                type="text"
                placeholder="Cari email..."
                class="form-control mx-4"
                v-model="searchEmail"
                @input="filterData"
              />
              <select id="emailUser" class="form-select mx-4" v-model="selectedUserEmail">
                <option value="">Semua email pengguna</option>
                <option v-for="user in filteredUsers" :key="user.id_user" :value="user.email">{{ user.email }}</option>
              </select>
            </div>

            <label for="tanggal" class="fw-bold">Filter berdasarkan tanggal pinjam:</label>
            <input type="date" id="tanggal" class="form-control mx-4" style="max-width: 200px;" v-model="selectedDate" />

            <label for="tanggalKembali" class="fw-bold">Filter berdasarkan tanggal kembali:</label>
            <input type="date" id="tanggalKembali" class="form-control mx-4" style="max-width: 200px;" v-model="selectedReturnDate" />

            <label for="status" class="fw-bold">Filter berdasarkan status:</label>
            <select id="status" class="form-select mx-4" v-model="selectedStatus">
              <option value="">Semua Status Pengembalian</option>
              <option value="Pending">Pending</option>
              <option value="Approved">Approved</option>
              <option value="Rejected">Rejected</option>
            </select>

            <div class="d-flex justify-content-between">
              <button class="btn bg-ijomuda mt-3 text-white" @click="applyFilters">Terapkan Filter</button>
              <button class="btn bg-kuning mt-3" @click="resetFilter">Reset Filter</button>
            </div>
          </div>

          <button class="btn bg-merah" @click="closeFilterModal">Tutup</button>
        </div>
      </div>

      <!-- Modal Update Status Pengembalian -->
      <div v-if="isUpdateModalOpen" class="modal-overlay" @click="closeUpdateModal">
        <div class="modal-content" @click.stop>
          <h4 class="text-center">Update Status Pengembalian</h4>
          <div class="filter-section">
            <label for="updateStatusPengembalian" class="fw-bold">Status Pengembalian:</label>
            <select id="updateStatusPengembalian" class="form-select mx-4" v-model="updateStatusPengembalian">
              <option value="Pending">Pending</option>
              <option value="Approved">Approved</option>
              <option value="Rejected">Rejected</option>
            </select>

            <div class="d-flex justify-content-between">
              <button class="btn bg-ijomuda mt-3 text-white" @click="submitUpdateStatus">Simpan Perubahan</button>
              <button class="btn bg-kuning mt-3" @click="closeUpdateModal">Tutup</button>
            </div>
          </div>
        </div>
      </div>

      <table class="table table-striped table-bordered table-hover">
        <thead class="bg-ijomuda text-white">
          <tr>
            <th>#</th>
            <th>Buku</th>
            <th>Jumlah Buku</th>
            <th>Tanggal Pinjam</th>
            <th>Tanggal Kembali</th>
            <th>Status Pengembalian</th>
            <th>Denda</th>
            <th>Aksi</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(item, index) in filteredData" :key="item.id_peminjaman">
            <td>{{ index + 1 }}</td>
            <td>{{ item.buku.judul }}</td>
            <td>{{ item.jumlah_pinjam }}</td>
            <td>{{ new Date(item.tanggal_pinjam).toLocaleDateString() }}</td>
            <td>{{ new Date(item.tanggal_kembali).toLocaleDateString() }}</td>
            <td>
              <span :class="getStatusClass(item.status_pengembalian)">
                {{ item.status_pengembalian }}
              </span>
            </td>
            <td>Rp {{ item.denda.toLocaleString() }}</td>
            <td class="d-flex align-items-center">
              <button @click="openUpdateModal(item)" class="btn bg-kuning btn-sm"><b-icon-pencil></b-icon-pencil></button>
              <button @click="deletePeminjaman(item.id_peminjaman)" class="btn bg-merah btn-sm"><b-icon-trash></b-icon-trash></button>
            </td>
          </tr>
          <tr v-if="filteredData.length === 0">
            <td colspan="8" class="text-center">Tidak ada data peminjaman untuk filter yang dipilih</td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Success Notification Modal for Edit -->
    <NotificationModal
      v-if="showSuccessEditModal"
      :isVisible="showSuccessEditModal"
      :messageTitle="'Edit Berhasil'"
      :messageBody="successEditMessage"
      @close="closeSuccessEditModal"
    />

    <!-- Success Notification Modal for Delete -->
    <NotificationModal
      v-if="showSuccessModal"
      :isVisible="showSuccessModal"
      :messageTitle="successTitle"
      :messageBody="successMessage"
      @close="closeSuccessModal"
    />

    <!-- Delete Confirmation Modal -->
    <NotificationModal
      v-if="showDeleteConfirmation"
      :isVisible="showDeleteConfirmation"
      :messageTitle="'Konfirmasi Hapus'"
      :messageBody="'Apakah Anda yakin ingin menghapus peminjaman ini?'"
      @close="cancelDelete"
    >
      <template #footer>
        <button @click="confirmDelete" class="btn btn-abu">Ya, Hapus</button>
        <button @click="cancelDelete" class="btn btn-cancel">Batal</button>
      </template>
    </NotificationModal>
  </div>
</template>

<script>
import { getAllPeminjaman, updatePeminjaman, deletePeminjaman as deletePeminjamanApi } from "@/api/peminjaman";
import { getBukuList } from "@/api/buku";
import { getUsers } from "@/api/user"; // Import data pengguna
import AppNavbar from "~/components/AppNavbar.vue";
import * as XLSX from 'xlsx';

export default {
  components: {
    AppNavbar,
  },
  data() {
    return {
      dataPeminjaman: [],
      books: [], // Daftar buku untuk dropdown
      users: [], // Daftar pengguna untuk dropdown email
      selectedDate: "",
      selectedStatus: "",
      selectedReturnDate: "",
      filterStatus: "",
      selectedBookName: "", // Filter berdasarkan nama buku
      selectedUserId: "", // Filter berdasarkan ID User
      selectedUserEmail: "", // Filter berdasarkan email pengguna
      filteredData: [],
      isModalOpen: false,
      searchEmail: "", // State untuk pencarian email
      filteredUsers: [], // Users setelah pencarian
      isUpdateModalOpen: false,
      updateStatusPengembalian: "", // Status pengembalian yang akan diupdate
      currentPeminjamanId: null, // ID peminjaman yang sedang diupdate
      showSuccessModal: false, // State untuk modal notifikasi sukses hapus
      successTitle: '', // Judul untuk modal notifikasi sukses hapus
      successMessage: '', // Pesan untuk modal notifikasi sukses hapus
      showSuccessEditModal: false, // State untuk modal notifikasi sukses edit
      successEditMessage: 'Data berhasil diperbarui.', // Pesan untuk modal notifikasi sukses edit
      showDeleteConfirmation: false, // State untuk modal konfirmasi hapus
    };
  },
  async mounted() {
    try {
      // Ambil data peminjaman
      this.dataPeminjaman = await getAllPeminjaman();
      this.filteredData = this.dataPeminjaman;

      // Ambil data buku untuk dropdown
      const bukuList = await getBukuList();
      this.books = bukuList;

      // Ambil data pengguna untuk dropdown email
      const userList = await getUsers(); // Ambil data pengguna
      this.users = userList;
    } catch (error) {
      console.error("Gagal memuat data:", error);
    }
  },
  watch: {
    searchEmail(value) {
      this.filteredUsers = this.users.filter(user =>
        user.email.toLowerCase().includes(value.toLowerCase())
      );
      // Re-filter data berdasarkan email pencarian
      this.filterData();
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
          item.buku.judul.toLowerCase() === this.selectedBookName.toLowerCase()
        );
      }

      // Filter berdasarkan ID User
      if (this.selectedUserId) {
        filtered = filtered.filter((item) => item.id_user === parseInt(this.selectedUserId));
      }

      // Filter berdasarkan email pengguna
      if (this.selectedUserEmail) {
        filtered = filtered.filter((item) => item.user.email === this.selectedUserEmail);
      }

      // Filter berdasarkan tanggal pinjam
      if (this.selectedDate) {
        filtered = filtered.filter(
          (item) => new Date(item.tanggal_pinjam).toLocaleDateString() === new Date(this.selectedDate).toLocaleDateString()
        );
      }

      // Filter berdasarkan tanggal kembali
      if (this.selectedReturnDate) {
        filtered = filtered.filter(
          (item) => new Date(item.tanggal_kembali).toLocaleDateString() === new Date(this.selectedReturnDate).toLocaleDateString()
        );
      }

      // Filter berdasarkan status pengembalian
      if (this.selectedStatus) {
        filtered = filtered.filter(
          (item) => item.status_pengembalian === this.selectedStatus
        );
      }

      // Filter berdasarkan status peminjaman
      if (this.filterStatus) {
        filtered = filtered.filter(
          (item) => item.status === this.filterStatus
        );
      }

      this.filteredData = filtered;
    },
    resetFilter() {
      this.selectedBookName = "";
      this.selectedUserEmail = "";
      this.selectedDate = "";
      this.selectedReturnDate = "";
      this.selectedStatus = "";
      this.filterStatus = "";
      this.searchEmail = "";
      this.filteredData = this.dataPeminjaman;
    },
    exportToExcel() {
      // Membuat array baru untuk menyimpan data yang akan diekspor
      const exportData = this.filteredData.map((item, index) => ({
        No: index + 1,
        Buku: item.buku.judul,
        Jumlah: item.jumlah_pinjam,
        TanggalPinjam: new Date(item.tanggal_pinjam).toLocaleDateString(),
        TanggalKembali: new Date(item.tanggal_kembali).toLocaleDateString(),
        StatusPengembalian: item.status_pengembalian,
        Denda: `Rp ${item.denda.toLocaleString()}`,
      }));

      const worksheet = XLSX.utils.json_to_sheet(exportData);
      const workbook = XLSX.utils.book_new();
      XLSX.utils.book_append_sheet(workbook, worksheet, 'Data Peminjaman');
      XLSX.writeFile(workbook, 'Data Peminjaman.xlsx');
    },
    getStatusClass(status) {
      if (status === "Pending") return "text-warning";
      if (status === "Approved") return "text-success";
      if (status === "Rejected") return "text-danger";
      return "";
    },
    openUpdateModal(item) {
      this.currentPeminjamanId = item.id_peminjaman;
      this.updateStatusPengembalian = item.status_pengembalian; // Mengisi status pengembalian saat modal dibuka
      this.isUpdateModalOpen = true;
    },
    closeUpdateModal() {
      this.isUpdateModalOpen = false;
      this.resetUpdateForm();
    },
    resetUpdateForm() {
      this.updateStatusPengembalian = "";
      this.currentPeminjamanId = null;
    },
    async submitUpdateStatus() {
      try {
        // Kirim pembaruan status pengembalian ke API
        const updatedData = await updatePeminjaman(this.currentPeminjamanId, {
          status_pengembalian: this.updateStatusPengembalian,
        });

        console.log("Respons dari server setelah pembaruan:", updatedData);

        // Pastikan respons berisi data yang diharapkan
        if (updatedData) {
          // Temukan item yang diperbarui dalam filteredData dan perbarui statusnya
          const index = this.filteredData.findIndex(item => item.id_peminjaman === this.currentPeminjamanId);
          if (index !== -1) {
            // Ganti item yang diperbarui dengan data dari respons
            this.$set(this.filteredData, index, updatedData); // Perbarui item dengan data terbaru dari server

            // Ubah status peminjaman dan status kembali berdasarkan status pengembalian
            if (updatedData.status_pengembalian === 'Approved' || updatedData.status_pengembalian === 'Pending') {
              console.log("Mengubah status peminjaman menjadi 'Dikembalikan' dan status kembali menjadi true");
              this.filteredData[index].status = 'Dikembalikan'; // Ubah status peminjaman menjadi 'Dikembalikan'
              this.filteredData[index].status_kembali = true; // Ubah status kembali menjadi true
            } else if (updatedData.status_pengembalian === 'Rejected') {
              console.log("Mengubah status peminjaman menjadi 'Dipinjam' dan status kembali menjadi false");
              this.filteredData[index].status = 'Dipinjam'; // Ubah status peminjaman menjadi 'Dipinjam'
              this.filteredData[index].status_kembali = false; // Ubah status kembali menjadi false
            }
          }

          // Tampilkan modal notifikasi sukses edit
          this.showSuccessEditModal = true; // Tampilkan modal sukses edit

          // Ambil kembali data peminjaman setelah pembaruan
          this.dataPeminjaman = await getAllPeminjaman();
          this.filteredData = this.dataPeminjaman; // Perbarui filteredData
        }

        this.closeUpdateModal();
      } catch (error) {
        console.error("Gagal memperbarui status pengembalian:", error);
      }
    },
    async deletePeminjaman(id) {
      this.showDeleteConfirmation = true; // Tampilkan modal konfirmasi hapus
      this.currentPeminjamanId = id; // Simpan ID peminjaman yang akan dihapus
    },
    async confirmDelete() {
      try {
        await deletePeminjamanApi(this.currentPeminjamanId);
        this.filteredData = this.filteredData.filter(item => item.id_peminjaman !== this.currentPeminjamanId);
        this.showSuccessModal = true; // Tampilkan modal sukses setelah penghapusan
        this.successTitle = 'Success';
        this.successMessage = 'Data berhasil dihapus.';
      } catch (error) {
        console.error("Gagal menghapus peminjaman:", error);
      } finally {
        this.showDeleteConfirmation = false; // Tutup modal konfirmasi hapus
      }
    },
    cancelDelete() {
      this.showDeleteConfirmation = false; // Tutup modal konfirmasi hapus
    },
    closeSuccessModal() {
      this.showSuccessModal = false; // Tutup modal notifikasi sukses
    },
    closeSuccessEditModal() {
      this.showSuccessEditModal = false; // Tutup modal notifikasi sukses edit
    },
  },
};
</script>

<style scoped>
.admin-peminjaman {
  max-width: 1200px;
  padding: 20px;
  background-color: #f9f9f9;
  border-radius: 8px;
}

.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.4);
  z-index: 999999;
  display: flex;
  justify-content: center;
  align-items: center;
  transition: opacity 0.3s ease;
}

.modal-content {
  background-color: #fff;
  padding: 30px;
  width: 500px;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2);
  animation: modalFadeIn 0.3s ease-out;
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

label {
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

.filter-section input,
.filter-section select {
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

.d-flex button {
  margin-left: 10px;
}
</style>
