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

      <div v-else-if="buku" class="card shadow-lg border-0 rounded-lg p-4">
        <div class="row no-gutters align-items-center">
          <div class="col-md-4">
            <img v-if="buku.gambar"
              :src="buku.gambar.startsWith('http') ? buku.gambar : `http://localhost:8080${buku.gambar}`"
              :alt="buku.judul" class="card-img rounded-lg" style="object-fit: cover; height: 100%;" />
            <div v-else class="text-center p-3 text-muted">Tidak ada gambar</div>
          </div>
          <div class="col-md-4">
            <div class="card-body text-center">
              <h3 class="card-title text-primary font-irish">{{ buku.judul }}</h3>
              <p class="card-text"><strong>Penulis:</strong> {{ buku.penulis.nama || 'Tidak diketahui' }}</p>
              <p class="card-text"><strong>Genre:</strong> {{ buku.genre.nama || 'Tidak diketahui' }}</p>
              <p class="card-text"><strong>Jumlah:</strong> {{ buku.jumlah }}</p>
              <button @click="showModal = true" class="btn bg-ijomuda mt-3 px-4 py-2 rounded-3 shadow-sm custom-hover">
                Pinjam Buku
              </button>
            </div>
          </div>
          <div class="col-md-4 text-center">
            <button @click="showDeskripsiModal = true"
              class="btn bg-ijotua mt-2 px-4 py-2 rounded-3 shadow-sm custom-hover">
              Lihat Deskripsi Buku <b-icon-eye></b-icon-eye>
            </button>
            <button @click="showBiografiModal = true"
              class="btn bg-kuning mt-2 px-4 py-2 rounded-3 shadow-sm custom-hover">
              Lihat Biografi Penulis <b-icon-eye></b-icon-eye>
            </button>
          </div>
        </div>
      </div>

      <!-- Modal Peminjaman -->
      <div v-if="showModal" class="modal fade show d-block rounded-lg" tabindex="-1" role="dialog"
        style="background: rgba(0, 0, 0, 0.5);">
        <div class="modal-dialog mt-8" role="document">
          <div class="modal-content">
            <div class="modal-header bg-ijomuda text-white">
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
                  <input type="number" v-model.number="form.durasiHari" placeholder="Maksimal 7 hari"
                    class="form-control" required min="1" max="7" />
                </div>
                <div class="form-group">
                  <label for="jumlahPinjam">Jumlah Buku</label>
                  <input type="number" id="jumlahPinjam" v-model.number="form.jumlahPinjam"
                    placeholder="Masukkan jumlah buku" class="form-control" required min="1" :max="buku.jumlah" />
                </div>
                <div v-if="user.role === 'petugas'" class="form-group">
                  <div class="form-group">
                    <label for="searchUser">Cari User (berdasarkan email)</label>
                    <div class="input-group">
                      <input type="text" id="searchUser" v-model="searchUser" placeholder="Masukkan email user"
                        class="form-control" @input="searchUserByEmail" />
                    </div>
                    <!-- Dropdown hasil pencarian -->
                    <div v-if="filteredUsers.length > 0" class="dropdown-menu show mt-1">
                      <button class="dropdown-item" v-for="user in filteredUsers" :key="user.id_user"
                        @click="selectUser(user)">
                        {{ user.email }}
                      </button>
                    </div>
                    <!-- Pesan error jika tidak ada user ditemukan -->
                    <div v-if="
                      searchUser.trim() &&
                      filteredUsers.length === 0 &&
                      !form.selectedUser" class="mt-2 text-danger">
                      Tidak ditemukan user dengan email tersebut.
                    </div>
                  </div>
                </div>
                <button type="submit"
                  class="btn bg-ijomuda text-white btn-block py-2 mt-3 rounded-3 shadow-sm">Kirim</button>
              </form>
            </div>
          </div>
        </div>
      </div>

      <!-- Modal Deskripsi Buku -->
      <div v-if="showDeskripsiModal" class="modal fade show d-block" tabindex="-1" role="dialog"
        style="background: rgba(0, 0, 0, 0.5);">
        <div class="modal-dialog mt-8" role="document">
          <div class="modal-content">
            <div class="modal-header bg-info text-white">
              <h5 class="modal-title">Deskripsi Buku</h5>
              <button type="button" class="close" @click="showDeskripsiModal = false" aria-label="Close">
                <span aria-hidden="true">&times;</span>
              </button>
            </div>
            <div class="modal-body">
              <p>{{ buku.deskripsi || 'Deskripsi tidak tersedia' }}</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Modal Biografi Penulis -->
      <div v-if="showBiografiModal" class="modal fade show d-block" tabindex="-1" role="dialog"
        style="background: rgba(0, 0, 0, 0.5);">
        <div class="modal-dialog mt-8" role="document">
          <div class="modal-content">
            <div class="modal-header bg-secondary text-white">
              <h5 class="modal-title">Biografi Penulis</h5>
              <button type="button" class="close" @click="showBiografiModal = false" aria-label="Close">
                <span aria-hidden="true">&times;</span>
              </button>
            </div>
            <div class="modal-body">
              <p>{{ buku.penulis.biografi || 'Biografi tidak tersedia' }}</p>
            </div>
          </div>
        </div>
      </div>
    </div>
    <!-- Notification Modal for Success or Failure -->
    <NotificationModal v-if="showSuccessModal" :isVisible="showSuccessModal" :messageTitle="notificationTitle"
      :messageBody="notificationMessage" @close="closeNotificationModal" />
  </div>
</template>

<script>
import { getBukuById } from "@/api/buku";
import { createPeminjaman } from "@/api/peminjaman";
import { getUsers } from "@/api/user";
import AppNavbar from "~/components/AppNavbar.vue";
import NotificationModal from "~/components/NotificationModal.vue";

export default {
  components: {
    AppNavbar,
    NotificationModal,
  },
  data() {
    return {
      buku: null,
      loading: false,
      error: null,
      showModal: false,
      showDeskripsiModal: false,
      showBiografiModal: false,
      showSuccessModal: false,
      notificationTitle: "",
      notificationMessage: "",
      users: [],
      form: {
        durasiHari: "",
        jumlahPinjam: "",
        selectedUser: "",
      },
      searchUser: "",
      filteredUsers: [],
      userNotFound: false,
    };
  },
  computed: {
    user() {
      return this.$store.getters.getUser;
    },
  },
  methods: {
    async fetchBuku() {
      try {
        this.loading = true;
        this.error = null;
        const id = this.$route.params.id_buku;
        if (!id) {
          throw new Error("ID buku tidak diberikan");
        }
        const data = await getBukuById(id);
        this.buku = data;
      } catch (error) {
        console.error(error);
        this.error = "Gagal memuat detail buku. Coba lagi nanti.";
      } finally {
        this.loading = false;
      }
    },
    async fetchCurrentUser() {
      try {
        if (this.user.role === "petugas") {
          const data = await getUsers();
          console.log("Fetched users:", data); // Log hasil pengambilan data
          this.users = data || [];
        }
      } catch (error) {
        console.error("Error fetching users:", error);
      }
    },

    searchUserByEmail() {
      // Jika input kosong, kosongkan filteredUsers dan keluar
      if (!this.searchUser.trim()) {
        this.filteredUsers = [];
        return;
      }
      // Filter user berdasarkan email
      this.filteredUsers = this.users.filter((user) =>
        user.email.toLowerCase().includes(this.searchUser.toLowerCase())
      );
    },
    selectUser(user) {
      // Simpan user yang dipilih ke dalam form
      this.form.selectedUser = user.id_user;
      this.searchUser = user.email; // Perbarui input dengan email user yang dipilih
      this.filteredUsers = []; // Kosongkan dropdown untuk menyembunyikannya
      this.error = null; // Pastikan pesan error di-reset
    },
    async submitPeminjaman() {
      try {
        if (!this.form.durasiHari || !this.form.jumlahPinjam) {
          this.showFailureNotification("Error", "Semua kolom harus diisi");
          return;
        }

        if (this.form.jumlahPinjam > this.buku.jumlah) {
          this.showFailureNotification(
            "Error",
            "Jumlah buku melebihi stok tersedia"
          );
          return;
        }

        // Tentukan ID user yang meminjam
        const idUserPeminjam = this.form.selectedUser || this.user.id_user;

        const peminjamanData = {
          id_user: idUserPeminjam, // Gunakan ID user yang terpilih
          id_buku: this.buku.id_buku,
          durasi_hari: this.form.durasiHari,
          jumlah_pinjam: this.form.jumlahPinjam,
        };

        this.loading = true;

        const response = await createPeminjaman(peminjamanData);

        if (response && response.id_peminjaman) {
          this.resetPeminjamanForm(); // Reset form jika peminjaman berhasil
          this.showSuccessNotification(
            "Peminjaman Berhasil",
            "Peminjaman buku berhasil dilakukan."
          );
          this.showModal = false; // Tutup modal peminjaman
        } else {
          this.showFailureNotification(
            "Error",
            "Terjadi kesalahan saat membuat peminjaman"
          );
        }
      } catch (error) {
        console.error(error);
        this.showFailureNotification("Error", "Terjadi kesalahan saat peminjaman");
      } finally {
        this.loading = false;
      }
    },

    showSuccessNotification(title, message) {
      this.notificationTitle = title;
      this.notificationMessage = message;
      this.showSuccessModal = true;
    },
    showFailureNotification(title, message) {
      this.notificationTitle = title;
      this.notificationMessage = message;
      this.showSuccessModal = false;
      this.showModal = false;
    },
    closeNotificationModal() {
      this.showSuccessModal = false;
      const userRole = this.$store.state.user.role;
      if (userRole === "petugas") {
        this.$router.push("/admin/datapeminjaman");
      } else {
        this.$router.push("/BukuPinjaman");
      }
    },

    resetPeminjamanForm() {
      this.form.durasiHari = "";
      this.form.jumlahPinjam = "";
      this.form.user = ""; // Reset user form
    },
  },
  watch: {
    $route: "fetchBuku",
  },
  created() {
    this.fetchBuku();
    this.fetchCurrentUser();
  },
};
</script>

<style scoped>
.card {
  border-radius: 16px;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.2);
  overflow: hidden;
}

.card-img {
  border-radius: 16px;
}

.spinner-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 300px;
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


.modal {
  display: flex;
  align-items: center;
  justify-content: center;
  animation: fadeIn 0.3s ease-in-out;
  z-index: 99999;
}

.modal-content {
  position: relative;
  border-radius: 16px;
  box-shadow: 0 8px 20px rgba(0, 0, 0, 0.3);
}

@keyframes fadeIn {
  from {
    opacity: 0;
  }

  to {
    opacity: 1;
  }
}

.card-title {
  font-size: 60px;
}
</style>
