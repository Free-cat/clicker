<template>
  <div class="mt-6 p-8 clickable-image">
    <div @click.stop="handleButtonClick" :class="{'click-animation': isAnimating}" class="w-full h-full rounded-full mx-auto gradient-border flex items-center justify-center overflow-hidden">
      <img src="/images/level1.webp" alt="Hamster" class="w-full h-full object-cover rounded-full">
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      isPulsing: true,
      isVibrating: false,
      isAnimating: false  // Controls the click animation
    };
  },
  methods: {
    handleButtonClick(event) {
      this.isPulsing = false;
      this.isAnimating = true;  // Start the animation
      setTimeout(() => {
        this.isAnimating = false;  // End the animation
      }, 100);  // Short duration for click effect

      if (this.canVibrate()) {
        this.triggerVibration();
      }
      this.$emit('click', event);
    },
    triggerVibration() {
      if (navigator.vibrate) {
        navigator.vibrate(500);
      }
    },
    canVibrate() {
      return 'vibrate' in navigator;
    }
  }
};
</script>

<style scoped>
@keyframes pulse {
  0%, 100% {
    transform: scale(1);
  }
  50% {
    transform: scale(1.05);
  }
}
@keyframes click-effect {
  0% {
    transform: scale(1);
  }
  50% {
    transform: scale(0.98);
  }
  100% {
    transform: scale(1);
  }
}
.pulse {
  animation: pulse 2s infinite;
}
.click-animation {
  animation: click-effect 80ms ease;
}
.gradient-border {
  background: linear-gradient(to right, #3b82f6, #8b5cf6);
  border-radius: 9999px;
  padding: 10px;
}
</style>