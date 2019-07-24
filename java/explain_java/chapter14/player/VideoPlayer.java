// -- 视频播放器
public class VideoPlayer implements Player {
	private int id;		// 制造编号
	private static int count = 0;	// 到目前为止已经赋的制造编号

	public VideoPlayer() {		// 构造函数
		id = ++count;
	}

	public void play() {
		System.out.println("视频播放开始!");
	}

	public void stop() {
		System.out.println("视频播放结束!");
	}

	public void printInfo() {
		System.out.println("该机器的制造编号为[" + id + "]。");
	}
}
