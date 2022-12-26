import { formatDistanceToNow, formatISO } from "date-fns";

export function formatDate(dateString) {
  let date = new Date(dateString);
  let distanceFromNow = Math.floor((new Date() - date) / (1000 * 60 * 60 * 24));
  if (distanceFromNow > 7) return formatISO(date, { representation: "date" });
  return formatDistanceToNow(date) + " ago";
}
