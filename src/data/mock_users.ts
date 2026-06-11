import type { Person } from '../types/person'

/** Stands in for the account API's "all users" response. Won't be many. */
export const mockUsers: Person[] = [
  {
    id: 'person_dave',
    name: 'Dave',
    type: 'adult',
    ghic_id: 'GHIC-90011223344',
    dob: '1985-04-12',
    documents: {
      passport: { fileName: 'passport.png', dataUrl: '/test-docs/passport.png', uploadedAt: '2026-06-01T00:00:00Z' },
    },
  },
  {
    id: 'person_adult2',
    name: 'Adult 2',
    type: 'adult',
    ghic_id: 'GHIC-90011223355',
    dob: '1987-09-30',
    documents: {
      passport: { fileName: 'passport.png', dataUrl: '/test-docs/passport.png', uploadedAt: '2026-06-01T00:00:00Z' },
    },
  },
  { id: 'person_child1', name: 'Child 1', type: 'child', dob: '2015-02-20', documents: {} },
  { id: 'person_child2', name: 'Child 2', type: 'child', dob: '2017-06-08', documents: {} },
  { id: 'person_child3', name: 'Child 3', type: 'child', dob: '2019-11-15', documents: {} },
  // extra account users beyond this one trip's party
  {
    id: 'person_sue',
    name: 'Grandma Sue',
    type: 'adult',
    ghic_id: 'GHIC-90011229999',
    dob: '1956-01-03',
    documents: {
      passport: { fileName: 'passport.png', dataUrl: '/test-docs/passport.png', uploadedAt: '2026-06-01T00:00:00Z' },
    },
  },
  { id: 'person_tom', name: 'Friend Tom', type: 'adult', dob: '1990-07-22', documents: {} },
]
