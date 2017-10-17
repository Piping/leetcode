# Backtrack
def readBinaryWatch(num):
   ans = []
   #bits pattern has length 10, h:4 m:6
   def bt(n,bits,usedn):
      h = bits >> 6 # get first four bit
      m = bits & ((1<<6)-1) #get last six bit
      if h > 11 or m > 59:
         return
      if n == 0:
         ans.append("%d:%02d" % (h,m))
         return
      for i in range(usedn,10):
         bt(n-1,bits | (1<<i), i+1)
   
   bt(num,0,0)
   return ans
