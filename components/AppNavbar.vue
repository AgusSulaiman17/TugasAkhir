<template>
  <div :class="{ 'hidden': isHidden }" class="navbar-wrapper">
    <LoadingSpinner v-if="isLoading" />
    <b-navbar toggleable="lg" class="content p-2 p-md-0 rounded-5">
      <div class="container-fluid">
        <b-navbar-brand href="#" class="judul bg-white p-2 rounded-3 font-weight-bold">
          <img src="../static/images/logoFinal.svg" alt="" style="width: 60px">
        </b-navbar-brand>
        <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

        <b-collapse v-if="!user" id="nav-collapse" is-nav>
          <b-navbar-nav class="ml-auto">
            <b-nav-item>
              <b-button @click.prevent="navigateWithLoading('/login')" class="px-3 bg-dark">Login
                <b-icon-door-open-fill></b-icon-door-open-fill></b-button>
            </b-nav-item>
          </b-navbar-nav>
        </b-collapse>

        <b-collapse v-if="user" id="nav-collapse" is-nav>
          <b-navbar-nav class="ml-auto">
            <b-nav-item>
              <a @click="home" :class="{ 'active-link': isActive('/admin/dashboard') || isActive('/user/dashboard') }"
                class="nav-link btn btn-white text-dark">
                Home
              </a>
            </b-nav-item>
            <b-nav-item>
              <button @click.prevent="navigateWithLoading('/BukuPinjaman')"
                :class="{ 'active-link': isActive('/BukuPinjaman') }" class="nav-link btn btn-white text-dark">
                Buku Pinjaman
              </button>
            </b-nav-item>
            <b-nav-item class="mr-5">
              <button @click.prevent="navigateWithLoading('/listBuku')" :class="{ 'active-link': isActive('/listBuku') }"
                class="nav-link btn btn-white text-dark">
                List Buku
              </button>
            </b-nav-item>
          </b-navbar-nav>

          <b-navbar-nav>
            <b-nav-item-dropdown right>
              <template #button-content>
                <b-icon-person-circle class="mr-2"></b-icon-person-circle>
                <em>{{ user.nama }}</em>
              </template>
              <b-dropdown-item @click="$router.push(`/Profile/${user.id_user}`)">
                Profile
              </b-dropdown-item>
              <b-dropdown-item @click.prevent="logout">Sign Out</b-dropdown-item>
            </b-nav-item-dropdown>
          </b-navbar-nav>
        </b-collapse>
      </div>
    </b-navbar>
  </div>
</template>

<script>
import LoadingSpinner from './LoadingSpinner.vue'; // Impor komponen LoadingSpinner

export default {
  name: 'AppNavbar',
  components: {
    LoadingSpinner, // Daftarkan komponen LoadingSpinner
  },
  computed: {
    user() {
      return this.$store.getters.getUser;
    },
    isLoading() {
      return this.$store.getters.isLoading; // Ambil status loading dari store
    }
  },
  methods: {
    logout() {
      if (confirm("Apakah Anda Yakin Akan Logout")) {
        this.$store.dispatch('setLoading', true); // Set loading saat logout
        setTimeout(() => {
          this.$store.dispatch('logout');
          this.$router.push('/');
          this.$store.dispatch('setLoading', false); // Matikan loading setelah logout
        },1000 );
      }
    },
    home() {
      this.$store.dispatch('setLoading', true); // Setel loading ke true

      setTimeout(() => {
        if (this.user && this.user.role === 'admin') {
          this.$router.push('/admin/dashboard');
        } else {
          this.$router.push('/user/dashboard');
        }

        this.$store.dispatch('setLoading', false); // Matikan loading setelah navigasi
      }, 1000);
    },
    navigateWithLoading(path) {
      this.$store.dispatch('navigateWithLoading', { path, router: this.$router });
    },
    isActive(path) {
      return this.$route.path === path; 
    }
  }
};
</script>

<style scoped>
.navbar-wrapper {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  z-index: 9999;
  transition: transform 0.3s ease-in-out, opacity 0.3s ease-in-out;
  padding: 20px;
}

.content {
  background-color: rgba(255, 255, 255, 0.8);
  border-radius: 20px;
  font-size: 20px;
  transition: box-shadow 0.3s ease-in-out;
  /* Tambahkan transisi untuk bayangan */
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
  /* Menggeser navbar ke atas saat di-scroll */
  opacity: 0;
  /* Menyembunyikan navbar */
}

.active-link {
  font-weight: bold;
  color: #000;
  /* Sesuaikan dengan warna yang diinginkan */
}
</style>
