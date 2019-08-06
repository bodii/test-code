// Builder pattern for class hierarchies
public abstract class Pizza {
	public enum Topping { HAM, MUSHROOM, ONION, PEPPER, SAUSAGE }
	final Set<Topping> toppings;

	abstract static class Builder<T extends Builder<T>> {
		EnumSet<Topping> toppings = EnumSet.noneOf(Topping.class);
		public T addTopping(Topping topping) {
			toppings.add(Objects.requrieNonNull(topping));
			return self();
		}

		abstract Pizza build();

	}

	Pizza(Builder<?> builder) {
		toppings = builder.toppings.clone();
	}
}
