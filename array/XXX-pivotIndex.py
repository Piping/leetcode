#%%
def pivotIndex(nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        n = sum(nums)
        s = 0
        for i,v in enumerate(nums):
            s += v
            n -= v
            if s-v == n:
                return i
        return -1

pivotIndex([1,2,3,4,5])
