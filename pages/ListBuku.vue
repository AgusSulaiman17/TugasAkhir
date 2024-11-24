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
            <option value="">Pilih Genre</option>
            <option v-for="genre in genres" :key="genre.id_genre" :value="genre.id_genre">
              {{ genre.nama }}
            </option>
          </select>
        </div>
        <!-- Dropdown Penulis -->
        <div class="col-md-3">
          <select v-model="selectedPenulis" class="form-control">
            <option value="">Pilih Penulis</option>
            <option v-for="penulis in penulisList" :key="penulis.id_penulis" :value="penulis.id_penulis">
              {{ penulis.nama }}
            </option>
          </select>
        </div>
        <!-- Button Search -->
        <div class="col-md-2">
          <button @click="searchBuku" class="btn btn-primary w-100">Search</button>
        </div>
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
          <div class="card h-100">
            <img v-if="buku.gambar"
              :src="buku.gambar.startsWith('http') ? buku.gambar : `http://localhost:8080${buku.gambar}`"
              :alt="buku.judul" class="card-img-top" style="height: 300px; object-fit: cover;" />
            <div v-else class="text-center p-3">Tidak ada gambar</div>
            <div class="card-body">
              <h5 class="card-title text-center">{{ buku.judul }}</h5>
              <p class="card-text m-0">{{ buku.penulis && buku.penulis.nama ? buku.penulis.nama : 'Tidak diketahui' }}
              </p>
              <p class="card-text m-0">{{ buku.genre && buku.genre.nama ? buku.genre.nama : 'Tidak diketahui' }}</p>
              <button @click="$router.push(`/DetailBuku/${buku.id_buku}`)" class="btn btn-info mt-2 w-100">
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
.card {
  border-radius: 10px;
  overflow: hidden;
}

.card-img-top {
  border-top-left-radius: 10px;
  border-top-right-radius: 10px;
  object-fit: cover;
}

.card-body {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  align-items: center;
}
</style>
