package chapter06;

// Marker annotation type declaration
import java.lang.annotation.*;
import java.lang.Throwable;

/**
 * Indicates that the annotated method is a test method.
 * Use only on parameterless static methods.
 */
@Retention(RetentionPolicy.RUNTIME)
@Target(ElementType.METHOD)
public @interface IExceptionTest {
    Class<? extends Throwable> value();
}