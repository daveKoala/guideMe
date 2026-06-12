<script setup lang="ts">
import { computed, inject } from 'vue'
import Accordion from 'primevue/accordion'
import AccordionPanel from 'primevue/accordionpanel'
import AccordionHeader from 'primevue/accordionheader'
import AccordionContent from 'primevue/accordioncontent'
import FileUpload from 'primevue/fileupload'
import DocPreview from '@/components/common/DocPreview.vue'
import StageFields from '../StageFields.vue'
import { partyKey, emptyParty, usePartyPeople } from '../tripContext'
import type { Stage } from '@/types/stage'

const props = defineProps<{ stage: Stage }>()

const party = inject(partyKey, emptyParty)
const people = usePartyPeople(party)

// pair each person with their boarding pass (if any) so the template narrows cleanly
const rows = computed(() =>
  people.value.map((p) => ({ p, pass: props.stage.boarding_passes?.[p.id] })),
)
</script>

<template>
  <Accordion :value="['0']" multiple>
    <AccordionPanel value="0">
      <AccordionHeader>Flight details</AccordionHeader>
      <AccordionContent>
        <StageFields :stage="stage" />
      </AccordionContent>
    </AccordionPanel>

    <AccordionPanel value="1">
      <AccordionHeader>Boarding passes ({{ party.passengers.length }})</AccordionHeader>
      <AccordionContent>
        <ul class="passes">
          <li v-for="row in rows" :key="row.p.id" class="pass-row">
            <span class="pass-name">
              {{ row.p.name }}
              <small class="pass-type">{{ row.p.type }}</small>
            </span>

            <DocPreview v-if="row.pass" :file-name="row.pass.fileName" :data-url="row.pass.dataUrl" />
            <FileUpload
              v-else
              mode="basic"
              :auto="false"
              custom-upload
              choose-label="Upload"
              accept="image/*,application/pdf"
            />
          </li>
        </ul>
      </AccordionContent>
    </AccordionPanel>
  </Accordion>
</template>

<style scoped>
.passes {
  list-style: none;
  margin: 0;
  padding: 0;
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.pass-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 1rem;
}

.pass-name {
  font-weight: 600;
}

.pass-type {
  font-weight: 400;
  opacity: 0.6;
  margin-left: 0.4rem;
}
</style>
