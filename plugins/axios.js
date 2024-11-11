export default function ({ $axios, store }) {
  // Menambahkan token dari Vuex store ke setiap request
  $axios.onRequest(config => {
    const token = store.state.token;
    if (token) {
      config.headers['Authorization'] = `Bearer ${token}`;
    }
    return config;
  });
}
