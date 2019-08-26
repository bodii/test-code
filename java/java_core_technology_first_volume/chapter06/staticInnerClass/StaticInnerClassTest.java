package staticInnerClass;

/**
 * This program demonstrates the use of static inner class.
 * @version 1.1 2019-08
 * @author wong
 */

 public class StaticInnerClassTest {
     public static void main(String[] args) {
         double[] d = new double[20];

         for (int i = 0; i < d.length; i++)
            d[i] = 100 * Math.random();
        
        ArrayAlg.Pair p = ArrayAlg.minmax(d);
        System.out.println("min = " + p.getFirst());
        System.out.println("max = " + p.getSecond());
     }
 }

 class ArrayAlg {
     /**
      * A pair of floating-point numbers
      */

    public static class Pair {
        private double first;
        private double second;

        public Pair(double f, double s) {
            first = f;
            second = s;
        }

        public double getFirst() {
            return first;
        }

        public double getSecond() {
            return second;
        }
    }

    public static Pair minmax(double[] values) {
        double min = Double.POSITIVE_INFINITY;
        double max = Double.NEGATIVE_INFINITY;

        for (double v : values) {
            if (min > v) min = v;
            if (max < v) max = v;
        }
        return new Pair(min, max);
    }
 }