<template>
  <template v-if="href">
    <a :href="href" :class="`button ${classString}`">
      <slot />
    </a>
  </template>
  <template v-else>
    <button :class="`button ${classString}`">
      <slot />
    </button>
  </template>
</template>

<script setup>
import { computed, onMounted } from 'vue'

const props = defineProps({
  href: String,
  variant: String,
  size: String,
})

const classString = computed(() => {
  const classes = {
    'bg-sky-500/80 hover:bg-sky-400 text-white border-sky-400': props.variant === 'primary',
    'bg-white/80 hover:bg-white text-slate-900 border-white': props.variant === 'secondary',
    'bg-red-700/80 hover:bg-red-700 text-white border-red-800': props.variant === 'danger',
    'bg-lime-500/80 hover:bg-lime-500 text-white border-lime-600': props.variant === 'success',
    'button__subtle bg-transparent hover:bg-white/10 text-white border-transparent':
      props.variant === 'subtle',
    'button__ghost bg-transparent text-white/80 border-transparent hover:text-white':
      props.variant === 'ghost',
    'button__lg px-5 py-3 text-lg gap-4': props.size === 'lg',
    'button__sm px-3 py-1.5 text-sm gap-2': props.size === 'sm',
    'px-4 py-2 gap-3': props.size !== 'sm' && props.size !== 'lg',
  }
  return Object.keys(classes)
    .filter((key) => classes[key])
    .join(' ')
})

onMounted(() => {
  console.log(classString)
})
</script>

<style>
@reference "@/assets/main.css";

button,
.button {
  @apply flex 
  items-center 
  h-fit 
  border
  rounded-lg
  font-semibold
  shadow-md
  shadow-transparent
  transition
  cursor-pointer;

  &:not(.button__subtle, .button__ghost):hover {
    @apply shadow-black;
  }

  svg {
    @apply size-5 stroke-1;
  }

  &.button__sm svg {
    @apply size-4;
  }

  &.button__lg svg {
    @apply size-6;
  }
}
</style>
