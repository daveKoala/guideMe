<script setup lang="ts">
import { RouterLink } from 'vue-router'
</script>

<template>
	<div class="layout">
		<header class="topbar">
			<div class="topbar-inner">
				<nav class="nav" aria-label="Primary">
					<RouterLink to="/">Home</RouterLink>
					<RouterLink to="/about">About</RouterLink>
				</nav>
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
	display: flex;
	gap: var(--space-6);
}

.nav :deep(a) {
	color: var(--color-link);
	text-decoration: none;
	font-weight: 600;
}

.nav :deep(a.router-link-active) {
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
</style>
