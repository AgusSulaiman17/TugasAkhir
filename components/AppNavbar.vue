<template>
  <div :class="{'hidden': isHidden}" class="navbar-wrapper">
    <b-navbar toggleable="lg" class="content p-2 p-md-0 rounded-5">
      <div class="container">
        <b-navbar-brand href="#" class="judul btn btn-white font-weight-bold">Lentera</b-navbar-brand>

        <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>
        <b-collapse v-if="!user" id="nav-collapse" is-nav>
          <b-navbar-nav class="ml-auto">
            <b-nav-item>
              <router-link to="/login">
                <b-button class="px-3 bg-dark">Login <b-icon-door-open-fill></b-icon-door-open-fill></b-button>
              </router-link>
            </b-nav-item>
          </b-navbar-nav>
        </b-collapse>

        <b-collapse v-if="user" id="nav-collapse" is-nav>
          <b-navbar-nav class="ml-auto">
            <b-nav-item>
              <a @click="home" class="nav-link btn btn-white text-dark" to="/">Home</a>
            </b-nav-item>
            <b-nav-item>
              <router-link class="nav-link btn btn-white text-dark" to="/hasilCheckout">Pesanan
                <b-icon-cart></b-icon-cart></router-link>
            </b-nav-item>
            <b-nav-item>
              <router-link class="nav-link btn btn-white text-dark" to="/film">Detail Film
                <b-icon-film></b-icon-film></router-link>
            </b-nav-item>
          </b-navbar-nav>
          <b-navbar-nav>
            <b-nav-item>
              <nuxt-link :to="{ name: 'profile', query: { id: user.id_user } }">
                <b-button class="px-3 text-white bg-dark">
                  {{ user.nama }} <b-icon-person></b-icon-person>
                </b-button>
              </nuxt-link>
              <b-button variant="outline-danger" class="ml-2" @click="logout">Logout
                <b-icon-door-closed></b-icon-door-closed></b-button>
            </b-nav-item>
          </b-navbar-nav>
        </b-collapse>
      </div>
    </b-navbar>
  </div>
</template>

<script>
export default {
  name: 'AppNavbar',

  data() {
    return {
      isHidden: false,  // Menyembunyikan navbar
      lastScrollY: 0,   // Menyimpan posisi scroll sebelumnya
      scrollTimeout: null, // Menyimpan timeout untuk scroll berhenti
    };
  },
  computed: {
    user() {
      return this.$store.getters.getUser;
    },
  },
  methods: {
    logout() {
      if (confirm("Apakah Anda Yakin Akan Logout")) {
        this.$store.dispatch('logout');
        this.$router.push('/');
      }
    },
    home() {
      if (this.user && this.user.role === 'admin') {
        this.$router.push('/admin/dashboard');
      } else {
        this.$router.push('/user/dashboard');
      }
    },
    handleScroll() {
      // Menyembunyikan navbar saat scroll ke bawah
      if (window.scrollY > this.lastScrollY) {
        this.isHidden = true; // Menyembunyikan navbar
      } else {
        this.isHidden = false; // Menampilkan navbar
      }
      this.lastScrollY = window.scrollY;

      // Membatalkan timeout sebelumnya jika ada scroll baru
      if (this.scrollTimeout) {
        clearTimeout(this.scrollTimeout);
      }

      // Set timeout untuk menampilkan navbar setelah scroll berhenti
      this.scrollTimeout = setTimeout(() => {
        if (window.scrollY === this.lastScrollY) {
          this.isHidden = false;  // Munculkan navbar setelah scroll berhenti
        }
      }, 150);  // Tunggu 150ms setelah scroll berhenti
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
}

.content {
  background-color: rgba(255, 255, 255, 0.8);
  border-radius: 20px;
}

.rounded-5 {
  border-radius: 20px !important;
}

.judul {
  font-family: 'Life Savers', cursive;
}

.btn-white {
  background-color: rgba(255, 255, 255, 0.8);
}

.btn-white:hover {
  background-color: rgba(240, 240, 240, 0.8);
}

.hidden {
  transform: translateY(-100%); /* Menggeser navbar ke atas saat di-scroll */
  opacity: 0; /* Menyembunyikan navbar */
}
</style>
