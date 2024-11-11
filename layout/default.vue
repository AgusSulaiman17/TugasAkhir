<template>
  <div>
    <!-- Menampilkan loading spinner jika isLoading true -->
    <LoadingSpinner :isLoading="$store.state.isLoading" />
    <nuxt /> <!-- Halaman aktif akan ditampilkan di sini -->
  </div>
</template>

<script>
import LoadingSpinner from '@/components/LoadingSpinner.vue';

export default {
  components: {
    LoadingSpinner
  }
};
</script>
