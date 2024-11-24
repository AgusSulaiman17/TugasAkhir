export default {
  // Global page headers: https://go.nuxtjs.dev/config-head
  head: {
    title: 'frontend',
    htmlAttrs: {
      lang: 'en',
    },
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: '' },
      { name: 'format-detection', content: 'telephone=no' },
    ],
    link: [{ rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }],
  },

  // Global CSS: https://go.nuxtjs.dev/config-css
  css: ['~/assets/css/main.css'],

  // Auto import components: https://go.nuxtjs.dev/config-components
  components: true,

  // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
  buildModules: [
    // https://go.nuxtjs.dev/eslint
  ],

  // Modules: https://go.nuxtjs.dev/config-modules
  modules: [
    // https://go.nuxtjs.dev/bootstrap
    'bootstrap-vue/nuxt',
    // https://go.nuxtjs.dev/axios
    '@nuxtjs/axios',
  ],

  bootstrapVue: {
    icons: true, // Aktifkan ikon di BootstrapVue
  },

  plugins: [
    '~/plugins/auth.js',  // Menambahkan plugin auth
    '~/plugins/axios.js'
  ],

  ssr: false,

  axios: {
    proxy: true
  },
  proxy: {
    '/api/': { target: 'http://localhost:8080', pathRewrite: { '^/api/': '' } }
  },

  build: {
    transpile: ['vuex-persist'],
  },
  router: {
    extendRoutes(routes, resolve) {
      routes.push(
        {
          name: 'admin-genres',
          path: '/admin/genres',
          component: resolve(__dirname, 'pages/admin/genres.vue')
        },
        {
          name: 'detail-buku',
          path: '/DetailBuku/:id_buku',
          component: resolve(__dirname, 'pages/DetailBuku.vue'),
        },
      );
    }
  }
}
