import java.util.HashMap;
import java.util.Map;

public class main {
    public static void main(String[] args) {

        int[] nums = {5, 7, 11, 2};
        int target = 9;
        int[] result = twoSum(nums, target);
        System.out.println("Output: [" + result[0] + ", " + result[1] + "]"); // Output: [0, 1]
    }

    // The twoSum function is implemented in Go, but here is a Java equivalent for reference.
    public static int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> map = new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
            int complement = target - nums[i];
            if (map.containsKey(complement)) {
                return new int[]{map.get(complement), i};
            }
            map.put(nums[i], i);
        }
        throw new IllegalArgumentException("No two sum solution");
    }
}