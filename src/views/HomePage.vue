<script setup lang="ts">
import { computed, provide, ref } from 'vue'
import SelectButton from 'primevue/selectbutton'
import StageList from '@/components/timeline/StageList.vue'
import TripOverview from '@/components/trip/TripOverview.vue'
import { partyKey, modeKey, emptyParty, type Mode } from '@/components/timeline/tripContext'
import { useTripsStore } from '@/stores/trips'

const store = useTripsStore()
store.loadTrips()

// Show whichever trip is "active" in the store. Edits mutate the store copy directly,
// so they persist and stay in sync with the Trips page.
const trip = computed(() => store.currentTrip)

// TripOverview reads `insurance!`; guarantee the array on the active trip.
if (trip.value && !trip.value.insurance) trip.value.insurance = []

const mode = ref<Mode>('edit')
const modeOptions: { label: string; value: Mode }[] = [
  { label: 'Edit', value: 'edit' },
  { label: 'Read', value: 'read' },
]

provide(partyKey, trip.value?.party ?? emptyParty)
provide(modeKey, mode)
</script>

<template>
  <section class="home">
    <template v-if="trip">
      <header class="home__head">
        <h2>{{ trip.trip.name }}</h2>
        <SelectButton
          v-model="mode"
          :options="modeOptions"
          option-label="label"
          option-value="value"
          :allow-empty="false"
          aria-label="Edit or read mode"
        />
      </header>

      <TripOverview :trip="trip" />

      <h3 class="home__section">Itinerary</h3>
      <StageList :stages="trip.stages" />
    </template>

    <p v-else class="home__empty">No active trip. Activate one on the Trips page.</p>
  </section>
</template>

<style scoped>
.home__head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 1rem;
  flex-wrap: wrap;
  margin-bottom: 1.25rem;
}

.home__section {
  max-width: 640px;
  margin: 1.75rem auto 0.75rem;
}

.home__empty {
  opacity: 0.6;
}
</style>
