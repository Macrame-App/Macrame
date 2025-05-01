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
  <div class="input-group form-select">
    <label v-if="label">
      {{ label }}
    </label>
    <div class="select__container">
      <template v-if="search">
        <div class="select__search-bar">
          <input
            type="search"
            ref="selectSearch"
            :list="`${name}-search__options`"
            v-model="select.search"
            @change="selectSearchValue($event)"
            :disabled="!select.searchActive"
            autocomplete="on"
          />
          <datalist :id="`${name}-search__options`">
            <option v-for="option in select.options" :key="option.value" :value="option.value">
              {{ option.label }}
            </option>
          </datalist>
          <ButtonComp v-if="!select.searchActive" variant="ghost" size="sm" @click="initSearch">
            <IconSearch />
          </ButtonComp>
          <ButtonComp v-else variant="ghost" size="sm" @click="resetSearch">
            <IconSearchOff />
          </ButtonComp>
        </div>
      </template>
      <select :name="name" ref="selectEl" v-model="select.value" @change="changeSelect($event)">
        <option value="" disabled>- Select {{ label.toLocaleLowerCase() }} -</option>
        <option v-for="option in select.options" :key="option.value" :value="option.value">
          {{ option.label }}
        </option>
      </select>
    </div>
  </div>
</template>

<script setup>
import { IconSearch, IconSearchOff } from '@tabler/icons-vue'
import { onMounted, onUpdated, reactive, ref } from 'vue'
import ButtonComp from '../base/ButtonComp.vue'

const emit = defineEmits(['change'])

const props = defineProps({
  label: String,
  name: String,
  options: [Array, Object],
  search: Boolean,
  value: String,
})

const select = reactive({
  options: [],
  search: '',
  searchActive: false,
  changed: false,
  value: '',
})

const selectEl = ref(null)
const selectSearch = ref(null)

onMounted(() => {
  setValue()

  if (typeof props.options == 'object') select.options = Object.values(props.options)
})

onUpdated(() => {
  setValue()
})

const setValue = () => {
  if ((select.value == '' && props.value) || (!select.changed && props.value != select.value)) {
    select.value = props.value
  }

  select.changed = false
}

const initSearch = () => {
  select.searchActive = true
  select.search = ''
  selectEl.value.classList = 'search__is-active'
  setTimeout(() => {
    selectSearch.value.focus()
  }, 50)
}

const resetSearch = () => {
  select.search = ''
  select.searchActive = false
  selectEl.value.classList = ''
}

const selectSearchValue = (event) => {
  changeSelect(event)
  resetSearch()
}

const changeSelect = (event) => {
  select.changed = true
  select.value = event.target.value

  emit('change', select.value)
}
</script>

<style scoped>
@reference "@/assets/main.css";

.select__container {
  @apply relative
  h-8;

  select,
  .select__search-bar {
    @apply absolute top-0 h-8;
  }
}
.select__search-bar {
  @apply right-0 
  grid
  grid-cols-[1fr_auto]
  items-center
  w-full
  pr-4
  z-10
  pointer-events-none;

  button {
    @apply pointer-events-auto;
  }

  input {
    @apply border-0 bg-transparent pointer-events-auto px-2 py-0 focus:outline-0;

    &[disabled] {
      @apply pointer-events-none;
    }
  }
  datalist {
    @apply absolute
    top-full left-0;
  }
}

select.search__is-active {
  @apply text-transparent;
}
</style>
