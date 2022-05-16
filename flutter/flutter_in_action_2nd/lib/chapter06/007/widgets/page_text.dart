import 'dart:developer';

import 'package:flutter/material.dart';

class PageText extends StatefulWidget {
  final String text;
  const PageText({Key? key, required this.text}) : super(key: key);

  @override
  _PageState createState() => _PageState();
}

class _PageState extends State<PageText> with AutomaticKeepAliveClientMixin {
  @override
  bool get wantKeepAlive => true;

  @override
  Widget build(BuildContext context) {
    super.build(context); // add
    log("build ${widget.text}");

    return Center(
      child: Text(
        widget.text,
        textScaleFactor: 5,
      ),
    );
  }
}
