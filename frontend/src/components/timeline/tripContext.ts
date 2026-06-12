import { computed, type ComputedRef, type InjectionKey, type Ref } from 'vue'
import type { Party } from '@/types/trip'
import type { Person } from '@/types/person'
import { useAccountStore } from '@/stores/account'

/** Provided by the itinerary view so stage details (e.g. boarding passes) can read the party. */
export const partyKey: InjectionKey<Party> = Symbol('party')

export const emptyParty: Party = { lead_passenger: '', passengers: [] }

/** Resolve a party's person ids to account People (skips ids not found). The single
 *  source of truth for name/type/ghic/documents — so docs follow the person. */
export function usePartyPeople(party: Party): ComputedRef<Person[]> {
  const account = useAccountStore()
  return computed<Person[]>(() =>
    party.passengers.map((id) => account.getPerson(id)).filter((p): p is Person => !!p),
  )
}

/** Owner-edit vs read-only share. Drives view-vs-upload affordances. */
export type Mode = 'edit' | 'read'

export const modeKey: InjectionKey<Ref<Mode>> = Symbol('mode')
