<template>
  <div v-if="isVisible" class="modal-overlay">
    <div class="modal-content">
      <h2>Anda bukan admin!</h2>
      <button @click="closeModal">Tutup</button>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      isVisible: false,
    };
  },
  methods: {
    showModal() {
      this.isVisible = true;
    },
    closeModal() {
      this.isVisible = false;
      this.$router.push('/'); // Arahkan pengguna ke halaman home
    }
  },
  watch: {
    '$store.state.showModal'(newVal) {
      this.isVisible = newVal;
    }
  }
}
</script>

<style scoped>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
}
.modal-content {
  background-color: white;
  padding: 20px;
  border-radius: 5px;
  text-align: center;
}
</style>
