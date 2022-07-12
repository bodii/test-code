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
      home: const MyHomePage(title: 'Flutter Demo Home Page'),
    );
  }
}

class MyHomePage extends StatefulWidget {
  const MyHomePage({Key? key, required this.title}) : super(key: key);

  // This widget is the home page of your application. It is stateful, meaning
  // that it has a State object (defined below) that contains fields that affect
  // how it looks.

  // This class is the configuration for the state. It holds the values (in this
  // case the title) provided by the parent (in this case the App widget) and
  // used by the build method of the State. Fields in a Widget subclass are
  // always marked "final".

  final String title;

  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage>
    with SingleTickerProviderStateMixin {
  List<Snowflake> snowflakes = List.generate(
    300,
    (index) => Snowflake(),
  );

  late AnimationController _controller;

  @override
  void initState() {
    _controller = AnimationController(
      vsync: this,
      duration: const Duration(seconds: 1),
    );
    _controller.repeat();
    super.initState();
  }

  @override
  void dispose() {
    _controller.dispose();
    super.dispose();
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
        title: Text(widget.title),
      ),
      body: Container(
        width: double.infinity,
        height: double.infinity,
        decoration: const BoxDecoration(
          gradient: LinearGradient(
            begin: Alignment.topCenter,
            end: Alignment.bottomCenter,
            colors: [Colors.blue, Colors.lightBlue, Colors.white],
            stops: [0.0, 0.7, 1.0],
          ),
        ),
        child: AnimatedBuilder(
            animation: _controller,
            builder: (context, child) {
              for (Snowflake s in snowflakes) {
                s.fall();
              }
              return CustomPaint(
                painter: MyPainter(snowflakes),
              );
            }),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          setState(() {
            if (_controller.status == AnimationStatus.dismissed) {
              _controller.repeat();
            } else {
              _controller.reset();
            }
          });
        },
        tooltip: 'refresh',
        child: const Icon(Icons.refresh),
      ), // This trailing comma makes auto-formatting nicer for build methods.
    );
  }
}

class MyPainter extends CustomPainter {
  final List<Snowflake> _snowflakes;

  MyPainter(this._snowflakes);

  @override
  void paint(Canvas canvas, Size size) {
    Paint paint = Paint()..color = Colors.white;

    canvas.drawCircle(
      size.center(const Offset(0.0, 110.0)),
      60.0,
      paint,
    );

    canvas.drawOval(
        Rect.fromCenter(
          center: size.center(const Offset(0.0, 290.0)),
          width: 240.0,
          height: 260.0,
        ),
        paint);

    for (var s in _snowflakes) {
      canvas.drawCircle(s.offset, s.radius, paint);
    }
  }

  @override
  bool shouldRepaint(covariant CustomPainter oldDelegate) => true;
}

class Snowflake {
  double radius = Random().nextDouble() * 2 + 1;
  double velocity = Random().nextDouble() * 4 + 2;
  Offset offset = Offset(
    Random().nextDouble() * 600,
    Random().nextDouble() * 10,
  );

  void fall() {
    offset += Offset(0.0, velocity);
    if (offset.dy > 800) {
      offset = Offset(
        Random().nextDouble() * 590 + 10,
        Random().nextDouble() * 10,
      );
      radius = Random().nextDouble() * 2 + 2;
      velocity = Random().nextDouble() * 4 + 2;
    }
  }
}
