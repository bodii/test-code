package chapter06;

import java.lang.annotation.Annotation;
import java.lang.reflect.*;

public class RunTests2 {
    public static void main(String[] args) throws Exception {
        int tests = 0;
        int passed = 0;

        Class<?> testClass = Class.forName(args[0]);
        for (Method m : testClass.getDeclaredMethods()) {
            if (m.isAnnotationPresent(IExceptionTest.class)) {
                tests++;
                try {
                    m.invoke(null);
                    passed++:
                } catch(InvocationTargetException warppedExc) {
                    Throwable exc = wrappedExc.getCause();
                    Class<? extends Throwable> excType = m.getAnnotation(IExceptionTest.class).value();
                    if (excType.isInstance(exc)) {
                        passed++;
                    } else {
                        System.out.printlf(
                            "Test %s failed: expected %s, got %s%n", m, excType.getName(), exc
                            );
                    }
                } catch (Exception exc) {
                    System.out.println("Invalid @Test: " + m);
                }
            }
        }
        System.out.printf("Passed: %d, Failed: %d%n", passed, tests - passed);
    }
}