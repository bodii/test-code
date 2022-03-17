import 'package:app_web/flutter_layout_article/button.dart';
import 'package:flutter/material.dart';
import 'package:flutter/rendering.dart';
import 'example.dart';

import 'package:app_web/flutter_layout_article/examples/example1.dart';
import 'package:app_web/flutter_layout_article/examples/example2.dart';
import 'package:app_web/flutter_layout_article/examples/example3.dart';
import 'package:app_web/flutter_layout_article/examples/example4.dart';
import 'package:app_web/flutter_layout_article/examples/example5.dart';
import 'package:app_web/flutter_layout_article/examples/example6.dart';
import 'package:app_web/flutter_layout_article/examples/example7.dart';
import 'package:app_web/flutter_layout_article/examples/example8.dart';
import 'package:app_web/flutter_layout_article/examples/example9.dart';
import 'package:app_web/flutter_layout_article/examples/example10.dart';
import 'package:app_web/flutter_layout_article/examples/example11.dart';
import 'package:app_web/flutter_layout_article/examples/example12.dart';
import 'package:app_web/flutter_layout_article/examples/example13.dart';
import 'package:app_web/flutter_layout_article/examples/example14.dart';
import 'package:app_web/flutter_layout_article/examples/example15.dart';
import 'package:app_web/flutter_layout_article/examples/example16.dart';
import 'package:app_web/flutter_layout_article/examples/example17.dart';
import 'package:app_web/flutter_layout_article/examples/example18.dart';
import 'package:app_web/flutter_layout_article/examples/example19.dart';
import 'package:app_web/flutter_layout_article/examples/example20.dart';
import 'package:app_web/flutter_layout_article/examples/example21.dart';
import 'package:app_web/flutter_layout_article/examples/example22.dart';
import 'package:app_web/flutter_layout_article/examples/example23.dart';
import 'package:app_web/flutter_layout_article/examples/example24.dart';
import 'package:app_web/flutter_layout_article/examples/example25.dart';
import 'package:app_web/flutter_layout_article/examples/example26.dart';
import 'package:app_web/flutter_layout_article/examples/example27.dart';
import 'package:app_web/flutter_layout_article/examples/example28.dart';
import 'package:app_web/flutter_layout_article/examples/example29.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return const FlutterLayoutAricle([
      Example1(),
      Example2(),
      Example3(),
      Example4(),
      Example5(),
      Example6(),
      Example7(),
      Example8(),
      Example9(),
      Example10(),
      Example11(),
      Example12(),
      Example13(),
      Example14(),
      Example15(),
      Example16(),
      Example17(),
      Example18(),
      Example19(),
      Example20(),
      Example21(),
      Example22(),
      Example23(),
      Example24(),
      Example25(),
      Example26(),
      Example27(),
      Example28(),
      Example29(),
    ]);
  }
}

class FlutterLayoutAricle extends StatefulWidget {
  final List<Example> examples;

  const FlutterLayoutAricle(this.examples, {Key? key}) : super(key: key);

  @override
  _FlutterLayoutAricleState createState() => _FlutterLayoutAricleState();
}

class _FlutterLayoutAricleState extends State<FlutterLayoutAricle> {
  late int count;
  late Widget example;
  late String code;
  late String explanation;

  @override
  void initState() {
    count = 1;
    code = const Example1().code;
    explanation = const Example1().explanation;

    super.initState();
  }

  @override
  void didUpdateWidget(FlutterLayoutAricle oldWidget) {
    super.didUpdateWidget(oldWidget);
    var example = widget.examples[count - 1];
    code = example.code;
    explanation = example.explanation;
  }

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      title: 'Flutter layout article',
      home: SafeArea(
        child: Material(
          color: Colors.black,
          child: FittedBox(
            child: Container(
              width: 400,
              height: 670,
              color: const Color(0xFFCCCCCC),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.center,
                children: [
                  Expanded(
                    child: ConstrainedBox(
                      constraints: const BoxConstraints.tightFor(
                        height: double.infinity,
                        width: double.infinity,
                      ),
                      child: widget.examples[count - 1],
                    ),
                  ),
                  Container(
                    height: 50,
                    width: double.infinity,
                    color: Colors.black,
                    child: SingleChildScrollView(
                      scrollDirection: Axis.horizontal,
                      child: Row(
                        mainAxisSize: MainAxisSize.min,
                        children: [
                          for (int i = 0; i < widget.examples.length; i++)
                            Container(
                              width: 58,
                              padding: const EdgeInsets.only(left: 4.0),
                              child: button(i + 1),
                            ),
                        ],
                      ),
                    ),
                  ),
                  Container(
                    height: 273,
                    color: Colors.grey[50],
                    child: Scrollbar(
                      child: SingleChildScrollView(
                        key: ValueKey(count),
                        child: Padding(
                          padding: const EdgeInsets.all(10.0),
                          child: Column(
                            children: [
                              Center(child: Text(code)),
                              const SizedBox(
                                height: 13,
                              ),
                              Text(
                                explanation,
                                style: TextStyle(
                                  color: Colors.blue[900],
                                  fontStyle: FontStyle.italic,
                                ),
                              ),
                            ],
                          ),
                        ),
                      ),
                    ),
                  ),
                ],
              ),
            ),
          ),
        ),
      ),
    );
  }

  Widget button(int exampleNumber) {
    return Button(
      key: ValueKey('button$exampleNumber'),
      isSelected: count == exampleNumber,
      exampleNumber: exampleNumber,
      onPressed: () {
        showExample(
          exampleNumber,
          widget.examples[exampleNumber - 1].code,
          widget.examples[exampleNumber - 1].explanation,
        );
      },
    );
  }

  void showExample(int exampleNumber, String code, String explanation) {
    setState(() {
      count = exampleNumber;
      this.code = code;
      this.explanation = explanation;
    });
  }
}
