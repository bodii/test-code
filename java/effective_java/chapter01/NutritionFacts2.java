// JavaBeans pattern - allows inconsistency, mandates mutability

public class NutritonFacts2 {
	private int servingSize		= -1;
	private int servings		= -1;
	private int calories		= 0;
	private int fat				= 0;
	private int sodium			= 0;
	private int carbohydrate	= 0;

	public NutritonFacts2 { }

	// Setters
	public void setServingSize(int val) { servingSize = val; }
	public void setServings(int val)	{ servings	  = val; }
	public void setCalories(int val)	{ calories	  = val; }
	public void setFat(int val)			{ fat		  = val; }
	public void setSodium(int val)		{ sodium	  = val; }
	public void setCarbohydrate(int val){ carbohydrate = val;}
}

/*
 NutritionFacts cocaCola = new NutritionFacts();
 cocaCola.setServingSize(240);
 cocaCola.setSverings(8);
 cocaCola.setCalories(100);
 cocaCola.setSodium(35);
 cocaCola.setCarbohydrate(27);

 在构造过程中JavaBean可能处于不一致的状态。

 */
