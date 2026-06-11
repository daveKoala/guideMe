<script setup lang="ts">
import { computed, inject, ref, onBeforeUnmount } from 'vue'
import Button from 'primevue/button'
import DocPreview from '@/components/common/DocPreview.vue'
import { modeKey, type Mode } from '@/components/timeline/tripContext'

defineProps<{ label: string }>()

const mode = inject(modeKey, ref<Mode>('read'))

// In-memory only — no backend yet. Resets on refresh.
const file = ref<File | null>(null)
const input = ref<HTMLInputElement | null>(null)

const objectUrl = computed(() => (file.value ? URL.createObjectURL(file.value) : ''))
onBeforeUnmount(() => {
  if (objectUrl.value) URL.revokeObjectURL(objectUrl.value)
})

function pick() {
  input.value?.click()
}
function onChange(e: Event) {
  file.value = (e.target as HTMLInputElement).files?.[0] ?? null
}
function remove() {
  file.value = null
  if (input.value) input.value.value = ''
}
</script>

<template>
  <div class="doc-slot">
    <span class="doc-label">{{ label }}</span>

    <div class="doc-actions">
      <template v-if="file">
        <DocPreview :file-name="file.name" :data-url="objectUrl" />
        <Button
          v-if="mode === 'edit'"
          icon="pi pi-times"
          size="small"
          text
          severity="danger"
          aria-label="Remove document"
          @click="remove"
        />
      </template>

      <template v-else>
        <Button
          v-if="mode === 'edit'"
          label="Upload"
          icon="pi pi-upload"
          size="small"
          @click="pick"
        />
        <span v-else class="doc-empty">Not provided</span>
      </template>
    </div>

    <input
      ref="input"
      type="file"
      accept="image/*,application/pdf"
      hidden
      @change="onChange"
    />
  </div>
</template>

<style scoped>
.doc-slot {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 0.75rem;
}

.doc-label {
  font-weight: 500;
}

.doc-actions {
  display: flex;
  align-items: center;
  gap: 0.25rem;
}

.doc-empty {
  opacity: 0.55;
  font-size: 0.875rem;
}
</style>
