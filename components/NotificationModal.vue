<template>
  <div v-if="isVisible" class="modal-overlay">
    <div class="modal-content">
      <h2 class="modal-title">{{ messageTitle }}</h2>
      <p class="modal-body">{{ messageBody }}</p>
      <div class="modal-actions">
        <!-- Slot untuk tombol tambahan -->
        <slot name="footer">
          <button @click="closeModal" class="btn bg-ijomuda">Close</button>
        </slot>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "NotificationModal",
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
      this.$emit("close"); // Emit event to parent to close modal
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
/* Overlay dengan efek blur */
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  backdrop-filter: blur(5px); /* Efek blur */
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 999999;
}

/* Konten modal */
.modal-content {
  background-color: #ffffff;
  padding: 30px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.2);
  border-radius: 10px;
  text-align: center;
  width: 90%;
  max-width: 400px; /* Ukuran maksimum modal */
  animation: fadeIn 0.3s ease-out;
}

/* Animasi untuk modal */
@keyframes fadeIn {
  from {
    transform: scale(0.8);
    opacity: 0;
  }
  to {
    transform: scale(1);
    opacity: 1;
  }
}

/* Judul modal */
.modal-title {
  font-size: 1.5rem;
  font-weight: bold;
  margin-bottom: 15px;
  color: #333333;
}

/* Isi modal */
.modal-body {
  font-size: 1rem;
  margin-bottom: 20px;
  color: #555555;
}

/* Tombol */
.modal-actions button {
  padding: 10px 20px;
  margin: 0 5px;
  border: none;
  border-radius: 5px;
  font-size: 1rem;
  cursor: pointer;
  transition: all 0.3s ease;
}

.modal-actions .btn-close {
  background-color:#70a799;
  color: white;
}

.modal-actions .btn-close:hover {
  background-color: #3a4b47;
}

.modal-actions .btn-cancel {
  background-color: #6c757d;
  color: white;
}

.modal-actions .btn-cancel:hover {
  background-color: #5a6268;
}

.modal-actions .btn-confirm {
  background-color: #007bff;
  color: white;
}

.modal-actions .btn-confirm:hover {
  background-color: #0056b3;
}
</style>
