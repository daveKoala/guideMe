<script setup lang="ts">
import Accordion from 'primevue/accordion'
import AccordionPanel from 'primevue/accordionpanel'
import AccordionHeader from 'primevue/accordionheader'
import AccordionContent from 'primevue/accordioncontent'
import FileUpload from 'primevue/fileupload'
import StageFields from '../StageFields.vue'
import type { Stage } from '@/types/stage'

defineProps<{ stage: Stage }>()
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
          mode="basic"
          :auto="false"
          custom-upload
          choose-label="Add booking confirmation"
          accept="image/*,application/pdf"
        />
      </AccordionContent>
    </AccordionPanel>
  </Accordion>
</template>
