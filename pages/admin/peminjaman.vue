<template>
  <div>
    <AppNavbar />
    <div class="mt-8">
    <h1>Admin Peminjaman</h1>

    <!-- Form untuk membuat peminjaman baru -->
    <form @submit.prevent="handleCreatePeminjaman">
      <div>
        <label for="id_buku">ID Buku</label>
        <input type="number" v-model="newPeminjaman.id_buku" required />
      </div>
      <div>
        <label for="durasi_hari">Durasi Hari</label>
        <input type="number" v-model="newPeminjaman.durasi_hari" required />
      </div>
      <div>
        <label for="jam_kembali">Jam Kembali</label>
        <input type="time" v-model="newPeminjaman.jam_kembali" required />
      </div>
      <button type="submit">Buat Peminjaman</button>
    </form>

    <hr />

    <!-- Daftar peminjaman -->
    <h2>Daftar Peminjaman</h2>
    <table>
      <thead>
        <tr>
          <th>ID Peminjaman</th>
          <th>ID User</th>
          <th>ID Buku</th>
          <th>Tanggal Pinjam</th>
          <th>Durasi Hari</th>
          <th>Jam Kembali</th>
          <th>Tanggal Kembali</th>
          <th>Status Kembali</th>
          <th>Denda</th>
          <th>Aksi</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="peminjaman in peminjamanList" :key="peminjaman.id_peminjaman">
          <td>{{ peminjaman.id_peminjaman }}</td>
          <td>{{ peminjaman.id_user }}</td>
          <td>{{ peminjaman.id_buku }}</td>
          <td>{{ peminjaman.tanggal_pinjam }}</td>
          <td>{{ peminjaman.durasi_hari }} hari</td>
          <td>{{ peminjaman.jam_kembali }}</td>
          <td>{{ peminjaman.tanggal_kembali }}</td>
          <td>{{ peminjaman.status_kembali ? 'Kembali' : 'Belum Kembali' }}</td>
          <td>{{ peminjaman.denda }}</td>
          <td>
            <button @click="handleReturnBook(peminjaman.id_peminjaman)" :disabled="peminjaman.status_kembali">
              Kembalikan Buku
            </button>
            <button @click="handleDeletePeminjaman(peminjaman.id_peminjaman)">Hapus</button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
  </div>
</template>

<script>
import {
  getAllPeminjaman,
  createPeminjaman,
  returnBook,
  deletePeminjaman,
} from '@/api/peminjaman';
import AppNavbar from '~/components/AppNavbar.vue';

export default {
  components:{
    AppNavbar
  },
  data() {
    return {
      peminjamanList: [],
      newPeminjaman: {
        id_buku: null,
        durasi_hari: null,
        jam_kembali: '',
      },
    };
  },
  methods: {
    async fetchPeminjaman() {
      try {
        this.peminjamanList = await getAllPeminjaman();
      } catch (error) {
        console.error('Error fetching peminjaman:', error);
      }
    },
    async handleCreatePeminjaman() {
      try {
        // Tambahkan id_user otomatis berdasarkan token yang ada
        const token = localStorage.getItem('token');
        const user = token ? JSON.parse(atob(token.split('.')[1])) : null;
        if (user) {
          this.newPeminjaman.id_user = user.id_user;
        }

        console.log('Sending data:', this.newPeminjaman); // Debugging log

        await createPeminjaman(this.newPeminjaman);
        this.fetchPeminjaman();
        this.newPeminjaman = { id_buku: null, durasi_hari: null, jam_kembali: '' };
      } catch (error) {
        console.error('Error creating peminjaman:', error);
      }
    },
    async handleReturnBook(id) {
      try {
        await returnBook(id);
        this.fetchPeminjaman();
      } catch (error) {
        console.error('Error returning book:', error);
      }
    },
    async handleDeletePeminjaman(id) {
      try {
        await deletePeminjaman(id);
        this.fetchPeminjaman();
      } catch (error) {
        console.error('Error deleting peminjaman:', error);
      }
    },
  },
  created() {
    this.fetchPeminjaman();
  },
};
</script>
