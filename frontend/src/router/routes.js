import HomeView from "@/views/HomeView";
import PodcastsListView from "@/views/podcasts/PodcastsListView";
import SinglePodcastView from "@/views/podcasts/SinglePodcastView";
import SinglePodcastItemView from "@/views/podcasts/SinglePodcastItemView";

const routes = [
  { path: "/", name: "home", component: HomeView },
  { path: "/podcasts", name: "podcasts", component: PodcastsListView },
  {
    path: "/podcasts/:id",
    name: "single-podcast",
    component: SinglePodcastView,
  },
  {
    path: "/podcasts/:id/items/:itemId",
    name: "single-item",
    component: SinglePodcastItemView,
  },
];

export default routes;
