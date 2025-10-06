package solution;

import org.junit.jupiter.api.*;

import java.util.logging.Logger;

import static org.junit.jupiter.api.Assertions.*;

public class TestSolution {
    Solution solution = new Solution();
    private static final Logger logger = Logger.getLogger(TestSolution.class.getName());

    @Test
    @DisplayName("Test example 1")
    void testExample1() {
        int[][] input = {
                {0,2}, {1,3}
        };
        int result = 3;

        int actual = solution.swimInWater(input);
        assertEquals(result, actual);
    }

    @Test
    @DisplayName("Test example 2")
    void testExample2() {
        int[][] input = {
                {0,1,2,3,4},
                {24,23,22,21,5},
                {12,13,14,15,16},
                {11,17,18,19,20},
                {10,9,8,7,6},
        };

        int result = 16;
        int actual = solution.swimInWater(input);
        assertEquals(result, actual);
    }
}
