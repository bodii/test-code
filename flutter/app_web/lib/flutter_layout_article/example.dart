import "package:flutter/material.dart";

abstract class Example extends StatelessWidget {
  const Example({Key? key}) : super(key: key);

  String get code;
  String get explanation;
}

const red = Colors.red;

const green = Colors.green;

const blue = Colors.blue;

const big = TextStyle(fontSize: 30);
