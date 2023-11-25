import { formatDuration, intervalToDuration } from 'date-fns'

export function getTimePlayed(howlerData, podcastProgress) {
  if (!howlerData)
    return '00:00'
  const duration = intervalToDuration({
    start: 0,
    end: podcastProgress * 1000,
  })
  const zeroPad = num => String(num).padStart(2, '0')

  return formatDuration(duration, {
    format: getFormat(podcastProgress),
    zero: true,
    delimiter: ':',
    locale: {
      formatDistance: (_token, count) => zeroPad(count),
    },
  })
}

export function getTotalTime(podcastData) {
  if (!podcastData)
    return 0

  const duration = intervalToDuration({
    start: 0,
    end: podcastData.Duration * 1000,
  })

  const zeroPad = num => String(num).padStart(2, '0')

  return formatDuration(duration, {
    format: getFormat(podcastData.Duration),
    zero: true,
    delimiter: ':',
    locale: {
      formatDistance: (_token, count) => zeroPad(count),
    },
  })
}

function getFormat(duration) {
  return duration > 3600
    ? ['hours', 'minutes', 'seconds']
    : ['minutes', 'seconds']
}
