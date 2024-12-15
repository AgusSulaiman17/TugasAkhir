<template>
  <div>
    <AppNavbar />
    <div class="user-dashboard mt-8">
      <div class="welcome-card">
        <div class="card-body">
          <h2 class="card-title">Dashboard Peminjaman Buku</h2>
          <p v-if="user.role === 'user'" class="card-text">
            Selamat datang, <span class="user-name"> {{ user.nama }} </span>!<br />
            Lihat status peminjaman buku Anda dan temukan buku menarik lainnya.
          </p>
          <p v-if="user.role === 'petugas'" class="card-text">
            Selamat datang Petugas, <span class="user-name"> {{ user.nama }} </span>!<br />
            Buat Peminjaman, Verifikasi Peminjaman, dan Lihat Peminjaman Disini
          </p>
        </div>
      </div>
    </div>
    <SliderBuku />
    <AppFooter />
  </div>
</template>

<script>
import AppFooter from "~/components/AppFooter.vue";
import AppNavbar from "~/components/AppNavbar.vue";
import LoadingSpinner from "~/components/LoadingSpinner.vue";
import SliderBuku from "~/components/SliderBuku.vue";

export default {
  middleware: "auth",
  name: "UserDashboard",
  components: {
    AppNavbar,
    LoadingSpinner,
    SliderBuku,
    AppFooter
  },
  computed: {
    user() {
      return this.$store.getters.getUser;
    },
    isLoading() {
    return this.$store.getters.isLoading;
  }
  },
};
</script>

<style scoped>
.welcome-card {
  background: linear-gradient(to right, #70a799, #a8d5ba); /* Hijau Pastel ke Hijau Mint */
  color: #ffffff;
  border-radius: 16px;
  padding: 2.5rem;
  margin-bottom: 2.5rem;
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.1);
  text-align: center;
  transition: transform 0.3s, box-shadow 0.3s;
}

.welcome-card:hover {
  transform: scale(1.02);
  box-shadow: 0 12px 24px rgba(0, 0, 0, 0.15);
}

.user-dashboard {
  padding: 2rem;
  margin-top: 100px;
}

.user-name {
  font-weight: bold;
  text-transform: capitalize;
}

.btn {
  display: inline-block;
  margin-top: 1rem;
  padding: 0.75rem 1.5rem;
  font-size: 1rem;
  font-weight: bold;
  border-radius: 8px;
  text-decoration: none;
  transition: background-color 0.3s;
}

.primary-btn {
  background-color: #70a799; /* Hijau Pastel */
  color: #ffffff;
}

.primary-btn:hover {
  background-color: #559f86; /* Hijau Lebih Gelap */
}

.secondary-btn {
  background-color: #a8d5ba; /* Hijau Mint */
  color: #ffffff;
}

.secondary-btn:hover {
  background-color: #8cc5a7; /* Hijau Mint Lebih Gelap */
}

.action-cards {
  display: flex;
  gap: 2rem;
  justify-content: space-between;
  flex-wrap: wrap;
}

.action-card {
  background-color: #ffffff;
  border-radius: 16px;
  padding: 1.5rem;
  flex: 1;
  min-width: 300px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  transition: transform 0.3s, box-shadow 0.3s;
  text-align: center;
}

.action-card:hover {
  transform: translateY(-8px);
  box-shadow: 0 8px 20px rgba(0, 0, 0, 0.15);
}

.action-card img {
  max-width: 100px;
  margin-bottom: 1rem;
}

.action-card h3 {
  font-size: 1.25rem;
  margin-bottom: 0.5rem;
}

.action-card p {
  font-size: 1rem;
  color: #666666;
  margin-bottom: 1.5rem;
}
</style>
