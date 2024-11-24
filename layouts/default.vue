<template>
  <div>
    <!-- Tampilkan modal jika showModal true -->
    <AdminAlert v-if="showModal" />
    <!-- Konten halaman lainnya -->
    <Nuxt />
  </div>
</template>

<script>
import AdminAlert from '~/components/AdminAlert.vue';

export default {
  components: {
    AdminAlert
  },
  computed: {
    showModal() {
      // Memeriksa status modal dari Vuex store
      console.log('Modal status:', this.$store.state.showModal); // Log untuk memverifikasi status
      return this.$store.state.showModal;
    }
  }
}
</script>
