function gridChallenge(grid) {
  // Write your code here
  for (let i = 0; i < grid.length; i++) {
    const element = grid[i];
    grid[i] = element.split('').sort().join('');
  }
  for (let i = 0; i < grid.length; i++) {
    for (let j = 0; j < grid.length - 1; j++) {
      if (grid[j][i] > grid[j + 1][i]) {
        return "NO";
      }
    }
  }
  return "YES";
}

let grid = ['mpxz', 'abcd', 'wlmf'];

console.log(gridChallenge(grid));
