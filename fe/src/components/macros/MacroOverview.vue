<template>
  <div class="macro-overview mcrm-block block__dark">
    <h4 class="border-b-2 border-transparent">Saved Macros</h4>
    <div class="macro-overview__list">
      <div class="macro-item" v-for="(macro, i) in macros.list" :key="i">
        <ButtonComp variant="dark" class="w-full" size="sm" @click.prevent="runMacro(macro)">
          <IconKeyboard /> {{ macro }}
        </ButtonComp>
      </div>
    </div>
  </div>
</template>

<script setup>
import { IconKeyboard } from '@tabler/icons-vue'
import ButtonComp from '../base/ButtonComp.vue'
import { onMounted, reactive } from 'vue'
import axios from 'axios'
import { appUrl } from '@/services/ApiService'

const macros = reactive({
  list: [],
})

onMounted(() => {
  axios.post(appUrl() + '/macro/list').then((data) => {
    if (data.data.length > 0) macros.list = data.data
  })
})

function runMacro(macro) {
  console.log(macro)
  axios.post(appUrl() + '/macro/play', { macro: macro }).then((data) => {
    console.log(data)
  })
}
</script>

<style scoped>
@reference "@/assets/main.css";

.macro-overview {
  @apply relative
  grid
  grid-rows-[auto_1fr];

  &::after {
    @apply content-['']
    absolute
    top-0
    left-full
    h-full
    w-px
    bg-slate-600;
  }

  .macro-overview__list {
    @apply grid 
    gap-1
    content-start;
  }

  .macro-item {
    @apply flex items-center;

    button {
      @apply w-full;
    }
  }
}
</style>
