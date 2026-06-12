import { createRouter, createWebHistory } from 'vue-router'


const routes = [
  {
    path: '/',
    name: 'home',
    component: () => import('../views/HomePage.vue'),
    meta: {
      template: 'BasicLayout',
      title: 'Home',
    },
  },
  {
    path: '/landing-page',
    name: 'landing-page',
    component: () => import('../views/LandingPage.vue'),
    meta: {
      template: 'BasicLayout',
      title: 'Landing Page',
    },
  },
  {
    path: '/people',
    name: 'people',
    component: () => import('../views/PeoplePage.vue'),
    meta: {
      template: 'BasicLayout',
      title: 'People',
    },
  },
  {
    path: '/trips',
    name: 'trips',
    component: () => import('../views/TripsPage.vue'),
    meta: {
      template: 'BasicLayout',
      title: 'Trips',
    },
  },
  {
    path: '/trips/new',
    name: 'new-trip',
    component: () => import('../views/NewTripPage.vue'),
    meta: {
      template: 'BasicLayout',
      title: 'New trip',
    },
  },
  {
    path: '/about',
    name: 'about',
    component: () => import('../views/LandingPage.vue'),
    meta: {
      template: 'BasicLayout',
      title: 'About',
    },
  },
  // default redirect to home page
  {
    path: '/:pathMatch(.*)*', 
    redirect: '/',
  },
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
})

const DEFAULT_TITLE = 'Guide Me'

router.afterEach((to) => {
  const title = to.meta.title as string | undefined
  document.title = title ? `${title} | ${DEFAULT_TITLE}` : DEFAULT_TITLE
})

export default router
