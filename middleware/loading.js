export default function({ store }) {
  // Set isLoading menjadi true saat navigasi mulai
  store.dispatch('setLoading', true);
}
