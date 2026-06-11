<script setup lang="ts">
import { RouterLink } from 'vue-router'
import Menubar from 'primevue/menubar'
import type { MenuItem } from 'primevue/menuitem'

const items: MenuItem[] = [
	{ label: 'Home', route: '/' },
	{ label: 'About', route: '/about' },
]
</script>

<template>
	<div class="layout">
		<header class="topbar">
			<div class="topbar-inner">
				<Menubar :model="items" class="nav" aria-label="Primary">
					<template #item="{ item, props }">
						<RouterLink v-slot="{ href, navigate, isActive }" :to="item.route" custom>
							<a
								:href="href"
								v-bind="props.action"
								:class="{ 'nav-link-active': isActive }"
								@click="navigate"
							>
								{{ item.label }}
							</a>
						</RouterLink>
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
	background: transparent;
	border: none;
	border-radius: 0;
	padding: 0;
}

.nav :deep(a) {
	color: var(--color-link);
	text-decoration: none;
	font-weight: 600;
}

.nav :deep(a.nav-link-active) {
	text-decoration: underline;
	text-underline-offset: 3px;
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
