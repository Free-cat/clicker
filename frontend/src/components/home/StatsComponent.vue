<template>
  <div class="flex justify-around mb-8 mt-2">
    <div class="text-center">
      <div class="text-sm text-gray-400">Earned Today</div>
      <div class="flex items-center justify-center space-x-2 mt-1">
        <img src="/images/dollar.png" alt="Icon" class="w-6 h-6">
        <span class="text-lg">{{ dailyTokens }}</span>
      </div>
    </div>
    <div class="text-center">
      <div class="text-sm text-gray-400">Total earned</div>
      <div class="flex items-center justify-center space-x-2 mt-1">
        <img src="/images/dollar.png" alt="Icon" class="w-6 h-6">
        <span class="text-lg">{{ totalClicks }}</span>
      </div>
    </div>
    <div class="text-center">
      <div class="text-sm text-gray-400">Per click</div>
      <div class="flex items center justify-center space-x-2 mt-1">
        <img src="/images/dollar.png" alt="Icon" class="w-6 h-6">
        <span class="text-lg">+{{ tokensByClick }}</span>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    dailyTokens: Number,
    totalClicks: Number,
    tokensByClick: Number,
  }
};
</script>
