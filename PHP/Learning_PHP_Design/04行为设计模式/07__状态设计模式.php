<?php 

/*
	【 状态设计模式 】

	什么是状态模式
	状态（State)设计模式是GoF提出的最吸引人的模式之一，也是一种最有用的模式。
	游戏通常就采用状态模式，因为游戏中的对象往往会非常频繁地改变状态。

	状态模式的作用就是允许对象在状态改变时改变其行为。

	对于状态设计模式，每个状态都有自己的具体类，它们实现一个公共接口。我们不打
	算查看对象的控制流，而是从另一个角度来考虑，即对象的状态。

	状态机是一个模型，其重点包括不同的状态，一个状态到另一个状态的变迁，以用导致状态改变的触发器。


	状态模型的本质：
	状态 （关灯和开灯）
	变迁 （从关灯到开灯，以及从开灯到关灯）
	触发器 （灯开关）
*/

