<template>
  <div>
    <AppNavbar />
    <div class="admin-user mt-8 mx-auto" style="max-width: 1200px;">
      <h1 class="text-center mb-4">Manage Users</h1>

      <!-- Tombol untuk menambah pengguna baru -->
      <div class="text-center mb-4">
        <button class="btn btn-primary" @click="openCreateModal">Add User</button>
      </div>

      <!-- Tabel untuk menampilkan pengguna -->
      <table class="table table-striped table-bordered table-hover">
        <thead class="bg-ijomuda text-white">
          <tr>
            <th>#</th>
            <th>ID User</th>
            <th>Name</th>
            <th>Email</th>
            <th>Role</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(user, index) in users" :key="user.id_user">
            <td>{{ index + 1 }}</td>
            <td>{{ user.id_user }}</td>
            <td>{{ user.nama }}</td>
            <td>{{ user.email }}</td>
            <td>{{ user.role }}</td>
            <td>
              <button class="btn btn-warning btn-sm" @click="editUser(user.id_user)">Edit</button>
              <button class="btn btn-danger btn-sm" @click="confirmDelete(user.id_user)">Delete</button>
            </td>
          </tr>
          <tr v-if="users.length === 0">
            <td colspan="6" class="text-center">No users available</td>
          </tr>
        </tbody>
      </table>

      <!-- Modal untuk tambah/edit pengguna -->
      <div v-if="editingUser !== null || newUser">
        <div class="modal-overlay" @click="closeEditModal"></div>
        <div class="modal">
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
            <button type="submit" class="btn btn-success">{{ newUser ? 'Create' : 'Save' }}</button>
            <button type="button" class="btn btn-secondary" @click="closeEditModal">Cancel</button>
          </form>
        </div>
      </div>

      <!-- Modal konfirmasi hapus -->
      <div v-if="deleteModal">
        <div class="modal-overlay" @click="deleteModal = false"></div>
        <div class="modal">
          <h2>Confirm Delete</h2>
          <p>Are you sure you want to delete this user?</p>
          <button class="btn btn-danger" @click="deleteUser(deleteUserId)">Yes</button>
          <button class="btn btn-secondary" @click="deleteModal = false">No</button>
        </div>
      </div>

      <!-- Modal notifikasi sukses -->
      <div v-if="successModal">
        <div class="modal-overlay" @click="closeSuccessModal"></div>
        <div class="modal">
          <h2>Success</h2>
          <p>{{ successMessage }}</p>
          <button class="btn btn-primary" @click="closeSuccessModal">OK</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { getUsers, updateUser, deleteUser, createUser } from '@/api/user';
import AppNavbar from '~/components/AppNavbar.vue';

export default {
  components: {
    AppNavbar,
  },
  data() {
    return {
      users: [],
      editingUser: null,
      newUser: false,
      deleteModal: false,
      deleteUserId: null,
      successModal: false,
      successMessage: '',
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
    openCreateModal() {
      this.newUser = true;
      this.editingUser = { nama: '', email: '', role: 'user' };
    },
    editUser(id) {
      const user = this.users.find((u) => u.id_user === id);
      if (user) {
        this.editingUser = { ...user };
        this.newUser = false;
      }
    },
    async saveUser() {
      try {
        if (this.newUser) {
          await createUser(this.editingUser);
          this.users.push({ ...this.editingUser });
          this.successMessage = 'User created successfully';
        } else {
          await updateUser(this.editingUser.id_user, this.editingUser);
          const index = this.users.findIndex((user) => user.id_user === this.editingUser.id_user);
          this.users[index] = { ...this.editingUser };
          this.successMessage = 'User updated successfully';
        }
        this.successModal = true;
        this.closeEditModal();
      } catch (error) {
        console.error(error);
      }
    },
    confirmDelete(id) {
      this.deleteUserId = id;
      this.deleteModal = true;
    },
    async deleteUser(id) {
      try {
        await deleteUser(id);
        this.users = this.users.filter((user) => user.id_user !== id);
        this.successMessage = 'User deleted successfully';
        this.successModal = true;
      } catch (error) {
        console.error(error);
      } finally {
        this.deleteModal = false;
      }
    },
    closeEditModal() {
      this.editingUser = null;
      this.newUser = false;
    },
    closeSuccessModal() {
      this.successModal = false;
      this.successMessage = '';
    },
  },
};
</script>

<style scoped>
.admin-user {
  padding: 20px;
  background-color: #f9f9f9;
  border-radius: 8px;
  box-shadow: 0px 4px 6px rgba(0, 0, 0, 0.1);
}

.table th,
.table td {
  text-align: center;
  vertical-align: middle;
}

h1 {
  font-size: 1.8rem;
  font-weight: bold;
  color: #333;
}

button {
  margin: 5px;
}

.modal {
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  background-color: #ffffff;
  padding: 20px;
  border-radius: 12px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.2);
  z-index: 2000;
  width: 400px;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.6);
  z-index: 1000;
}

.modal h2 {
  margin-bottom: 20px;
  color: #333;
}

.modal form {
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: 15px;
}

.modal label {
  font-size: 14px;
  font-weight: 600;
  color: #555;
}

.modal input,
.modal select {
  width: 100%;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 8px;
}

.modal button[type="submit"],
.modal button[type="button"] {
  padding: 10px 20px;
  border-radius: 8px;
  cursor: pointer;
}

.modal button[type="submit"] {
  background-color: #4caf50;
  color: white;
}

.modal button[type="button"] {
  background-color: #f44336;
  color: white;
}

.modal button[type="submit"]:hover {
  background-color: #45a049;
}

.modal button[type="button"]:hover {
  background-color: #e41e1e;
}

@media (max-width: 480px) {
  .modal {
    width: 90%;
  }
}
</style>
