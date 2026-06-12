<script setup lang="ts">
import { computed, onMounted, provide, reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import Button from 'primevue/button'
import InputText from 'primevue/inputtext'
import Select from 'primevue/select'
import PartyPeoplePicker from '@/components/people/PartyPeoplePicker.vue'
import { useTripsStore, type TripDraft } from '@/stores/trips'
import { useAccountStore } from '@/stores/account'
import { modeKey } from '@/components/timeline/tripContext'
import type { TripType, TripStatus } from '@/types/trip'

const router = useRouter()
const trips = useTripsStore()
const account = useAccountStore()

// The picker's doc slots need edit affordances.
provide(modeKey, ref<'edit' | 'read'>('edit'))

onMounted(() => account.fetchPeople())

const draft = reactive<TripDraft>({
  name: '',
  type: 'outbound',
  status: 'planned',
  timezone: 'Europe/London',
  party: { lead_passenger: '', passengers: [] },
})

const typeOptions: { label: string; value: TripType }[] = [
  { label: 'Outbound', value: 'outbound' },
  { label: 'Return', value: 'return' },
]
const statusOptions: { label: string; value: TripStatus }[] = [
  { label: 'Planned', value: 'planned' },
  { label: 'Booked', value: 'booked' },
  { label: 'Completed', value: 'completed' },
  { label: 'Cancelled', value: 'cancelled' },
]

// Lead options follow whoever is on the party.
const leadOptions = computed(() =>
  draft.party.passengers.map((id) => account.getPerson(id)).filter((p) => !!p),
)

const canCreate = computed(() => draft.name.trim().length > 0 && draft.party.passengers.length > 0)

function create() {
  if (!canCreate.value) return
  // Default lead to first person if not explicitly chosen / no longer present.
  if (!draft.party.passengers.includes(draft.party.lead_passenger)) {
    draft.party.lead_passenger = draft.party.passengers[0] ?? ''
  }
  trips.saveTrip(draft) // sets currentTripId
  router.push('/')
}
</script>

<template>
  <section class="new-trip">
    <header class="new-trip__head">
      <h2>New trip</h2>
    </header>

    <div class="field">
      <label for="trip-name">Trip name</label>
      <InputText id="trip-name" v-model="draft.name" placeholder="e.g. Family trip to Barcelona" fluid />
    </div>

    <div class="row">
      <div class="field">
        <label for="trip-type">Direction</label>
        <Select id="trip-type" v-model="draft.type" :options="typeOptions" option-label="label" option-value="value" fluid />
      </div>
      <div class="field">
        <label for="trip-status">Status</label>
        <Select id="trip-status" v-model="draft.status" :options="statusOptions" option-label="label" option-value="value" fluid />
      </div>
      <div class="field">
        <label for="trip-tz">Timezone</label>
        <InputText id="trip-tz" v-model="draft.timezone" fluid />
      </div>
    </div>

    <h3 class="new-trip__section">People</h3>
    <p class="hint">Add existing people or create new ones. Documents (passport, GHIC) follow the person.</p>
    <PartyPeoplePicker v-model="draft.party.passengers" />

    <div v-if="leadOptions.length" class="field">
      <label for="trip-lead">Lead passenger</label>
      <Select
        id="trip-lead"
        v-model="draft.party.lead_passenger"
        :options="leadOptions"
        option-label="name"
        option-value="id"
        placeholder="Defaults to first person"
        fluid
      />
    </div>

    <div class="actions">
      <Button label="Cancel" text @click="router.push('/trips')" />
      <Button label="Create trip" icon="pi pi-check" :disabled="!canCreate" @click="create" />
    </div>
  </section>
</template>

<style scoped>
.new-trip {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.new-trip__head {
  margin-bottom: 0.25rem;
}

.new-trip__section {
  margin: 0.75rem 0 0;
}

.hint {
  opacity: 0.7;
  font-size: 0.9rem;
  margin: 0;
}

.row {
  display: flex;
  gap: 1rem;
  flex-wrap: wrap;
}

.row .field {
  flex: 1;
  min-width: 160px;
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

.actions {
  display: flex;
  justify-content: flex-end;
  gap: 0.5rem;
  margin-top: 0.5rem;
}
</style>
