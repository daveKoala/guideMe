import type { StoredDoc } from './person'

export type StageKind = 'flight' | 'travel_to_airport' | 'accommodation' | 'note'

export interface Stage {
  id: string
  kind: StageKind
  subkind?: string // e.g. 'taxi' for travel_to_airport
  start: string // ISO datetime copied from the type's startKey field; '' allowed
  values: Record<string, string> // field key -> value
  boarding_passes?: Record<string, StoredDoc> // person id -> boarding pass (flight only)
}
