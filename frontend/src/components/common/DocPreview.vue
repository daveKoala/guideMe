<script setup lang="ts">
import { computed } from 'vue'
import Image from 'primevue/image'
import Button from 'primevue/button'

/**
 * View + download a document, PWA-safe (no new tab).
 * Images render as a thumbnail that opens a fullscreen, zoomable overlay on click.
 * Non-images (PDF) show a file icon + download.
 */
const props = defineProps<{ fileName: string; dataUrl: string }>()

const isImage = computed(() => {
  if (props.dataUrl.startsWith('data:image')) return true
  return /\.(png|jpe?g|webp|avif|gif|svg)$/i.test(props.fileName)
})
</script>

<template>
  <span class="doc-preview">
    <Image
      v-if="isImage"
      :src="dataUrl"
      :alt="fileName"
      preview
      image-class="doc-preview__thumb"
    />
    <i v-else class="pi pi-file doc-preview__icon" aria-hidden="true" />

    <a :href="dataUrl" :download="fileName">
      <Button icon="pi pi-download" size="small" text :aria-label="`Download ${fileName}`" />
    </a>
  </span>
</template>

<style scoped>
.doc-preview {
  display: inline-flex;
  align-items: center;
  gap: 0.25rem;
}

:deep(.doc-preview__thumb) {
  display: block;
  width: 44px;
  height: 44px;
  object-fit: cover;
  border-radius: 6px;
  border: 1px solid var(--color-border);
  cursor: zoom-in;
}

.doc-preview__icon {
  font-size: 1.5rem;
  opacity: 0.7;
}
</style>
