import 'package:flutter/material.dart';

class StatefulBuilder2 extends StatefulWidget {
  const StatefulBuilder2({
    Key? key,
    required this.builder,
  })  : assert(builder != null),
        super(key: key);

  final StatefulWidgetBuilder? builder;

  @override
  _StatefulBuilderState createState() => _StatefulBuilderState();
}

class _StatefulBuilderState extends State<StatefulBuilder2> {
  @override
  Widget build(BuildContext context) => widget.builder!(context, setState);
}
