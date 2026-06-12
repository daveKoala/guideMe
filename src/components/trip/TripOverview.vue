<script setup lang="ts">
import PeopleCard from './PeopleCard.vue'
import GhicCard from './GhicCard.vue'
import InsuranceCard from './InsuranceCard.vue'
import type { Trip } from '@/types/trip'

const props = defineProps<{ trip: Trip }>()

// Guaranteed present by the view (initialised before provide); `!` keeps the child prop non-optional.
const policies = props.trip.insurance!
</script>

<template>
  <div class="trip-overview">
    <PeopleCard :party="trip.party" />
    <GhicCard :default-open="!!trip.trip.needs_ghic" />
    <InsuranceCard :policies="policies" :default-open="!!trip.trip.needs_insurance" />
  </div>
</template>

<style scoped>
.trip-overview {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  width: 100%;
  max-width: 640px;
  margin: 0 auto;
}
</style>
