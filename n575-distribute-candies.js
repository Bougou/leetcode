function distributeCandies(candies) {
    var candiesSet = new Set(candies);
    return candies.length/2 > candiesSet.size ? candiesSet.size : candies.length/2
}
