<script setup lang="ts">
import { inject, ref } from 'vue'
import Panel from 'primevue/panel'
import PartyPeoplePicker from '@/components/people/PartyPeoplePicker.vue'
import PersonDocumentSlot from '@/components/people/PersonDocumentSlot.vue'
import { modeKey, usePartyPeople, type Mode } from '@/components/timeline/tripContext'
import type { Party } from '@/types/trip'

const props = defineProps<{ party: Party }>()

const mode = inject(modeKey, ref<Mode>('read'))

// Start collapsed; user expands on demand.
const collapsed = ref(true)

const people = usePartyPeople(props.party)
</script>

<template>
  <Panel toggleable v-model:collapsed="collapsed">
    <template #header>
      <span class="ov-title"><span class="ov-icon">👥</span> People</span>
    </template>

    <PartyPeoplePicker v-if="mode === 'edit'" v-model="party.passengers" />

    <template v-else>
      <ul class="people">
        <li v-for="p in people" :key="p.id" class="person">
          <div class="person__head">
            <span class="person__name-text">{{ p.name || 'Unnamed' }}</span>
            <span class="person__type-text">{{ p.type }}</span>
          </div>
          <div class="person__docs">
            <PersonDocumentSlot :person-id="p.id" kind="passport" label="Passport" />
          </div>
        </li>
      </ul>

      <p v-if="!people.length" class="empty">No people yet.</p>
    </template>
  </Panel>
</template>

<style scoped>
.people {
  list-style: none;
  margin: 0;
  padding: 0;
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.person {
  border: 1px solid var(--color-border);
  border-radius: var(--radius-md);
  padding: 0.75rem 1rem;
}

.person__head {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  margin-bottom: 0.5rem;
}

.person__name-text {
  font-weight: 600;
  flex: 1;
}

.person__type-text {
  opacity: 0.6;
  text-transform: capitalize;
}

.person__docs {
  display: flex;
  flex-direction: column;
  gap: 0.4rem;
}

.empty {
  opacity: 0.6;
}

.ov-icon {
  margin-right: 0.4rem;
}
</style>
