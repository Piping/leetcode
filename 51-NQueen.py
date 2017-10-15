#%%
def solveNQueens(n):
        """
        :type n: int
        :rtype: List[List[str]]
        """
        #backtracking steps
        # start with root node, in dfs order
        # check if node is completed, return answer if it is
        # check if node can be completed
        #   ==> no, whole subtree is skipped, return
        #   ==> yes, iteractively get next sibiling,record it if necessary
        #       ==> recursively enumerates all sub-trees of the sibiling
        #       ==> backtrack   
        ans = []
        one = [[]]
        def bitToStr(np):
            f = '{:0'+str(n)+'b}'
            return ''.join(map(lambda x: '.' if x is '0' else 'Q',f.format(np)))
        
        # row, ld, rd are 3 bit patterns constrain possible positions
        upper = (1<<n) - 1
        def bt(row,ld,rd):
            if row == upper:
                print(one)
                ans.append(one)
                return
            pos = ~(row|ld|rd) & upper
            while pos != 0:
                # to get rightmost bit for next position
                np = pos & (~pos + 1) & upper # or pos & -pos 
                one.append(bitToStr(np))
                print(bitToStr(np))
                bt(row|np, (ld|np)<< 1, (rd|np)>>1)
                pos ^= np
                one.pop()
                print('backtrack:',bitToStr(np),'next',bitToStr(pos))
        bt(0,0,0)
        return ans
solveNQueens(5)