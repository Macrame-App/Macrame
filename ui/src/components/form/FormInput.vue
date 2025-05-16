<template>
  <div :class="`form-input ${horizontal ? 'horizontal' : ''}`">
    <template v-if="label">
      <label :for="name">
        {{ label }}
      </label>
    </template>
    <template v-if="type != 'radio' && type != 'checkbox'">
      <input
        :type="type"
        :name="name"
        :id="name"
        :value="value"
        :placeholder="placeholder"
        :disabled="disabled"
        :readonly="readonly"
        :required="required"
        :autofocus="autofocus"
        @change="$emit('onChange', $event.target.value)"
        @input="$emit('onInput', $event.target.value)"
      />
    </template>
    <template v-else-if="options">
      <div class="boolean-group">
        <template v-for="option in options" :key="option.value">
          <label :for="`${name}-${option.value}`">
            {{ option.label }}
          </label>
          <input
            :type="type"
            :name="name"
            :id="`${name}-${option.value}`"
            :disabled="disabled"
            :readonly="readonly"
            :required="required"
            :checked="value == option.value"
            :value="option.value"
            @change="$emit('onChange', $event.target.value)"
            @input="$emit('onInput', $event.target.value)"
          />
        </template>
      </div>
    </template>
    <template v-else>
      <input
        :type="type"
        :name="name"
        :id="name"
        :value="value"
        :disabled="disabled"
        :readonly="readonly"
        :required="required"
        :checked="value === true || value === 'true'"
        @change="$emit('onChange', $event.target)"
        @input="$emit('onInput', $event.target)"
      />
    </template>
  </div>
</template>

<script setup>
defineProps({
  type: String,
  name: String,
  value: [String, Number, Boolean],
  options: Array,
  label: String,
  placeholder: String,
  disabled: Boolean,
  readonly: Boolean,
  required: Boolean,
  autofocus: Boolean,
  horizontal: Boolean,
})

defineEmits(['onChange', 'onInput'])
</script>

<style scoped>
@reference "@/assets/main.css";

.form-input {
  @apply flex
  flex-col
  items-center
  gap-2;

  &.horizontal {
    @apply flex-row 
    justify-between;

    input {
      @apply max-w-2/3;
    }
  }
}

.boolean-group {
  @apply flex
  items-center
  gap-2;
}

input[type='checkbox'],
input[type='radio'] {
  @apply size-5
  cursor-pointer;
}

input[disabled] {
  @apply opacity-60
  cursor-not-allowed;
}
</style>
