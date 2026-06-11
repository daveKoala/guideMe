<script setup lang="ts">
import Panel from 'primevue/panel'
import Toolbar from 'primevue/toolbar'
import { detailFor } from './stageDetails'
import { getStageDef, getStageTitle } from '@/data/stageRegistry'
import type { Stage } from '@/types/stage'

const props = defineProps<{ stage: Stage; expanded: boolean }>()
defineEmits<{ toggle: [] }>()

function formatStart(iso: string): string {
  if (!iso) return ''
  const d = new Date(iso)
  if (Number.isNaN(d.getTime())) return iso
  return new Intl.DateTimeFormat(undefined, { dateStyle: 'medium', timeStyle: 'short' }).format(d)
}

const icon = () => getStageDef(props.stage.kind).icon
</script>

<template>
  <Panel
    class="stage-card"
    toggleable
    :collapsed="!expanded"
    @update:collapsed="$emit('toggle')"
  >
    <template #header>
      <Toolbar class="stage-card__bar" @click="$emit('toggle')">
        <template #start>
          <span class="stage-icon">{{ icon() }}</span>
          <span class="stage-card__title">{{ getStageTitle(stage) }}</span>
        </template>
        <template #end>
          <time class="stage-card__when">{{ formatStart(stage.start) }}</time>
        </template>
      </Toolbar>
    </template>

    <div class="stage-card__body">
      <component :is="detailFor(stage.kind)" :stage="stage" />
    </div>
  </Panel>
</template>

<style scoped>
/* de-chrome the Toolbar so the Panel header reads as one surface */
.stage-card :deep(.p-toolbar) {
  background: transparent;
  border: 0;
  padding: 0;
  width: 100%;
  cursor: pointer;
}

.stage-icon {
  margin-right: 0.5rem;
  font-size: 1.1rem;
}

.stage-card__title {
  font-weight: 600;
}

.stage-card__when {
  color: var(--color-text);
  opacity: 0.7;
  font-size: 0.875rem;
  margin-right: 0.5rem;
  white-space: nowrap;
}

/* 80vh internal scroll on the Panel's content region */
.stage-card :deep(.p-panel-content) {
  max-height: 80vh;
  overflow-y: auto;
}
</style>
