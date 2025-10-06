package solution;

import java.util.*;

public class Solution {
    private final int[][] directions = {
            {1,0}, {-1,0}, {0,1}, {0,-1}
    };
    private record Node (int row, int col, int highestElevation) implements Comparable<Node> {
        @Override
        public int compareTo(Node o) {
            return Integer.compare(highestElevation, o.highestElevation);
        }

        @Override
        public boolean equals(Object o) {
            if (this == o) return true;
            if (!(o instanceof Node obj)) return false;
            return this.row == obj.row && this.col == obj.col;
        }

        @Override
        public int hashCode() {
            return Objects.hash(row, col);
        }
    }

    public int swimInWater(int[][] grid) {
        int size = grid.length;
        HashSet<Node> visited = new HashSet<>();
        PriorityQueue<Node> queue = new PriorityQueue<>();
        queue.add(new Node(0,0,grid[0][0]));
        Node result = new Node(0,0,0);

        while (!queue.isEmpty()) {
            Node current = queue.poll();
            visited.add(current);

            for (List<Integer> direction: getDirections(current.row, current.col, size, size)){
                int newHeight = Integer.max(current.highestElevation,
                        grid[direction.get(0)][direction.get(1)]);

                Node neighbor = new Node(direction.get(0), direction.get(1), newHeight);
                if (visited.contains(neighbor)) continue;

                if (direction.get(0) == size - 1 && direction.get(1) == size - 1) {
                    return newHeight;
                }
                queue.add(new Node(direction.get(0), direction.get(1), newHeight));
            }

        }

        return result.highestElevation;
    }

    private List<List<Integer>> getDirections(int row, int col, int maxRow, int maxCol) {
        List<List<Integer>> list = new ArrayList<>();

        for (int[] direction : directions) {
            if (row + direction[0] < 0 || row +direction[0] >= maxRow || col + direction[1] < 0
                    || col +direction[1] >= maxCol) {
                continue;
            }
            list.add(Arrays.asList(row + direction[0], col + direction[1]));
        }

        return list;
    }
}
