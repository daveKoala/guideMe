<script setup lang="ts">
import type { Person } from '@/types/person'

defineProps<{ people: Person[]; selectedId: string | null }>()
defineEmits<{ select: [id: string] }>()

function docCount(p: Person): number {
  return Object.keys(p.documents).length
}
</script>

<template>
  <ul class="person-list">
    <li v-for="p in people" :key="p.id">
      <button
        type="button"
        class="row"
        :class="{ 'row--active': p.id === selectedId }"
        @click="$emit('select', p.id)"
      >
        <span class="row__name">{{ p.name || 'Unnamed' }}</span>
        <span class="row__type">{{ p.type }}</span>
        <span class="row__docs">{{ docCount(p) }} doc{{ docCount(p) === 1 ? '' : 's' }}</span>
      </button>
    </li>
  </ul>
</template>

<style scoped>
.person-list {
  list-style: none;
  margin: 0;
  padding: 0;
  display: flex;
  flex-direction: column;
  gap: 0.4rem;
}

.row {
  width: 100%;
  display: flex;
  align-items: center;
  gap: 0.6rem;
  padding: 0.6rem 0.85rem;
  border: 1px solid var(--color-border);
  border-radius: var(--radius-md);
  background: var(--color-surface);
  font: inherit;
  color: inherit;
  text-align: left;
  cursor: pointer;
}

.row--active {
  border-color: var(--color-link);
  box-shadow: 0 0 0 1px var(--color-link) inset;
}

.row__name {
  font-weight: 600;
  flex: 1;
}

.row__type {
  opacity: 0.6;
  text-transform: capitalize;
  font-size: 0.875rem;
}

.row__docs {
  opacity: 0.55;
  font-size: 0.8rem;
  white-space: nowrap;
}
</style>
