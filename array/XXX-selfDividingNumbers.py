#%%
def selfDividingNumbers(left, right):
        """
        :type left: int
        :type right: int
        :rtype: List[int]
        """
        ans = []
        for i in range(left,right+1):
            ok = True
            l = str(i)
            if '0' in l:
               continue
            for n in str(i):
                if i % int(n) != 0:
                    ok = False
                    break
            if ok :
                ans.append(i)
        return ans

selfDividingNumbers(1,22)