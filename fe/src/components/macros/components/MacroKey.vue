<template>
  <kbd :class="active ? 'active' : ''">
    <sup v-if="keyObj.loc">
      {{ keyObj.loc }}
    </sup>
    <span :innerHTML="keyObj.str" />
    <span class="dir">{{ dir.value === 'down' ? '&darr;' : '&uarr;' }}</span>
  </kbd>
</template>

<script setup>
import { onMounted, reactive } from 'vue'

const props = defineProps({
  keyObj: Object,
  direction: String,
  active: Boolean,
})

const dir = reactive({
  value: false,
})

onMounted(() => {
  if (props.direction) dir.value = props.direction
  else dir.value = props.keyObj.direction
})
</script>

<style>
@reference "@/assets/main.css";

kbd {
  @apply flex 
  items-center
  gap-2
  pl-4 pr-2 py-1 
  h-9 
  bg-slate-700
  font-sans
  font-bold
  text-lg
  text-white
  uppercase
  rounded-md 
  border
  border-slate-500
  transition-all
  shadow-slate-500;
  box-shadow: 0 0.2rem 0 0.2rem var(--tw-shadow-color);

  &:has(sup) {
    @apply pl-2;
  }

  sup {
    @apply text-slate-200 text-xs font-light mt-1;
  }

  span.dir {
    @apply text-slate-200 pl-1;
  }
}

:has(kdb):not(.edit) kbd {
  @apply pointer-events-none cursor-default;
}

.edit kbd {
  @apply cursor-pointer pointer-events-auto;

  &:hover,
  &.active {
    @apply bg-sky-900 border-sky-400 shadow-sky-700;
  }
}
</style>
