<script setup lang="ts">
import { computed, provide, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import Button from 'primevue/button'
import TripOverview from '@/components/trip/TripOverview.vue'
import StageList from '@/components/timeline/StageList.vue'
import AddStage from '@/components/timeline/AddStage.vue'
import { modeKey, partyKey, emptyParty, type Mode } from '@/components/timeline/tripContext'
import { useTripsStore } from '@/stores/trips'
import { useAccountStore } from '@/stores/account'

const route = useRoute()
const router = useRouter()
const store = useTripsStore()
const account = useAccountStore()

store.loadTrips()
account.fetchPeople()

// This is the one place trips are edited; everything inside renders editors.
provide(modeKey, ref<Mode>('edit'))

const tripId = computed(() => String(route.params.id))
const trip = computed(() => store.trips.find((t) => t.trip.id === tripId.value) ?? null)

// TripOverview reads `insurance!`; guarantee the array before it renders.
if (trip.value && !trip.value.insurance) trip.value.insurance = []

provide(partyKey, trip.value?.party ?? emptyParty)

// Show stages in time order; blank-start stages (just added) sink to the bottom.
const orderedStages = computed(() =>
  trip.value
    ? [...trip.value.stages].sort((a, b) => (a.start || '￿').localeCompare(b.start || '￿'))
    : [],
)

function done() {
  router.push('/')
}
</script>

<template>
  <section class="builder">
    <template v-if="trip">
      <header class="builder__head">
        <div>
          <p class="builder__eyebrow">Trip builder</p>
          <h2>{{ trip.trip.name }}</h2>
        </div>
        <Button label="Done" icon="pi pi-check" @click="done" />
      </header>

      <p class="hint">Set up people, GHIC, insurance and build the itinerary. Changes save automatically.</p>

      <TripOverview :trip="trip" />

      <h3 class="builder__section">Itinerary</h3>
      <AddStage />
      <StageList :stages="orderedStages" />
    </template>

    <p v-else class="builder__empty">
      Trip not found. <RouterLink to="/trips">Back to trips</RouterLink>
    </p>
  </section>
</template>

<style scoped>
.builder {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.builder__head {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 1rem;
  flex-wrap: wrap;
}

.builder__eyebrow {
  margin: 0;
  font-size: 0.75rem;
  text-transform: uppercase;
  letter-spacing: 0.06em;
  opacity: 0.6;
}

.builder__head h2 {
  margin: 0.1rem 0 0;
}

.builder__section {
  max-width: 640px;
  margin: 1.25rem auto 0.25rem;
  width: 100%;
}

.hint {
  opacity: 0.7;
  font-size: 0.9rem;
  margin: 0;
}

.builder__empty {
  opacity: 0.6;
}
</style>
