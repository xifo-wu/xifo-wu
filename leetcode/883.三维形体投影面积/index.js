/**
 * @param {number[][]} grid
 * @return {number}
 */
var projectionArea = function (grid) {
  const gLen = grid.length;

  let heng = 0;
  let shu = 0;
  let di = gLen * gLen;
  for (let i = 0; i < gLen; ++i) {
    let henMax = 0;
    for (let j = 0; j < gLen; ++j) {
      if (grid[i][j] === 0) {
        di--;
      }

      if (grid[i][j] >= henMax) {
        henMax = grid[i][j];
      }

      if (i == 0) {
        let shuMax = 0;
        for (let k = 0; k < gLen; ++k) {
          if (grid[k][j] >= shuMax) {
            shuMax = grid[k][j];
          }
        }

        shu += shuMax;
      }
    }

    heng += henMax;
  }

  return heng + shu + di;
};

console.log(
  projectionArea([
    [1, 2],
    [3, 4],
  ])
);

console.log(projectionArea([[2]]));

console.log(
  projectionArea([
    [1, 0],
    [0, 2],
  ])
);
