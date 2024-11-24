<template>
  <div>
    <LoadingSpinner v-if="isLoading" />
    <AppNavbar />
    <!-- Container dengan Layout Flexbox -->
    <div class="container-fluid mt-8">
      <div class="row d-flex flex-wrap ">
        <div class="card col-lg-8 mr-4 mb-4">
          <div class="card-body">
            <h5 class="card-title">Selamat datang Di Lentera</h5>
            <h6 class="card-subtitle mb-2 text-muted">Card subtitle</h6>
            <p class="card-text">Some quick example text to build on the card title and make up the bulk of the card's content.</p>
            <a href="#" class="card-link">Card link</a>
            <a href="#" class="card-link">Another link</a>
          </div>
        </div>

        <div v-for="(item, index) in content" :key="index" class="card col-lg-3 mb-4 mx-2">
          <div class="card-body">
            <h5 class="card-title">{{ item.judul }}</h5>
            <button @click="next(item.link)" class="btn btn-abu">Lihat {{ item.judul }}</button>
            <p class="card-text">Some quick example text to build on the card title and make up the bulk of the card's content.</p>
            <a href="#" class="card-link">Card link</a>
            <a href="#" class="card-link">Another link</a>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import AppNavbar from '~/components/AppNavbar.vue';
import LoadingSpinner from '~/components/LoadingSpinner.vue';

export default {
  middleware: 'admin',
  name: 'AdminDashboard',
  components: {
    AppNavbar,
    LoadingSpinner
  },
  data() {
    return {
      content: [
        { judul: 'Genre', link: 'genres' },
        { judul: 'Penulis', link: 'penulis' },
        { judul: 'Buku', link: 'buku' },
        { judul: 'Peminjaman', link: 'peminjaman'}
      ],
    };
  },
  computed: {
    isLoading() {
      return this.$store.state.isLoading;
    }
  },
  methods: {
    next(link) {
      // Mengarahkan ke halaman sesuai dengan link yang dipilih dan menampilkan loading state
      if (link) {
        this.$store.dispatch('navigateWithLoading', { path: `/admin/${link}`, router: this.$router });
      }
    }
  }
}
</script>

<style scoped>
.card {
  border-radius: 8px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.card-body {
  padding: 1.5rem;
}

.card-title {
  font-size: 1.5rem;
  font-weight: bold;
  margin-bottom: 10px;
}

.card-subtitle {
  font-size: 1rem;
}

.card-text {
  font-size: 1rem;
  color: #555;
}

.card-link {
  color: #007bff;
  text-decoration: none;
}

.card-link:hover {
  text-decoration: underline;
}

.btn-abu {
  background-color: #ccc;
  color: #333;
  border: none;
  padding: 0.5rem 1rem;
  border-radius: 5px;
  cursor: pointer;
}

.btn-abu:hover {
  background-color: #bbb;
}

/* Responsif */
@media (max-width: 768px) {
  .card {
    margin-bottom: 15px;
    width: 100%;
  }
}
</style>
