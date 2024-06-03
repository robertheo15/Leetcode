class Solution:
    def isIsomorphic(self, s: str, t: str) -> bool:
        t_to_s_charmap = {}
        s_to_t_charmap = {}

        for i, t in enumerate(t):
            if t in t_to_s_charmap.keys():
                if t_to_s_charmap[t] != s[i]:
                    return False
                continue
            if s[i] in s_to_t_charmap.keys():
                if s_to_t_charmap[s[i]] != t:
                    return False
                continue

            t_to_s_charmap[t] = s[i]
            s_to_t_charmap[s[i]] = t
            
        return True