package chapter10;


@FunctionalInterface public interface SetObserver<E> {
    void added(ObservableSet<E> set, E element);
}