<script setup lang="ts">
import { inject, ref } from 'vue'
import Panel from 'primevue/panel'
import Button from 'primevue/button'
import InputText from 'primevue/inputtext'
import Select from 'primevue/select'
import DocumentSlot from './DocumentSlot.vue'
import { modeKey, type Mode } from '@/components/timeline/tripContext'
import type { Party, PassengerType } from '@/types/trip'

const props = defineProps<{ party: Party }>()

const mode = inject(modeKey, ref<Mode>('read'))

const typeOptions: { label: string; value: PassengerType }[] = [
  { label: 'Adult', value: 'adult' },
  { label: 'Child', value: 'child' },
  { label: 'Infant', value: 'infant' },
]

function addPerson() {
  props.party.passengers.push({
    id: `passenger_${crypto.randomUUID()}`,
    name: '',
    type: 'adult',
  })
}
function removePerson(id: string) {
  props.party.passengers = props.party.passengers.filter((p) => p.id !== id)
}
</script>

<template>
  <Panel toggleable>
    <template #header>
      <span class="ov-title"><span class="ov-icon">👥</span> People</span>
    </template>
    <template #icons>
      <Button
        v-if="mode === 'edit'"
        label="Add"
        icon="pi pi-plus"
        size="small"
        @click="addPerson"
      />
    </template>

    <ul class="people">
      <li v-for="p in party.passengers" :key="p.id" class="person">
        <div class="person__head">
          <template v-if="mode === 'edit'">
            <InputText v-model="p.name" placeholder="Name" class="person__name" />
            <Select
              v-model="p.type"
              :options="typeOptions"
              option-label="label"
              option-value="value"
              class="person__type"
            />
            <Button
              icon="pi pi-trash"
              size="small"
              text
              severity="danger"
              aria-label="Remove person"
              @click="removePerson(p.id)"
            />
          </template>
          <template v-else>
            <span class="person__name-text">{{ p.name || 'Unnamed' }}</span>
            <span class="person__type-text">{{ p.type }}</span>
          </template>
        </div>

        <div class="person__docs">
          <DocumentSlot label="Passport" />
        </div>
      </li>
    </ul>

    <p v-if="!party.passengers.length" class="empty">No people yet.</p>
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

.person__name {
  flex: 1;
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
