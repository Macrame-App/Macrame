<template>
  <div :class="`mcrm-toast mcrm-block block__${toastOptions.variant}`" ref="toast">
    <ButtonComp v-if="closable" variant="subtle" size="sm" @click="closeToast()">
      <IconX />
    </ButtonComp>
    <h4>{{ title }}</h4>
    <p>{{ message }}</p>
  </div>
</template>

<script setup>
import { useNoticationStore } from '@/stores/notifications'
import { onMounted, reactive, ref } from 'vue'
import ButtonComp from './ButtonComp.vue'
import { IconX } from '@tabler/icons-vue'

const props = defineProps({
  title: String,
  message: String,
  variant: String,
  time: Number,
  closable: Boolean,
  notification: [String, Number],
})

const notifications = useNoticationStore()

const toast = ref(null)

const toastOptions = reactive({
  variant: props.variant,
})

onMounted(() => {
  if (toastOptions.variant == 'info') toastOptions.variant = 'primary'

  setTimeout(() => {
    closeToast()
  }, props.time)
})

const closeToast = () => {
  toast.value.classList.add('closing')

  setTimeout(() => {
    toast.value.remove()
    if (props.notification) notifications.remove(props.notification)
  }, 500)
}
</script>

<style scoped>
@reference "@/assets/main.css";

.mcrm-toast {
  @apply relative 
  grid
  gap-2
  p-4
  transition-opacity
  duration-400;

  h4 {
    @apply pr-6
    text-base;
  }

  p {
    @apply text-sm opacity-80;
  }

  &.closing {
    @apply opacity-0;
  }

  button.btn {
    @apply absolute
    top-2 right-2
    p-2;
  }
}
</style>
