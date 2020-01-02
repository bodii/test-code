package chapter04;

// Typical use of a nonstatic member class
public class MySet<E> extends AbstractSet<E> {
	@Override public Iterator<E> iterator() {
		return new MyIterator();
	}

	private class MyIterator implements Iterator<E> {
	}
}
