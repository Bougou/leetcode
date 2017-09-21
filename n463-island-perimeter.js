var islandPerimeter = function(grid) {
    var rowLen = grid.length
    var colLen = grid[0].length
    
    var numLan = 0
    var numNeigh = 0
    
    for (var r = 0; r < rowLen; r++) {
        for (var c = 0; c < colLen; c++) {
            if (grid[r][c] == 1) {
                numLan++
                
                if ((r+1 < rowLen) && grid[r+1][c] == 1) {
                    numNeigh++
                }
                
                if ((c+1 < colLen) && grid[r][c+1] == 1) {
                    numNeigh++
                }
            }
        }
    }
    
    console.log(numLan)
    console.log(numNeigh)
    return numLan*4 - numNeigh*2
};

c = [[0,1,0,0],[1,1,1,0],[0,1,0,0],[1,1,0,0]]

console.log(islandPerimeter(c))
