<template>
  <div>
    <AppNavbar />
    <div class="admin-user mt-8 mx-auto" style="max-width: 1200px;">
      <h1 class="text-center mb-4">Manage Users</h1>

      <!-- Tombol untuk menambah pengguna baru -->
      <div class="text-center mb-4">
        <button class="btn bg-ijotua" @click="openCreateModal">
          Add User <b-icon-plus></b-icon-plus>
        </button>
      </div>

      <!-- Tabel untuk menampilkan pengguna -->
      <table class="table table-striped table-bordered table-hover">
        <thead class="bg-ijomuda text-white">
          <tr>
            <th>#</th>
            <th>Name</th>
            <th>Email</th>
            <th>Role</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(user, index) in users" :key="user.id_user">
            <td>{{ index + 1 }}</td>
            <td>{{ user.nama }}</td>
            <td>{{ user.email }}</td>
            <td>{{ user.role }}</td>
            <td>
              <button class="btn bg-kuning btn-sm" @click="editUser(user.id_user)">
                <b-icon-pencil></b-icon-pencil>
              </button>
              <button class="btn bg-merah btn-sm" @click="confirmDelete(user.id_user)">
                <b-icon-trash></b-icon-trash>
              </button>
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
          <h2>{{ newUser ? "Add User" : "Edit User" }}</h2>
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
              <label>Password</label>
              <input v-model="editingUser.kata_sandi" type="password" :required="newUser" />
            </div>
            <div>
              <label>Role</label>
              <select v-model="editingUser.role" required>
                <option value="admin">Admin</option>
                <option value="petugas">Petugas</option>
                <option value="user">User</option>
              </select>
            </div>
            <button type="submit" class="btn btn-success" :disabled="isLoading">
              <b-icon :spin="isLoading" name="spinner"></b-icon> {{ newUser ? "Create" : "Save" }}
            </button>
            <button type="button" class="btn btn-secondary" @click="closeEditModal">
              Cancel
            </button>
          </form>
        </div>
      </div>

      <!-- Modal notifikasi sukses -->
      <NotificationModal v-if="showSuccessModal" :isVisible="showSuccessModal" :messageTitle="successTitle"
        :messageBody="successMessage" @close="closeSuccessModal" />

      <!-- Modal konfirmasi hapus -->
      <NotificationModal v-if="showDeleteConfirmModal" :isVisible="showDeleteConfirmModal"
        :messageTitle="'Delete Confirmation'" :messageBody="'Are you sure you want to delete this user?'"
        @close="closeDeleteConfirmModal">
        <template #footer>
          <button @click="deleteUser(deleteUserId)" class="btn bg-merah">
            Yes, Delete
          </button>
          <button @click="closeDeleteConfirmModal" class="btn bg-ijomuda">Cancel</button>
        </template>
      </NotificationModal>
    </div>
  </div>
</template>

<script>
import { getUsers, updateUser, deleteUser, createUser } from "@/api/user";
import AppNavbar from "~/components/AppNavbar.vue";
import NotificationModal from "~/components/NotificationModal.vue";

export default {
  components: {
    AppNavbar,
    NotificationModal,
  },
  data() {
    return {
      users: [],
      editingUser: null,
      newUser: false,
      showDeleteConfirmModal: false,
      showSuccessModal: false,
      deleteUserId: null,
      successTitle: "",
      successMessage: "",
      isLoading: false, // state loading
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
      this.editingUser = { nama: "", email: "", role: "user" };
    },
    editUser(id) {
      const user = this.users.find((u) => u.id_user === id);
      if (user) {
        this.editingUser = { ...user };
        this.newUser = false;
      }
    },
    async saveUser() {
      this.isLoading = true; // mulai loading
      try {
        if (this.newUser) {
          await createUser(this.editingUser);
          this.users.push({ ...this.editingUser }); // tambah pengguna baru di local state
          this.successTitle = "User Created";
          this.successMessage = "User created successfully.";
        } else {
          await updateUser(this.editingUser.id_user, this.editingUser);
          const index = this.users.findIndex(
            (user) => user.id_user === this.editingUser.id_user
          );
          this.users[index] = { ...this.editingUser };
          this.successTitle = "User Updated";
          this.successMessage = "User updated successfully.";
        }
        this.showSuccessModal = true;
        this.closeEditModal();
        await this.getUsers(); // refresh data pengguna setelah perubahan
      } catch (error) {
        console.error(error);
      } finally {
        this.isLoading = false; // hentikan loading
      }
    },
    confirmDelete(id) {
      this.deleteUserId = id;
      this.showDeleteConfirmModal = true;
    },
    async deleteUser(id) {
      try {
        await deleteUser(id);
        this.users = this.users.filter((user) => user.id_user !== id);
        this.successTitle = "User Deleted";
        this.successMessage = "User deleted successfully.";
        this.showSuccessModal = true;
      } catch (error) {
        console.error(error);
      } finally {
        this.showDeleteConfirmModal = false;
      }
    },
    closeEditModal() {
      this.editingUser = null;
      this.newUser = false;
    },
    closeSuccessModal() {
      this.showSuccessModal = false;
      this.successTitle = "";
      this.successMessage = "";
    },
    closeDeleteConfirmModal() {
      this.showDeleteConfirmModal = false;
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
  z-index: 99999;
  width: 400px;
  height: 600px;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.modal-overlay {
  position: fixed;
  height: 100%;
  top: 0;
  left: 0;
  width: 100%;
  background-color: rgba(0, 0, 0, 0.6);
  z-index: 9999;
}

.modal h2 {
  margin-bottom: 20px;
  color: #333;
}

.modal form {
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

.modal button[type="submit"]:disabled,
.modal button[type="button"]:disabled {
  background-color: #ddd;
  cursor: not-allowed;
}

.notification-modal {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.notification-modal .message-title {
  font-size: 1.2rem;
  font-weight: bold;
  color: #333;
}

.notification-modal .message-body {
  font-size: 1rem;
  color: #666;
  text-align: center;
}
</style>
