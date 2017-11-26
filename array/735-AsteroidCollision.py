class Solution:
    def asteroidCollision(self, asteroids):
        """
        :type asteroids: List[int]
        :rtype: List[int]
        """
        ans, pos = [],[]
        for a in asteroids:
            if a > 0:
                pos.append(a)
            else:
                exist = True
                while pos:
                    val = pos[-1]+a
                    if val == 0:
                        pos.pop()
                        exist = False
                        break
                    elif val < 0:
                        pos.pop()
                    else:
                        exist = False
                        break
                if exist:
                    ans.append(a)
        ans.extend(pos)
        return ans