<template>
  <div class="accordion">
    <header>
      <slot name="title" />
    </header>
    <section :class="`accordion__content ${open ? 'open' : ''}`">
      <div>
        <slot name="content" />
      </div>
    </section>
  </div>
</template>

<script setup>
defineProps({
  open: Boolean,
})
</script>

<style scoped>
@reference "@/assets/main.css";

.accordion {
  @apply grid;

  .accordion__content {
    @apply grid
    grid-rows-[0fr]
    overflow-hidden
    duration-300
    ease-in-out;

    div {
      @apply grid
      grid-rows-[0fr];
    }
  }
}
</style>
