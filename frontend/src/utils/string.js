export function stringContainsAnyOf(str, arr) {
  return arr
    .filter((el) => el.length > 0)
    .some((el) => {
      return str.toLowerCase().includes(el.toLowerCase());
    });
}
