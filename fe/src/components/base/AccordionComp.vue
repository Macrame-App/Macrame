<!--
Macrame is a program that enables the user to create keyboard macros and button panels. 
The macros are saved as simple JSON files and can be linked to the button panels. The panels can 
be created with HTML and CSS.

Copyright (C) 2025 Jesse Malotaux

This program is free software: you can redistribute it and/or modify 
it under the terms of the GNU General Public License as published by 
the Free Software Foundation, either version 3 of the License, or 
(at your option) any later version.

This program is distributed in the hope that it will be useful, 
but WITHOUT ANY WARRANTY; without even the implied warranty of 
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the 
GNU General Public License for more details.

You should have received a copy of the GNU General Public License 
along with this program. If not, see <https://www.gnu.org/licenses/>.
-->

<template>
  <div class="accordion">
    <header @click="toggleAccordion(!accordionOpen)">
      <h4>{{ title }}</h4>
      <ButtonComp variant="ghost" size="sm" class="!px-1">
        <IconChevronDown v-if="!accordionOpen" />
        <IconChevronUp v-else />
      </ButtonComp>
    </header>
    <section :class="`accordion__wrapper ${accordionOpen ? 'open' : ''}`">
      <div class="accordion__content">
        <slot />
      </div>
    </section>
  </div>
</template>

<script setup>
import { onMounted, onUpdated, ref } from 'vue'
import ButtonComp from './ButtonComp.vue'
import { IconChevronDown, IconChevronUp } from '@tabler/icons-vue'

const emit = defineEmits(['onOpen', 'onClose', 'onToggle'])

defineExpose({ toggleAccordion })

const props = defineProps({
  title: String,
  open: Boolean,
})

const accordionOpen = ref(false)

onMounted(() => {
  if (props.open) toggleAccordion(props.open)
})

onUpdated(() => {
  if (props.open) toggleAccordion(props.open)
})

function toggleAccordion(open = false) {
  if (open) {
    accordionOpen.value = true
    emit('onOpen')
  } else {
    accordionOpen.value = false
    emit('onClose')
  }

  emit('onToggle')
}
</script>

<style scoped>
@reference "@/assets/main.css";

.accordion {
  @apply grid;

  header {
    @apply grid
    grid-cols-[1fr_auto]
    px-4 py-2
    cursor-pointer;
  }

  .accordion__wrapper {
    @apply grid
    grid-rows-[0fr]
    border-y
    border-b-white/60
    border-t-transparent
    duration-300
    ease-in-out;

    .accordion__content {
      @apply grid
      grid-rows-[0fr]
      overflow-hidden
      opacity-0
      transition-opacity
      delay-0;
    }

    &.open {
      @apply grid-rows-[1fr]
      border-t-white/20;

      .accordion__content {
        @apply grid-rows-[1fr] 
        opacity-100
        delay-200;
      }
    }
  }
}
</style>
