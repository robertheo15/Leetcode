class Solution {
   public boolean isValid(String s) {
        HashMap<Character, Character> Hmap = new HashMap<>();
        Hmap.put(')','(');
        Hmap.put('}','{');
        Hmap.put(']','[');
        Stack<Character> stack = new Stack<>();

        for (int idx = 0; idx < s.length(); idx++){
            if (s.charAt(idx) == '(' || s.charAt(idx) == '{' || s.charAt(idx) == '[') {
                stack.push(s.charAt(idx));
                continue;
            }

            if (stack.isEmpty() || Hmap.get(s.charAt(idx)) != stack.pop()) {
                return false;
            }
        }

        return stack.isEmpty();
    }
}