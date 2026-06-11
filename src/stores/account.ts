import { ref, watch } from 'vue'
import { defineStore } from 'pinia'
import type { Person, PersonDocKind, StoredDoc } from '../types/person'
import { mockUsers } from '../data/mock_users'

const STORAGE_KEY = 'guide-me:people:v3'

function seed(): Person[] {
  return mockUsers.map((u) => ({ ...u, documents: { ...u.documents } }))
}

function delay(ms: number): Promise<void> {
  return new Promise((resolve) => setTimeout(resolve, ms))
}

function loadPersisted(): Person[] | null {
  try {
    const raw = localStorage.getItem(STORAGE_KEY)
    if (!raw) return null
    const parsed = JSON.parse(raw)
    return Array.isArray(parsed) ? (parsed as Person[]) : null
  } catch {
    return null
  }
}

function slugify(value: string): string {
  return value
    .toLowerCase()
    .trim()
    .replace(/[^a-z0-9]+/g, '_')
    .replace(/^_+|_+$/g, '')
}

export const useAccountStore = defineStore('account', () => {
  const people = ref<Person[]>([])
  const loading = ref(false)
  const loaded = ref(false)

  function persist() {
    try {
      // NOTE: document dataUrls are base64 and count against the ~5MB localStorage quota.
      // Fine for a few small scans in this prototype. TODO: IndexedDB for real blobs.
      localStorage.setItem(STORAGE_KEY, JSON.stringify(people.value))
    } catch {
      // quota exceeded / unavailable — ignore, keep working in-memory
    }
  }

  watch(people, persist, { deep: true })

  /** Simulated account API: returns ALL users. Hydrates from storage, else seeds mock. */
  async function fetchPeople() {
    if (loaded.value || loading.value) return
    loading.value = true
    try {
      await delay(400)
      const persisted = loadPersisted()
      people.value = persisted ?? seed()
      loaded.value = true
    } finally {
      loading.value = false
    }
  }

  function getPerson(id: string): Person | undefined {
    return people.value.find((p) => p.id === id)
  }

  function updatePerson(person: Person) {
    const i = people.value.findIndex((p) => p.id === person.id)
    if (i !== -1) people.value[i] = person
  }

  function setDocument(personId: string, kind: PersonDocKind, doc: StoredDoc) {
    const p = getPerson(personId)
    if (p) p.documents[kind] = doc
  }

  function removeDocument(personId: string, kind: PersonDocKind) {
    const p = getPerson(personId)
    if (p) delete p.documents[kind]
  }

  function addPerson(): Person {
    const person: Person = {
      id: `person_${slugify('new')}_${crypto.randomUUID().slice(0, 8)}`,
      name: '',
      type: 'adult',
      documents: {},
    }
    people.value.push(person)
    return person
  }

  function removePerson(id: string) {
    people.value = people.value.filter((p) => p.id !== id)
  }

  /** Drop any stored edits/uploads and re-seed from mock. Handy while swapping test docs. */
  function reset() {
    try {
      localStorage.removeItem(STORAGE_KEY)
    } catch {
      // ignore
    }
    people.value = seed()
    loaded.value = true
  }

  return {
    people,
    loading,
    loaded,
    fetchPeople,
    getPerson,
    updatePerson,
    setDocument,
    removeDocument,
    addPerson,
    removePerson,
    reset,
  }
})
