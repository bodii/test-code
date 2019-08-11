import java.util.Date;

public class FormatDate {
	public static void main(String[] args) {
		System.out.printf("%1$s %2$tB %2$tY", "Due date: ", new Date());
		System.out.printf("%s %tB %<te, %<tY", "Due date: \n", new Date());

		System.out.println();
	}
}
