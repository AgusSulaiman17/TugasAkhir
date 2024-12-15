<template>
  <div :class="{ 'hidden': isHidden }" class="navbar-wrapper">
    <LoadingSpinner v-if="isLoading" /> <!-- Tampilkan spinner saat loading -->
    <b-navbar toggleable="lg" class="content p-2 p-md-0 rounded-5">
      <div class="container-fluid">
        <b-navbar-brand href="#" class="judul bg-white p-2 rounded-3 font-weight-bold">
          <img src="../static/images/logoFinal.svg" alt="" style="width: 60px">
        </b-navbar-brand>
        <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

        <b-collapse v-if="!user" id="nav-collapse" is-nav>
          <b-navbar-nav class="ml-auto">
            <b-nav-item>
              <b-button @click.prevent="navigateWithLoading('/login')" class="px-3 f-20 bg-dark">Login
                <b-icon-door-open-fill></b-icon-door-open-fill></b-button>
            </b-nav-item>
          </b-navbar-nav>
        </b-collapse>

        <b-collapse v-if="user" id="nav-collapse" is-nav>
          <b-navbar-nav class="ml-auto">
            <b-nav-item v-if="user.role === 'admin'">
              <button @click="toggleSidebar" class="nav-link btn btn-white text-dark">
                SideMenu
              </button>
            </b-nav-item>
            <b-nav-item>
              <a @click="home" :class="{ 'active-link': isActive('/admin/dashboard') || isActive('/user/dashboard') }"
                class="nav-link btn btn-white text-dark">
                Home
              </a>
            </b-nav-item>
            <b-nav-item v-if="user.role === 'user'">
              <button @click.prevent="navigateWithLoading('/BukuPinjaman')"
                :class="{ 'active-link': isActive('/BukuPinjaman') }" class="nav-link btn btn-white text-dark">
                Buku Pinjaman
              </button>
            </b-nav-item>
            <b-nav-item v-if="user.role === 'user' || user.role === 'petugas'">
              <button @click.prevent="navigateWithLoading('/listBuku')" :class="{ 'active-link': isActive('/listBuku') }"
                class="nav-link btn btn-white text-dark">
                List Buku
              </button>
            </b-nav-item>
            <b-nav-item v-if="user.role ==='petugas'">
              <button @click.prevent="navigateWithLoading('/admin/datapeminjaman')" :class="{ 'active-link': isActive('/admin/datapeminjaman') }"
                class="nav-link btn btn-white text-dark">
                Data Peminjaman
              </button>
            </b-nav-item>
            <b-nav-item v-if="user.role === 'petugas'">
              <button @click.prevent="navigateWithLoading('/admin/pengembalian')" :class="{ 'active-link': isActive('/admin/pengembalian') }"
                class="nav-link btn btn-white text-dark">
                Data Pengembalian
              </button>
            </b-nav-item>
          </b-navbar-nav>

          <b-navbar-nav class="ml-5">
            <b-nav-item-dropdown right>
              <template #button-content>
                <img
                  :src="'https://ui-avatars.com/api/?name=' + user.nama + '&background=random'"
                  alt="Avatar"
                  class="rounded-circle avatar font-irish"
                />
                <em class="text-ijotua">{{ user.nama }}</em>
              </template>
              <b-dropdown-item @click="$router.push(`/Profile/${user.id_user}`)">
                Profile <b-icon-person></b-icon-person>
              </b-dropdown-item>
              <b-dropdown-item @click.prevent="logout" :disabled="isLoading"> <!-- Tambahkan disabled -->
                <span v-if="!isLoading">Sign Out</span>
                <span v-else>
                  <div class="spinner"></div> <!-- Spinner saat loading -->
                </span>
                <b-icon-box-arrow-right></b-icon-box-arrow-right>
              </b-dropdown-item>
            </b-nav-item-dropdown>
          </b-navbar-nav>
        </b-collapse>
      </div>
    </b-navbar>

    <Sidebar :isOpen="isSidebarOpen" :items="sidebarItems" @close="isSidebarOpen = false" @navigate="navigate" />

    <!-- Notification Modal for Logout Confirmation -->
    <NotificationModal
      v-if="showLogoutConfirmationModal"
      :isVisible="showLogoutConfirmationModal"
      :messageTitle="'Konfirmasi Logout'"
      :messageBody="'Apakah Anda yakin ingin logout?'"
      @close="cancelLogout"
      class="notif"
    >
      <template #footer>
        <button @click="confirmLogout" class="btn btn-abu" :disabled="isLoading">
          <span v-if="isLoading" class="spinner-border spinner-border-sm" role="status" aria-hidden="true"></span>
          <span v-else> Ya, Logout </span>
        </button>
        <button @click="cancelLogout" class="btn btn-cancel">Batal</button>
      </template>
    </NotificationModal>
  </div>
</template>

<script>
import LoadingSpinner from './LoadingSpinner.vue';
import Sidebar from './Sidebar.vue';
import NotificationModal from './NotificationModal.vue';

export default {
  name: 'AppNavbar',
  components: {
    LoadingSpinner,
    Sidebar,
    NotificationModal,
  },
  data() {
    return {
      isHidden: false,
      lastScrollTop: 0,
      isSidebarOpen: false,
      sidebarItems: [
        { judul: "Genre", link: "admin/genres" },
        { judul: "Penulis", link: "admin/penulis" },
        { judul: "Buku", link: "admin/buku" },
        { judul: "Data Peminjaman", link: "admin/dataPeminjaman" },
        { judul: "Users", link: "admin/Users" },
      ],
      showLogoutConfirmationModal: false,
      isLoading: false, // Tambahkan state loading
    };
  },
  computed: {
    user() {
      return this.$store.getters.getUser;
    },
    isLoading() {
      return this.$store.getters.isLoading;
    }
  },
  methods: {
    handleScroll() {
      const currentScroll = window.pageYOffset || document.documentElement.scrollTop;
      if (currentScroll > this.lastScrollTop) {
        this.isHidden = true;
      } else {
        this.isHidden = false;
      }
      this.lastScrollTop = currentScroll <= 0 ? 0 : currentScroll;
    },
    logout() {
      this.showLogoutConfirmationModal = true; // Tampilkan modal konfirmasi
    },
    confirmLogout() {
      this.isLoading = true; // Tampilkan loading saat logout
      setTimeout(() => {
        this.$store.dispatch('logout'); // Dispatch logout action
        this.$router.push('/'); // Redirect to home
        this.showLogoutConfirmationModal = false; // Tutup modal
        this.isLoading = false; // Sembunyikan spinner
      }, 2000);
    },
    cancelLogout() {
      this.showLogoutConfirmationModal = false; // Tutup modal
    },
    home() {
      this.$store.dispatch('setLoading', true);

      setTimeout(() => {
        if (this.user && this.user.role === 'admin') {
          this.$router.push('/admin/dashboard');
        } else {
          this.$router.push('/user/dashboard');
        }
        this.$store.dispatch('setLoading', false);
      });
    },
    toggleSidebar() {
      this.isSidebarOpen = !this.isSidebarOpen;
    },
    navigate(link) {
      this.$router.push(`/${link}`);
    },
    navigateWithLoading(path) {
      this.$store.dispatch('navigateWithLoading', { path, router: this.$router });
    },
    isActive(path) {
      return this.$route.path === path;
    },
  },
  mounted() {
    window.addEventListener('scroll', this.handleScroll);
  },
  beforeDestroy() {
    window.removeEventListener('scroll', this.handleScroll);
  },
};
</script>

<style scoped>
.notif {
  z-index: 9999999;
  height: 100vh;
}

.f-20 {
  font-size: 20px;
}
.navbar-wrapper {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  z-index: 9999;
  transition: transform 0.3s ease-in-out, opacity 0.3s ease-in-out;
  padding: 20px;
  transform: translateY(0);
  opacity: 1;
}

.content {
  background-color: rgba(255, 255, 255, 0.8);
  border-radius: 20px;
  font-size: 20px;
  transition: box-shadow 0.3s ease-in-out;
  transition: background-color 0.3s ease-in-out;
}

.content:hover {
  background-color: white;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

.rounded-5 {
  border-radius: 20px !important;
}

.btn-white {
  background-color: rgba(255, 255, 255, 0.8);
  font-size: 20px;
}

.btn-white:hover {
  background-color: rgba(240, 240, 240, 0.8);
}

.hidden {
  transform: translateY(-100%);
  opacity: 0;
  transition: transform 0.3s ease-in-out, opacity 0.3s ease-in-out;
}

.active-link {
  font-weight: bold;
  color: #000;
  pointer-events: none;
  border-bottom: 2px solid #000;
  padding-bottom: 4px;
}

.avatar {
  width: 32px;
  height: 32px;
  object-fit: cover;
  border: 3px solid  #70a799;
  margin-right: 5px;
}

.spinner {
  border: 3px solid #f3f3f3;
  border-top: 3px solid #000;
  border-radius: 50%;
  width: 15px;
  height: 15px;
  animation: spin 1s linear infinite;
  display: inline-block;
}
</style>
