<?php


# 利用观察者模式构建多设备CMS
#
# 内容管理系统(content management system, CMS)
#
# PHP内置特性
# SplObserver接口和SqlSubject以及SplObjectStorage接口，利用这些接口，构建
# PHP观察者模式简直易如反掌。“SPL”是标准PHP库（Standard PHP Library)的简写
# ，这个库中包括一组解决标准问题的接口和类。
# 
# 观察者模式
# 观察者模式的类图很详细，模型-视图-控制器（Model-View-Controller,MVC)模式
# 中推崇的很多特性都可以在观察者模式中找到（你甚至可能会认为这个模式是MVC
# 的一种替代模式）。观察者模式的核心在于Subject和Observer接口。Subject包含
# 一个给定的状态，观察者“订阅”这个主题，将主题的当前状态通知观察者。
#
# 观察者模式很直观。何必让多个对象创建或跟踪一个给定的状态呢？如果由一个对
# 象完成这个工作，然后通知其他可能用到这个状态的对象，这样会合理得多。
#
