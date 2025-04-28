<template>
  <div class="macro-overview mcrm-block block__dark">
    <h4 class="border-b-2 border-transparent">Saved Macros</h4>
    <div class="macro-overview__list">
      <LoadComp :loading="macros.loading" text="Loading macros..." />
      <div class="macro-item" v-for="(macro, i) in macros.list" :key="i">
        <ButtonComp variant="dark" class="w-full" size="sm">
          <IconKeyboard /> {{ macro.name }}
        </ButtonComp>
      </div>
    </div>
  </div>
</template>

<script setup>
// TODO
// - load macro on click
// - delete macro

import { IconKeyboard } from '@tabler/icons-vue'
import ButtonComp from '../base/ButtonComp.vue'
import { onMounted, reactive } from 'vue'
import axios from 'axios'
import { appUrl, isLocal } from '@/services/ApiService'
import { AuthCall } from '@/services/EncryptService'
import { GetMacroList, RunMacro } from '@/services/MacroService'
import LoadComp from '../base/LoadComp.vue'

const macros = reactive({
  loading: true,
  list: [],
})

onMounted(() => {
  loadMacroList()
})

const loadMacroList = async () => {
  const list = await GetMacroList()
  macros.list = list
  macros.loading = false
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
    @apply flex
    flex-col 
    pr-1
    gap-1
    h-[calc(100vh-11.7rem)]
    overflow-auto;
  }

  .macro-item {
    @apply flex items-center;

    button {
      @apply w-full;
    }
  }
}
</style>
