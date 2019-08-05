// Telescoping constructor pattern - does not scale well !

public class NutritionFacts {
	private final int servingSize;
	private final int servings;
	private final int calories;
	private final int fat;
	private final int sodium;
	private final int carbohydrate;

	public NutritionFacts(int servingSize, int servings) {
		this(servingSize, servings, 0);
	}

	public NutritionFacts(int servingSize, int servings, int calories) {
		this(servingSize, servings, calories, 0);
	}

	public NutritonFacts(int servingSize, int servings, int calcories, 
			int fat) {
		this(servingSize, servings, calories, fat, 0);
	}

	public NutrionFacts(int servingSize, int servings, int calcories,
			int fat, int sodium) {
		this(servingSize, servings, calories, fat, sodium, 0);
	}

	public NutrionFacts(int servingSize, int servings, int calcories,
			int fat, int sodium, int carbohydrate) {
		this.servingSize = servingSize;
		this.servings	 = servings;
		this.calories	 = calories;
		this.fat		 = fat;
		this.sodium		 = sodium;
		this.carbohydate = carbohydrate;
	}
}

/*
NutritionFacts cocaCola = new NutritionFacts(240, 8, 100, 0, 35, 27);

重叠构造器模式可行，但是当有许多参数的时候，客户端代码会很难编写，并且仍然较难以阅读。
*/
