import type { Component } from 'vue'
import type { StageKind } from '@/types/stage'
import FlightDetail from './details/FlightDetail.vue'
import AccommodationDetail from './details/AccommodationDetail.vue'
import TravelDetail from './details/TravelDetail.vue'
import NoteDetail from './details/NoteDetail.vue'

/** kind → expanded detail component. Add a stage type = add one line here. */
const MAP: Record<StageKind, Component> = {
  flight: FlightDetail,
  travel_to_airport: TravelDetail,
  accommodation: AccommodationDetail,
  note: NoteDetail,
}

export const detailFor = (kind: StageKind): Component => MAP[kind]
