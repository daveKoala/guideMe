<script setup lang="ts">
import Accordion from 'primevue/accordion'
import AccordionPanel from 'primevue/accordionpanel'
import AccordionHeader from 'primevue/accordionheader'
import AccordionContent from 'primevue/accordioncontent'
import FileUpload from 'primevue/fileupload'
import { inject, ref } from 'vue'
import StageFields from '../StageFields.vue'
import { modeKey, type Mode } from '../tripContext'
import type { Stage } from '@/types/stage'

defineProps<{ stage: Stage }>()

const mode = inject(modeKey, ref<Mode>('read'))
</script>

<template>
  <Accordion :value="['0']" multiple>
    <AccordionPanel value="0">
      <AccordionHeader>Booking details</AccordionHeader>
      <AccordionContent>
        <StageFields :stage="stage" />
      </AccordionContent>
    </AccordionPanel>

    <AccordionPanel value="1">
      <AccordionHeader>Documents</AccordionHeader>
      <AccordionContent>
        <FileUpload
          v-if="mode === 'edit'"
          mode="basic"
          :auto="false"
          custom-upload
          choose-label="Add booking confirmation"
          accept="image/*,application/pdf"
        />
        <span v-else class="docs-empty">No documents yet</span>
      </AccordionContent>
    </AccordionPanel>
  </Accordion>
</template>

<style scoped>
.docs-empty {
  opacity: 0.5;
  font-size: 0.875rem;
}
</style>
