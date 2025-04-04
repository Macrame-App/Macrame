<template>
  <div class="dialog-container">
    <div class="trigger" @click="toggleDialog(true)">
      <slot name="trigger" />
    </div>
    <dialog ref="dialog" class="mcrm-block block__dark">
      <ButtonComp
        class="dialog__close p-0"
        variant="ghost"
        size="sm"
        tabindex="-1"
        @click="toggleDialog(false)"
      >
        <IconX />
      </ButtonComp>
      <slot name="content" />
    </dialog>
  </div>
</template>

<script setup>
import ButtonComp from './ButtonComp.vue'
import { IconX } from '@tabler/icons-vue'
import { onMounted, onUpdated, ref } from 'vue'

const dialog = ref(null)
const openDialog = ref()

const emit = defineEmits(['onOpen', 'onClose', 'onToggle'])

defineExpose({ toggleDialog })

const props = defineProps({
  open: Boolean,
})

onMounted(() => {
  if (props.open === true) toggleDialog(props.open)
})

onUpdated(() => {
  if (props.open === true) toggleDialog(props.open)
})

function toggleDialog(openToggle) {
  if (openToggle) {
    dialog.value.showModal()
    emit('onOpen')
  } else {
    dialog.value.close()
    emit('onClose')
  }

  openDialog.value = openToggle
  emit('onToggle')
}

onMounted(() => {
  openDialog.value = props.open

  if (dialog.value.innerHTML.includes('form')) {
    dialog.value.querySelector('form').addEventListener('submit', () => {
      toggleDialog()
    })
  }
})
</script>

<style>
@reference "@/assets/main.css";

.dialog-container {
  @apply relative;

  dialog {
    @apply fixed
    top-1/2 left-1/2
    -translate-x-1/2 -translate-y-1/2
    max-w-[calc(100vw-2rem)]
    text-slate-200
    /* shadow-md */
    /* shadow-black  */
    z-50
    pointer-events-none;

    &[open] {
      @apply pointer-events-auto;
    }

    &::backdrop {
      @apply bg-black/50 backdrop-blur-xs transition;
    }

    .dialog__close {
      @apply absolute
      top-4 right-4
      p-0
      text-white;

      svg {
        @apply size-5;
      }
    }
  }
}
.dialog__content {
  > *:first-child {
    @apply pr-8;
  }
}
</style>
