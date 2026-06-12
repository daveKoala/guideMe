<script setup lang="ts">
import { RouterLink } from 'vue-router'
import Menubar from 'primevue/menubar'
import SelectButton from 'primevue/selectbutton'
import type { MenuItem } from 'primevue/menuitem'
import { useUiStore } from '@/stores/ui'
import type { Mode } from '@/components/timeline/tripContext'

const items: MenuItem[] = [
	{ label: 'Home', route: '/', icon: 'pi pi-home' },
	{ label: 'About', route: '/about', icon: 'pi pi-info-circle' },
	{ label: 'People', route: '/people', icon: 'pi pi-users' },
	{ label: 'Trips', route: '/trips', icon: 'pi pi-map' },
]

const ui = useUiStore()
const modeOptions: { label: string; value: Mode }[] = [
	{ label: 'Edit', value: 'edit' },
	{ label: 'Read', value: 'read' },
]
</script>

<template>
	<div class="layout">
		<header class="topbar">
			<div class="topbar-inner">
				<Menubar :model="items" class="nav" :breakpoint="'768px'" aria-label="Primary">
					<template #start>
						<RouterLink to="/" class="brand">
							<i class="pi pi-compass brand-icon" aria-hidden="true" />
							<span class="brand-text">Guide Me</span>
						</RouterLink>
					</template>
					<template #item="{ item, props }">
						<RouterLink v-slot="{ href, navigate, isActive }" :to="item.route" custom>
							<a
								:href="href"
								v-bind="props.action"
								:class="{ 'nav-link-active': isActive }"
								@click="navigate"
							>
								<span v-if="item.icon" :class="item.icon" class="nav-icon" aria-hidden="true" />
								<span class="nav-label">{{ item.label }}</span>
							</a>
						</RouterLink>
					</template>
					<template #end>
						<SelectButton
							:model-value="ui.mode"
							:options="modeOptions"
							option-label="label"
							option-value="value"
							:allow-empty="false"
							size="small"
							aria-label="Edit or read mode"
							@update:model-value="ui.setMode($event)"
						/>
					</template>
				</Menubar>
			</div>
		</header>

		<main class="main">
			<div class="content">
				<slot name="main">
					<slot />
				</slot>
			</div>
		</main>
	</div>
</template>

<style scoped>
.layout {
	min-height: 100dvh;
	background: var(--color-bg);
	color: var(--color-text);
}

.topbar {
	position: fixed;
	top: 0;
	right: 0;
	left: 0;
	height: var(--topbar-height);
	border-bottom: 1px solid var(--color-border);
	background: var(--color-nav-bg);
	backdrop-filter: blur(8px);
	z-index: 10;
}

.topbar-inner {
	width: min(100% - (var(--space-4) * 2), var(--layout-max-width));
	height: 100%;
	margin: 0 auto;
	display: flex;
	align-items: center;
}

.nav {
	width: 100%;
	background: transparent;
	border: none;
	border-radius: 0;
	padding: 0;
	gap: var(--space-4);
}

.brand {
	display: inline-flex;
	align-items: center;
	gap: 8px;
	color: var(--color-text);
	text-decoration: none;
	font-weight: 700;
	font-size: 1.125rem;
}

.brand-icon {
	color: var(--color-link);
	font-size: 1.25rem;
}

.nav :deep(a) {
	display: flex;
	align-items: center;
	gap: 8px;
	color: var(--color-link);
	text-decoration: none;
	font-weight: 600;
}

.nav-icon {
	font-size: 0.95rem;
}

.nav :deep(a.nav-link-active) {
	text-decoration: underline;
	text-underline-offset: 3px;
}

/* Responsive popup (below breakpoint) gets a solid panel + spacing */
.nav :deep(.p-menubar-root-list) {
	gap: 4px;
}

.nav :deep(.p-menubar-mobile-active .p-menubar-root-list) {
	background: var(--color-surface);
	border: 1px solid var(--color-border);
	border-radius: var(--radius-md);
	padding: 8px;
	box-shadow: 0 8px 24px rgb(0 0 0 / 0.08);
}

.main {
	padding: calc(var(--topbar-height) + var(--space-6)) var(--space-4) var(--space-6);
}

.content {
	width: min(100%, var(--layout-max-width));
	margin: 0 auto;
	background: var(--color-surface);
	border: 1px solid var(--color-border);
	border-radius: var(--radius-md);
	padding: var(--space-6);
}

/* Reclaim width on hand-held screens (360px floor) */
@media (max-width: 480px) {
	.main {
		padding: calc(var(--topbar-height) + var(--space-4)) var(--space-4) var(--space-4);
	}

	.content {
		padding: var(--space-4);
	}
}
</style>
