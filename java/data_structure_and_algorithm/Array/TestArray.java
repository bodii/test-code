package array;

/**
 * TestArray class
 * 测试Array数组泛型类
 */
public class TestArray {
	public static void main(String[] args) {
		Array<String> sa = new Array<>();
		System.out.println("添加0－32到数组中");
		for (int i = 0; i < 33; i++)
			sa.addLast(i + "");
		System.out.println(sa + "\n");
		System.out.println("在数组头部添加55: ");
		sa.addFirst("55");
		System.out.println(sa + "\n");
		System.out.println("在数组尾部添加99: ");
		sa.addLast("99");
		System.out.println(sa + "\n");


		System.out.println("删除索引为2的元素： ");
		sa.remove(2);
		System.out.println(sa + "\n");
		System.out.println("删除索引为7的元素,且删除13次，每次索引加1： ");
		for (int i = 7; i < 20; i++)
			sa.remove(i);
		System.out.println(sa + "\n");
		System.out.println("删除第一个元素： " + sa.remvoeFirst());
		System.out.println(sa + "\n");
		System.out.println("删除最后一个元素： " + sa.removeLast());
		System.out.println(sa + "\n");
		System.out.println("删除指定元素: 3");
		sa.removeElement("3");
		System.out.println(sa + "\n");

		System.out.println("查询20是否存在：" + sa.contains("20"));
		System.out.println();

		System.out.println("设置索引6的值为41:");
		sa.set(6, "41");
		System.out.println(sa + "\n");

	}
}
