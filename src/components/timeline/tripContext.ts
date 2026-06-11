import type { InjectionKey, Ref } from 'vue'
import type { Party } from '@/types/trip'

/** Provided by the itinerary view so stage details (e.g. boarding passes) can read the party. */
export const partyKey: InjectionKey<Party> = Symbol('party')

export const emptyParty: Party = { lead_passenger: '', passengers: [] }

/** Owner-edit vs read-only share. Drives view-vs-upload affordances. */
export type Mode = 'edit' | 'read'

export const modeKey: InjectionKey<Ref<Mode>> = Symbol('mode')
