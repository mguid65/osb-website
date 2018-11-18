/**
 * Fetches the data from the leaderboard's scores table.
 * @returns {Object[]} The JSON representation of scores data.
 */
async function getData() {
  const res = await fetch("https://opensystembench.com/results");
  return res.json();
}

/**
 * Compares two elements in descending order.
 * @param {*} a An element.
 * @param {*} b Another element.
 * @param {string} orderBy The property to order the elements by.
 * @returns {number}
 */
function desc(a, b, orderBy) {
  if (b[orderBy] < a[orderBy]) return -1;
  if (b[orderBy] > a[orderBy]) return 1;
  return 0;
}

/**
 * Sorts an array while retaining the order of elements.
 * @param {array} array The array to sort.
 * @param {Function} cmp A callback for comparing elements.
 * @returns {array} The sorted elements.
 */
function stableSort(array, cmp) {
  let stabilized = array.map((el, index) => [el, index]);
  stabilized.sort((a, b) => {
    const order = cmp(a[0], b[0]);
    if (order !== 0) return order;
    return a[1] - b[1];
  });
  return stabilized.map(el => el[0]);
}

/**
 * Compare two elements for sorting in a specific order.
 * @param {string} order The sort direction.
 * @param {string} orderBy The propperty to order by.
 * @return {number}
 */
function getSorting(order, orderBy) {
  return order === "desc"
    ? (a, b) => desc(a, b, orderBy)
    : (a, b) => -desc(a, b, orderBy);
}

/**
 * Sleep for the specified amount of time.
 * @param {number} time The time to sleep in milliseconds.
 * @returns {Promise}
 */
function sleep(time) {
  return new Promise(resolve => setTimeout(resolve, time));
}

export { sleep, getData, desc, stableSort, getSorting };
