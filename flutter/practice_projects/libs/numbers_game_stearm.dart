import 'dart:async';
import 'dart:math';

import 'package:flutter/material.dart';


void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Demo',
      debugShowCheckedModeBanner: false,
      theme: ThemeData(
        // This is the theme of your application.
        //
        // Try running your application with "flutter run". You'll see the
        // application has a blue toolbar. Then, without quitting the app, try
        // changing the primarySwatch below to Colors.green and then invoke
        // "hot reload" (press "r" in the console where you ran "flutter run",
        // or simply save your changes to "hot reload" in a Flutter IDE).
        // Notice that the counter didn't reset back to zero; the application
        // is not restarted.
        primarySwatch: Colors.blue,
      ),
      home: const MyHomePage(),
    );
  }
}

class MyHomePage extends StatefulWidget {
  const MyHomePage({Key? key}) : super(key: key);

  // This widget is the home page of your application. It is stateful, meaning
  // that it has a State object (defined below) that contains fields that affect
  // how it looks.

  // This class is the configuration for the state. It holds the values (in this
  // case the title) provided by the parent (in this case the App widget) and
  // used by the build method of the State. Fields in a Widget subclass are
  // always marked "final".

  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage>
    with SingleTickerProviderStateMixin {
  StreamController<int> inputStream = StreamController.broadcast();
  StreamController<int> scoreStream = StreamController.broadcast();

  @override
  void initState() {
    scoreStream.stream.listen((event) {
      debugPrint(event.toString());
    });
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    // This method is rerun every time setState is called, for instance as done
    // by the _incrementCounter method above.
    //
    // The Flutter framework has been optimized to make rerunning build methods
    // fast, so that you can just rebuild anything that needs updating rather
    // than having to individually change instances of widgets.
    return Scaffold(
      appBar: AppBar(
        // Here we take the value from the MyHomePage object that was created by
        // the App.build method, and use it to set our appbar title.
        title: Center(
          child: StreamBuilder(
            stream: scoreStream.stream.transform<int>(TallyTransformer()),
            builder: ((context, AsyncSnapshot<int> snapshot) {
              if (snapshot.hasData) {
                return Text('score: ${snapshot.data}');
              }
              return const Text('score: 0');
            }),
          ),
        ),
      ),
      body: Stack(
        children: [
          ...List.generate(
            5,
            (index) => Topic(
              inputStream: inputStream,
              scoreStream: scoreStream,
            ),
          ),
          Align(
            alignment: Alignment.bottomCenter,
            child: KeyPad(
              stream: inputStream,
              scoreStream: scoreStream,
            ),
          ),
        ],
      ),
    );
  }
}

class TallyTransformer implements StreamTransformer<int, int> {
  int sum = 0;
  final StreamController<int> _controller = StreamController();

  @override
  Stream<int> bind(Stream<int> stream) {
    stream.listen((event) {
      sum += event;
      _controller.add(sum);
    });

    return _controller.stream;
  }

  @override
  StreamTransformer<RS, RT> cast<RS, RT>() => StreamTransformer.castFrom(this);
}

class Topic extends StatefulWidget {
  Topic({
    Key? key,
    required this.inputStream,
    required this.scoreStream,
  }) : super(key: key);

  StreamController<int> inputStream;
  StreamController<int> scoreStream;

  @override
  _TopicState createState() => _TopicState();
}

class _TopicState extends State<Topic> with SingleTickerProviderStateMixin {
  late int a, b;
  late double x;
  late Color color;
  late AnimationController _controller;

  @override
  void initState() {
    _controller = AnimationController(
      vsync: this,
    );

    reset();

    _controller.addStatusListener((status) {
      if (status == AnimationStatus.completed) {
        reset();
      }
    });

    widget.inputStream.stream.listen((event) {
      if (event == a + b) {
        reset();
        widget.scoreStream.add(5);
      }
    });

    super.initState();
  }

  @override
  void dispose() {
    _controller.dispose();
    super.dispose();
  }

  void reset() {
    a = Random().nextInt(5) + 1;
    b = Random().nextInt(4);
    x = Random().nextDouble() * 350;
    color = Colors.primaries[Random().nextInt(Colors.primaries.length)];
    _controller.duration = const Duration(seconds: 10);
    _controller.forward(from: Random().nextDouble());
  }

  @override
  Widget build(BuildContext context) {
    return AnimatedBuilder(
      animation: _controller,
      builder: (context, child) {
        return Positioned(
          top: _controller.value * 700 - 100,
          left: x,
          child: Container(
            width: 55.0,
            height: 28.0,
            decoration: BoxDecoration(
              color: color.withOpacity(.3),
              borderRadius: BorderRadius.circular(30.0),
              border: Border.all(color: Colors.black87, width: 1),
            ),
            child: Center(
              child: Text(
                '$a+$b',
                style: const TextStyle(fontSize: 22.0),
              ),
            ),
          ),
        );
      },
    );
  }
}

class KeyPad extends StatelessWidget {
  KeyPad({
    Key? key,
    required this.stream,
    required this.scoreStream,
  }) : super(key: key);

  StreamController<int> stream;
  StreamController<int> scoreStream;

  @override
  Widget build(BuildContext context) {
    return GridView.count(
      crossAxisCount: 3,
      childAspectRatio: 5 / 3,
      shrinkWrap: true,
      physics: const NeverScrollableScrollPhysics(),
      padding: const EdgeInsets.all(0.0),
      children: List.generate(9, (index) {
        final int val = index + 1;
        return GestureDetector(
          child: Container(
            color: Colors.primaries[index][300],
            child: Center(
              child: Text(
                '$val',
                textAlign: TextAlign.center,
                style: const TextStyle(
                  fontSize: 18.0,
                ),
              ),
            ),
          ),
          onTap: () {
            stream.add(val);
            scoreStream.add(-2);
          },
        );
      }),
    );
  }
}
