import 'package:flutter/material.dart';

class UseThemeWidget extends StatefulWidget {
  const UseThemeWidget({Key? key}) : super(key: key);

  @override
  _UseThemeWidgetState createState() => _UseThemeWidgetState();
}

class _UseThemeWidgetState extends State<UseThemeWidget> {
  Color _themeColor = Colors.teal;

  @override
  Widget build(BuildContext context) {
    ThemeData themeData = Theme.of(context);
    return Theme(
      data: ThemeData(
        primarySwatch: _themeColor as MaterialColor,
        iconTheme: IconThemeData(color: _themeColor),
        textTheme: TextTheme(bodyText2: TextStyle(color: _themeColor)),
      ),
      child: Scaffold(
        appBar: AppBar(
          title: const Text('Theme 主题测试'),
          centerTitle: true,
        ),
        body: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Row(
              mainAxisAlignment: MainAxisAlignment.center,
              children: const [
                Icon(Icons.favorite),
                Icon(Icons.airport_shuttle),
                Text(" 颜色跟随主题"),
              ],
            ),
            Theme(
              data: themeData.copyWith(
                iconTheme: themeData.iconTheme.copyWith(color: Colors.black),
              ),
              child: Row(
                mainAxisAlignment: MainAxisAlignment.center,
                children: const [
                  Icon(Icons.favorite),
                  Icon(Icons.airport_shuttle),
                  Text(" 颜色固定黑色", style: TextStyle(color: Colors.black)),
                ],
              ),
            ),
          ],
        ),
        floatingActionButton: FloatingActionButton(
          child: const Icon(Icons.palette),
          onPressed: () => setState(
            () => _themeColor =
                _themeColor == Colors.teal ? Colors.blue : Colors.teal,
          ),
        ),
      ),
    );
  }
}
