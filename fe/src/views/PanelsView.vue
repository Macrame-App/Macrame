<template>
  <div id="panels" class="panel">
    <h1 class="flex items-end justify-between !w-full panel__title">
      <div>
        Panels
        <span class="text-sm">{{ isLocal() ? 'remote' : 'servers' }}</span>
      </div>
      <ButtonComp
        v-if="panel.function != 'overview'"
        variant="subtle"
        size="sm"
        @click="router.push('/panels')"
      >
        <IconArrowLeft /> Overview
      </ButtonComp>
    </h1>
    <div :class="`panel__content !p-0 !pt-4 ${panel.function == 'overview' ?? '!pr-4'}`">
      <PanelsOverview v-if="panel.function == 'overview'" />
      <PanelEdit v-if="panel.function == 'edit'" :dirname="panel.dirname" />
      <PanelView v-if="panel.function == 'preview'" :dirname="panel.dirname" />
    </div>
  </div>
</template>

<script setup>
import ButtonComp from '@/components/base/ButtonComp.vue'
import PanelEdit from '@/components/panels/PanelEdit.vue'
import PanelView from '@/components/panels/PanelView.vue'
import PanelsOverview from '@/components/panels/PanelsOverview.vue'
import { isLocal } from '@/services/ApiService'
import { IconArrowLeft } from '@tabler/icons-vue'
import { onMounted, onUpdated, reactive } from 'vue'
import { useRoute, useRouter } from 'vue-router'

const route = useRoute()
const router = useRouter()

const panel = reactive({
  function: '',
  dirname: '',
})

onMounted(() => {
  setVarsByRoute()
})

onUpdated(() => {
  setVarsByRoute()
})

const setVarsByRoute = () => {
  if (route.name.includes('panel-')) {
    panel.function = route.name == 'panel-edit' ? 'edit' : 'preview'
  } else {
    panel.function = 'overview'
  }

  panel.dirname = route.params.dirname
}
</script>

<style scoped>
@reference "@/assets/main.css";
</style>
