<template>
  <footer class="app-footer bg-ijotua">
    <div class="footer-container">
      <!-- Logo dan Deskripsi -->
      <div class="footer-logo">
        <h3>Lentera</h3>
        <p>Platform peminjaman buku yang mudah dan cepat. Jelajahi dunia dengan membaca!</p>
      </div>

      <!-- Navigasi Cepat -->
      <div class="footer-links">
        <h4>Navigasi Cepat</h4>
        <ul>
          <li><nuxt-link to="/">Beranda</nuxt-link></li>
          <li><nuxt-link to="/user/dashboard">Dashboard</nuxt-link></li>
          <li><nuxt-link to="/user/cariBuku">Cari Buku</nuxt-link></li>
          <li><nuxt-link to="/user/peminjaman">Buku Dipinjam</nuxt-link></li>
        </ul>
      </div>

      <!-- Informasi Kontak -->
      <div class="footer-contact">
        <h4>Hubungi Kami</h4>
        <p>Email: support@lentera.com</p>
        <p>Telepon: +62 812-3456-7890</p>
        <p>Alamat: Jl. Literasi No. 123, Jakarta</p>
      </div>
    </div>
    <div class="footer-bottom">
      <p>© {{ currentYear }} Lentera. All Rights Reserved.</p>
    </div>
  </footer>
</template>

<script>
export default {
  data() {
    return {
      currentYear: new Date().getFullYear(),
    };
  },
};
</script>

<style scoped>
.app-footer {
  color: #fff;
  padding: 2rem 1rem;
  margin-top: 3rem;
  font-family: Arial, sans-serif;
}

.footer-container {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;
  gap: 2rem;
}

.footer-logo {
  flex: 1;
}

.footer-logo h3 {
  font-size: 1.8rem;
  margin-bottom: 0.5rem;
}

.footer-logo p {
  font-size: 1rem;
  line-height: 1.6;
  color: #dfe6f1;
}

.footer-links {
  flex: 1;
}

.footer-links h4 {
  font-size: 1.2rem;
  margin-bottom: 1rem;
}

.footer-links ul {
  list-style: none;
  padding: 0;
}

.footer-links li {
  margin-bottom: 0.5rem;
}

.footer-links a {
  color: #fff;
  text-decoration: none;
  transition: color 0.3s;
}

.footer-links a:hover {
  color: #ffeb99;
}

.footer-contact {
  flex: 1;
}

.footer-contact h4 {
  font-size: 1.2rem;
  margin-bottom: 1rem;
}

.footer-contact p {
  font-size: 1rem;
  line-height: 1.6;
  color: #dfe6f1;
}

.footer-bottom {
  text-align: center;
  margin-top: 2rem;
  border-top: 1px solid rgba(255, 255, 255, 0.2);
  padding-top: 1rem;
  font-size: 0.9rem;
  color: #dfe6f1;
}

/* Responsif */
@media (max-width: 768px) {
  .footer-container {
    flex-direction: column;
    align-items: center;
  }

  .footer-links ul {
    text-align: center;
  }
}
</style>
