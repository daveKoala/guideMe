<script setup lang="ts">
import { ref, watch } from 'vue'
import StageCard from './StageCard.vue'
import type { Stage } from '@/types/stage'

const props = defineProps<{ stages: Stage[] }>()

const STORAGE_KEY = 'guide-me:stage-list:expanded'

// Restore last-open stage across reloads/remounts. Only if it still exists in this trip.
function restore(): string | null {
  try {
    const id = localStorage.getItem(STORAGE_KEY)
    return id && props.stages.some((s) => s.id === id) ? id : null
  } catch {
    return null
  }
}

const expandedId = ref<string | null>(restore())

watch(expandedId, (id) => {
  try {
    if (id) localStorage.setItem(STORAGE_KEY, id)
    else localStorage.removeItem(STORAGE_KEY)
  } catch {
    // storage unavailable — keep working in-memory
  }
})

function toggle(id: string) {
  expandedId.value = expandedId.value === id ? null : id // tap open card = close
}
</script>

<template>
  <ul class="stage-list">
    <li v-for="stage in stages" :key="stage.id">
      <StageCard
        :stage="stage"
        :expanded="expandedId === stage.id"
        @toggle="toggle(stage.id)"
      />
    </li>
  </ul>
</template>

<style scoped>
.stage-list {
  list-style: none;
  margin: 0 auto;
  padding: 0;
  display: flex;
  flex-direction: column;
  gap: 1rem;
  width: 100%;
  max-width: 640px;
}
</style>
