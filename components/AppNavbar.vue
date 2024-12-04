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
            <b-nav-item >
              <button @click.prevent="navigateWithLoading('/listBuku')" :class="{ 'active-link': isActive('/listBuku') }"
                class="nav-link btn btn-white text-dark">
                List Buku
              </button>
            </b-nav-item>
            <b-nav-item v-if="user.role === 'admin'" class="mr-5">
              <button @click="toggleSidebar" class="nav-link btn btn-white text-dark">
                Menu
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

    <Sidebar :isOpen="isSidebarOpen" :items="sidebarItems" @close="isSidebarOpen = false" @navigate="navigate" />
  </div>
</template>

<script>
import LoadingSpinner from './LoadingSpinner.vue';
import Sidebar from './Sidebar.vue';

export default {
  name: 'AppNavbar',
  components: {
    LoadingSpinner,
    Sidebar,
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
        { judul: "Pengembalian", link: "admin/pengembalian" },
        { judul: "Users", link: "admin/Users" },
      ],
    };
  },
  computed: {
    user() {
      return this.$store.getters.getUser;
    },
    isLoading() {
      return this.$store.getters.isLoading;
    },
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
      if (confirm('Apakah Anda Yakin Akan Logout')) {
        this.$store.dispatch('setLoading', true);
        setTimeout(() => {
          this.$store.dispatch('logout');
          this.$router.push('/');
          this.$store.dispatch('setLoading', false);
        }, 2000);
      }
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
</style>
