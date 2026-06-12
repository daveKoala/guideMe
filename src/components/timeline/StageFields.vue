<script setup lang="ts">
import { computed, inject, ref } from 'vue'
import InputText from 'primevue/inputtext'
import Textarea from 'primevue/textarea'
import Select from 'primevue/select'
import DatePicker from 'primevue/datepicker'
import { getFields } from '@/data/stageRegistry'
import { modeKey, type Mode } from './tripContext'
import type { Stage } from '@/types/stage'

const props = defineProps<{ stage: Stage }>()

// Read-only on calm views (Home); editable only in the trip builder.
const mode = inject(modeKey, ref<Mode>('read'))

const fields = computed(() => getFields(props.stage.kind, props.stage.subkind))

/** Date fields are stored as ISO strings; convert for the DatePicker and back. */
function asDate(key: string): Date | null {
  const v = props.stage.values[key]
  if (!v) return null
  const d = new Date(v)
  return Number.isNaN(d.getTime()) ? null : d
}
function setDate(key: string, d: unknown) {
  props.stage.values[key] = d instanceof Date ? d.toISOString() : ''
}

/** Human-readable value for read mode; dates formatted, empties skipped. */
function readValue(key: string, type: string): string {
  const v = props.stage.values[key]
  if (!v) return ''
  if (type === 'date' || type === 'datetime-local') {
    const d = new Date(v)
    if (Number.isNaN(d.getTime())) return v
    return new Intl.DateTimeFormat(undefined, {
      dateStyle: 'medium',
      ...(type === 'datetime-local' ? { timeStyle: 'short' } : {}),
    }).format(d)
  }
  return v
}
</script>

<template>
  <div class="stage-fields">
    <!-- Read mode: calm static rows; skip empty fields entirely. -->
    <template v-if="mode === 'read'">
      <div v-for="f in fields" :key="f.key" class="read-row" v-show="readValue(f.key, f.type)">
        <span class="read-label">{{ f.label }}</span>
        <span class="read-value" :class="{ 'read-value--multiline': f.type === 'textarea' }">{{
          readValue(f.key, f.type)
        }}</span>
      </div>
    </template>

    <template v-else>
    <div v-for="f in fields" :key="f.key" class="field">
      <label :for="`${stage.id}-${f.key}`">{{ f.label }}</label>

      <Textarea
        v-if="f.type === 'textarea'"
        :id="`${stage.id}-${f.key}`"
        v-model="stage.values[f.key]"
        :placeholder="f.placeholder"
        rows="3"
        autoResize
        fluid
      />

      <Select
        v-else-if="f.type === 'select'"
        :input-id="`${stage.id}-${f.key}`"
        v-model="stage.values[f.key]"
        :options="f.options"
        option-label="label"
        option-value="value"
        :placeholder="f.placeholder ?? 'Choose…'"
        fluid
      />

      <DatePicker
        v-else-if="f.type === 'date' || f.type === 'datetime-local'"
        :input-id="`${stage.id}-${f.key}`"
        :model-value="asDate(f.key)"
        :show-time="f.type === 'datetime-local'"
        hour-format="24"
        fluid
        @update:model-value="(d) => setDate(f.key, d)"
      />

      <InputText
        v-else
        :id="`${stage.id}-${f.key}`"
        v-model="stage.values[f.key]"
        :placeholder="f.placeholder"
        fluid
      />
    </div>
    </template>
  </div>
</template>

<style scoped>
.stage-fields {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.field {
  display: flex;
  flex-direction: column;
  gap: 0.35rem;
}

.field label {
  font-weight: 600;
  font-size: 0.875rem;
}

.read-row {
  display: flex;
  gap: 0.75rem;
  align-items: baseline;
}

.read-label {
  font-weight: 600;
  font-size: 0.875rem;
  opacity: 0.6;
  min-width: 7rem;
}

.read-value {
  flex: 1;
}

.read-value--multiline {
  white-space: pre-wrap;
}
</style>
