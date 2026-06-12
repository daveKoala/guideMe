<script setup lang="ts">
import { computed, ref } from 'vue'
import Button from 'primevue/button'
import MultiSelect from 'primevue/multiselect'
import InputText from 'primevue/inputtext'
import Select from 'primevue/select'
import PersonDocumentSlot from './PersonDocumentSlot.vue'
import { useAccountStore } from '@/stores/account'
import type { Person } from '@/types/person'
import type { PassengerType } from '@/types/trip'

const props = defineProps<{ modelValue: string[] }>()
const emit = defineEmits<{ 'update:modelValue': [string[]] }>()

const store = useAccountStore()

const typeOptions: { label: string; value: PassengerType }[] = [
  { label: 'Adult', value: 'adult' },
  { label: 'Child', value: 'child' },
  { label: 'Infant', value: 'infant' },
]

/** Resolved account People for the ids currently on the party. */
const people = computed<Person[]>(() =>
  props.modelValue.map((id) => store.getPerson(id)).filter((p): p is Person => !!p),
)

/** Account people not yet on the party — candidates for "add existing". */
const available = computed<Person[]>(() =>
  store.people.filter((p) => !props.modelValue.includes(p.id)),
)

// Buffer for the add-existing MultiSelect before the user commits with "Add".
const toAdd = ref<string[]>([])

// Track ids created in this picker so we can reveal inline name/type editors for them.
const justCreated = ref<Set<string>>(new Set())

function addExisting() {
  if (!toAdd.value.length) return
  emit('update:modelValue', [...props.modelValue, ...toAdd.value])
  toAdd.value = []
}

function addNew() {
  const p = store.addPerson() // real account Person — also appears on /people
  justCreated.value.add(p.id)
  emit('update:modelValue', [...props.modelValue, p.id])
}

function remove(id: string) {
  // Unlink from the trip only; the account Person is kept.
  emit('update:modelValue', props.modelValue.filter((x) => x !== id))
  justCreated.value.delete(id)
}
</script>

<template>
  <div class="picker">
    <ul v-if="people.length" class="people">
      <li v-for="p in people" :key="p.id" class="person">
        <div class="person__head">
          <template v-if="justCreated.has(p.id)">
            <InputText :model-value="p.name" placeholder="Name" class="person__name" @update:model-value="p.name = $event ?? ''" />
            <Select
              :model-value="p.type"
              :options="typeOptions"
              option-label="label"
              option-value="value"
              class="person__type"
              @update:model-value="p.type = $event"
            />
          </template>
          <template v-else>
            <span class="person__name-text">{{ p.name || 'Unnamed' }}</span>
            <span class="person__type-text">{{ p.type }}</span>
          </template>
          <Button
            icon="pi pi-trash"
            size="small"
            text
            severity="danger"
            aria-label="Remove from trip"
            @click="remove(p.id)"
          />
        </div>

        <div class="person__docs">
          <PersonDocumentSlot :person-id="p.id" kind="passport" label="Passport" />
        </div>
      </li>
    </ul>

    <p v-else class="empty">No people on this trip yet.</p>

    <div class="add-row">
      <MultiSelect
        v-model="toAdd"
        :options="available"
        option-label="name"
        option-value="id"
        placeholder="Add existing people"
        display="chip"
        filter
        class="add-existing"
      />
      <Button label="Add" icon="pi pi-plus" size="small" :disabled="!toAdd.length" @click="addExisting" />
      <Button label="New person" icon="pi pi-user-plus" size="small" outlined @click="addNew" />
    </div>
  </div>
</template>

<style scoped>
.picker {
  display: flex;
  flex-direction: column;
  gap: 1rem;
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

.add-row {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  flex-wrap: wrap;
}

.add-existing {
  flex: 1;
  min-width: 200px;
}

.empty {
  opacity: 0.6;
}
</style>
