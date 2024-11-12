<template>
  <div class="loader-overlay" v-if="$store.getters.isLoading">
    <div class="loader">
      <span class="bar"></span>
      <span class="bar"></span>
      <span class="bar"></span>
    </div>
  </div>
</template>

<script>
export default {
  name: 'LoadingSpinner'
};
</script>

<style scoped>
.loader-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  background-color: #f5f5f5; 
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
  opacity: 1;
  transition: opacity 0.3s ease;
  visibility: visible; /* Untuk menjaga agar spinner selalu terlihat */
}


.loader {
  display: flex;
  align-items: center;
}

.bar {
  display: inline-block;
  width: 3px;
  height: 20px;
  background-color: #000000;
  border-radius: 10px;
  animation: scale-up 1s linear infinite;
}

.bar:nth-child(2) {
  height: 35px;
  margin: 0 5px;
  animation-delay: 0.25s;
}

.bar:nth-child(3) {
  animation-delay: 0.5s;
}

@keyframes scale-up {
  20% {
    background-color: #fff;
    transform: scaleY(1.5);
  }
  40% {
    transform: scaleY(1);
  }
}
</style>
