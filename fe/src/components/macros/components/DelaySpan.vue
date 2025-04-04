<template>
  <span :class="`delay ${active ? 'active' : ''} ${preset ? 'preset' : ''}`">
    <template v-if="value < 10000"> {{ value }} <i>ms</i> </template>
    <template v-else> >10 <i>s</i> </template>
  </span>
</template>

<script setup>
import { IconTimeDuration10 } from '@tabler/icons-vue'

defineProps({
  value: Number,
  active: Boolean,
  preset: Boolean,
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
  font-sans
  font-semibold
  rounded-sm
  text-sm
  cursor-default;

  &.preset {
    @apply text-amber-400
    border-amber-300/80
    bg-amber-100/60;
  }

  i {
    @apply pl-1
    font-normal 
    not-italic 
    opacity-80;
  }
}

.edit span.delay {
  @apply cursor-pointer;

  &:hover,
  &.active {
    @apply bg-lime-700 border-lime-500 text-lime-200;
  }
}
</style>
