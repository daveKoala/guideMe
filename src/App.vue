<script setup lang="ts">
import { provide } from 'vue'
import { storeToRefs } from 'pinia'
import { RouterView } from 'vue-router'
import BasicLayout from './templates/BasicLayout.vue'
import { useUiStore } from './stores/ui'
import { modeKey } from './components/timeline/tripContext'

// App-wide edit/read mode, provided once so every page + the nav control share it.
const { mode } = storeToRefs(useUiStore())
provide(modeKey, mode)

const refreshStamp = Date.now()

const layoutByName = {
  BasicLayout,
} as const

function getLayout(templateName: unknown) {
  if (typeof templateName !== 'string') {
    return layoutByName.BasicLayout
  }

  return layoutByName[templateName as keyof typeof layoutByName] ?? layoutByName.BasicLayout
}
</script>

<template>
  <RouterView v-slot="{ Component, route }">
    <component :is="getLayout(route.meta.template)">
      <template #main>
        <component :is="Component" :key="`${route.fullPath}-${refreshStamp}`" />
      </template>
    </component>
  </RouterView>
</template>

<style scoped></style>
