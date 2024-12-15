<template>
  <div class="container">
    <div class="row justify-content-between align-items-center mb-5">
      <div class="col-auto">
        <h3>Buku Yang Ada Disini</h3>
      </div>
      <div class="col-auto">
        <nuxt-link to="/ListBuku" class="btn bg-ijomuda">Lihat Semua  <b-icon-eye></b-icon-eye></nuxt-link>
      </div>
    </div>
    <div class="slider">
      <div v-if="bukuList.length" class="slides-container" :style="{ transform: `translateX(-${currentSlideIndex * 100}%)` }">
        <div class="card" v-for="buku in bukuList" :key="buku.id_buku" @click="$router.push(`/DetailBuku/${buku.id_buku}`)">
          <img v-if="buku.gambar"
               :src="buku.gambar.startsWith('http') ? buku.gambar : `http://localhost:8080${buku.gambar}`"
               :alt="buku.judul" class="card-img-top" style="height: 200px;" />
          <div v-else class="text-center p-3">Tidak ada gambar</div>
          <div class="card-info">
            <h3>{{ buku.judul }}</h3>
          </div>
        </div>
      </div>
      <div class="controls">
        <button @click="prevSlide" :disabled="currentSlideIndex === 0">&lt;</button>
        <button @click="nextSlide" :disabled="isLastSlide"> &gt; </button>
      </div>
      <div class="indicators">
        <span
          v-for="(dot, index) in totalSlides"
          :key="index"
          :class="{ active: index === currentSlideIndex }"
          @click="goToSlide(index)"
        ></span>
      </div>
    </div>
  </div>
</template>


<script>
import { getBukuList } from '@/api/buku';

export default {
  data() {
    return {
      bukuList: [],
      currentSlideIndex: 0,
      maxSlideIndex: 0,
      autoSlideInterval: null,
      totalSlides: 0,
      waiting: false, // Untuk menunggu 5 detik di slide terakhir
    };
  },
  async created() {
    try {
      this.bukuList = await getBukuList();
      if (this.bukuList && this.bukuList.length) {
        this.calculateMaxSlideIndex();
        this.startAutoSlide();
      } else {
        console.warn("Tidak ada buku untuk ditampilkan");
      }
    } catch (error) {
      console.error("Error fetching buku list:", error);
    }
  },
  computed: {
    isLastSlide() {
      return this.currentSlideIndex === this.maxSlideIndex;
    },
  },
  methods: {
    prevSlide() {
      if (this.currentSlideIndex > 0) {
        this.currentSlideIndex -= 1;
      }
    },
    async nextSlide() {
      if (this.isLastSlide) {
        this.stopAutoSlide(); // Berhenti sementara
        this.waiting = true;
        await this.delay(5000); // Tunggu 5 detik
        this.waiting = false;
        this.currentSlideIndex = 0; // Kembali ke slide pertama
        this.startAutoSlide(); // Mulai ulang auto-slide
      } else {
        this.currentSlideIndex += 1;
      }
    },
    goToSlide(index) {
      this.currentSlideIndex = index;
    },
    startAutoSlide() {
      this.autoSlideInterval = setInterval(() => {
        if (!this.waiting) {
          this.nextSlide();
        }
      }, 4000);
    },
    stopAutoSlide() {
      clearInterval(this.autoSlideInterval);
    },
    delay(ms) {
      return new Promise(resolve => setTimeout(resolve, ms));
    },
    calculateMaxSlideIndex() {
      const visibleCards = 4; // Jumlah kartu per slide
      this.totalSlides = Math.ceil(this.bukuList.length / visibleCards);
      this.maxSlideIndex = this.totalSlides - 1;
    },
  },
  beforeDestroy() {
    this.stopAutoSlide();
  },
};
</script>

<style scoped>
.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
  background-color: #f9f9f9;
  border-radius: 8px;
}

.lihat-semua {
  display: inline-block;
  background-color: #70a799;
  color: white;
  text-decoration: none;
  padding: 0.5rem 1rem;
  border-radius: 5px;
  font-size: 1rem;
  transition: background-color 0.3s ease;
}

.lihat-semua:hover {
  background-color: #2a594d;
}


.slider {
  position: relative;
  width: 100%;
  overflow: hidden;
}

.slides-container {
  display: flex;
  justify-content: flex-start;
  flex-wrap: nowrap;
  gap: 1rem;
  width: 100%;
  transition: transform 0.5s ease;
}

.card {
  flex: 1 0 calc(25% - 1rem); /* 4 kartu per slide */
  background-color: white;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  border-radius: 5px;
  overflow: hidden;
  text-align: center;
  transition: transform 0.3s ease;
}

.card:hover {
  transform: translateY(-5px);
}

.card-img-top {
  width: 100%;
  height: auto;
}

.card-info {
  padding: 1rem;
}

.controls {
  display: flex;
  justify-content: space-between;
  position: absolute;
  top: 50%;
  width: 100%;
  transform: translateY(-50%);
}

button {
  background-color: rgba(0, 0, 0, 0.5);
  color: white;
  border: none;
  padding: 0.5rem 1rem;
  cursor: pointer;
}

button:disabled {
  background-color: rgba(0, 0, 0, 0.2);
  cursor: not-allowed;
}

.indicators {
  display: flex;
  justify-content: center;
  margin-top: 1rem;
}

.indicators span {
  display: inline-block;
  width: 10px;
  height: 10px;
  margin: 0 5px;
  background-color: rgba(0, 0, 0, 0.5);
  border-radius: 50%;
  cursor: pointer;
  transition: background-color 0.3s;
}

.indicators span.active {
  background-color: rgba(0, 0, 0, 0.9);
}
</style>
