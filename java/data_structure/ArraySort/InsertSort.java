/**
 * 直接插入排序算法
 */

public class InsertSort {

	public static void main(String[] args) {
		int[] data = { 23, 45, 16, 7, 42 };
		int len = data.length;

		for (int i = 1; i < len; i++) {
			int currentElement = data[i];
			int temp = i;
			while (temp > 0 && data[temp - 1] > currentElement) {
				data[temp] = data[temp - 1];
				temp--;
			}
			data[temp] = currentElement;
		}

		System.out.println("sorted:");
		for (int j = 0; j < len; j++) {
			System.out.println(data[j]);
		}
	}
}
