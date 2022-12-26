const routes = [
  {
    path: "/",
    name: "home",
    component: () => import("@/views/HomeView"),
  },
  {
    path: "/podcasts",
    name: "podcasts",
    component: () => import("@/views/podcasts/PodcastsListView"),
  },
  {
    path: "/podcasts/latest",
    name: "latest-podcasts",
    component: () => import("@/views/podcasts/LatestPodcastItemsView"),
  },
  {
    path: "/podcasts/:id",
    name: "single-podcast",
    component: () => import("@/views/podcasts/SinglePodcastView"),
  },
  {
    path: "/podcasts/:id/items/:itemId",
    name: "single-item",
    component: () => import("@/views/podcasts/SinglePodcastItemView"),
  },
];

export default routes;
