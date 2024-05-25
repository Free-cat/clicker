<template>
  <div class="relative w-full max-w-md mx-auto mt-8">
    <div class="absolute left-0 -top-6 text-sm text-gray-400">{{ tier }}</div>
    <div class="absolute right-0 -top-6 text-sm text-gray-400">Level {{ currentLevel }}/{{ maxLevel }}</div>
    <div class="h-2 bg-gray-700 rounded-full mt-2">
      <div class="h-full bg-purple-500 rounded-full" :style="{ width: progress + '%' }"></div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    tier: {
      type: String,
      default: 'Silver'
    },
    currentLevel: {
      type: Number,
      required: true
    },
    maxLevel: {
      type: Number,
      required: true
    },
    progress: {
      type: Number,
      required: true
    },
    levelImage: {
      type: String,
      default: ''
    }
  }
};
</script>

<style scoped>
.relative {
  position: relative;
}

.absolute {
  position: absolute;
}
</style>
