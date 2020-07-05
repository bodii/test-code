package chapter12;

public class RealResult<T> extends Result<T> {
    private final T resultValue;

    public RealResult(T value) {
        resultValue = value;
    }

    public T getResultValue() {
        return resultValue;
    }
}