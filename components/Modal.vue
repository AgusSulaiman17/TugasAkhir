<template>
  <div v-if="visible" class="modal fade show" tabindex="-1" style="display: block;" aria-labelledby="exampleModalLabel" aria-hidden="true">
    <div class="modal-dialog">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title">{{ title }}</h5>
          <button type="button" class="btn-close" @click="closeModal" aria-label="Close"></button>
        </div>
        <div class="modal-body">
          <p>{{ message }}</p>
        </div>
        <div class="modal-footer">
          <button type="button" class="btn btn-secondary" @click="closeModal">Tutup</button>
          <button type="button" class="btn btn-primary" @click="confirmAction">Konfirmasi</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    visible: {
      type: Boolean,
      required: true
    },
    title: {
      type: String,
      required: true
    },
    message: {
      type: String,
      required: true
    }
  },
  methods: {
    closeModal() {
      this.$emit('update:visible', false); // Emit untuk menutup modal
    },
    confirmAction() {
      this.$emit('confirm'); // Emit konfirmasi aksi
      this.closeModal();
    }
  }
};
</script>
