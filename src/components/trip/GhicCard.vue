<script setup lang="ts">
import { inject, ref } from 'vue'
import Panel from 'primevue/panel'
import InputText from 'primevue/inputtext'
import PersonDocumentSlot from '@/components/people/PersonDocumentSlot.vue'
import { modeKey, partyKey, emptyParty, usePartyPeople, type Mode } from '@/components/timeline/tripContext'

const mode = inject(modeKey, ref<Mode>('read'))
const party = inject(partyKey, emptyParty)

// Start collapsed; user expands on demand.
const collapsed = ref(true)

// Resolve party ids to account People; ghic_id is edited on the person so it follows them.
const people = usePartyPeople(party)
</script>

<template>
  <Panel toggleable v-model:collapsed="collapsed">
    <template #header>
      <span class="ov-title"><span class="ov-icon">🩺</span> GHIC</span>
    </template>

    <p class="hint">
      UK Global Health Insurance Card — necessary state healthcare in the EEA and some other
      countries. One per person.
    </p>

    <ul class="people">
      <li v-for="p in people" :key="p.id" class="person">
        <div class="person__head">
          <span class="person__name-text">{{ p.name || 'Unnamed' }}</span>
          <span class="person__type-text">{{ p.type }}</span>
        </div>

        <div class="field">
          <label :for="mode === 'edit' ? `${p.id}-ghic` : undefined">GHIC number</label>
          <InputText
            v-if="mode === 'edit'"
            :id="`${p.id}-ghic`"
            v-model="p.ghic_id"
            placeholder="UK Global Health Insurance Card"
            fluid
          />
          <span v-else class="value">{{ p.ghic_id || '—' }}</span>
        </div>

        <PersonDocumentSlot :person-id="p.id" kind="ghic_card" label="GHIC card" />
      </li>
    </ul>

    <p v-if="!people.length" class="empty">No people yet.</p>
  </Panel>
</template>

<style scoped>
.hint {
  opacity: 0.7;
  font-size: 0.875rem;
  margin-bottom: 1rem;
}

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
  display: flex;
  flex-direction: column;
  gap: 0.6rem;
}

.person__head {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.person__name-text {
  font-weight: 600;
  flex: 1;
}

.person__type-text {
  opacity: 0.6;
  text-transform: capitalize;
}

.field {
  display: flex;
  flex-direction: column;
  gap: 0.3rem;
}

.field label {
  font-weight: 600;
  font-size: 0.875rem;
}

.value {
  color: var(--color-text);
}

.empty {
  opacity: 0.6;
}

.ov-icon {
  margin-right: 0.4rem;
}
</style>
