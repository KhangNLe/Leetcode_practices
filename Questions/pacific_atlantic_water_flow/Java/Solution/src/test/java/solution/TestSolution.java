package solution;

import java.util.*;
import java.util.logging.Logger;

import org.junit.jupiter.api.*;
import static org.junit.jupiter.api.Assertions.*;

public class TestSolution {
    Solution solution = new Solution();
    private static Logger logger = Logger.getLogger(TestSolution.class.getName());

    @Test
    @DisplayName("Test Example 1")
    void testExample1() {
        int[][] inputs = {
                {1,2,2,3,5},
                {3,2,3,4,4},
                {2,4,5,3,1},
                {6,7,1,4,5},
                {5,1,1,2,4}
        };

        int[][] result = {
                {0,4},
                {1,3},
                {1,4},
                {2,2},
                {3,0},
                {3,1},
                {4,0},
        };

        List<List<Integer>> list = solution.pacificAtlantic(inputs);
        sort2DArrays(list);

        int[][] ans = convertList(list);

        logger.info(Arrays.deepToString(ans));
        logger.info(Arrays.deepToString(result));
        assertArrayEquals(result, ans);

    }

    @Test
    @DisplayName("Test Example 2")
    void testExample2() {
        int[][] inputs = {
                {1},
        };
        int[][] result = {{0,0}};
        int[][] ans = convertList(solution.pacificAtlantic(inputs));

        assertArrayEquals(result, ans);
    }

    private int[][] convertList(List<List<Integer>> list){
        int[][] result = new int[list.size()][list.getFirst().size()];

        for (int i = 0; i < list.size(); i++) {
            for (int j = 0; j < list.getFirst().size(); j++) {
                result[i][j] = list.get(i).get(j);
            }
        }
        return result;
    }

    private void sort2DArrays(List<List<Integer>> list){
        list.sort(Comparator.comparing((List<Integer> x) -> x.getFirst())
                .thenComparing(x -> x.get(1)));
    }
}
