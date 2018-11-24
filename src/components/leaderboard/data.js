/**
 * Fetches data from the database.
 */
async function getData() {
  const resultsRes = await fetch("https://opensystembench.com/api/results");
  const usersRes = await fetch("https://opensystembench.com/api/users");
  const specsRes = await fetch("https://opensystembench.com/api/specs");
  // const resultsRes = await fetch("http://localhost:8080/api/results");
  // const usersRes = await fetch("http://localhost:8080/api/users");
  // const specsRes = await fetch("http://localhost:8080/api/specs");
  const results = await resultsRes.json();
  const users = await usersRes.json();
  const specs = await specsRes.json();

  const d = results.map(result => {
    const total = result.scores.find(score => score.name === "Total");
    console.log(result.scores)
    return {
      id: result.ID,
      totalTime: total.time,
      totalScore: total.score,
      user: users.find(user => user.ID === result.UserID).Name,
      scores: result.scores,
      specs: specs.find(spec => spec.ResultID === result.ID)
    };
  });

  const sorted = d.sort((a, b) => {
    if (a.totalScore < b.totalScore) return 1;
    else if (b.totalScore < a.totalScore) return -1;
    else return 0;
  });

  const ranked = sorted.map((result, index) => {
    result.rank = index + 1;
    return result;
  });

  return ranked;
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
  const stabilizedThis = array.map((el, index) => [el, index]);
  stabilizedThis.sort((a, b) => {
    const order = cmp(a[0], b[0]);
    if (order !== 0) return order;
    return a[1] - b[1];
  });
  return stabilizedThis.map(el => el[0]);
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
