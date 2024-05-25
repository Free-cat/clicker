<template>
  <div class="mt-6 text-xl flex justify-between items-center">
    <div class="flex items-center space-x-2">
      <span :class="{'text-yellow-500': true, 'animate-flash': shouldAnimate}">⚡</span>
      <span>{{ energy }} / {{ energyLimit }}</span>
    </div>
    <router-link to="/boost" class="flex items-center space-x-2 cursor-pointer">
      <span class="text-blue-500">🚀</span>
      <span>Boost</span>
    </router-link>
  </div>
</template>

<script>
export default {
  props: {
    energy: {
      type: Number,
      required: true
    },
    energyLimit: {
      type: Number,
      required: true
    }
  },
  data() {
    return {
      shouldAnimate: false
    };
  },
  watch: {
    energy(newValue) {
      if (newValue === this.energyLimit) {
        this.triggerAnimation();
      }
    }
  },
  methods: {
    triggerAnimation() {
      this.shouldAnimate = true;
      setTimeout(() => {
        this.shouldAnimate = false;
      }, 1000); // Duration of the animation
    },
    boostEnergy() {
      this.$emit('boost');
    }
  }
};
</script>

<style scoped>
.flex {
  display: flex;
  align-items: center;
}
.cursor-pointer {
  cursor: pointer;
}

@keyframes flash {
  0%, 100% {
    opacity: 1;
  }
  50% {
    opacity: 0;
  }
}

.animate-flash {
  animation: flash 1s;
}
</style>
