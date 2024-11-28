<template>
  <div>
    <AppNavbar />
    <HeroLanding class="mt-8 mb-8" />
    <AppFooter />
  </div>
</template>

<script>
import AppFooter from '~/components/AppFooter.vue';
import AppNavbar from '~/components/AppNavbar.vue';
import HeroLanding from '~/components/HeroLanding.vue';

export default{
  name:'Landing',
  components: {
    AppNavbar,
    HeroLanding,
    AppFooter
  }
}
</script>
<style scoped>
.mb-8{
  margin-bottom: 150px;
}
</style>
