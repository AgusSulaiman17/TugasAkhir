<template>
  <div>
    <LoadingSpinner v-if="isLoading" />
    <AppNavbar />
    <div class="admin-dashboard">
      <div class="welcome-card">
        <div class="card-body">
          <h2 class="card-title">✨ Selamat Datang di Lentera ✨</h2>
          <p class="card-subtitle">
            Platform peminjaman buku yang <span class="highlight">mudah</span> dan
            <span class="highlight">cepat</span>
          </p>
          <p class="card-text">
            Kelola buku, data peminjaman, dan pengembalian dengan mudah melalui
            dashboard ini.
          </p>
        </div>
      </div>

      <div class="cards-container">
        <div
          v-for="(item, index) in content"
          :key="index"
          class="card feature-card"
        >
          <div class="card-body">
            <h5 class="card-title">{{ item.judul }}</h5>
            <button @click="next(item.link)" class="btn btn-primary">
              ➡️ Lihat {{ item.judul }}
            </button>
          </div>
        </div>
      </div>
    </div>
    <AppFooter />
  </div>
</template>

<script>
import AppFooter from "~/components/AppFooter.vue";
import AppNavbar from "~/components/AppNavbar.vue";
import LoadingSpinner from "~/components/LoadingSpinner.vue";

export default {
  middleware: "admin",
  name: "AdminDashboard",
  components: {
    AppNavbar,
    LoadingSpinner,
    AppFooter
  },
  data() {
    return {
      content: [
        { judul: "Genre", link: "genres" },
        { judul: "Penulis", link: "penulis" },
        { judul: "Buku", link: "buku" },
        { judul: "Data Peminjaman", link: "dataPeminjaman" },
        { judul: "Pengembalian", link: "pengembalian" },
      ],
    };
  },
  computed: {
    isLoading() {
      return this.$store.state.isLoading;
    },
  },
  methods: {
    next(link) {
      if (link) {
        this.$store.dispatch("navigateWithLoading", {
          path: `/admin/${link}`,
          router: this.$router,
        });
      }
    },
  },
};
</script>

<style scoped>
.admin-dashboard {
  padding: 2rem;
  margin-top: 100px;
  background: linear-gradient(135deg, #f6f8ff, #eef2ff);
  min-height: 100vh;
}

.welcome-card {
  background-color: #ffffff;
  border-radius: 16px;
  padding: 2rem;
  margin-bottom: 2.5rem;
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.1);
  text-align: center;
  transition: transform 0.3s, box-shadow 0.3s;
}

.welcome-card:hover {
  transform: scale(1.02);
  box-shadow: 0 12px 24px rgba(0, 0, 0, 0.15);
}

.highlight {
  color: #007bff;
  font-weight: bold;
}

.cards-container {
  display: flex;
  flex-wrap: wrap;
  gap: 2rem;
  justify-content: center;
}

.feature-card {
  width: 280px;
  background: #ffffff;
  border-radius: 12px;
  padding: 1.5rem;
  text-align: center;
  transition: transform 0.3s, box-shadow 0.3s;
  position: relative;
  overflow: hidden;
}

.feature-card:before {
  content: "";
  position: absolute;
  top: -50%;
  left: -50%;
  width: 200%;
  height: 200%;
  background: radial-gradient(circle, rgba(0, 123, 255, 0.2), transparent);
  transform: scale(0);
  transition: transform 0.5s;
  z-index: -1;
}

.feature-card:hover:before {
  transform: scale(1);
}

.feature-card:hover {
  transform: translateY(-8px);
  box-shadow: 0 16px 32px rgba(0, 0, 0, 0.15);
}

.card-title {
  font-size: 1.5rem;
  margin-bottom: 1.5rem;
  color: #333;
}

.btn-primary {
  background-color: #007bff;
  color: #ffffff;
  padding: 0.75rem 1.5rem;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  font-size: 1rem;
  transition: background-color 0.3s, transform 0.3s;
}

.btn-primary:hover {
  background-color: #0056b3;
  transform: scale(1.05);
}

/* Responsif */
@media (max-width: 768px) {
  .cards-container {
    flex-direction: column;
    align-items: center;
  }

  .feature-card {
    width: 90%;
  }
}
</style>
