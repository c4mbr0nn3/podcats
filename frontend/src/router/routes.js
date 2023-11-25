const routes = [
  {
    path: '/',
    redirect: { name: 'login' },
  },
  {
    path: '/login',
    name: 'login',
    component: () => import('@/views/login/LoginView'),
    meta: {
      requiresFillHeight: true,
    },
  },
  {
    path: '/login/change-password',
    name: 'change-password',
    component: () => import('@/views/login/LoginChangePassword'),
    meta: {
      requiresFillHeight: true,
      requiresAuth: true,
    },
  },
  {
    path: '/home',
    name: 'home',
    component: () => import('@/views/HomeView'),
    meta: {
      requiresFillHeight: true,
      requiresAuth: true,
      requiresNavbar: true,
    },
  },
  {
    path: '/podcasts',
    name: 'podcasts',
    component: () => import('@/views/podcasts/PodcastsListView'),
    meta: {
      requiresAuth: true,
      requiresNavbar: true,
      name: 'My Podcasts',
    },
  },
  {
    path: '/podcasts/latest',
    name: 'latest-podcasts',
    component: () => import('@/views/podcasts/LatestPodcastItemsView'),
    meta: {
      requiresAuth: true,
      requiresNavbar: true,
      name: 'Latest Podcasts',
    },
  },
  {
    path: '/podcasts/favorites',
    name: 'favorites-podcasts',
    component: () => import('@/views/podcasts/FavouritesPodcastItemsView'),
    meta: {
      requiresAuth: true,
      requiresNavbar: true,
      name: 'Favorites Podcasts',
    },
  },
  {
    path: '/podcasts/:id',
    name: 'single-podcast',
    component: () => import('@/views/podcasts/SinglePodcastView'),
    meta: {
      requiresAuth: true,
      requiresNavbar: true,
    },
  },
  {
    path: '/podcasts/:id/items/:itemId',
    name: 'single-item',
    component: () => import('@/views/podcasts/SinglePodcastItemView'),
    meta: {
      requiresAuth: true,
      requiresNavbar: true,
    },
  },
  {
    path: '/admin/settings',
    name: 'settings-admin',
    component: () => import('@/views/settings/SettingsAdminView'),
    meta: {
      requiresAuth: true,
      requiresNavbar: true,
      requiresRoot: true,
      name: 'Settings',
    },
  },
]

export default routes
