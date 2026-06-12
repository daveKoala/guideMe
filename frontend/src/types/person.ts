import type { PassengerType } from './trip' // reuse 'adult' | 'child' | 'infant'

export type PersonDocKind = 'passport' | 'ghic_card'

export interface StoredDoc {
  fileName: string
  dataUrl: string // base64 — persists + powers view/download across refresh
  uploadedAt: string
}

/** Canonical account-level person. Trips reference these by id (party.passengers). */
export interface Person {
  id: string
  name: string
  type: PassengerType
  ghic_id?: string
  dob?: string // YYYY-MM-DD
  documents: Partial<Record<PersonDocKind, StoredDoc>>
}
