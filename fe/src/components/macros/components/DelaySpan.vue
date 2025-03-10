<template>
  <span :class="`delay ${active ? 'active' : ''}`">
    {{ value < 10000 ? value + 'ms' : '>10s' }}
  </span>
</template>

<script setup>
defineProps({
  value: Number,
  active: Boolean,
})
</script>

<style scoped>
@reference "@/assets/main.css";

span.delay {
  @apply flex 
  items-center
  px-2 py-1
  bg-slate-500
  border
  border-slate-400
  text-slate-950
  font-semibold
  rounded-sm
  text-sm
  cursor-default;
}

.edit span.delay {
  @apply cursor-pointer;

  &:hover,
  &.active {
    @apply bg-lime-700 border-lime-500 text-lime-200;
  }
}
</style>
