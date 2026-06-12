<script setup lang="ts">
import { computed, onMounted, provide, reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import Button from 'primevue/button'
import InputText from 'primevue/inputtext'
import Select from 'primevue/select'
import Checkbox from 'primevue/checkbox'
import PartyPeoplePicker from '@/components/people/PartyPeoplePicker.vue'
import { useTripsStore, type TripDraft } from '@/stores/trips'
import { useAccountStore } from '@/stores/account'
import { modeKey } from '@/components/timeline/tripContext'

const router = useRouter()
const trips = useTripsStore()
const account = useAccountStore()

// The picker's doc slots need edit affordances.
provide(modeKey, ref<'edit' | 'read'>('edit'))

onMounted(() => account.fetchPeople())

const draft = reactive<TripDraft>({
  name: '',
  // Direction/status/timezone aren't asked at create — sensible defaults; status is
  // editable later. Kept on the draft so the data model stays complete.
  type: 'outbound',
  status: 'planned',
  timezone: 'Europe/London',
  party: { lead_passenger: '', passengers: [] },
  needs_insurance: false,
  needs_ghic: false,
})

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
  const trip = trips.saveTrip(draft) // sets currentTripId
  // Hand off to the builder to fill in itinerary, insurance, GHIC.
  router.push(`/trips/${trip.trip.id}/build`)
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

    <h3 class="new-trip__section">What to set up</h3>
    <p class="hint">We'll open these for you in the builder next.</p>
    <div class="checks">
      <div class="check">
        <Checkbox v-model="draft.needs_insurance" :binary="true" input-id="needs-insurance" />
        <label for="needs-insurance">Travel insurance</label>
      </div>
      <div class="check">
        <Checkbox v-model="draft.needs_ghic" :binary="true" input-id="needs-ghic" />
        <label for="needs-ghic">GHIC (EU health cover)</label>
      </div>
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

.checks {
  display: flex;
  flex-direction: column;
  gap: 0.6rem;
}

.check {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.check label {
  font-weight: 500;
}

.actions {
  display: flex;
  justify-content: flex-end;
  gap: 0.5rem;
  margin-top: 0.5rem;
}
</style>
