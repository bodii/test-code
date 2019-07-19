// 宠物类的使用示例（验证多态)

class PetTester01 {
	public static void main(String[] args) {
		Pet kurt = new Pet("Kurt", "艾一");
		kurt.introduce();
		System.out.println();

		RobotPet r2d2 = new RobotPet("R2D2", "卢克");
		r2d2.introduce();
		System.out.println();

		Pet p = r2d2;
		p.introduce();
	}
}
