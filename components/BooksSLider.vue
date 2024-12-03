<template>
  <div class="book-slider">
    <swiper
      :space-between="10"
      :slides-per-view="4"
      :loop="true"
      :pagination="{ clickable: true }"
      :navigation="true"
    >
      <swiper-slide v-for="buku in bukuList" :key="buku.id_buku">
        <div class="card">
          <img
            v-if="buku.gambar"
            :src="buku.gambar.startsWith('http') ? buku.gambar : `http://localhost:8080${buku.gambar}`"
            :alt="buku.judul"
            class="card-img-top"
            style="height: 200px;"
          />
          <div v-else class="text-center p-3">Tidak ada gambar</div>
          <div class="card-body">
            <h5 class="card-title text-center">{{ buku.judul }}</h5>
            <p class="card-text m-0">{{ buku.penulis && buku.penulis.nama ? buku.penulis.nama : 'Tidak diketahui' }}</p>
            <p class="card-text m-0">{{ buku.genre && buku.genre.nama ? buku.genre.nama : 'Tidak diketahui' }}</p>
            <button @click="$router.push(`/DetailBuku/${buku.id_buku}`)" class="btn bg-ijomuda mt-2 w-100">
              Detail
            </button>
          </div>
        </div>
      </swiper-slide>
    </swiper>
  </div>
</template>

<script>
import { Swiper, SwiperSlide } from 'swiper/vue';
import 'swiper/swiper-bundle.css';
import { getBukuList } from '@/api/buku'; // pastikan path sesuai dengan lokasi file

export default {
  components: {
    Swiper,
    SwiperSlide,
  },
  data() {
    return {
      bukuList: [], // Menyimpan data buku
    };
  },
  async mounted() {
    try {
      // Memanggil API untuk mendapatkan daftar buku
      this.bukuList = await getBukuList();
    } catch (error) {
      console.error('Error fetching buku list:', error);
    }
  },
};
</script>

<style scoped>
.book-slider {
  margin-top: 20px;
}

.card {
  border-radius: 10px;
  overflow: hidden;
  border: none;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  background-color: #fff;
  transition: transform 0.3s ease, box-shadow 0.3s ease;
}

.card:hover {
  transform: translateY(-5px);
  box-shadow: 0 8px 12px rgba(0, 0, 0, 0.2);
}

.card-img-top {
  border-top-left-radius: 10px;
  border-top-right-radius: 10px;
  height: 200px;
  background-color: #f7f7f7;
}

.card-body {
  padding: 20px;
}

.card-title {
  font-size: 25px;
  font-weight: bold;
  color: #333;
  margin-bottom: 0.5rem;
}

.card-text {
  font-size: 0.9rem;
  color: #555;
  margin-bottom: 0.5rem;
}

.btn.bg-ijomuda {
  background-color: #70a799;
  color: #fff;
  padding: 0.75rem 1.5rem;
  border: none;
  border-radius: 8px;
  font-size: 1rem;
  cursor: pointer;
  transition: all 0.3s ease-in-out;
  position: relative;
  overflow: hidden;
}

.btn.bg-ijomuda:hover {
  background-color: #5e9278;
  box-shadow: 0 12px 24px rgba(0, 0, 0, 0.2);
}

.swiper-container {
  width: 100%;
  padding-bottom: 50px;
}
</style>
