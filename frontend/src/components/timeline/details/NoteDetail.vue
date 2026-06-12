<script setup lang="ts">
import StageFields from '../StageFields.vue'
import type { Stage } from '@/types/stage'

defineProps<{ stage: Stage }>()
</script>

<template>
  <StageFields :stage="stage" />
</template>
