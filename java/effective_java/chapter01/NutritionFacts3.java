// Builder pattern

public class NutritionFacts3 {
	private final int servingSize;
	private final int servings;
	private final int calories;
	private final int fat;
	private final int sodium;
	private final int carbohydrate;

	public static class Builder {
		// Required parameters
		private final int servingSize;
		private final int servings;

		// Optional parameters - initialied to default values
		private int calories		= 0;
		private int fat				= 0;
		private int sodium			= 0;
		private int carbohydrate	= 0;

		public Builder(int servingSize, int servings) {
			this.servingSize = servingSize;
			this.servings	 = servings;
		}

		public Builder calories(int val) {
			calories = val;
			return this;
		}

		public Builder fat(int val) {
			fat = val; 
			return this;
		}

		public Builder sodium(int val) {
			sodium = val;
			return this;
		}

		public Builder carbohydrate(int val) {
			carbohydrate = val;
			return this;
		}

		public NutritionFacts3 build() {
			return new NutritionFact3(this);
		}
	}

	private NutritionFacts3(Builder builder) {
		servingSize		= builder.servingSize;
		servings		= builder.servings;
		calories		= builder.calories;
		fat				= builder.fat;
		sodium			= builder.sodium;
		carbohydrate	= builder.carbohydrate;
	}
}


/*
 NutritionFacts3 cocaCola = new NutritionFacts3.Builder(240, 8)
		.calories(100).sodium(25).carbohydrate(27).build();

	Builder 模式模拟了具名的可选参数。
*/


