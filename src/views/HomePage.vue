<script setup lang="ts">
import { computed, provide } from 'vue'
import StageList from '@/components/timeline/StageList.vue'
import TripOverview from '@/components/trip/TripOverview.vue'
import { partyKey, emptyParty } from '@/components/timeline/tripContext'
import { useTripsStore } from '@/stores/trips'
import { useAccountStore } from '@/stores/account'

const store = useTripsStore()
store.loadTrips()

// Party cards resolve people from the account store, so people must be loaded.
const account = useAccountStore()
account.fetchPeople()

// Show whichever trip is "active" in the store. Edits mutate the store copy directly,
// so they persist and stay in sync with the Trips page.
const trip = computed(() => store.currentTrip)

// TripOverview reads `insurance!`; guarantee the array on the active trip.
if (trip.value && !trip.value.insurance) trip.value.insurance = []

provide(partyKey, trip.value?.party ?? emptyParty)
</script>

<template>
  <section class="home">
    <template v-if="trip">
      <header class="home__head">
        <h2>{{ trip.trip.name }}</h2>
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
