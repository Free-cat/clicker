<template>
  <div v-if="!isLoading">
    <router-view></router-view>
  </div>
  <div v-else class="loader-container">
    <div class="loader"></div>
    <p class="text-white">Loading...</p>
  </div>
</template>

<script>
import { mapActions, mapState } from 'vuex';

export default {
  name: 'App',
  data() {
    return {
      isLoading: true
    };
  },
  computed: {
    ...mapState('websocket', ['socket'])
  },
  methods: {
    ...mapActions('websocket', ['connectWebSocket']),
    async initializeWebSocket() {
      try {
        await this.connectWebSocket();
      } catch (error) {
        console.error('Failed to connect to WebSocket:', error);
      }
      this.isLoading = false;
    }
  },
  async created() {
    await this.initializeWebSocket();
  }
};
</script>

<style scoped>
.loader-container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: #1f2937; /* Background color matching the app */
}

.loader {
  border: 8px solid #f3f3f3; /* Light grey */
  border-top: 8px solid #3498db; /* Blue */
  border-radius: 50%;
  width: 60px;
  height: 60px;
  animation: spin 2s linear infinite;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

.text-white {
  color: white;
}
</style>
