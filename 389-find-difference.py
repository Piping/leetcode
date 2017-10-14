def best_practice(s,t):
   st = 0
   n = len(s)
   for i in range(n):
       st ^= ord(s[i])
       st ^= ord(t[i])
   st ^= ord(t[-1])
   return chr(st)

def clever(s,t):
   return chr(reduce(operator.xor, map(ord, s + t)))