package union_find;

import java.util.Random;

import union_find.UnionFindable;
import union_find.UnionFind;
import union_find.UnionFind2;
import union_find.UnionFind3;
import union_find.UnionFind4;
import union_find.UnionFind5;
import union_find.UnionFind6;

public class Test {

    private static double testUnionFind(UnionFindable u, int m) {
        int size = u.getSize();
        Random random = new Random();

        long startTime = System.nanoTime();

        for (int i = 0 ; i < m; i ++ ) {
            int a = random.nextInt(size);
            int b = random.nextInt(size);
            u.unionElements(a, b);
        }

        for (int i = 0; i < m; i ++) {
            int a = random.nextInt(size);
            int b = random.nextInt(size);
            u.isConnected(a, b);
        }

        long endTime = System.nanoTime();

        return (endTime - startTime) / 1000000000.0;
    }

    public static void main(String[] args) {
        int size = 10000;
        int m = 10000;

        UnionFind u1 = new UnionFind(size);
        System.out.println("union find 1:  " + testUnionFind(u1, m));

        UnionFind2 u2 = new UnionFind2(size);
        System.out.println("union find 2: " + testUnionFind(u2, m));

        UnionFind3 u3 = new UnionFind3(size);
        System.out.println("union find 3: " + testUnionFind(u3, m));

        UnionFind4 u4 = new UnionFind4(size);
        System.out.println("union find 4: " + testUnionFind(u4, m));

        // UnionFind5 u5 = new UnionFind5(size);
        // System.out.println("union find 5: " + testUnionFind(u5, m));

        UnionFind6 u6 = new UnionFind6(size);
        System.out.println("union find 6: " + testUnionFind(u6, m));
    }
}
