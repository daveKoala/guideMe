import type { Trip } from '@/types/trip'
import { mockTrip } from './mock_trip'

/** Stands in for the account API's "all trips" response. */
export const mockTrips: Trip[] = [
  mockTrip,
  {
    trip: {
      id: 'trip_2026_09_14_lhr_jfk',
      name: 'New York city break',
      type: 'outbound',
      status: 'booked',
      timezone: 'Europe/London',
      created_at: '2026-05-20T09:00:00+01:00',
      updated_at: '2026-06-05T18:30:00+01:00',
    },
    party: {
      lead_passenger: 'Dave',
      passengers: [
        { id: 'passenger_1', name: 'Dave', type: 'adult', ghic_id: 'GHIC-90011223344' },
        { id: 'passenger_2', name: 'Adult 2', type: 'adult', ghic_id: 'GHIC-90011223355' },
      ],
    },
    sharing: {
      owner_edit_url: '/trips/trip_2026_09_14_lhr_jfk/edit/secret-edit-token',
      read_only_url: '/trips/trip_2026_09_14_lhr_jfk/share/secret-share-token',
      offline_enabled: true,
    },
    insurance: [
      {
        id: 'ins_couple',
        policy_number: 'AXA-77120034',
        emergency_contact: '+44 20 7946 0000',
        account_url: 'https://example-insurer.com/account',
        covers: ['passenger_1', 'passenger_2'],
        medical: {
          assist_id: 'MA-66102',
          phone: '+44 20 7946 1111',
          url: 'https://example-assist.com/members',
        },
      },
    ],
    stages: [
      {
        id: 'stage_ny_1',
        kind: 'flight',
        start: '2026-09-14T11:00:00',
        values: {
          airline: 'British Airways',
          flight_no: 'BA177',
          from: 'LHR',
          to: 'JFK',
          depart: '2026-09-14T11:00:00',
          arrive: '2026-09-14T14:05:00',
        },
      },
      {
        id: 'stage_ny_2',
        kind: 'accommodation',
        start: '2026-09-14T16:00:00',
        values: {
          name: 'The Standard, High Line',
          address: '848 Washington St, New York, NY 10014',
          check_in: '2026-09-14T16:00:00',
          check_out: '2026-09-18T11:00:00',
        },
      },
    ],
  },
  {
    trip: {
      id: 'trip_2025_12_22_man_gva',
      name: 'Christmas ski week',
      type: 'outbound',
      status: 'completed',
      timezone: 'Europe/London',
      created_at: '2025-10-01T10:00:00+01:00',
      updated_at: '2026-01-05T12:00:00+00:00',
    },
    party: {
      lead_passenger: 'Dave',
      passengers: [
        { id: 'passenger_1', name: 'Dave', type: 'adult', ghic_id: 'GHIC-90011223344' },
        { id: 'passenger_2', name: 'Grandma Sue', type: 'adult', ghic_id: 'GHIC-90011229999' },
        { id: 'passenger_3', name: 'Child 1', type: 'child' },
      ],
    },
    sharing: {
      owner_edit_url: '/trips/trip_2025_12_22_man_gva/edit/secret-edit-token',
      read_only_url: '/trips/trip_2025_12_22_man_gva/share/secret-share-token',
      offline_enabled: false,
    },
    stages: [
      {
        id: 'stage_ski_1',
        kind: 'travel_to_airport',
        subkind: 'taxi',
        start: '2025-12-22T04:30:00',
        values: {
          company: 'Addison Lee',
          pickup_time: '2025-12-22T04:30:00',
          pickup_location: '12 Elm Road, Manchester',
        },
      },
      {
        id: 'stage_ski_2',
        kind: 'flight',
        start: '2025-12-22T07:00:00',
        values: {
          airline: 'easyJet',
          flight_no: 'U21923',
          from: 'MAN',
          to: 'GVA',
          depart: '2025-12-22T07:00:00',
          arrive: '2025-12-22T10:15:00',
        },
      },
    ],
  },
  {
    trip: {
      id: 'trip_2027_03_10_man_kef',
      name: 'Iceland northern lights',
      type: 'outbound',
      status: 'planned',
      timezone: 'Europe/London',
      created_at: '2026-06-08T20:00:00+01:00',
      updated_at: '2026-06-08T20:00:00+01:00',
    },
    party: {
      lead_passenger: 'Dave',
      passengers: [{ id: 'passenger_1', name: 'Dave', type: 'adult', ghic_id: 'GHIC-90011223344' }],
    },
    sharing: {
      owner_edit_url: '/trips/trip_2027_03_10_man_kef/edit/secret-edit-token',
      read_only_url: '/trips/trip_2027_03_10_man_kef/share/secret-share-token',
      offline_enabled: true,
    },
    stages: [
      {
        id: 'stage_isl_1',
        kind: 'note',
        start: '2026-06-08T20:00:00',
        values: {
          title: 'Book aurora tour',
          when: '2026-06-08T20:00:00',
          notes: 'Compare operators in Reykjavík; want a small-group minibus tour.',
        },
      },
    ],
  },
]
