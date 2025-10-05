package solution;
import java.util.*;

public class Solution {
    private final int[][] directions = {
            {1,0}, {-1,0}, {0,1}, {0, -1}
    };

    private record Node(int height, int row, int col) implements Comparable<Node> {
        @Override
        public int compareTo(Node o) {
            return Integer.compare(height, o.height);
        }

        @Override
        public boolean equals(Object o) {
            if (this == o) return true;
            if (!(o instanceof Node obb)) return false;
            return  height == obb.height &&  row == obb.row && col == obb.col;
        }

        @Override
        public int hashCode() {
            return Objects.hash(height, row, col);
        }
    }

    public List<List<Integer>> pacificAtlantic(int[][] heights){
        HashSet<Node> pacific = getPacificFlows(heights);
        HashSet<Node> atlantic = getAtlanticFlows(heights);
        List<List<Integer>> result = new ArrayList<>();

        for (Node node : pacific) {
            if (atlantic.contains(node)) {
                result.add(Arrays.asList(node.row, node.col));
            }
        }

        return result;
    }

    private HashSet<Node> getPacificFlows(int[][] heights){
        Deque<Node> nodes = new ArrayDeque<>();
        int rows = heights.length;
        int cols = heights[0].length;

        for (int i = 0; i < cols; i++) {
            nodes.add(new Node(heights[0][i], 0, i));
        }

        for (int i = 0; i < rows; i++) {
            nodes.add(new Node(heights[i][0], i, 0));
        }

        return getVisitedNodes(nodes, rows, cols, heights);
    }

    private HashSet<Node> getAtlanticFlows(int[][] heights){
        Deque<Node> nodes = new ArrayDeque<>();
        int cols = heights[0].length;
        int rows = heights.length;

        for (int i = 0; i < cols; i++) {
            nodes.add(new Node(heights[rows-1][i], rows-1, i));
        }

        for (int i = 0; i < rows; i++) {
            nodes.add(new Node(heights[i][cols-1], i, cols-1));
        }

        return getVisitedNodes(nodes, rows, cols, heights);
    }


    private HashSet<Node> getVisitedNodes(Deque<Node> nodes, int rows, int cols, int[][] heights){
        HashSet<Node> visited = new HashSet<>();

        while(!nodes.isEmpty()){
            Node node = nodes.removeFirst();
            visited.add(node);

            List<List<Integer>> neighbors = getNeighbors(node.row, node.col, rows, cols);
            for (List<Integer> neighbor :neighbors){
                int neighborHeight = heights[neighbor.get(0)][neighbor.get(1)];
                if (node.height <= neighborHeight){
                    Node newNeighbor = new Node(neighborHeight, neighbor.get(0), neighbor.get(1));
                    if (visited.contains(newNeighbor)){
                        continue;
                    }
                    nodes.addLast(newNeighbor);
                }
            }
        }
        return  visited;
    }

    private List<List<Integer>> getNeighbors(int row, int col, int maxRow, int maxCol){
        List<List<Integer>> ans = new  ArrayList<>();

        for (int[] direction : directions){
            int newRow = direction[0] + row;
            int newCol = direction[1] + col;
            if (newRow < 0 || newRow >= maxRow || newCol < 0 || newCol >= maxCol){
                continue;
            }
            ans.add(Arrays.asList(newRow, newCol));
        }
        return ans;
    }
}
