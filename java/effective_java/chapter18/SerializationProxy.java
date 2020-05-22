package chapter18;

import java.util.EnumSet;

public static class SerializationProxy <E extends Enum<E>>
        implements Serializable {
    private final Class<E> elementType;

    private final Enum<?>[] elements;

    SerializationProxy(EnumSet<E> set) {
        elementType = set.elementType;
        elements = set.toArray(new Enum<?>[0]);
    }

    private Object readResolve() {
        EnumSet<E> result = EnumSet.noneOf(elementType);
        for (Enum<?> e : elements)
            result.add((E) e);

        return result;
    }

    private static final long serialVersionUID = 362491234563181265L;
}