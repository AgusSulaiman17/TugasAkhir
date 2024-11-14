<template>
  <div>
    <LoadingSpinner />
    <AppNavbar />
    <div class="container mt-8">
      <div class="row">
        <div class="card col-lg-8 mr-4 mb-4">
          <div class="card-body">
            <h5 class="card-title">Selamat datang Di Lentera</h5>
            <h6 class="card-subtitle mb-2 text-muted">Card subtitle</h6>
            <p class="card-text">Some quick example text to build on the card title and make up the bulk of the card's
              content.</p>
            <a href="#" class="card-link">Card link</a>
            <a href="#" class="card-link">Another link</a>
          </div>
        </div>
        <div class="card col-lg-3 mb-4 mx-2 mx-md-0">
          <div class="card-body">
            <h5 class="card-title">Genre</h5>
            <button @click="next" class="btn btn-abu">Lihat Genre</button>
            <p class="card-text">Some quick example text to build on the card title and make up the bulk of the card's
              content.</p>
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
  middleware: 'auth',
  name: 'AdminDashboard',
  components: {
    AppNavbar,
    LoadingSpinner
  },
  computed: {
    isLoading() {
      return this.$store.state.isLoading;
    }
  },
  methods: {
    next() {
      this.$store.dispatch('navigateWithLoading', { path: '/admin/genres', router: this.$router });
    }
  }
}
</script>
