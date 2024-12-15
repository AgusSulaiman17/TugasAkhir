<template>
  <div>
    <AppNavbar />
    <div class="profile-page container mt-8">
      <div class="profile-card card shadow-lg">
        <div class="card-body text-center">
          <h2 class="card-title">Profil Pengguna</h2>
          <div v-if="user" class="user-details">
            <img
              :src="'https://ui-avatars.com/api/?name=' + user.nama + '&background=random'"
              alt="Avatar"
              class="rounded-circle avatar font-irish"
            />
            <h3 class="mt-3">{{ user.nama }}</h3>
            <p><i class="fas fa-envelope"></i> {{ user.email }}</p>
            <button class="btn bg-ijomuda mt-3" @click="openEditModal">
              <i class="fas fa-edit"></i> Edit Profil
            </button>
          </div>
          <div v-else>
            <p>Memuat data...</p>
          </div>
        </div>
      </div>

      <!-- Modal Edit Profile -->
      <div v-if="isModalOpen" class="modal mt-8 fade show" tabindex="-1" role="dialog" style="display: block;">
        <div class="modal-dialog" role="document">
          <div class="modal-content">
            <div class="modal-header">
              <h5 class="modal-title">Edit Profil</h5>
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
                <button type="submit" class="btn bg-ijomuda w-100">Simpan Perubahan</button>
              </form>
            </div>
          </div>
        </div>
      </div>
      <div v-if="isModalOpen" class="modal-backdrop fade show"></div>
    </div>

    <!-- Success Notification Modal -->
    <NotificationModal
      v-if="showSuccessModal"
      :isVisible="showSuccessModal"
      :messageTitle="successTitle"
      :messageBody="successMessage"
      @close="closeSuccessModal"
    />
  </div>
</template>


<script>
import { getUser, updateUser } from '@/api/user';
import AppNavbar from '~/components/AppNavbar.vue';
import NotificationModal from '~/components/NotificationModal.vue';

export default {
  components: {
    AppNavbar,
    NotificationModal,
  },
  data() {
    return {
      user: null,
      formData: {
        nama: '',
        email: '',
      },
      isModalOpen: false,
      showSuccessModal: false,
      successTitle: '',
      successMessage: '',
    };
  },
  async created() {
    const userId = this.$route.params.id_user;
    if (!userId) {
      console.error('User ID is missing');
      return;
    }
    await this.fetchUserProfile(userId);
  },
  methods: {
    async fetchUserProfile(id_user) {
      try {
        this.user = await getUser(id_user);
        this.formData = { ...this.user };
      } catch (error) {
        console.error('Error fetching user profile:', error);
      }
    },
    openEditModal() {
      this.isModalOpen = true;
    },
    closeEditModal() {
      this.isModalOpen = false;
    },
    async updateUserProfile() {
      try {
        const id_user = this.user.id_user;
        await updateUser(id_user, this.formData);
        this.$store.commit('setUser', { ...this.formData });
        localStorage.setItem('user', JSON.stringify(this.formData));
        this.user = { ...this.formData };
        this.closeEditModal();

        // Tampilkan notifikasi sukses
        this.successTitle = 'Berhasil!';
        this.successMessage = 'Profil Anda telah diperbarui.';
        this.showSuccessModal = true;
      } catch (error) {
        console.error('Error updating profile:', error);
      }
    },
    closeSuccessModal() {
      this.showSuccessModal = false;
    },
  },
};

</script>

<style scoped>
.profile-page {
  max-width: 600px;
}

.profile-card {
  margin-top: 50px;
  padding: 20px;
  border-radius: 15px;
  background-color: #fff;
}

.avatar {
  width: 120px;
  height: 120px;
  object-fit: cover;
  border: 3px solid  #70a799;
}

.modal{
  z-index: 999999;
}
.modal-backdrop {
  background-color: rgba(0, 0, 0, 0.5);
  z-index: 99999;
}
</style>
