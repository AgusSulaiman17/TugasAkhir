<template>
  <div>
    <AppNavbar />
    <div class="container mt-8">
      <!-- Filter Section -->
      <div class="row mb-4">
        <!-- Search bar -->
        <div class="col-md-4">
          <input type="text" v-model="searchQuery" placeholder="Cari buku..." class="form-control" />
        </div>
        <!-- Dropdown Genre -->
        <div class="col-md-3">
          <select v-model="selectedGenre" class="form-control">
            <option value="">Semua Genre</option>
            <option v-for="genre in genres" :key="genre.id_genre" :value="genre.id_genre">
              {{ genre.nama }}
            </option>
          </select>
        </div>
        <!-- Dropdown Penulis -->
        <div class="col-md-3">
          <select v-model="selectedPenulis" class="form-control">
            <option value="">Semua Penulis</option>
            <option v-for="penulis in penulisList" :key="penulis.id_penulis" :value="penulis.id_penulis">
              {{ penulis.nama }}
            </option>
          </select>
        </div>
        <!-- Button Search -->
      </div>

      <div class="row">
        <!-- Tampilkan pesan error jika gagal mendapatkan data -->
        <div v-if="error" class="alert alert-danger">{{ error }}</div>

        <!-- Jika tidak ada buku -->
        <div v-if="filteredBuku.length === 0 && !error" class="alert alert-info">
          Tidak ada buku untuk ditampilkan.
        </div>

        <!-- Looping untuk setiap buku -->
        <div v-for="buku in filteredBuku" :key="buku.id_buku" class="col-md-3 col-sm-6 mb-4">
          <div class="card ">
            <img v-if="buku.gambar"
              :src="buku.gambar.startsWith('http') ? buku.gambar : `http://localhost:8080${buku.gambar}`"
              :alt="buku.judul" class="card-img-top" style="height: 200px;" />
            <div v-else class="text-center p-3">Tidak ada gambar</div>
            <div class="card-body">
              <h5 class="card-title text-center">{{ buku.judul }}</h5>
              <p class="card-text m-0">{{ buku.penulis && buku.penulis.nama ? buku.penulis.nama : 'Tidak diketahui' }}
              </p>
              <p class="card-text m-0">{{ buku.genre && buku.genre.nama ? buku.genre.nama : 'Tidak diketahui' }}</p>
              <button @click="$router.push(`/DetailBuku/${buku.id_buku}`)" class="btn bg-ijomuda mt-2 w-100">
                Detail
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { getBukuList } from '@/api/buku';
import { getGenres } from '@/api/genres';
import { getPenulis } from '@/api/penulis';
import AppNavbar from '~/components/AppNavbar.vue';

export default {
  components: {
    AppNavbar,
  },
  data() {
    return {
      bukuList: [],
      genres: [],
      penulisList: [],
      searchQuery: '',
      selectedGenre: '',
      selectedPenulis: '',
      error: null,
    };
  },
  computed: {
    filteredBuku() {
      let result = this.bukuList;

      // Filter berdasarkan query pencarian
      if (this.searchQuery) {
        result = result.filter((buku) =>
          buku.judul.toLowerCase().includes(this.searchQuery.toLowerCase())
        );
      }

      // Filter berdasarkan genre
      if (this.selectedGenre) {
        result = result.filter((buku) => buku.id_genre === parseInt(this.selectedGenre));
      }

      // Filter berdasarkan penulis
      if (this.selectedPenulis) {
        result = result.filter((buku) => buku.id_penulis === parseInt(this.selectedPenulis));
      }

      return result;
    },
  },
  methods: {
    async fetchBuku() {
      try {
        this.error = null;
        const data = await getBukuList();
        console.log('Data Buku:', data);
        this.bukuList = data;
      } catch (error) {
        console.error(error);
        this.error = 'Gagal memuat data buku. Coba lagi nanti.';
      }
    },
    async fetchGenres() {
      try {
        this.error = null;
        const data = await getGenres();
        console.log('Data Genre:', data);
        this.genres = data;
      } catch (error) {
        console.error(error);
        this.error = 'Gagal memuat data genre.';
      }
    },
    async fetchPenulis() {
      try {
        this.error = null;
        const data = await getPenulis();
        console.log('Data Penulis:', data);
        this.penulisList = data;
      } catch (error) {
        console.error(error);
        this.error = 'Gagal memuat data penulis.';
      }
    },
    pesan(id_buku) {
      alert(`Pesan buku dengan ID: ${id_buku}`);
      // Implementasi pemesanan buku bisa ditambahkan di sini
    },
    searchBuku() {
      console.log(
        'Pencarian dilakukan:',
        this.searchQuery,
        this.selectedGenre,
        this.selectedPenulis
      );
    },
  },
  async mounted() {
    await this.fetchBuku();
    await this.fetchGenres();
    await this.fetchPenulis();
  },
};
</script>

<style scoped>
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

.btn.bg-ijomuda::before {
  content: '';
  position: absolute;
  top: 50%;
  left: 50%;
  width: 300%;
  height: 300%;
  background-color: rgba(255, 255, 255, 0.2);
  transition: all 0.4s ease-in-out;
  border-radius: 50%;
  transform: translate(-50%, -50%) scale(0);
}

.btn.bg-ijomuda:hover {
  background-color: #5e9278;
  box-shadow: 0 12px 24px rgba(0, 0, 0, 0.2);
  transform: translateY(-4px); /* Mengangkat tombol sedikit */
}

.btn.bg-ijomuda:hover::before {
  transform: translate(-50%, -50%) scale(1);
}

.btn.bg-ijomuda:active {
  transform: translateY(2px); /* Memberikan efek tekan tombol */
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
  transform: translateY(-5px); /* Efek melayang saat hover */
  box-shadow: 0 8px 12px rgba(0, 0, 0, 0.2); /* Perbesar bayangan saat hover */
}

.card-img-top {
  border-top-left-radius: 10px;
  border-top-right-radius: 10px;
  height: 200px; /* Tinggi gambar */
  background-color: #f7f7f7; /* Warna latar jika gambar tidak tersedia */
}

.card-body {
  padding: 20px; /* Padding konten */
}

.card-title {
  font-size: 25px; /* Ukuran judul */
  font-weight: bold;
  color: #333; /* Warna teks */
  margin-bottom: 0.5rem;
}

.card-text {
  font-size: 0.9rem; /* Ukuran teks deskripsi */
  color: #555; /* Warna teks deskripsi */
  margin-bottom: 0.5rem;
}

.btn-info {
  background-color: #ff6b00; /* Warna tombol */
  border: none; /* Hapus border tombol */
  color: white;
  font-weight: bold;
}

.btn-info:hover {
  background-color: #e55a00; /* Warna tombol saat hover */
}

.alert {
  border-radius: 10px; /* Sesuaikan dengan kartu */
}
</style>
