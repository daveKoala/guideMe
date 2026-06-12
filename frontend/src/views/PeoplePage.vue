<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import Button from 'primevue/button'
import PersonList from '@/components/people/PersonList.vue'
import PersonDetail from '@/components/people/PersonDetail.vue'
import { useAccountStore } from '@/stores/account'

const store = useAccountStore()

const selectedId = ref<string | null>(null)

const selected = computed(() => (selectedId.value ? store.getPerson(selectedId.value) : undefined))

onMounted(async () => {
  await store.fetchPeople()
  const first = store.people[0]
  if (!selectedId.value && first) selectedId.value = first.id
})

function reset() {
  store.reset()
  selectedId.value = store.people[0]?.id ?? null
}
</script>

<template>
  <section class="people-page">
    <header class="people-page__head">
      <h2>People</h2>
      <div class="people-page__controls">
        <Button label="Reset demo data" icon="pi pi-refresh" text size="small" @click="reset" />
      </div>
    </header>

    <p class="hint">Everyone in your account. Select a person to view their details and documents.</p>

    <p v-if="store.loading" class="status">Loading people…</p>

    <div v-else class="directory">
      <PersonList
        class="directory__list"
        :people="store.people"
        :selected-id="selectedId"
        @select="selectedId = $event"
      />
      <PersonDetail v-if="selected" :key="selected.id" class="directory__detail" :person="selected" />
      <p v-else class="status">No people yet.</p>
    </div>
  </section>
</template>

<style scoped>
.people-page__head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 1rem;
  flex-wrap: wrap;
  margin-bottom: 0.5rem;
}

.people-page__controls {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  flex-wrap: wrap;
}

.hint {
  opacity: 0.7;
  font-size: 0.9rem;
  margin-bottom: 1.25rem;
}

.status {
  opacity: 0.6;
}

.directory {
  display: grid;
  grid-template-columns: minmax(220px, 1fr) 2fr;
  gap: 1.25rem;
  align-items: start;
}

/* Stack on hand-helds: list above detail, single column */
@media (max-width: 720px) {
  .directory {
    grid-template-columns: 1fr;
  }
}
</style>
