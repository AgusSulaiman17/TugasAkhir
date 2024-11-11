<template>
  <div>
    <AppNavbar />
    <HeroLanding class="mt-8" />
  </div>
</template>

<script>
import AppNavbar from '~/components/AppNavbar.vue';
import HeroLanding from '~/components/HeroLanding.vue';

export default{
  name:'Landing',
  components: {
    AppNavbar,
    HeroLanding
  }
}
</script>
<style>
.mt-8{
  margin-top: 100px;
}
</style>
