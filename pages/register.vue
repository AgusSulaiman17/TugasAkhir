<template>
  <div class="register-container d-flex flex-wrap">
    <LoadingSpinner />
    <div class="image-container col-lg-5">
      <img src="../static/images/bg2.jpg" alt="Register Image" />
    </div>
    <div class="form-container col-lg-7">
      <form @submit.prevent="handleRegister">
        <div class="card">
          <a class="register">Register</a>
          <div class="inputBox">
            <input v-model="user.nama" type="text" id="nama" required />
            <span class="label">Nama</span>
          </div>
          <div class="inputBox">
            <input v-model="user.email" type="email" id="email" required />
            <span class="label">Email</span>
          </div>
          <div class="inputBox">
            <input v-model="user.kata_sandi" type="password" id="kata_sandi" required />
            <span class="label">Kata Sandi</span>
          </div>
          <button class="enter" type="submit">Register</button>
        </div>
        <div class="login-link">
        <p>Sudah punya akun? <a @click="login">Login</a></p>
      </div>
      </form>
      <p v-if="errorMessage" style="color: red">{{ errorMessage }}</p>
    </div>
  </div>
</template>
<script>
import LoadingSpinner from '~/components/LoadingSpinner.vue';
export default {
  components:{
    LoadingSpinner
  },
  data() {
    return {
      user: {
        nama: '',
        email: '',
        kata_sandi: '',
      },
      errorMessage: '',
    };
  },
  methods: {
    async handleRegister() {
      try {
        await this.$axios.post('/register', this.user);
        this.$router.push('/login');
      } catch (error) {
        this.errorMessage = error.response.data.message;
      }
    },
    login() {
      this.$store.dispatch('navigateWithLoading', { path: '/login', router: this.$router });
    }
  }
};
</script>

<style scoped>
.register-container {
  display: flex;
  height: 100vh;
}

.image-container {
  flex: 1;
  background: url('https://via.placeholder.com/600x600') no-repeat center center;
  background-size: cover;
}

.image-container img{
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

.register {
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
  transform: translateX(150px) translateY(-20px);
  font-size: 10px;
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
