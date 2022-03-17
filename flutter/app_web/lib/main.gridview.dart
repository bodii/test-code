import "package:flutter/material.dart";
import 'package:flutter/rendering.dart';

void main() {
  debugPaintSizeEnabled = true;
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter demo',
      home: Scaffold(
        appBar: AppBar(
          title: const Text('Grid view learning'),
        ),
        body: _buildGrid(),
      ),
    );
  }
}

Widget _buildGrid() => GridView.extent(
      maxCrossAxisExtent: 150,
      padding: const EdgeInsets.all(4),
      mainAxisSpacing: 4,
      crossAxisSpacing: 4,
      children: _buildGridTileList(30),
    );

List<Container> _buildGridTileList(int count) => List.generate(
      count,
      (i) => Container(
        color: Colors.green[500],
        child: Text('$i', style: const TextStyle(color: Colors.white)),
        alignment: Alignment.center,
      ),
    );
