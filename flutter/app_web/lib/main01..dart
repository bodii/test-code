import 'package:flutter/material.dart';

Widget _buildGrid() => GridView.extent(
      maxCrossAxisExtent: 150,
      padding: const EdgeInsets.all(4),
      mainAxisSpacing: 4,
      crossAxisSpacing: 4,
      children: _buildGridTitleList(30),
    );

List<Container> _buildGridTitleList(int count) => List.generate(
      count,
      (i) => Container(
        decoration: BoxDecoration(
          color: Colors.green[500],
        ),
      ),
    );

Widget _buildImageColumn() {
  return Container(
    decoration: const BoxDecoration(
      color: Colors.black26,
    ),
    child: Column(
      children: [
        _buildImageRow(1),
        _buildImageRow(3),
      ],
    ),
  );
}

Widget _buildImageRow(int imageIndex) => Row(
      children: [
        _buildDecoratedImage(imageIndex),
        _buildDecoratedImage(imageIndex + 1),
      ],
    );

Widget _buildDecoratedImage(int imageIndex) => Expanded(
      child: Container(
        decoration: BoxDecoration(
          border: Border.all(width: 10, color: Colors.black38),
          borderRadius: const BorderRadius.all(Radius.circular(8)),
          color: Colors.green[500],
        ),
        margin: const EdgeInsets.all(4),
        child: Center(
          child: Text(
            '$imageIndex',
            textDirection: TextDirection.ltr,
            style: const TextStyle(
              fontSize: 32,
              color: Colors.black87,
            ),
          ),
        ),
      ),
    );

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter layout demo',
      home: Scaffold(
        appBar: AppBar(),
        body: _buildImageColumn(),
      ),
    );
  }
}

void main() => runApp(const MyApp());
