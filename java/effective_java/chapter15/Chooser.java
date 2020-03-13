package chapter15;

import java.util.ArrayList;
import java.util.Collection;
import java.util.List;
import java.util.Random;
import java.util.concurrent.ThreadLocalRandom;

/**
 * List-based Chooser - typesafe
 * @param <T>
 */
public class Chooser<T> {
    private final List<T> choiceList;

    public Chooser(Collection<T> choices) {
        choiceList = new ArrayList<>(choices);
    }

    public T choose() {
        Random rand = ThreadLocalRandom.current();
        return choiceList.get(rand.nextInt(choiceList.size()));
    }
}