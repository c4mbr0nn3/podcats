import DOMPurify from "dompurify";
import { marked } from "marked";

function markdownToHtml(text) {
  if (text == undefined || text == null) return;
  let dirty = marked(text.toString());
  let clean = DOMPurify.sanitize(dirty);
  return clean;
}

export { markdownToHtml };
