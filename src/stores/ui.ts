import { ref, watch } from 'vue'
import { defineStore } from 'pinia'
import type { Mode } from '@/components/timeline/tripContext'

const STORAGE_KEY = 'guide-me:ui:mode'

function loadMode(): Mode {
  try {
    const raw = localStorage.getItem(STORAGE_KEY)
    return raw === 'read' || raw === 'edit' ? raw : 'edit'
  } catch {
    return 'edit'
  }
}

/** App-wide edit/read switch. Owner-edit vs read-only share, shared across every view. */
export const useUiStore = defineStore('ui', () => {
  const mode = ref<Mode>(loadMode())

  watch(mode, (value) => {
    try {
      localStorage.setItem(STORAGE_KEY, value)
    } catch {
      // ignore — keep working in-memory
    }
  })

  function setMode(value: Mode) {
    mode.value = value
  }

  function toggle() {
    mode.value = mode.value === 'edit' ? 'read' : 'edit'
  }

  return { mode, setMode, toggle }
})
