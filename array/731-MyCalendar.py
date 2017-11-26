#%%
class MyCalendar:

    def __init__(self):
        self.map = {}
        
''' brutal force/ binary search condition
e1 <= s2 or e2 <= s1
'''
    def book(self, start, end):
        """
        :type start: int
        :type end: int
        :rtype: bool
        """
        self.map[start] = self.map.get(start,0) - 1
        self.map[end] = self.map.get(end,0) + 1 
        booked = 0
        for key in sorted(self.map.keys()):
           booked += self.map[key]
           if booked >= 2:
              self.map[start] -= 1
              self.map[end] += 1
              return False
        return True

obj = MyCalendar()
print(obj.book(10, 20))
print(obj.book(40, 50))
print(obj.book(45, 60))