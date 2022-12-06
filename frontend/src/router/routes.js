import HomeView from "@/views/HomeView";
import AboutView from "@/views/AboutView";

const routes = [
  { path: "/", name: "home", component: HomeView },
  { path: "/about", name: "about", component: AboutView },
];

export default routes;
