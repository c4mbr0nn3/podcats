import DOMPurify from 'dompurify'
import { marked } from 'marked'

const options = {
  mangle: false,
  headerIds: false,
}
marked.setOptions(options)

function markdownToHtml(text) {
  if (text === undefined || text === null)
    return
  const dirty = marked(text.toString())
  const clean = DOMPurify.sanitize(dirty)
  return clean
}

export { markdownToHtml }
