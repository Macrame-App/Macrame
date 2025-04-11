<template>
  <div id="panels-overview">
    <AlertComp v-if="Object.keys(panels.list).length == 0" variant="info">
      No panels found
    </AlertComp>
    <div class="panel-list">
      <div class="panel-item mcrm-block block__dark" v-for="(panel, i) in panels.list" :key="i">
        <div class="panel-item__content" @click="panelItemClick(panel.dir)">
          <div class="thumb">
            <img v-if="panel.thumb" :src="`data:image/jpeg;base64,${panel.thumb}`" alt="" />
            <IconLayoutGrid v-else />
          </div>
          <h4>{{ panel.name }}</h4>
          <div class="description" v-if="isLocal()">
            <div class="content">
              <strong class="block mb-1 text-slate-400">{{ panel.name }}</strong>
              <hr class="mb-2 border-slate-600" />
              <p v-if="panel.description != 'null'" class="text-slate-200">
                {{ panel.description }}
              </p>
            </div>
            <footer>
              <ButtonComp variant="subtle" size="sm" :href="`/panel/view/${panel.dir}`">
                <IconEye /> Preview
              </ButtonComp>
              <ButtonComp variant="primary" size="sm" :href="`/panel/edit/${panel.dir}`">
                <IconPencil /> Edit
              </ButtonComp>
            </footer>
          </div>
        </div>
        <template v-if="!isLocal()"> </template>
      </div>
    </div>
  </div>
</template>

<script setup>
import { usePanelStore } from '@/stores/panel'
import { onMounted, reactive } from 'vue'
import AlertComp from '../base/AlertComp.vue'
import { IconEye, IconLayoutGrid, IconPencil } from '@tabler/icons-vue'
import ButtonComp from '../base/ButtonComp.vue'
import { isLocal } from '@/services/ApiService'
import { useRouter } from 'vue-router'

const panel = usePanelStore()

const panels = reactive({
  list: {},
})

const router = useRouter()

onMounted(async () => {
  const panelList = await panel.getList()
  // console.log(panelList)

  panels.list = panelList
})

function panelItemClick(dir) {
  if (isLocal()) return

  router.push(`/panel/view/${dir}`)
}
</script>

<style scoped>
@reference "@/assets/main.css";

.panel-list {
  @apply grid
  grid-cols-2
  md:grid-cols-4
  lg:grid-cols-6
  gap-4
  w-full h-fit;
}

.panel-item {
  @apply p-px 
  overflow-hidden;

  .thumb {
    @apply flex 
    justify-center
    items-center
    w-full
    aspect-[4/3];

    img {
      @apply size-full
      object-cover;
    }

    &:not(:has(img)) {
      @apply bg-sky-950;
    }

    svg {
      @apply size-12;
    }
  }

  h4 {
    @apply px-4 py-2
    h-12
    truncate;
  }

  &:hover .description {
    @apply opacity-100;
  }

  .description {
    @apply absolute
    inset-0
    size-full
    pt-2
    pr-1
    pb-13
    bg-slate-900/60
    backdrop-blur-md
    text-slate-100
    opacity-0
    transition-opacity
    cursor-default
    z-10;

    .content {
      @apply h-full
      p-4
      pt-2
      overflow-y-auto;
    }

    footer {
      @apply absolute
      bottom-0 left-0
      w-full
      h-12
      grid
      grid-cols-2
      bg-slate-900
      border-t
      border-slate-600;

      .btn {
        @apply size-full
        rounded-none
        justify-center
        border-0;

        &:last-child {
          @apply border-l
          border-slate-600;
        }
      }
    }
  }
}
</style>
