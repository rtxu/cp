type NumArray struct {
    nums []int
    bit []int
}


func Constructor(nums []int) NumArray {
    array := NumArray{
        nums: nums,
        bit : make([]int, len(nums)+1),
    }
    for i := 0; i < len(nums); i++ {
        array._update(i, nums[i])
    }
    return array
}


func (this NumArray) _update(i int, delta int)  {
    for i++; i < len(this.bit); i += i&-i {
        this.bit[i] += delta
    }
}

func (this NumArray) prefixSum(i int) int {
    var result int
    for i++; i > 0; i -= i&-i {
        result += this.bit[i]
    }
    return result
}

func (this *NumArray) Update(i int, val int)  {
    delta := val - this.nums[i]
    this.nums[i] = val
    this._update(i, delta)
}


func (this *NumArray) SumRange(i int, j int) int {
    return this.prefixSum(j) - this.prefixSum(i-1)
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(i,val);
 * param_2 := obj.SumRange(i,j);
 */
