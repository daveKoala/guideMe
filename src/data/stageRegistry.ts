import type { Stage, StageKind } from '../types/stage'

export interface FieldDef {
  key: string
  label: string
  type: 'text' | 'datetime-local' | 'date' | 'textarea' | 'select'
  options?: { value: string; label: string }[]
  placeholder?: string
}

export interface SubOption {
  value: string
  label: string
  startKey: string
  fields: FieldDef[]
}

export interface StageTypeDef {
  kind: StageKind
  label: string
  icon: string
  startKey?: string // field whose value seeds Stage.start (flat types)
  subkindLabel?: string
  subkinds?: SubOption[]
  fields?: FieldDef[]
}

export const STAGE_TYPES: StageTypeDef[] = [
  {
    kind: 'flight',
    label: 'Flight',
    icon: '✈️',
    startKey: 'depart',
    fields: [
      { key: 'airline', label: 'Airline', type: 'text', placeholder: 'e.g. Vueling' },
      { key: 'flight_no', label: 'Flight number', type: 'text', placeholder: 'e.g. VY7821' },
      { key: 'from', label: 'From', type: 'text', placeholder: 'Departure airport' },
      { key: 'to', label: 'To', type: 'text', placeholder: 'Arrival airport' },
      { key: 'depart', label: 'Departs', type: 'datetime-local' },
      { key: 'arrive', label: 'Arrives', type: 'datetime-local' },
    ],
  },
  {
    kind: 'travel_to_airport',
    label: 'Travel to airport',
    icon: '🚕',
    subkindLabel: 'How?',
    subkinds: [
      {
        value: 'taxi',
        label: 'Taxi',
        startKey: 'pickup_time',
        fields: [
          { key: 'company', label: 'Company', type: 'text', placeholder: 'e.g. Addison Lee' },
          { key: 'pickup_time', label: 'Pickup time', type: 'datetime-local' },
          { key: 'pickup_location', label: 'Pickup location', type: 'text' },
        ],
      },
      {
        value: 'car_park',
        label: 'Car & park',
        startKey: 'arrive_time',
        fields: [
          { key: 'car_park', label: 'Car park', type: 'text', placeholder: 'e.g. T2 Long Stay' },
          { key: 'arrive_time', label: 'Arrival time', type: 'datetime-local' },
          { key: 'booking_ref', label: 'Booking ref', type: 'text' },
        ],
      },
      {
        value: 'drop_off',
        label: 'Drop-off',
        startKey: 'drop_time',
        fields: [
          { key: 'driver', label: 'Driver', type: 'text', placeholder: 'Who is dropping you?' },
          { key: 'drop_time', label: 'Drop-off time', type: 'datetime-local' },
        ],
      },
      {
        value: 'public_transport',
        label: 'Public transport',
        startKey: 'depart_time',
        fields: [
          { key: 'service', label: 'Service', type: 'text', placeholder: 'e.g. train, bus' },
          { key: 'depart_time', label: 'Departs', type: 'datetime-local' },
          { key: 'from', label: 'From', type: 'text' },
          { key: 'to', label: 'To', type: 'text' },
        ],
      },
    ],
  },
  {
    kind: 'accommodation',
    label: 'Accommodation',
    icon: '🏨',
    startKey: 'check_in',
    fields: [
      { key: 'name', label: 'Name', type: 'text', placeholder: 'Hotel / place name' },
      { key: 'address', label: 'Address', type: 'text' },
      { key: 'check_in', label: 'Check-in', type: 'datetime-local' },
      { key: 'check_out', label: 'Check-out', type: 'datetime-local' },
    ],
  },
  {
    kind: 'note',
    label: 'Note',
    icon: '📝',
    startKey: 'when',
    fields: [
      { key: 'title', label: 'Title', type: 'text' },
      { key: 'when', label: 'When', type: 'datetime-local' },
      { key: 'notes', label: 'Notes', type: 'textarea' },
    ],
  },
]

export function getStageDef(kind: StageKind): StageTypeDef {
  const def = STAGE_TYPES.find((t) => t.kind === kind)
  if (!def) throw new Error(`Unknown stage kind: ${kind}`)
  return def
}

/** Resolve the active field set for a stage, honouring its subkind. */
export function getFields(kind: StageKind, subkind?: string): FieldDef[] {
  const def = getStageDef(kind)
  if (def.subkinds) {
    const sub = def.subkinds.find((s) => s.value === subkind)
    return sub ? sub.fields : []
  }
  return def.fields ?? []
}

export function getStartKey(kind: StageKind, subkind?: string): string {
  const def = getStageDef(kind)
  if (def.subkinds) {
    return def.subkinds.find((s) => s.value === subkind)?.startKey ?? ''
  }
  return def.startKey ?? ''
}

/** Human label for a stage's subkind, e.g. 'Taxi'. */
export function getSubkindLabel(kind: StageKind, subkind?: string): string {
  const def = getStageDef(kind)
  return def.subkinds?.find((s) => s.value === subkind)?.label ?? ''
}

/** Best-effort display title for a stage in the timeline. */
export function getStageTitle(stage: Stage): string {
  const def = getStageDef(stage.kind)
  const sub = getSubkindLabel(stage.kind, stage.subkind)
  return sub ? `${def.label} · ${sub}` : def.label
}
