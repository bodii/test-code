// -- CD播放器
public class CDPlayer implements Player {
	public void play() {
		System.out.println("CD播放开始!");
	}

	public void stop() {
		System.out.println("CD播放结束!");
	}

	public void cleaning() {
		System.out.println("已清洗磁头。");
	}
}
