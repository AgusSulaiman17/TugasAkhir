<template>
  <div>
    <AppNavbar />
    <div class="mt-8">
      <h1>Buku Management</h1>
      <div>
        <button @click="openCreateModal">Tambah Buku</button>
        <table>
          <thead>
            <tr>
              <th>Judul</th>
              <th>Penulis</th>
              <th>Genre</th>
              <th>Jumlah</th>
              <th>Gambar</th>
              <th>Aksi</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="buku in bukuList" :key="buku.id_buku">
              <td>{{ buku.judul }}</td>
              <td>{{ buku.penulis && buku.penulis.nama ? buku.penulis.nama : 'Tidak diketahui' }}</td>
              <td>{{ buku.genre && buku.genre.nama ? buku.genre.nama : 'Tidak diketahui' }}</td>
              <td>{{ buku.jumlah }}</td>
              <td><img v-if="buku.gambar" :src="`http://localhost:8080${buku.gambar}`" :alt="buku.judul"
                  class="card-img-top" style="width: 100px" />
                <span v-else class="text-center">Tidak ada gambar</span>
              </td>
              <td>
                <button @click="openEditModal(buku)">Edit</button>
                <button @click="deleteBuku(buku.id_buku)">Hapus</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Modal untuk menambah dan mengedit buku -->
      <div v-if="isModalOpen" class="modal">
        <h2>{{ isEditMode ? 'Edit Buku' : 'Tambah Buku' }}</h2>
        <form @submit.prevent="submitForm">
          <div>
            <label for="judul">Judul</label>
            <input type="text" v-model="buku.judul" required>
          </div>
          <div>
            <label for="id_penulis">Penulis</label>
            <select v-model="buku.id_penulis" required>
              <option v-for="penulis in penulisList" :key="penulis.id_penulis" :value="penulis.id_penulis">
                {{ penulis.nama }}
              </option>
            </select>
          </div>
          <div>
            <label for="id_genre">Genre</label>
            <select v-model="buku.id_genre" required>
              <option v-for="genre in genreList" :key="genre.id_genre" :value="genre.id_genre">
                {{ genre.nama }}
              </option>
            </select>
          </div>
          <div>
            <label for="deskripsi">Deskripsi</label>
            <textarea v-model="buku.deskripsi"></textarea>
          </div>
          <div>
            <label for="jumlah">Jumlah</label>
            <input type="number" v-model="buku.jumlah" required>
          </div>
          <div>
            <label for="gambar">Gambar</label>
            <input type="file" @change="onFileChange">
          </div>
          <button type="submit">{{ isEditMode ? 'Update' : 'Tambah' }}</button>
          <button @click="closeModal">Batal</button>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import { getBukuList, createBuku, updateBuku, deleteBuku } from '~/api/buku.js';
import { getPenulis } from '@/api/penulis';
import { getGenres } from '@/api/genres';
import AppNavbar from '~/components/AppNavbar.vue';

export default {
  components: {
    AppNavbar
  },
  data() {
    return {
      bukuList: [],
      penulisList: [],
      genreList: [],
      buku: {
        id_buku: null,
        judul: '',
        id_penulis: '',
        id_genre: '',
        deskripsi: '',
        jumlah: 0,
        gambar: null,
      },
      isModalOpen: false,
      isEditMode: false,
    };
  },
  async created() {
    await this.getBukuList();
    // Fetch penulisList dan genreList juga
    await this.fetchPenulisList();
    await this.fetchGenreList();
  },
  methods: {
    async getBukuList() {
      try {
        this.bukuList = await getBukuList();
      } catch (error) {
        console.error('Error fetching buku list:', error);
      }
    },
    async fetchPenulisList() {
      try {
        const data = await getPenulis(); // Memanggil fungsi getPenulis dari api/penulis.js
        console.log('Penulis Data:', data); // Periksa hasilnya di console
        this.penulisList = data || []; // Simpan data penulis ke penulisList
      } catch (error) {
        console.error('Error fetching penulis:', error);
      }
    },
    async fetchGenreList() {
      try {
        const data = await getGenres(); // Memanggil fungsi getGenres dari api/genres.js
        console.log('Genre Data:', data); // Periksa hasilnya di console
        this.genreList = data || []; // Simpan data genre ke genreList
      } catch (error) {
        console.error('Error fetching genre:', error);
      }
    },
    openCreateModal() {
      this.isModalOpen = true;
      this.isEditMode = false;
      this.buku = {
        id_buku: null,
        judul: '',
        id_penulis: '',
        id_genre: '',
        deskripsi: '',
        jumlah: 0,
        gambar: null,
      };
    },
    openEditModal(buku) {
      this.isModalOpen = true;
      this.isEditMode = true;
      this.buku = { ...buku };
    },
    closeModal() {
      this.isModalOpen = false;
    },
    onFileChange(event) {
      const file = event.target.files[0];
      if (file) {
        this.buku.gambar = file;
      }
    },
    async submitForm() {
      try {
        if (this.isEditMode) {
          await updateBuku(this.buku);
        } else {
          await createBuku(this.buku);
        }
        this.getBukuList();
        this.closeModal();
      } catch (error) {
        console.error('Error while saving data:', error);
        alert('Error while saving data');
      }
    },
    async deleteBuku(id_buku) {
      if (confirm('Are you sure you want to delete this book?')) {
        try {
          await deleteBuku(id_buku);
          this.getBukuList();
        } catch (error) {
          console.error('Error deleting buku:', error);
        }
      }
    },
  },
};
</script>

<style scoped>
.modal {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
}

form {
  background-color: white;
  padding: 20px;
  border-radius: 5px;
}
</style>
