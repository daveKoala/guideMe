<script setup lang="ts">
import { computed } from 'vue'
import InputText from 'primevue/inputtext'
import Textarea from 'primevue/textarea'
import Select from 'primevue/select'
import DatePicker from 'primevue/datepicker'
import { getFields } from '@/data/stageRegistry'
import type { Stage } from '@/types/stage'

const props = defineProps<{ stage: Stage }>()

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
</script>

<template>
  <div class="stage-fields">
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
</style>
