<template>
  <div class="container mt-5">
    <div class="row">
      <div class="col-md-6">
        <h2>Profile</h2>
        <div v-if="user">
          <p><strong>Nama:</strong> {{ user.nama }}</p>
          <p><strong>Email:</strong> {{ user.email }}</p>
          <p><strong>Role:</strong> {{ user.role }}</p>
          <button class="btn btn-primary" @click="openEditModal">Edit Profile</button>
        </div>
        <div v-else>
          <p>Loading...</p>
        </div>
      </div>
    </div>

    <!-- Modal Edit Profile -->
    <div v-if="isModalOpen" class="modal fade show" tabindex="-1" role="dialog" style="display: block;">
      <div class="modal-dialog" role="document">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">Edit Profile</h5>
            <button type="button" class="close" @click="closeEditModal" aria-label="Close">
              <span aria-hidden="true">&times;</span>
            </button>
          </div>
          <div class="modal-body">
            <form @submit.prevent="updateUserProfile">
              <div class="form-group">
                <label for="nama">Nama</label>
                <input v-model="formData.nama" type="text" class="form-control" id="nama" required />
              </div>
              <div class="form-group">
                <label for="email">Email</label>
                <input v-model="formData.email" type="email" class="form-control" id="email" required />
              </div>
              <div class="form-group">
                <label for="role">Role</label>
                <input v-model="formData.role" type="text" class="form-control" id="role" required />
              </div>
              <button type="submit" class="btn btn-primary">Simpan Perubahan</button>
            </form>
          </div>
        </div>
      </div>
    </div>
    <div v-if="isModalOpen" class="modal-backdrop fade show"></div>
  </div>
</template>

<script>
import { getUser, updateUser } from '@/api/user'; // Sesuaikan dengan struktur proyek

export default {
  data() {
    return {
      user: null, // Data pengguna yang akan ditampilkan
      formData: { // Data untuk form edit profil
        nama: '',
        email: '',
        role: '',
      },
      isModalOpen: false, // Status modal
    };
  },
  async created() {
    const userId = this.$route.params.id_user; // Ambil id_user dari parameter route
    if (!userId) {
      console.error('User ID is missing');
      return;
    }
    await this.fetchUserProfile(userId); // Ambil data pengguna berdasarkan id_user
  },
  methods: {
    // Fungsi untuk mengambil data pengguna berdasarkan ID
    async fetchUserProfile(id_user) {
      try {
        this.user = await getUser(id_user); // Ambil data pengguna dari API
        this.formData = { ...this.user }; // Sinkronkan formData dengan data pengguna
      } catch (error) {
        console.error('Error fetching user profile:', error);
      }
    },
    // Fungsi untuk membuka modal edit profil
    openEditModal() {
      this.isModalOpen = true;
    },
    // Fungsi untuk menutup modal edit profil
    closeEditModal() {
      this.isModalOpen = false;
    },
    // Fungsi untuk memperbarui profil pengguna
    async updateUserProfile() {
      try {
        const id_user = this.user.id_user;
        await updateUser(id_user, this.formData);  // Update ke server

        // Update data di Vuex dan localStorage
        this.$store.commit('setUser', { ...this.formData });
        localStorage.setItem('user', JSON.stringify(this.formData));  // Simpan ke localStorage

        this.user = { ...this.formData };  // Update data pengguna setelah disimpan
        this.closeEditModal();  // Tutup modal setelah berhasil
      } catch (error) {
        console.error('Error updating profile:', error);
      }
    }
  },
};
</script>


<style scoped>
.modal-backdrop {
  background-color: rgba(0, 0, 0, 0.5);
}
</style>
