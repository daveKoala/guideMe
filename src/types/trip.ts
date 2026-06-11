import type { Stage } from './stage'

export type PassengerType = 'adult' | 'child' | 'infant'
export type TripType = 'outbound' | 'return'
export type TripStatus = 'planned' | 'booked' | 'completed' | 'cancelled'

export interface Passenger {
  id: string
  name: string
  type: PassengerType
  ghic_id?: string // UK Global Health Insurance Card number
}

export interface TripMeta {
  id: string
  name: string
  type: TripType
  status: TripStatus
  timezone: string
  created_at: string
  updated_at: string
}

export interface Party {
  lead_passenger: string
  passengers: Passenger[]
}

export interface Sharing {
  owner_edit_url: string
  read_only_url: string
  offline_enabled: boolean
}

export interface Insurance {
  id: string
  policy_number?: string
  emergency_contact?: string
  account_url?: string
  covers: string[] // passenger ids — one or many
  medical: MedicalAssist // 3rd-party medical assistance provider for this policy
}

export interface MedicalAssist {
  assist_id?: string
  phone?: string
  url?: string
}

export interface Trip {
  trip: TripMeta
  party: Party
  sharing: Sharing
  stages: Stage[]
  insurance?: Insurance[]
}
