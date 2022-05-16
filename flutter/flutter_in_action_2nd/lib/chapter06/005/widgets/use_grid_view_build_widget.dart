import 'package:flutter/material.dart';

class UseGridViewBuildWidget extends StatefulWidget {
  const UseGridViewBuildWidget({Key? key}) : super(key: key);

  @override
  State<UseGridViewBuildWidget> createState() => _UseGridViewBuildWidgetState();
}

class _UseGridViewBuildWidgetState extends State<UseGridViewBuildWidget> {
  final List<Icon> _icons = [];

  @override
  void initState() {
    _retrieveIcons();
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return GridView.builder(
      gridDelegate: const SliverGridDelegateWithFixedCrossAxisCount(
        crossAxisCount: 3,
        childAspectRatio: 1.0,
      ),
      itemCount: _icons.length,
      itemBuilder: (context, index) {
        if (index == _icons.length - 1 && _icons.length < 200) {
          _retrieveIcons();
        }
        return _icons[index];
      },
    );
  }

  void _retrieveIcons() {
    Future.delayed(const Duration(milliseconds: 200)).then((e) {
      setState(() {
        _icons.addAll(const [
          Icon(Icons.ac_unit),
          Icon(Icons.airport_shuttle),
          Icon(Icons.all_inclusive),
          Icon(Icons.beach_access),
          Icon(Icons.cake),
          Icon(Icons.free_breakfast),
        ]);
      });
    });
  }
}
