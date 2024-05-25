<template>
  <div class="text-4xl font-bold flex items-center justify-center space-x-4 mb-8">
    <img src="/images/dollar.png" alt="Coin" class="w-10 h-10">
    <span>{{ farmed }}</span>
  </div>
</template>

<script>
export default {
  props: {
    farmed: {
      type: Number,
      required: true
    }
  }
};
</script>

<style scoped>
.farmed-number {
  display: inline-block;
}
</style>
