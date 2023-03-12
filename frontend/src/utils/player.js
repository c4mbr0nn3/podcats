import { intervalToDuration, formatDuration } from "date-fns";

export function getTimePlayed(howlerData, podcastProgress) {
  if (!howlerData) return "00:00";
  const duration = intervalToDuration({
    start: 0,
    end: podcastProgress * 1000,
  });
  const zeroPad = (num) => String(num).padStart(2, "0");

  return formatDuration(duration, {
    format: ["minutes", "seconds"],
    // format: ["hours", "minutes", "seconds"],
    zero: true,
    delimiter: ":",
    locale: {
      formatDistance: (_token, count) => zeroPad(count),
    },
  });
}

export function getTotalTime(podcastData) {
  if (!podcastData) return 0;

  const duration = intervalToDuration({
    start: 0,
    end: podcastData.Duration * 1000,
  });

  const zeroPad = (num) => String(num).padStart(2, "0");

  return formatDuration(duration, {
    format: ["minutes", "seconds"],
    // format: ["hours", "minutes", "seconds"],
    zero: true,
    delimiter: ":",
    locale: {
      formatDistance: (_token, count) => zeroPad(count),
    },
  });
}
