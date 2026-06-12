<script setup lang="ts">
import { computed, ref } from 'vue'
import Button from 'primevue/button'
import Select from 'primevue/select'
import { STAGE_TYPES, getStageDef } from '@/data/stageRegistry'
import { useTripsStore } from '@/stores/trips'
import type { StageKind } from '@/types/stage'

const store = useTripsStore()

// Two-step picker: choose a kind, then a subkind when the kind has them.
const kind = ref<StageKind | null>(null)
const subkind = ref<string | null>(null)

const kindOptions = computed(() =>
  STAGE_TYPES.map((t) => ({ value: t.kind, label: `${t.icon}  ${t.label}` })),
)

const subOptions = computed(() => {
  if (!kind.value) return []
  return getStageDef(kind.value).subkinds ?? []
})

const subkindLabel = computed(() =>
  kind.value ? (getStageDef(kind.value).subkindLabel ?? 'Type') : 'Type',
)

// A subkind-bearing stage needs one chosen before it can be added.
const canAdd = computed(() => !!kind.value && (subOptions.value.length === 0 || !!subkind.value))

function reset() {
  kind.value = null
  subkind.value = null
}

function add() {
  if (!kind.value) return
  store.addStage({
    kind: kind.value,
    subkind: subkind.value ?? undefined,
    start: '',
    values: {},
  })
  reset()
}
</script>

<template>
  <div class="add-stage">
    <div class="add-stage__row">
      <Select
        v-model="kind"
        :options="kindOptions"
        option-label="label"
        option-value="value"
        placeholder="Add to itinerary…"
        @change="subkind = null"
        fluid
      />

      <Select
        v-if="subOptions.length"
        v-model="subkind"
        :options="subOptions"
        option-label="label"
        option-value="value"
        :placeholder="subkindLabel"
        fluid
      />

      <Button label="Add" icon="pi pi-plus" :disabled="!canAdd" @click="add" />
    </div>
  </div>
</template>

<style scoped>
.add-stage {
  width: 100%;
  max-width: 640px;
  margin: 0 auto;
}

.add-stage__row {
  display: flex;
  gap: 0.5rem;
  align-items: center;
  flex-wrap: wrap;
}

.add-stage__row :deep(.p-select) {
  flex: 1;
  min-width: 160px;
}
</style>
