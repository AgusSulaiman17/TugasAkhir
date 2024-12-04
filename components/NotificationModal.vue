<template>
  <div v-if="isVisible" class="modal-overlay">
    <div class="modal-content">
      <h2>{{ messageTitle }}</h2>
      <p>{{ messageBody }}</p>
      <button @click="closeModal">Close</button>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    isVisible: {
      type: Boolean,
      required: true
    },
    messageTitle: {
      type: String,
      required: true
    },
    messageBody: {
      type: String,
      required: true
    }
  },
  methods: {
    closeModal() {
      console.log("Closing modal");
      this.$emit('close'); // Emit event to parent to close modal
    }
  },
  watch: {
    isVisible(newValue) {
      console.log("Modal visibility changed:", newValue); // Log perubahan nilai isVisible
    }
  }
};
</script>

<style scoped>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.modal-content {
  background-color: white;
  padding: 20px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.5);
  border-radius: 5px;
  text-align: center;
}

button {
  padding: 10px 20px;
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}
</style>
