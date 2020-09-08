/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
    if isBadVersion(1) {
        return 1
    }
    // left always good, right always bad
    left, right := 1, n
    for left+1 != right {
        m := (left+right)>>1
        if isBadVersion(m) {
            right = m
        } else {
            left = m
        }
    }
    return right
}
