export default ({ store }) => {
  // Memeriksa apakah ada token dan user di localStorage saat aplikasi dimulai
  if (process.client) {
    const token = localStorage.getItem('token');
    const user = localStorage.getItem('user');

    if (token && user) {
      store.commit('setToken', token);
      store.commit('setUser', JSON.parse(user));
    }
  }
};
