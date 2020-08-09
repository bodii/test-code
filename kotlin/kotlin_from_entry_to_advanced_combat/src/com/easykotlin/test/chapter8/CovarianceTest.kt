package com.easykotlin.test.chapter8

open class Food
open class Fruit: Food()
class Apple: Fruit()
class Banana: Fruit()
class Grape: Fruit()

object GenericTypeDemo1 {
    fun addFruit(fruit: MutableList<Fruit>) {

    }

    fun getFruit(fruit: MutableList<Fruit>) {

    }
}

fun main(args: Array<String>) {
    val fruit: MutableList<Fruit> = mutableListOf(Fruit(), Fruit(), Fruit())
    GenericTypeDemo1.addFruit(fruit)
    GenericTypeDemo1.getFruit(fruit)

    val apples: MutableList<Apple> = mutableListOf(Apple(), Apple(), Apple())
//    GenericTypeDemo1.addFruit(apples)
//    GenericTypeDemo1.getFruit(apples)
}