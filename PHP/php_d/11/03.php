<?php

# SPL具体观察者
# 具体观察者既简单又复杂。在这个例子中，它实现了更新函数来更新关联的观察者
# 实例。update()方法的实现中包含一个代码提示，要求使用SplSubject接口作为方
# 法参数。这就要求开发人员在所有update()调用中都要包含一个SplSubject实例:
#

class ConcreteObserver implements SplObserver
{
	public function update(SplSubject $subject)
	{
		echo $subject->getData() . "<br>";
	}
}

# SPL客户
# "SPL"Client类只是一个标准客户。这个客户按照SPL接口向具体主题和观察者发出
# 多个请求，不过自己并没有实现SPL类或接口。
# 在这个例子中，Client创建了一个主题实例和3个具体观察者实例。然后使用setData
# ()方法设置一个新状态，并将这3个观察者关联到这个主题。最后，它调用具体主题
# 实例的notify()方法将当前状态发送给订阅的观察者：

function __autoload($class_name)
{
	include $class_name . '.php';
}

class Client
{
	public function __construct()
	{
		echo "<p>Create three new concrete observers, a new concrete subject:
			</p>";
		
		$ob1 = new ConcreteObserver();
		$ob2 = new ConcreteObserver();
		$ob3 = new ConcreteObserver();

		$subject = new ConcreteSubject();
		$subject->setObservers();
		$subject->setData("Here's your data!");
		$subject->attach($ob1);
		$subject->attach($ob2);
		$subject->attach($ob3);

		$subject->notify();

		echo "<p>Datach observer ob3. Now only ob1 and ob2 are notified:</p>";
		$subject->detach($ob3);
		$subject->notify();

		echo "<p>Reset data and reattach ob3 and detach ob2--only ob 1 and
			3 are notified:</p>";
		$subject->setData('More data that only ob1 and ob3 need.');
		$subject->attach($ob3);
		$subject->detach($ob2);
		$subject->notify();
	}
}

$worker = new Client();


# 有一点需要注意，要把ConcreteSubject->attach()和ConcreteSubject->detach()
# 方法与SplObjectStorage->attach()和SplObjectStorage->detach()方法区别开。
# ConcreteSubject类把SplObjectStorage的attach()和detach()方法包装在自己的
# attach()和detach()方法中。下面的伪代码显示了一个类版本，可以看到Concrete-
# Subject如何创建关联/解除关联（attach/detach)方法：

public function attach(SplObserver $observer)
{
	SplobjectStorage->attach($observer);
}
public function detach(SplObserver $observer)
{
	SplObjectStorage->detach($observer);
}
# 如果把具体主题类的attach/detach方法改为其他名字，这可能同样会让人糊涂，
# 所以只需要了解内置SPL attach/detach方法可以用来创建ConcreteSubject attach/
# detach方法。
