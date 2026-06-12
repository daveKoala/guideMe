<script setup lang="ts">
import { onMounted } from 'vue'
import Button from 'primevue/button'
import Tag from 'primevue/tag'
import { useTripsStore } from '@/stores/trips'
import type { Trip, TripStatus } from '@/types/trip'

const store = useTripsStore()

onMounted(() => store.loadTrips())

const statusSeverity: Record<TripStatus, string> = {
  planned: 'info',
  booked: 'success',
  completed: 'secondary',
  cancelled: 'danger',
}

function dateRange(trip: Trip): string {
  const starts = trip.stages.map((s) => s.start).filter(Boolean).sort()
  if (!starts.length) return 'No dates yet'
  const fmt = (iso: string | undefined) => (iso ?? '').slice(0, 10)
  const first = fmt(starts[0])
  const last = fmt(starts[starts.length - 1])
  return first === last ? first : `${first} → ${last}`
}
</script>

<template>
  <section class="trips-page">
    <header class="trips-page__head">
      <h2>Trips</h2>
      <Button label="Reset demo data" icon="pi pi-refresh" text size="small" @click="store.reset()" />
    </header>

    <p class="hint">All your trips. Activate one to make it the current trip.</p>

    <ul v-if="store.trips.length" class="trip-list">
      <li
        v-for="t in store.trips"
        :key="t.trip.id"
        class="row"
        :class="{ 'row--current': t.trip.id === store.currentTripId }"
      >
        <div class="row__main">
          <span class="row__name">{{ t.trip.name }}</span>
          <span class="row__meta">{{ dateRange(t) }} · {{ t.party.passengers.length }} travelling</span>
        </div>

        <Tag :value="t.trip.status" :severity="statusSeverity[t.trip.status]" />

        <Tag v-if="t.trip.id === store.currentTripId" value="Current" severity="contrast" />
        <Button
          v-else
          label="Activate"
          icon="pi pi-check"
          size="small"
          outlined
          @click="store.setCurrentTrip(t.trip.id)"
        />
      </li>
    </ul>

    <p v-else class="status">No trips yet.</p>
  </section>
</template>

<style scoped>
.trips-page__head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 1rem;
  flex-wrap: wrap;
  margin-bottom: 0.5rem;
}

.hint {
  opacity: 0.7;
  font-size: 0.9rem;
  margin-bottom: 1.25rem;
}

.status {
  opacity: 0.6;
}

.trip-list {
  list-style: none;
  margin: 0;
  padding: 0;
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.row {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  padding: 0.75rem 1rem;
  border: 1px solid var(--color-border);
  border-radius: var(--radius-md);
  background: var(--color-surface);
}

.row--current {
  border-color: var(--color-link);
  box-shadow: 0 0 0 1px var(--color-link) inset;
}

.row__main {
  display: flex;
  flex-direction: column;
  gap: 0.2rem;
  flex: 1;
  min-width: 0;
}

.row__name {
  font-weight: 600;
}

.row__meta {
  opacity: 0.6;
  font-size: 0.85rem;
}
</style>
