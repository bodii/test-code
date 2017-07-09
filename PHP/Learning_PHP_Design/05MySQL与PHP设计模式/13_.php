<?php 
	
/*
	【 使用SPL实现观察者模式 】 

	可用于观察者设计模式的3个SPL接口/类：
	SplSubject
	SplObserver
	SplObjectStorage
*/

/* SplObjectStorage类最重要的方法是内置attach()和detach()方法，利用这两个方法
	,可以很容易地指定哪些观察者将“订阅”和“解除订阅”更新通知
 */

class ConcreteSubject implements SplSubject
{
	private $observers;
	private $data;

	public function setObservers()
	{
		$this->observers = new SplObjectStorage;
	}

	public function attach(SplObserver $observer)
	{
		$this->observers->attach($observer);
	}

	public function detach(SplObserver $observer)
	{
		$this->observers->detach($observer);
	}

	public function notify()
	{
		foreach ($this->observers as $observer) {
			$observer->update($this);
		}
	}

	public function setData($dataNow)
	{
		$this->data = $dataNow;
	}

	public function getData()
	{
		return $this->data;
	}

	/*
		SplSubject接口不包括获取方法和设置方法，不过这是观察者设计模式中的
		一部分，所以需要增加获取方法和设置方法。设置方法setData()包含一个参
		数，这是要增加的任何类型的数据。获取方法getData()存储当前主题状态，
		由具体观察者用来更新观察者数据。

		另外还增加了setObservers()方法。并不是在构造函数中设置SplObjectStorage
		()实例（这需要ConcreteSubject类的一个新实例），也没有在setData()方法
		中设置观察者实例，这里实现了一个单独的setObservers()方法，可以提供更
		松的耦合，并允许有多组观察者。
	 */
}


/*
	SPL具体观察者
 */

class ConcreteObserver implements SplObserver
{
	public function update(SplSubject $subject)
	{
		echo $subject->getData() , '<br>';
	}
}

/*
	SPL客户
	"SPL"Client类只是一个标准客户.这个客户按照SPL接口向具体主题和观察者发出
	多个请求,不过自己并没有实现SPL类或接口.
	在这个例子中，Client创建了一个主题实例和3个具体观察者实例。然后使用setData
	()方法设置一个新状态，并将这个3个观察者关联到这个主题。最后，它调用具体主
	题实例的notify()方法将当前状态发送给订阅的观察者：
 */

class Client
{
	public function __construct()
	{
		echo '<p>Create three new concrete observers,a new concrete subject:</p>';
		$ob1 = new ConcreteObserver;
		$ob2 = new ConcreteObserver;
		$ob3 = new ConcreteObserver;

		$subject = new ConcreteSubject;

		$subject->setObservers();
		$subject->setData('Here\'s your data!');
		$subject->attach($ob1);
		$subject->attach($ob2);
		$subject->attach($ob3);

		$subject->notify();
		echo '<p> Detach observer ob3. Now only ob1 and ob2 are notified:</p>';
		$subject->detach($ob3);
		$subject->notify();
		$subject->setData('More data that only ob1 and ob3 need.');
		$subject->attach($ob3);
		$subject->detach($ob2);
		$subject->notify();
	}
}

$worker = new Client;