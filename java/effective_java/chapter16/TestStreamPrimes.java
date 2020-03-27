package chapter16;

import java.math.BigInteger;
import java.util.stream.Stream;

public class TestStreamPrimes {

    private static final byte[] one = {1};
    private static final byte[] two = {2};
    private static final BigInteger ONE = new BigInteger(one);
    private static final BigInteger TWO = new BigInteger(two);

    private static Stream<BigInteger> primes() {
        return Stream.iterate(TWO, BigInteger::nextProbablePrime);
    }

    public static void main(String[] args) {
        primes().map(p -> TWO.pow(p.intValueExact()).subtract(ONE))
            .filter(mersenne ->mersenne.isProbablePrime(50))
            .limit(20)
            .forEach(System.out::println);
    }
}