<script setup lang="ts">
import { computed, inject, ref } from 'vue'
import Panel from 'primevue/panel'
import InputText from 'primevue/inputtext'
import Select from 'primevue/select'
import DatePicker from 'primevue/datepicker'
import PersonDocumentSlot from './PersonDocumentSlot.vue'
import { modeKey, type Mode } from '@/components/timeline/tripContext'
import type { Person } from '@/types/person'
import type { PassengerType } from '@/types/trip'

const props = defineProps<{ person: Person }>()

const mode = inject(modeKey, ref<Mode>('read'))

const typeOptions: { label: string; value: PassengerType }[] = [
  { label: 'Adult', value: 'adult' },
  { label: 'Child', value: 'child' },
  { label: 'Infant', value: 'infant' },
]

// dob stored as YYYY-MM-DD; convert for the DatePicker and back.
const dobDate = computed<Date | null>(() => (props.person.dob ? new Date(props.person.dob) : null))
function setDob(d: unknown) {
  props.person.dob = d instanceof Date ? d.toISOString().slice(0, 10) : ''
}
function fmtDob(iso?: string): string {
  if (!iso) return '—'
  const d = new Date(iso)
  return Number.isNaN(d.getTime())
    ? iso
    : new Intl.DateTimeFormat(undefined, { dateStyle: 'medium' }).format(d)
}
</script>

<template>
  <Panel class="person-detail">
    <template #header>
      <span class="ov-title">{{ person.name || 'Unnamed' }}</span>
    </template>

    <div class="fields">
      <div class="field">
        <label :for="`${person.id}-name`">Name</label>
        <InputText v-if="mode === 'edit'" :id="`${person.id}-name`" v-model="person.name" fluid />
        <span v-else class="value">{{ person.name || '—' }}</span>
      </div>

      <div class="field">
        <label :for="`${person.id}-type`">Type</label>
        <Select
          v-if="mode === 'edit'"
          :input-id="`${person.id}-type`"
          v-model="person.type"
          :options="typeOptions"
          option-label="label"
          option-value="value"
          fluid
        />
        <span v-else class="value" style="text-transform: capitalize">{{ person.type }}</span>
      </div>

      <div class="field">
        <label :for="`${person.id}-dob`">Date of birth</label>
        <DatePicker
          v-if="mode === 'edit'"
          :input-id="`${person.id}-dob`"
          :model-value="dobDate"
          date-format="dd M yy"
          fluid
          @update:model-value="setDob"
        />
        <span v-else class="value">{{ fmtDob(person.dob) }}</span>
      </div>

      <div class="field">
        <label :for="`${person.id}-ghic`">GHIC number</label>
        <InputText
          v-if="mode === 'edit'"
          :id="`${person.id}-ghic`"
          v-model="person.ghic_id"
          placeholder="UK Global Health Insurance Card"
          fluid
        />
        <span v-else class="value">{{ person.ghic_id || '—' }}</span>
      </div>
    </div>

    <fieldset class="docs">
      <legend>Documents</legend>
      <PersonDocumentSlot :person-id="person.id" kind="passport" label="Passport" />
      <PersonDocumentSlot :person-id="person.id" kind="ghic_card" label="GHIC card" />
    </fieldset>
  </Panel>
</template>

<style scoped>
.fields {
  display: flex;
  flex-direction: column;
  gap: 0.85rem;
  margin-bottom: 1rem;
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

.docs {
  border: 1px solid var(--color-border);
  border-radius: var(--radius-md);
  padding: 0.75rem 1rem 1rem;
  margin: 0;
  display: flex;
  flex-direction: column;
  gap: 0.6rem;
}

.docs legend {
  font-weight: 700;
  font-size: 0.8rem;
  text-transform: uppercase;
  letter-spacing: 0.04em;
  opacity: 0.7;
  padding: 0 0.4rem;
}
</style>
