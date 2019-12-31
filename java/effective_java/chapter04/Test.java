package chapter04;

// Use of static import to avoid qualifying constants
import static chapter04.PhysicalConstants02.*;

public class Test {
    double atoms(double mols) {
        return AVOGADROS_NUMBER * mols;
    }
}