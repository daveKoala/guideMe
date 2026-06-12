import type { Trip } from '@/types/trip'

export const mockTrip: Trip = {
  trip: {
    id: 'trip_2026_07_01_man_bcn',
    name: 'Family trip to Barcelona',
    type: 'outbound',
    status: 'planned',
    timezone: 'Europe/London',
    created_at: '2026-06-10T12:00:00+01:00',
    updated_at: '2026-06-10T12:00:00+01:00',
  },

  party: {
    lead_passenger: 'person_dave',
    passengers: ['person_dave', 'person_adult2', 'person_child1', 'person_child2', 'person_child3'],
  },

  sharing: {
    owner_edit_url: '/trips/trip_2026_07_01_man_bcn/edit/secret-edit-token',
    read_only_url: '/trips/trip_2026_07_01_man_bcn/share/secret-share-token',
    offline_enabled: true,
  },

  insurance: [
    {
      id: 'ins_family',
      policy_number: 'AXA-99481726',
      emergency_contact: '+44 20 7946 0000',
      account_url: 'https://example-insurer.com/account',
      covers: ['person_dave', 'person_adult2', 'person_child1', 'person_child2', 'person_child3'],
      medical: {
        assist_id: 'MA-55012',
        phone: '+44 20 7946 1111',
        url: 'https://example-assist.com/members',
      },
    },
    {
      id: 'ins_dave_ski',
      policy_number: 'SKI-22310',
      emergency_contact: '+44 20 7946 2222',
      account_url: 'https://example-ski-insurer.com',
      covers: ['person_dave'],
      medical: {
        assist_id: 'SKI-ASSIST-77',
        phone: '+44 20 7946 3333',
        url: 'https://example-ski-assist.com',
      },
    },
  ],

  stages: [
    {
      id: 'stage_1',
      kind: 'travel_to_airport',
      subkind: 'taxi',
      start: '2026-07-01T05:00:00',
      values: {
        company: 'Addison Lee',
        pickup_time: '2026-07-01T05:00:00',
        pickup_location: '12 Elm Road, Manchester',
      },
    },
     {
      id: 'stage_1b',
      kind: 'note',
      start: '2026-07-02T09:00:00',
      values: {
        title: 'Sagrada Família tickets',
        when: '2026-07-02T09:00:00',
        notes: 'Pre-booked timed entry at 09:15. Bring passports for ID check.',
      },
    },
    {
      id: 'stage_2',
      kind: 'flight',
      start: '2026-07-01T07:30:00',
      values: {
        airline: 'Vueling',
        flight_no: 'VY7821',
        from: 'MAN',
        to: 'BCN',
        depart: '2026-07-01T07:30:00',
        arrive: '2026-07-01T10:50:00',
      },
      boarding_passes: {
        person_dave: {
          fileName: 'boardingpass.png',
          dataUrl: '/test-docs/boardingpass.png',
          uploadedAt: '2026-06-20T00:00:00Z',
        },
        person_adult2: {
          fileName: 'boardingpass.png',
          dataUrl: '/test-docs/boardingpass.png',
          uploadedAt: '2026-06-20T00:00:00Z',
        },
      },
    },
    {
      id: 'stage_3',
      kind: 'accommodation',
      start: '2026-07-01T14:00:00',
      values: {
        name: 'Hotel Arts Barcelona',
        address: "Marina, 19-21, 08005 Barcelona",
        check_in: '2026-07-01T14:00:00',
        check_out: '2026-07-08T11:00:00',
      },
    },
    {
      id: 'stage_4',
      kind: 'note',
      start: '2026-07-02T09:00:00',
      values: {
        title: 'Sagrada Família tickets',
        when: '2026-07-02T09:00:00',
        notes: 'Pre-booked timed entry at 09:15. Bring passports for ID check.',
      },
    },
  ],
}
