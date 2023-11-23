use std::thread;
use std::sync::{mpsc, Arc, Mutex};

#[allow(dead_code)]
pub struct ThreadPool {
    workers: Vec<Worker>,
    sender: mpsc::Sender<Job>,
}

#[allow(dead_code)]
struct Worker {
    id: usize,
    thread: thread::JoinHandle<()>,
}

// struct Job;
type Job = Box<dyn FnOnce() + Send + 'static>;

impl ThreadPool {
    /// create thread pool
    ///
    /// number of threads in the thread pool 
    /// 
    /// # Pnaics
    ///
    /// `new`函数在size为0时会panic
    pub fn new(size: usize) -> ThreadPool {
        assert!(size > 0);

        // 通道
        let (sender, receiver) = mpsc::channel();

        // Arc使得多个wroker拥有接收端，
        // 而Mutex则确保一次只有一个worker能从接收端得到任务
        let receiver = Arc::new(Mutex::new(receiver));

        let mut workers = Vec::with_capacity(size);

        for id in 0..size {
            // 对于每一个新worker，克隆Arc来增加引用计数，
            // 如此这些worker就可以共享接收端的所有权了
            workers.push(Worker::new(id, Arc::clone(&receiver)));
        }

        ThreadPool {
            workers,
            sender,
        }
    }

    /// 得到的闭包新建Job实例之后，
    /// 将这些任务从通道的发送端发出
    pub fn execute<F>(&self, f: F) 
        where F: FnOnce() + Send + 'static 
    {
        let job = Box::new(f);

        // 这里调用send上的unwrap，因为发送可能会失败，这可能发生于
        // 停止了所有线程执行的情况，这意味着接收端停止接收新消息了。
        self.sender.send(job).unwrap();
    }
}

impl Worker {
    fn new(id: usize, receiver: Arc<Mutex<mpsc::Receiver<Job>>>) -> Worker {
        let thread = thread::spawn(move || {
            // 向通道的接收端请求任务，并在得到任务时执行他们
            // 如果锁定了互斥器，接着调用recv从通道中接收Job。
            // 最后的unwrap也绕过了一些错误，
            // 这可能发生于持有通道发送端的线程停止的情况，
            // 调用recv会阻塞当前线程，所以如果还没有任务，其会等待直到有可用的任务。
            // Mutex<T>确保一次只有一个Worker线程尝试请求任务
            while let Ok(job) = receiver.lock().unwrap().recv() {
                println!("Worker {} got a job; execting.", id);

                job();
            }
        });

        Worker {
            id, 
            thread,
        }
    }
}
