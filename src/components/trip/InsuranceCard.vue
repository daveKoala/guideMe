<script setup lang="ts">
import { inject, ref } from 'vue'
import Panel from 'primevue/panel'
import Button from 'primevue/button'
import InputText from 'primevue/inputtext'
import MultiSelect from 'primevue/multiselect'
import DocumentSlot from './DocumentSlot.vue'
import { modeKey, partyKey, emptyParty, usePartyPeople, type Mode } from '@/components/timeline/tripContext'
import { useAccountStore } from '@/stores/account'
import type { Insurance } from '@/types/trip'

const props = withDefaults(defineProps<{ policies: Insurance[]; defaultOpen?: boolean }>(), {
  defaultOpen: false,
})

const mode = inject(modeKey, ref<Mode>('read'))
const party = inject(partyKey, emptyParty)
const account = useAccountStore()

// Collapsed by default. Only the builder (edit mode) auto-opens it when the trip
// flagged insurance for setup — Home stays collapsed and calm regardless.
const collapsed = ref(!(props.defaultOpen && mode.value === 'edit'))

// Only trip members are selectable as "covered people".
const people = usePartyPeople(party)

function nameFor(id: string): string {
  return account.getPerson(id)?.name || 'Unknown'
}
function coveredNames(ids: string[]): string {
  return ids.length ? ids.map(nameFor).join(', ') : '—'
}

function addPolicy() {
  props.policies.push({ id: `ins_${crypto.randomUUID()}`, covers: [], medical: {} })
}
function removePolicy(id: string) {
  const i = props.policies.findIndex((p) => p.id === id)
  if (i !== -1) props.policies.splice(i, 1)
}
</script>

<template>
  <Panel toggleable v-model:collapsed="collapsed">
    <template #header>
      <span class="ov-title"><span class="ov-icon">🛡️</span> Travel insurance</span>
    </template>
    <template #icons>
      <Button
        v-if="mode === 'edit'"
        label="Add policy"
        icon="pi pi-plus"
        size="small"
        @click="addPolicy"
      />
    </template>

    <ul class="policies">
      <li v-for="policy in policies" :key="policy.id" class="policy">
        <div class="field">
          <label :for="mode === 'edit' ? `${policy.id}-covers` : undefined">Covered people</label>
          <MultiSelect
            v-if="mode === 'edit'"
            :input-id="`${policy.id}-covers`"
            v-model="policy.covers"
            :options="people"
            option-label="name"
            option-value="id"
            placeholder="Who's covered?"
            display="chip"
            fluid
          />
          <span v-else class="value">{{ coveredNames(policy.covers) }}</span>
        </div>

        <div class="field">
          <label :for="mode === 'edit' ? `${policy.id}-policy` : undefined">Policy number</label>
          <InputText
            v-if="mode === 'edit'"
            :id="`${policy.id}-policy`"
            v-model="policy.policy_number"
            fluid
          />
          <span v-else class="value">{{ policy.policy_number || '—' }}</span>
        </div>

        <div class="field">
          <label :for="mode === 'edit' ? `${policy.id}-emergency` : undefined">Emergency contact</label>
          <InputText
            v-if="mode === 'edit'"
            :id="`${policy.id}-emergency`"
            v-model="policy.emergency_contact"
            fluid
          />
          <a
            v-else-if="policy.emergency_contact"
            class="value"
            :href="`tel:${policy.emergency_contact}`"
          >
            {{ policy.emergency_contact }}
          </a>
          <span v-else class="value">—</span>
        </div>

        <div class="field">
          <label :for="mode === 'edit' ? `${policy.id}-url` : undefined">Account URL</label>
          <InputText
            v-if="mode === 'edit'"
            :id="`${policy.id}-url`"
            v-model="policy.account_url"
            fluid
          />
          <a
            v-else-if="policy.account_url"
            class="value"
            :href="policy.account_url"
            target="_blank"
            rel="noopener"
          >
            Open account ↗
          </a>
          <span v-else class="value">—</span>
        </div>

        <div class="policy__docs">
          <DocumentSlot label="Policy document" />
        </div>

        <fieldset class="medical">
          <legend>Medical assistance</legend>

          <div class="field">
            <label :for="mode === 'edit' ? `${policy.id}-med-id` : undefined">Assist ID</label>
            <InputText
              v-if="mode === 'edit'"
              :id="`${policy.id}-med-id`"
              v-model="policy.medical.assist_id"
              fluid
            />
            <span v-else class="value">{{ policy.medical.assist_id || '—' }}</span>
          </div>

          <div class="field">
            <label :for="mode === 'edit' ? `${policy.id}-med-phone` : undefined">Phone</label>
            <InputText
              v-if="mode === 'edit'"
              :id="`${policy.id}-med-phone`"
              v-model="policy.medical.phone"
              fluid
            />
            <a v-else-if="policy.medical.phone" class="value" :href="`tel:${policy.medical.phone}`">
              {{ policy.medical.phone }}
            </a>
            <span v-else class="value">—</span>
          </div>

          <div class="field">
            <label :for="mode === 'edit' ? `${policy.id}-med-url` : undefined">URL</label>
            <InputText
              v-if="mode === 'edit'"
              :id="`${policy.id}-med-url`"
              v-model="policy.medical.url"
              fluid
            />
            <a
              v-else-if="policy.medical.url"
              class="value"
              :href="policy.medical.url"
              target="_blank"
              rel="noopener"
            >
              Open ↗
            </a>
            <span v-else class="value">—</span>
          </div>
        </fieldset>

        <div v-if="mode === 'edit'" class="policy__actions">
          <Button
            label="Remove policy"
            icon="pi pi-trash"
            size="small"
            text
            severity="danger"
            @click="removePolicy(policy.id)"
          />
        </div>
      </li>
    </ul>

    <p v-if="!policies.length" class="empty">No insurance policies yet.</p>
  </Panel>
</template>

<style scoped>
.policies {
  list-style: none;
  margin: 0;
  padding: 0;
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.policy {
  border: 1px solid var(--color-border);
  border-radius: var(--radius-md);
  padding: 1rem;
  display: flex;
  flex-direction: column;
  gap: 0.85rem;
}

.field {
  display: flex;
  flex-direction: column;
  gap: 0.3rem;
}

.field label {
  font-weight: 600;
  font-size: 0.875rem;
}

.value {
  color: var(--color-text);
}

.policy__docs {
  border-top: 1px solid var(--color-border);
  padding-top: 0.85rem;
}

.medical {
  border: 1px solid var(--color-border);
  border-radius: var(--radius-md);
  padding: 0.75rem 1rem 1rem;
  margin: 0;
  display: flex;
  flex-direction: column;
  gap: 0.85rem;
}

.medical legend {
  font-weight: 700;
  font-size: 0.8rem;
  text-transform: uppercase;
  letter-spacing: 0.04em;
  opacity: 0.7;
  padding: 0 0.4rem;
}

.policy__actions {
  display: flex;
  justify-content: flex-end;
}

.empty {
  opacity: 0.6;
}

.ov-icon {
  margin-right: 0.4rem;
}
</style>
