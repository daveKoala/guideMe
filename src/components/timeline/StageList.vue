<script setup lang="ts">
import { ref, watch } from 'vue'
import StageCard from './StageCard.vue'
import type { Stage } from '@/types/stage'

const props = defineProps<{ stages: Stage[] }>()

const STORAGE_KEY = 'guide-me:stage-list:expanded'

// Restore which stages were open across reloads/remounts. Each card is independent,
// so any number (including none) can be expanded. Default: all minimised.
function restore(): Set<string> {
  try {
    const raw = localStorage.getItem(STORAGE_KEY)
    const ids: unknown = raw ? JSON.parse(raw) : null
    if (!Array.isArray(ids)) return new Set()
    return new Set(ids.filter((id) => props.stages.some((s) => s.id === id)))
  } catch {
    return new Set()
  }
}

const expanded = ref<Set<string>>(restore())

watch(
  expanded,
  (set) => {
    try {
      localStorage.setItem(STORAGE_KEY, JSON.stringify([...set]))
    } catch {
      // storage unavailable — keep working in-memory
    }
  },
  { deep: true },
)

function toggle(id: string) {
  const next = new Set(expanded.value)
  if (next.has(id)) next.delete(id) // tap open card = close
  else next.add(id)
  expanded.value = next
}
</script>

<template>
  <ul class="stage-list">
    <li v-for="stage in stages" :key="stage.id">
      <StageCard
        :stage="stage"
        :expanded="expanded.has(stage.id)"
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
