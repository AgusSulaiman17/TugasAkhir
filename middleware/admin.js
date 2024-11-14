export default function ({ store, redirect }) {
  // Cek apakah pengguna sudah login dan perannya adalah admin
  if (!store.state.user || store.state.user.role !== 'admin') {
    // Jika bukan admin, arahkan ke halaman home atau halaman login
    return redirect('/')
  }
}
