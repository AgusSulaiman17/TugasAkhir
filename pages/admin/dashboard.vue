<template>
  <div>
    <LoadingSpinner v-if="isLoading" />
    <AppNavbar />
    <div class="admin-dashboard">
      <div class="welcome-card bg-ijomuda">
        <div class="card-body">
          <h2 class="card-title">Selamat Datang di Lentera</h2>
          <p class="card-subtitle">
            Platform peminjaman buku yang <span class="highlight">mudah</span> dan
            <span class="highlight ">cepat</span>
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
          <div class="card-details">

            <p class="text-title">{{ item.judul }}</p>

            <p class="text-body">
              Kelola {{ item.judul }} ada disini
            </p>
          </div>
          <button @click="next(item.link)" class="card-button">
            Lihat {{ item.judul }}
          </button>
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
import { BIcon } from 'bootstrap-icons-vue';

export default {
  middleware: "admin",
  name: "AdminDashboard",
  components: {
    AppNavbar,
    LoadingSpinner,
    AppFooter,
    BIcon
  },
  data() {
    return {
      content: [
        { judul: "Genre", link: "genres" },
        { judul: "Penulis", link: "penulis" },
        { judul: "Buku", link: "buku" },
        { judul: "Data Peminjaman", link: "dataPeminjaman" },
        { judul: "Pengembalian", link: "pengembalian" },
        { judul: "Users", link: "Users" },
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
    getIconForCategory(judul) {
      switch (judul) {
        case "Genre":
          return "bookmarks"; // Ikon untuk genre
        case "Penulis":
          return "pen"; // Ikon untuk penulis
        case "Buku":
          return "book"; // Ikon untuk buku
        case "Data Peminjaman":
          return "clipboard"; // Ikon untuk data peminjaman
        case "Pengembalian":
          return "reply"; // Ikon untuk pengembalian
        case "Users":
          return "person"; // Ikon untuk users
        default:
          return "file-earmark"; // Ikon default
      }
    },
  },
};
</script>


<style scoped>
.admin-dashboard {
  padding: 2rem;
  margin-top: 100px;
  min-height: 100vh;
}

.welcome-card {
  color: white;
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
  font-weight: bold;
}

.cards-container {
  display: flex;
  flex-wrap: wrap;
  gap: 2rem;
  justify-content: center;
}

.card {
  width: 190px;
  height: 254px;
  border-radius: 20px;
  background: #f5f5f5;
  position: relative;
  padding: 1.8rem;
  border: 2px solid #c3c6ce;
  transition: 0.5s ease-out;
  overflow: visible;
}

.card-details {
  color: black;
  height: 100%;
  gap: 0.5em;
  display: grid;
  place-content: center;
}

.card-button {
  transform: translate(-50%, 125%);
  width: 60%;
  border-radius: 1rem;
  border: none;
  background-color: #70a799;
  color: #fff;
  font-size: 1rem;
  padding: 0.5rem 1rem;
  position: absolute;
  left: 50%;
  bottom: 0;
  opacity: 0;
  transition: 0.3s ease-out;
}

.text-body {
  color: rgb(134, 134, 134);
}

.text-title {
  font-size: 1.5em;
  font-weight: bold;
}

/* Hover Effects */
.card:hover {
  border-color: #70a799;
  box-shadow: 0 4px 18px 0 rgba(0, 0, 0, 0.25);
}

.card:hover .card-button {
  transform: translate(-50%, 50%);
  opacity: 1;
}

/* Responsif */
@media (max-width: 768px) {
  .cards-container {
    flex-direction: column;
    align-items: center;
  }

  .card {
    width: 80%; /* Membuat card lebih lebar di layar kecil */
  }
}
</style>
