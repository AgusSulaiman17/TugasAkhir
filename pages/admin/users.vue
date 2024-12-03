<template>
  <div>
    <h1>Manage Users</h1>

    <!-- Tombol untuk menambah pengguna baru -->
    <button @click="openCreateModal">Add User</button>

    <!-- Tabel untuk menampilkan pengguna -->
    <table>
      <thead>
        <tr>
          <th>ID</th>
          <th>Name</th>
          <th>Email</th>
          <th>Role</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="user in users" :key="user.id_user">
          <td>{{ user.id_user }}</td>
          <td>{{ user.nama }}</td>
          <td>{{ user.email }}</td>
          <td>{{ user.role }}</td>
          <td>
            <button @click="editUser(user.id_user)">Edit</button>
            <button @click="deleteUser(user.id_user)">Delete</button>
          </td>
        </tr>
      </tbody>
    </table>

    <!-- Modal untuk tambah/edit pengguna -->
    <div v-if="editingUser !== null || newUser">
      <h2>{{ newUser ? 'Add User' : 'Edit User' }}</h2>
      <form @submit.prevent="saveUser">
        <div>
          <label>Name</label>
          <input v-model="editingUser.nama" type="text" required />
        </div>
        <div>
          <label>Email</label>
          <input v-model="editingUser.email" type="email" required />
        </div>
        <div>
          <label>Role</label>
          <select v-model="editingUser.role" required>
            <option value="admin">Admin</option>
            <option value="user">User</option>
          </select>
        </div>
        <button type="submit">{{ newUser ? 'Create' : 'Save' }}</button>
        <button @click="closeEditModal">Cancel</button>
      </form>
    </div>
  </div>
</template>

<script>
import { getUsers, updateUser, deleteUser, createUser } from '@/api/user';

export default {
  data() {
    return {
      users: [],
      editingUser: null,
      newUser: false, // Menandakan apakah kita sedang menambah pengguna baru
    };
  },
  async created() {
    try {
      const usersData = await getUsers();
      this.users = usersData;
    } catch (error) {
      console.error(error);
    }
  },
  methods: {
    // Fungsi untuk menambah pengguna baru
    openCreateModal() {
      this.newUser = true;
      this.editingUser = { nama: '', email: '', role: 'user' }; // Menginisialisasi data kosong untuk pengguna baru
    },
    // Fungsi untuk mengedit pengguna
    editUser(id) {
      const user = this.users.find((u) => u.id_user === id);
      if (user) {
        this.editingUser = { ...user }; // Menyalin data pengguna untuk diedit
        this.newUser = false;
      }
    },
    // Fungsi untuk menyimpan perubahan pengguna (baik create maupun update)
    async saveUser() {
      try {
        if (this.newUser) {
          // Jika newUser true, maka kita sedang membuat pengguna baru
          await createUser(this.editingUser);
          this.users.push({ ...this.editingUser }); // Menambah pengguna baru ke tabel
        } else {
          // Jika newUser false, maka kita sedang memperbarui pengguna
          await updateUser(this.editingUser.id_user, this.editingUser);
          const index = this.users.findIndex((user) => user.id_user === this.editingUser.id_user);
          this.users[index] = { ...this.editingUser }; // Memperbarui data pengguna di tabel
        }
        this.closeEditModal(); // Menutup modal setelah proses selesai
      } catch (error) {
        console.error(error);
      }
    },
    // Fungsi untuk menghapus pengguna
    async deleteUser(id) {
      try {
        await deleteUser(id);
        this.users = this.users.filter((user) => user.id_user !== id); // Menghapus pengguna dari tabel
      } catch (error) {
        console.error(error);
      }
    },
    // Fungsi untuk menutup modal edit/tambah
    closeEditModal() {
      this.editingUser = null;
      this.newUser = false;
    }
  }
};
</script>
