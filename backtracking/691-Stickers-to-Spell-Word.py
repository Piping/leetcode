#%%

sticker = ["these", "guess", "about", "garden", "him"]
target = "atomher"

def minStickers( sticker, target):
    """
    :type stickers: List[str]
    :type target: str
    :rtype: int
    """
    tc = collections.Counter(target)
    sc = [collections.Counter(s) & tc for s in sticker]  # Optimization 1: &tc

    # Optimization 2: Remove Dominated Stickers
    for i in range(len(sc) - 1, -1, -1):
        if any(sc[i] == sc[i] & sc[j] for j in range(len(sc)) if j != i):
            sc.pop(i)

    best = len(target) + 1

    def search(ans):
        nonlocal best
        if ans >= best:
            return
        if not sc:  # sc is empty
            if all(tc[c] <= 0 for c in target):
                best = ans
            return

        sticker = sc.pop()
        #math.ceil, (A-1)//B + 1,  A//B + A%B == 0
        used = max((tc[c] - 1) // sticker[c] + 1 for c in sticker)
        used = max(used, 0)

        for c in sticker:
            tc[c] -= used * sticker[c]

        search(ans + used)

        for i in range(used - 1, -1, -1):
            # backtracking target Counter
            for c in sticker:
                tc[c] += sticker[c]
            search(ans + i)
        # backtracking sticker Counter
        sc.append(sticker)
    search(0)
    return best if best <= len(target) else -1

minStickers(sticker, target)
