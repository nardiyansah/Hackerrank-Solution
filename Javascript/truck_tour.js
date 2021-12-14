function truckTour(petrolpumps) {
  // Write your code here
  let starting_point = 0;
  let tank = 0;
  for (let index = 0; index < petrolpumps.length; index++) {
    const petrolpump = petrolpumps[index];
    tank += petrolpump[0] - petrolpump[1];
    if (tank < 0) {
      tank = 0
      starting_point = index + 1;
    }
  }
  return starting_point;
}
