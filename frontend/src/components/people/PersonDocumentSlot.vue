<script setup lang="ts">
import { computed, inject, ref } from 'vue'
import Button from 'primevue/button'
import DocPreview from '@/components/common/DocPreview.vue'
import { useAccountStore } from '@/stores/account'
import { modeKey, type Mode } from '@/components/timeline/tripContext'
import type { PersonDocKind } from '@/types/person'

const props = defineProps<{ personId: string; kind: PersonDocKind; label: string }>()

const mode = inject(modeKey, ref<Mode>('read'))
const store = useAccountStore()
const input = ref<HTMLInputElement | null>(null)

const doc = computed(() => store.getPerson(props.personId)?.documents[props.kind])

const MAX_BYTES = 1_000_000 // ~1MB soft cap to stay within localStorage quota

function pick() {
  input.value?.click()
}

function onChange(e: Event) {
  const f = (e.target as HTMLInputElement).files?.[0]
  if (!f) return
  if (f.size > MAX_BYTES) {
    window.alert('File too large for this prototype — keep it under ~1MB.')
    if (input.value) input.value.value = ''
    return
  }
  const reader = new FileReader()
  reader.onload = () => {
    store.setDocument(props.personId, props.kind, {
      fileName: f.name,
      dataUrl: String(reader.result),
      uploadedAt: new Date().toISOString(),
    })
  }
  reader.readAsDataURL(f)
}

function remove() {
  store.removeDocument(props.personId, props.kind)
  if (input.value) input.value.value = ''
}
</script>

<template>
  <div class="doc-slot">
    <span class="doc-label">{{ label }}</span>

    <div class="doc-actions">
      <template v-if="doc">
        <DocPreview :file-name="doc.fileName" :data-url="doc.dataUrl" />
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
        <Button v-if="mode === 'edit'" label="Upload" icon="pi pi-upload" size="small" @click="pick" />
        <span v-else class="doc-empty">Not provided</span>
      </template>
    </div>

    <input ref="input" type="file" accept="image/*,application/pdf" hidden @change="onChange" />
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
