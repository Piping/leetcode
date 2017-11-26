#%%
class Solution:
    def areSentencesSimilar(self, words1, words2, pairs):
        """
        :type words1: List[str]
        :type words2: List[str]
        :type pairs: List[List[str]]
        :rtype: bool
        """
        if len(words1) != len(words2):
            return False
        if not pairs:
            for w1,w2 in zip(words1,words2):
                if w1 != w2:
                    return False
            return True
        else:
            for w1,w2 in zip(words1,words2):
                if [w1,w2] not in pairs and [w2,w1] not in pairs and w1 != w2:
                    return False
            return True
         
words1 = ["great","acting","skills"]
words2 = ["fine","drama","talent"]
pairs = [["great", "fine"], ["acting","drama"], ["skills","talent"]]
print(areSentencesSimilar(words1,words2,pairs))
#print(['great','fine'] in [["great", "fine"]])
["an","extraordinary","meal","meal"]
["one","good","dinner"]
[["great","good"],["extraordinary","good"],["well","good"],["wonderful","good"],["excellent","good"],["fine","good"],["nice","good"],["any","one"],["some","one"],["unique","one"],["the","one"],["an","one"],["single","one"],["a","one"],["truck","car"],["wagon","car"],["automobile","car"],["auto","car"],["vehicle","car"],["entertain","have"],["drink","have"],["eat","have"],["take","have"],["fruits","meal"],["brunch","meal"],["breakfast","meal"],["food","meal"],["dinner","meal"],["super","meal"],["lunch","meal"],["possess","own"],["keep","own"],["have","own"],["extremely","very"],["actually","very"],["really","very"],["super","very"]]