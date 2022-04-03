/**
 * @param {character[]} letters
 * @param {character} target
 * @return {character}
 */
var nextGreatestLetter = function (letters, target) {
  const len = letters.length;
  let res = "";
  for (let i = 0; i < len; i++) {
    if (letters[i] > target) {
      res = letters[i];
      break;
    }
  }
  if (res) {
    return res;
  }

  return letters[0];
};

console.log(nextGreatestLetter(["c", "f", "j"], "a"), "c");
console.log(nextGreatestLetter(["c", "f", "j"], "c"), "f");
console.log(nextGreatestLetter(["c", "f", "j"], "d"), "f");
console.log(nextGreatestLetter(["c", "f", "j"], "z"), "c");
