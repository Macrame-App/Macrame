<template>
  <div id="panels-overview">
    <!-- <AlertComp v-if="Object.keys(panels.list).length == 0" type="info">No panels found</AlertComp> -->
    <div class="panel-list">
      <div class="panel-item" v-for="(panel, i) in panels.list" :key="i">
        <!-- <router-link :to="'/panel/' + panel.id"> -->
        <div class="panel-item__content">
          <img :src="panel.thumb" alt="" />
          <h4>{{ panel.name }}</h4>
          <p>{{ panel.description }}</p>
        </div>
        <!-- </router-link> -->
      </div>
    </div>
  </div>
</template>

<script setup>
import { usePanelStore } from '@/stores/panel'
import { onMounted, reactive } from 'vue'
import AlertComp from '../base/AlertComp.vue'

const panel = usePanelStore()

const panels = reactive({
  list: {},
})

onMounted(async () => {
  const panelList = await panel.getPanels()
  console.log(panelList)

  panels.list = panelList
})
</script>

<style scoped>
@reference "@/assets/main.css";

.panel-list {
  @apply grid
  grid-cols-2
  md:grid-cols-4
  lg:grid-cols-6
  gap-4
  size-full;
}
</style>
