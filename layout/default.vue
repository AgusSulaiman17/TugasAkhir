<template>
  <div>
    <LoadingSpinner />
    <Nuxt />
  </div>
</template>

<script>
export default {
  mounted() {
    this.$router.beforeEach((to, from, next) => {
      this.$store.dispatch('setLoading', true);
      next();
    });

    this.$router.afterEach(() => {
      setTimeout(() => {
        this.$store.dispatch('setLoading', false);
      }, 500); // Delay opsional agar terlihat smooth
    });
  }
};
</script>
