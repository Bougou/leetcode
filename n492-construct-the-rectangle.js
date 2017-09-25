/**
 * @param {number} area
 * @return {number[]}
 */
 var constructRectangle = function(area) {
    //
    var L = Math.ceil(Math.sqrt(area))
    
    while (area % L !== 0) {
        L++
    }
    
    return [L, area/L]
};
