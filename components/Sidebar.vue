<template>
  <transition name="slide">
    <div v-if="isOpen" class="sidebar">
      <div class="sidebar-header">
        <h2 class="mt-2 ml-2">MENU</h2>
        <button class="close-btn" @click="closeSidebar">×</button>
      </div>
      <ul class="sidebar-menu">
        <li v-for="(item, index) in items" :key="index">
          <button
            :class="['btn', 'bg-ijomuda', 'text-white', { 'active-link': isActive(item.link) }]"
            @click="navigateWithLoading(item.link)"
          >
            {{ item.judul }}
          </button>
        </li>
      </ul>
      <!-- Loading Spinner -->
      <LoadingSpinner v-if="isLoading" />
    </div>
  </transition>
</template>

<script>
import LoadingSpinner from './LoadingSpinner.vue'; // Import the LoadingSpinner component

export default {
  name: "Sidebar",
  components: {
    LoadingSpinner,
  },
  props: {
    isOpen: {
      type: Boolean,
      required: true,
    },
    items: {
      type: Array,
      required: true,
    },
  },
  computed: {
    isLoading() {
      return this.$store.getters.isLoading;
    },
  },
  methods: {
    closeSidebar() {
      this.$emit("close");
    },
    navigate(link) {
      this.$emit("navigate", link);
    },
    isActive(link) {
      return this.$route.path === `/${link}`;
    },
    navigateWithLoading(link) {
      this.$store.dispatch('setLoading', true);
      setTimeout(() => {
        this.$emit("navigate", link);
        this.$store.dispatch('setLoading', false);
      }, );
    }
  },
};
</script>

<style scoped>
.sidebar {
  position: fixed;
  top: 0;
  left: 0;
  width: 250px;
  height: 100vh;
  background-color: #70a799;
  color: white;
  box-shadow: 2px 0 5px rgba(0, 0, 0, 0.2);
  z-index: 1000;
  padding: 1rem;
  display: flex;
  flex-direction: column;
  transition: transform 0.3s ease-in-out;
}

.sidebar-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.close-btn {
  background: none;
  border: none;
  color: white;
  font-size: 1.5rem;
  cursor: pointer;
}

.sidebar-menu {
  list-style: none;
  padding: 0;
  margin: 2rem 0 0;
}

.sidebar-menu li {
  margin-bottom: 1rem;
}

.sidebar-menu li a {
  color: white;
  text-decoration: none;
  font-size: 1.2rem;
  transition: color 0.3s ease-in-out;
}

.sidebar-menu li a:hover {
  color: #cce7d5;
}

.slide-enter-active,
.slide-leave-active {
  transition: transform 0.3s ease-in-out;
}

.slide-enter {
  transform: translateX(-100%);
}

.slide-leave-to {
  transform: translateX(-100%);
}

.btn.bg-ijomuda {
  background-color: #70a799;
  color: #fff;
  padding: 0.75rem 1.5rem;
  border: none;
  border-radius: 8px;
  font-size: 1rem;
  cursor: pointer;
  transition: all 0.3s ease-in-out;
  position: relative;
  overflow: hidden;
}

.btn.bg-ijomuda::before {
  content: '';
  position: absolute;
  top: 50%;
  left: 50%;
  width: 300%;
  height: 300%;
  background-color: rgba(255, 255, 255, 0.2);
  transition: all 0.4s ease-in-out;
  border-radius: 50%;
  transform: translate(-50%, -50%) scale(0);
}

.btn.bg-ijomuda:hover {
  background-color: #5e9278;
  box-shadow: 0 12px 24px rgba(0, 0, 0, 0.2);
  transform: translateY(-4px); 
}

.btn.bg-ijomuda:hover::before {
  transform: translate(-50%, -50%) scale(1);
}

.btn.bg-ijomuda:active {
  transform: translateY(2px);
}

.active-link {
  font-weight: bold;
  pointer-events: none;
}

</style>
