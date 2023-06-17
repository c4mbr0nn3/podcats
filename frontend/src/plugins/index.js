/**
 * plugins/index.js
 *
 * Automatically included in `./src/main.js`
 */

import { loadFonts } from "@/plugins/webfontloader";
import vuetify from "@/plugins/vuetify";
import { VueQueryPlugin } from "@tanstack/vue-query";

export function registerPlugins(app) {
  loadFonts();
  app.use(vuetify);
  app.use(VueQueryPlugin);
}
