<template>
  <div class="login">
    <div class="login-container">
      <div class="image-container col-lg-5 p-0">
        <img src="../static/images/bg1.jpg" alt="Login Image" />
      </div>
      <div class="form-container col-lg-7">
        <form @submit.prevent="handleLogin">
          <div class="card">
            <a class="login">Log in</a>
            <div class="inputBox">
              <input v-model="user.email" id="email" type="email" required />
              <span class="user">Email</span>
            </div>
            <div class="inputBox">
              <input v-model="user.kata_sandi" id="password" type="password" required />
              <span>Password</span>
            </div>
            <button class="enter" type="submit" :disabled="isLoading">Login</button>
          </div>
          <div class="login-link">
            <p v-if="errorMessage" style="color: red">{{ errorMessage }}</p>
            <p>Belum punya akun? <router-link to="/register">Register</router-link></p>
          </div>
        </form>
      </div>
    </div>
    <LoadingSpinner v-if="isLoading" />
  </div>
</template>

<script>
import LoadingSpinner from "@/components/LoadingSpinner.vue";

export default {
  components: {
    LoadingSpinner
  },
  data() {
    return {
      user: {
        email: "",
        kata_sandi: ""
      },
      errorMessage: "",
    };
  },
  computed: {
    isLoading() {
      return this.$store.state.isLoading;
    }
  },
  methods: {
    handleLogin() {
      this.$store.dispatch('setLoading', true); // Set loading true sebelum mulai login

      // Simulasi loading 3 detik
      setTimeout(() => {
        this.$axios.post("/login", this.user)
          .then(response => {
            const token = response.data.token;
            const user = response.data.user; // Anggap response mengembalikan user dan token

            // Menyimpan token dan user ke Vuex store
            this.$store.dispatch("login", { token, user });

            // Decode token untuk mendapatkan role
            const decodedToken = this.decodeToken(token);

            if (decodedToken.role === "admin") {
              this.$router.push("/admin/dashboard");
            } else {
              this.$router.push("/user/dashboard");
            }
          })
          .catch(error => {
            this.errorMessage = error.response?.data?.message || "Login failed";
          })
          .finally(() => {
            this.$store.dispatch('setLoading', false); // Set loading false setelah proses selesai
          });
      }, 3000); // Simulasi loading selama 3 detik
    },
    decodeToken(token) {
      const payload = token.split(".")[1];
      const decoded = atob(payload);
      return JSON.parse(decoded);
    }
  },
  watch: {
    '$route'(to, from) {
      // Jika sudah login dan mengunjungi halaman login, redirect ke dashboard
      if (this.$store.getters.isAuthenticated && (to.name === 'login')) {
        const decodedRole = this.decodeToken(this.$store.state.token).role;
        this.$router.push(decodedRole === 'admin' ? '/admin/dashboard' : '/user/dashboard');
      }
    }
  }
};
</script>

<style scoped>
.login-container {
  display: flex;
  height: 100vh; /* Menggunakan seluruh tinggi layar */
}

.image-container {
  flex: 1;
  background-size: cover;
}

.image-container img {
  width: 100vh;
  height: 100vh;
}

.form-container {
  flex: 1;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: #f5f5f5;
}

.card {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 450px;
  width: 400px;
  flex-direction: column;
  gap: 35px;
  background: #e3e3e3;
  box-shadow: 16px 16px 32px #c8c8c8, -16px -16px 32px #fefefe;
  border-radius: 8px;
}

.login {
  color: #000;
  text-transform: uppercase;
  letter-spacing: 2px;
  font-weight: bold;
  font-size: x-large;
  display: block;
}

.inputBox {
  position: relative;
  width: 250px;
}

.inputBox input {
  width: 100%;
  padding: 10px;
  outline: none;
  border: none;
  color: #000;
  font-size: 1em;
  background: transparent;
  border-left: 2px solid #000;
  border-bottom: 2px solid #000;
  transition: 0.1s;
  border-bottom-left-radius: 8px;
}

.inputBox span {
  margin-top: 5px;
  position: absolute;
  left: 0;
  transform: translateY(-4px);
  margin-left: 10px;
  padding: 10px;
  pointer-events: none;
  font-size: 12px;
  color: #000;
  text-transform: uppercase;
  transition: 0.5s;
  letter-spacing: 3px;
  border-radius: 8px;
}

.inputBox input:valid~span,
.inputBox input:focus~span {
  transform: translateX(113px) translateY(-15px);
  font-size: 0.8em;
  padding: 5px 10px;
  background: #000;
  letter-spacing: 0.2em;
  color: #fff;
}

.inputBox input:valid,
.inputBox input:focus {
  border: 2px solid #000;
  border-radius: 8px;
}

.enter {
  height: 45px;
  width: 100px;
  border-radius: 5px;
  border: 2px solid #000;
  cursor: pointer;
  background-color: transparent;
  transition: 0.5s;
  text-transform: uppercase;
  font-size: 10px;
  letter-spacing: 2px;
  margin-bottom: 1em;
}

.enter:hover {
  background-color: rgb(0, 0, 0);
  color: white;
}

.login-link {
  text-align: center;
  margin-top: 20px;
}

.login-link a {
  color: #000;
  text-decoration: none;
  font-weight: bold;
}

.login-link a:hover {
  color: #007bff;
}
</style>
