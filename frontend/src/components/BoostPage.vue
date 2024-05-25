<template>
  <div class="bg-gray-900 text-white min-h-screen p-4">
    <!-- Header -->
    <header class="flex items-center justify-between mb-4">
      <button class="text-white" @click="goBack">Back</button>
    </header>

    <!-- Balance Section -->
    <section class="balance-section text-center my-4">
      <coin-amount-component :farmed="347"></coin-amount-component>
      <div class="text-yellow-400 mt-2">How boosts work?</div>
    </section>

    <!-- Daily Boosters -->
    <section class="daily-boosters my-4">
      <h2 class="text-lg font-bold mb-2 text-gray-300">Free daily boosts</h2>
      <div class="booster-card bg-gray-800 p-4 rounded-lg flex items-center justify-between mb-2">
        <div class="flex items-center">
          <span class="text-yellow-500">⚡</span>
          <span class="ml-2 text-white">Full energy</span>
        </div>
        <div class="text-gray-400">4/6 available</div>
      </div>
      <div class="booster-card bg-gray-800 p-4 rounded-lg flex items-center justify-between mb-2">
        <div class="flex items-center">
          <span class="text-gray-500">🚀</span>
          <span class="ml-2 text-gray-500">Turbo</span>
        </div>
        <div class="text-gray-500">Soon...</div>
      </div>
    </section>

    <!-- Other Boosters -->
    <section class="other-boosters my-4">
      <h2 class="text-lg font-bold mb-2 text-gray-300">Boosts</h2>
      <div class="booster-card bg-gray-800 p-4 rounded-lg flex items-center justify-between mb-2">
        <div class="flex items-center">
          <span class="text-white">👆</span>
          <span class="ml-2 text-white">Multitap</span>
        </div>
        <div class="text-gray-400">8K • 4 lvl</div>
      </div>
      <div class="booster-card bg-gray-800 p-4 rounded-lg flex items-center justify-between mb-2">
        <div class="flex items-center">
          <span class="text-green-500">🔋</span>
          <span class="ml-2 text-white">Energy limit</span>
        </div>
        <div class="text-gray-400">2K • 2 lvl</div>
      </div>
    </section>
  </div>
</template>

<script>
import CoinAmountComponent from "@/components/home/CoinAmountComponent.vue";

export default {
  components: {CoinAmountComponent},
  methods: {
    goBack() {
      this.$router.go(-1); // Logic to go back
    }
  }
};
</script>

<style scoped>
.bg-gray-900 {
  background-color: #1f2937;
}

.bg-gray-800 {
  background-color: #2d3748;
}

.text-yellow-500 {
  color: #ecc94b;
}

.text-yellow-400 {
  color: #d69e2e;
}

.text-gray-400 {
  color: #a0aec0;
}

.text-gray-300 {
  color: #cbd5e0;
}

.text-gray-500 {
  color: #718096;
}

.booster-card {
  background-color: #2d3748;
  padding: 16px;
  border-radius: 8px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.nav-bar {
  background-color: #2d3748;
  padding: 16px;
  position: fixed;
  bottom: 0;
  width: 100%;
  display: flex;
  justify-content: space-around;
  align-items: center;
}
</style>
