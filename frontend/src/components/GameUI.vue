<template>
  <div class="bg-gray-900 text-white min-h-screen p-4">
    <div ref="container" class="text-center w-full max-w-2xl mx-auto flex-grow relative">
      <stats-component
          :dailyTokens="dailyClicks"
          :totalClicks="totalClicks"
          :tokensByClick="1">
      ></stats-component>
      <coin-amount-component :farmed="totalClicks"></coin-amount-component>
      <progress-bar-component
          :tier="'Silver'"
          :currentLevel="level"
          :maxLevel="9"
          :progress="progress"
          :levelImage="levelImage">
      </progress-bar-component>
      <hamster-component @click="handleClick"></hamster-component>
      <energy-boost-component
          :energy="tokensLeft"
          :energyLimit="tokensLimit"
          @boost="handleBoost">
      </energy-boost-component>
      <!-- Plus One Animation -->
      <div class="animation-container absolute inset-0 pointer-events-none">
        <div
            v-for="plus in plusOnes"
            :key="plus.id"
            class="absolute text-4xl font-bold text-white fly-up"
            :style="{ left: plus.x + 'px', top: plus.y + 'px', opacity: plus.opacity }"
        >+1</div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapActions, mapGetters } from 'vuex';
import StatsComponent from './home/StatsComponent.vue';
import CoinAmountComponent from './home/CoinAmountComponent';
import ProgressBarComponent from './home/ProgressBarComponent';
import HamsterComponent from './home/HamsterComponent';
import EnergyBoostComponent from './home/EnergyBoostComponent';

export default {
  components: {
    StatsComponent,
    CoinAmountComponent,
    ProgressBarComponent,
    HamsterComponent,
    EnergyBoostComponent
  },
  computed: {
    ...mapGetters('websocket', [
      'tokensLeft',
      'tokensLimit',
      'totalClicks',
      'level',
      'progress',
      'levelImage',
      'dailyClicks'
    ])
  },
  data() {
    return {
      stats: [
        { text: 'Text 1', icon: '/images/dollar.png', value: 100 },
        { text: 'Text 2', icon: '/images/dollar.png', value: 200 },
        { text: 'Text 3', icon: '/images/dollar.png', value: 300 }
      ],
      plusOnes: []
    };
  },
  methods: {
    ...mapActions('websocket', [
      'connectWebSocket',
      'sendClickEvent'
    ]),
    handleClick(event) {
      this.sendClickEvent(); // Dispatch the Vuex action to send the click event

      const id = Date.now();
      const container = this.$refs.container.getBoundingClientRect();
      const randomX = Math.random() * 50 - 25; // Random x offset
      const x = event.clientX - container.left + randomX;
      const y = event.clientY - container.top - 50; // Adjusted to account for cursor position

      this.plusOnes.push({ id, x, y, opacity: 0 });

      // Fade out the "+1" animation after 1 second
      setTimeout(() => {
        this.plusOnes = this.plusOnes.filter(plus => plus.id !== id);
      }, 1000);
    }
  },
  mounted() {
    this.connectWebSocket(); // Connect to WebSocket when the component is mounted
  }
};
</script>

<style>
@keyframes fly-up {
  0% {
    opacity: 1;
    transform: translateY(0);
  }
  100% {
    opacity: 0;
    transform: translateY(-300px);
  }
}

.fly-up {
  animation: fly-up 2s forwards;
}
</style>
