#%%
def floodFill(image, sr, sc, newColor):
   oldColor = image[sr][sc]
   def flood(sr, sc):
      if sr < 0 or sr >= len(image):
          return
      if sc < 0 or sc >= len(image[0]):
          return
      if image[sr][sc] == newColor:
         return
      if image[sr][sc] == oldColor:
          image[sr][sc] = newColor
          flood(sr+1,sc)
          flood(sr-1,sc)
          flood(sr,sc+1)
          flood(sr,sc-1)
   flood(sr,sc)
   return image
      
image = [[0,0,0],[0,1,1]]
#image = [[1, 1, 1], [1, 1, 0], [1, 0, 1]]
sr = 1
sc = 1
newColor = 1
floodFill(image,sr,sc,newColor)
