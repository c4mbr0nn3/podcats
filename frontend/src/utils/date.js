import { formatDistanceToNow, formatISO } from 'date-fns'

export function formatDate(dateString) {
  const date = new Date(dateString)
  const distanceFromNow = Math.floor((new Date() - date) / (1000 * 60 * 60 * 24))
  if (distanceFromNow > 7)
    return formatISO(date, { representation: 'date' })
  return `${formatDistanceToNow(date)} ago`
}

export function formatDateToIso(dateString) {
  const date = new Date(dateString)
  return formatISO(date, { representation: 'date' })
}
