<template>
  <template v-if="href">
    <RouterLink :to="href" :class="classString">
      <slot />
    </RouterLink>
  </template>
  <template v-else>
    <button :class="classString">
      <slot />
    </button>
  </template>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  href: String,
  variant: String,
  size: String,
})

const classString = computed(() => {
  let classes = 'btn'
  if (props.variant) classes += ` btn__${props.variant}`
  if (props.size) classes += ` btn__${props.size}`

  return classes
})
</script>

<style>
@reference "@/assets/main.css";

button,
.btn {
  @apply flex 
  items-center 
  gap-3
  h-fit 
  px-4 py-2
  border
  border-solid
  rounded-lg
  tracking-wide
  font-normal
  transition-all
  cursor-pointer;

  transition:
    border-color 0.1s ease-in-out,
    background-color 0.2s ease;

  &:not(.button__subtle, .button__ghost):hover {
    @apply shadow-black;
  }

  &[disabled],
  &.disabled {
    @apply opacity-50 pointer-events-none cursor-not-allowed;
  }

  svg {
    @apply size-5 transition-[stroke] duration-400 ease-in-out;
  }

  &.btn__sm {
    @apply px-3 py-1
    text-sm;

    svg {
      @apply size-4;
    }
  }

  &.btn__lg {
    @apply px-6 py-3
    text-lg;

    svg {
      @apply size-6;
    }
  }

  &:hover {
    @apply text-white;

    svg {
      @apply stroke-current;
    }
  }

  &.btn__primary {
    @apply bg-sky-100/10 border-sky-100 text-sky-100;

    svg {
      @apply stroke-sky-200;
    }

    &:hover {
      @apply bg-sky-400/40 border-sky-300;
    }
  }

  &.btn__secondary {
    @apply bg-amber-100/10 border-amber-100 text-amber-100;

    svg {
      @apply stroke-amber-300;
    }

    &:hover {
      @apply bg-amber-400/40 border-amber-400;
    }
  }

  &.btn__danger {
    @apply bg-rose-200/20 border-rose-100 text-rose-200;

    svg {
      @apply stroke-rose-400;
    }

    &:hover {
      @apply bg-rose-400/40 border-rose-500 text-white;
    }
  }

  &.btn__dark {
    /* @apply bg-slate-700/80 hover:bg-slate-700 text-white border-slate-600; */
    @apply bg-slate-200/10 border-slate-400 text-slate-100;

    svg {
      @apply stroke-slate-300;
    }

    &:hover {
      @apply bg-slate-400/40 border-slate-200 text-white;
    }
  }

  &.btn__success {
    /* @apply bg-lime-500/80 hover:bg-lime-500 text-white border-lime-600; */
    @apply bg-lime-200/10 border-lime-100 text-lime-100;

    svg {
      @apply stroke-lime-400;
    }

    &:hover {
      @apply bg-lime-400/40 border-lime-500 text-white;
    }
  }

  &.btn__subtle {
    @apply bg-transparent hover:bg-white/10 text-white border-transparent;

    &:hover {
      @apply bg-white/20 to-white/30 border-white/40;
    }
  }

  &.btn__ghost {
    @apply bg-transparent text-white/80 border-transparent hover:text-white;
  }
}
</style>
