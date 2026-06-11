import { ref, computed, watch } from 'vue'
import { defineStore } from 'pinia'
import type { Trip, TripType, TripStatus, Party, Sharing } from '../types/trip'
import type { Stage } from '../types/stage'
import { mockTrips } from '../data/mock_trips'

const STORAGE_KEY = 'guide-me:trips'

/** Editable fields the form produces. `id`/`created_at` present only when editing. */
export interface TripDraft {
  id?: string
  created_at?: string
  name: string
  type: TripType
  status: TripStatus
  timezone: string
  party: Party
}

function nowIso(): string {
  return new Date().toISOString()
}

function slugify(value: string): string {
  return value
    .toLowerCase()
    .trim()
    .replace(/[^a-z0-9]+/g, '_')
    .replace(/^_+|_+$/g, '')
}

function genId(name: string): string {
  const date = nowIso().slice(0, 10).replace(/-/g, '_')
  const slug = slugify(name) || 'trip'
  return `trip_${date}_${slug}`
}

let stageSeq = 0
function genStageId(): string {
  stageSeq += 1
  return `stage_${nowIso()}_${stageSeq}`
}

function makeSharing(id: string): Sharing {
  return {
    owner_edit_url: `/trips/${id}/edit/secret-edit-token`,
    read_only_url: `/trips/${id}/share/secret-share-token`,
    offline_enabled: true,
  }
}

interface PersistedState {
  trips: Trip[]
  currentTripId: string | null
}

function loadState(): PersistedState {
  try {
    const raw = localStorage.getItem(STORAGE_KEY)
    if (!raw) return { trips: [], currentTripId: null }
    const parsed = JSON.parse(raw) as Partial<PersistedState>
    // Migration: older trips were saved before `stages` existed.
    const trips = Array.isArray(parsed.trips)
      ? parsed.trips.map((t) => ({ ...t, stages: Array.isArray(t.stages) ? t.stages : [] }))
      : []
    return {
      trips,
      currentTripId: typeof parsed.currentTripId === 'string' ? parsed.currentTripId : null,
    }
  } catch {
    return { trips: [], currentTripId: null }
  }
}

export const useTripsStore = defineStore('trips', () => {
  const initial = loadState()
  const trips = ref<Trip[]>(initial.trips)
  const currentTripId = ref<string | null>(initial.currentTripId)

  const currentTrip = computed(
    () => trips.value.find((t) => t.trip.id === currentTripId.value) ?? null,
  )

  function persist() {
    try {
      localStorage.setItem(
        STORAGE_KEY,
        JSON.stringify({ trips: trips.value, currentTripId: currentTripId.value }),
      )
    } catch {
      // storage unavailable / quota — ignore, keep working in-memory
    }
  }

  watch([trips, currentTripId], persist, { deep: true })

  /** Deep-clone the mock set so the store owns its copy and HomePage's
   *  reactive mockTrip stays independent. */
  function seed(): Trip[] {
    return mockTrips.map((t) => structuredClone(t))
  }

  /** Seed from mock trips on first run (empty storage) and guarantee a valid
   *  current trip — falls back to the first trip if none/stale id selected. */
  function loadTrips() {
    if (trips.value.length === 0) trips.value = seed()
    const exists = trips.value.some((t) => t.trip.id === currentTripId.value)
    if (!exists) currentTripId.value = trips.value[0]?.trip.id ?? null
  }

  /** Drop persisted trips and re-seed from mock. Recovers from stale storage
   *  that predates a mock change (e.g. a missing demo trip). */
  function reset() {
    try {
      localStorage.removeItem(STORAGE_KEY)
    } catch {
      // ignore
    }
    trips.value = seed()
    currentTripId.value = trips.value[0]?.trip.id ?? null
  }

  function setCurrentTrip(id: string | null) {
    currentTripId.value = id
  }

  /** Upsert: edit existing trip (matched by draft.id) or create a new one. */
  function saveTrip(draft: TripDraft): Trip {
    const existing = draft.id ? trips.value.find((t) => t.trip.id === draft.id) : undefined

    if (existing) {
      existing.trip.name = draft.name
      existing.trip.type = draft.type
      existing.trip.status = draft.status
      existing.trip.timezone = draft.timezone
      existing.trip.updated_at = nowIso()
      existing.party = draft.party
      currentTripId.value = existing.trip.id
      return existing
    }

    const id = genId(draft.name)
    const created = draft.created_at ?? nowIso()
    const trip: Trip = {
      trip: {
        id,
        name: draft.name,
        type: draft.type,
        status: draft.status,
        timezone: draft.timezone,
        created_at: created,
        updated_at: nowIso(),
      },
      party: draft.party,
      sharing: makeSharing(id),
      stages: [],
    }
    trips.value.push(trip)
    currentTripId.value = id
    return trip
  }

  function deleteTrip(id: string) {
    trips.value = trips.value.filter((t) => t.trip.id !== id)
    if (currentTripId.value === id) currentTripId.value = null
  }

  /** Add a stage to the current trip. Generates an id if missing; bumps updated_at. */
  function addStage(stage: Omit<Stage, 'id'> & { id?: string }): Stage | null {
    const trip = currentTrip.value
    if (!trip) return null
    const full: Stage = { ...stage, id: stage.id ?? genStageId() }
    trip.stages.push(full)
    trip.trip.updated_at = nowIso()
    return full
  }

  /** Replace an existing stage on the current trip by id. */
  function updateStage(stage: Stage) {
    const trip = currentTrip.value
    if (!trip) return
    const i = trip.stages.findIndex((s) => s.id === stage.id)
    if (i === -1) return
    trip.stages[i] = stage
    trip.trip.updated_at = nowIso()
  }

  function removeStage(id: string) {
    const trip = currentTrip.value
    if (!trip) return
    trip.stages = trip.stages.filter((s) => s.id !== id)
    trip.trip.updated_at = nowIso()
  }

  return {
    trips,
    currentTripId,
    currentTrip,
    loadTrips,
    reset,
    setCurrentTrip,
    saveTrip,
    deleteTrip,
    addStage,
    updateStage,
    removeStage,
  }
})
