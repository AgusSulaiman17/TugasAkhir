<template>
  <div>
    <h1>Manage Users</h1>

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

    <!-- Modal untuk edit pengguna -->
    <div v-if="editingUser">
      <h2>Edit User</h2>
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
        <button type="submit">Save</button>
        <button @click="closeEditModal">Cancel</button>
      </form>
    </div>
  </div>
</template>

<script>
import { getUsers, updateUser, deleteUser } from '@/api/user';

export default {
  data() {
    return {
      users: [],
      editingUser: null,
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
    // Fungsi untuk mengedit pengguna
    editUser(id) {
      const user = this.users.find((u) => u.id_user === id);
      if (user) {
        this.editingUser = { ...user }; // Menyalin data pengguna untuk diedit
      }
    },
    // Fungsi untuk menyimpan perubahan pengguna
    async saveUser() {
      try {
        await updateUser(this.editingUser.id_user, this.editingUser);
        const index = this.users.findIndex((user) => user.id_user === this.editingUser.id_user);
        this.users[index] = { ...this.editingUser }; // Memperbarui data pengguna di tabel
        this.editingUser = null; // Menutup modal edit
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
    // Fungsi untuk menutup modal edit
    closeEditModal() {
      this.editingUser = null;
    }
  }
};
</script>

<style scoped>
table {
  width: 100%;
  border-collapse: collapse;
}
th, td {
  padding: 8px;
  text-align: left;
  border: 1px solid #ddd;
}
button {
  margin: 5px;
}
div {
  margin-top: 20px;
}
</style>
