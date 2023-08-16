import { GenericApi } from "@/api";
import { unref } from "vue";

export async function search(searchString) {
  const baseUrl = "https://itunes.apple.com/search";
  const params = unref(searchString).split(" ").join("+");
  const url = `${baseUrl}?term=${params}`;
  const { data } = await GenericApi.get(url);
  return data.results.filter((item) => item.kind === "podcast");
}
