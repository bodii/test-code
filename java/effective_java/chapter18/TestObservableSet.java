package chapter18;

import java.util.HashSet;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

import chapter18.ObservableSet;
import chapter18.SetObserver;

public class TestObservableSet {
    public static void main(String[] args) {
        ObservableSet<Integer> set = 
            new ObservableSet<>(new HashSet<>());

        set.addObserver((s, e) -> System.out.println(e));

        set.addObserver(new SetObserver<>() {
            public void added(ObservableSet<Integer> s, Integer e) {
                System.out.println(e);
                if (e == 23) {
                    ExecutorService exec = Executors.newSingleThreadExecutor();

                    try {
                        exec.submit(() -> s.removeObserver(this)).get();
                    } catch (ExecutionException | InterruptedException ex) {
                        throw new AssertionError(ex);
                    } finally {
                        exec.shutdown();
                    }
                }
            }
        });

        for (int i = 0; i < 100; i ++)
            set.add(i);
    }
}