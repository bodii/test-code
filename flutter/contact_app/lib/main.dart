import 'dart:convert';
import 'dart:io';

import 'package:flutter/material.dart';
import 'package:flutter/rendering.dart';

void main() {
  // debugPaintSizeEnabled = true;
  runApp(const MyApp());
}

const Color themeColor = Colors.black87;

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
        primaryColorDark: themeColor,
        bottomNavigationBarTheme: const BottomNavigationBarThemeData(
            backgroundColor: themeColor,
            unselectedItemColor: Colors.white38,
            selectedItemColor: Colors.white70),
      ),
      home: const MyHomePage(title: 'phone contact demo'),
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

class _MyHomePageState extends State<MyHomePage> {
  int _counter = 0;

  void _incrementCounter() {
    setState(() {
      // This call to setState tells the Flutter framework that something has
      // changed in this State, which causes it to rerun the build method below
      // so that the display can reflect the updated values. If we changed
      // _counter without calling setState(), then the build method would not be
      // called again, and so nothing would appear to happen.
      _counter++;
    });
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
      drawer: Drawer(
        backgroundColor: Colors.white,
        child: ListView(
          children: [
            DrawerHeader(
              decoration: const BoxDecoration(
                color: Colors.white,
                image: DecorationImage(
                  image: AssetImage('images/01.jpg'),
                  fit: BoxFit.cover,
                ),
              ),
              child: Row(
                children: [
                  const Expanded(
                    child: Center(
                      child: SizedBox(
                        width: 60.0,
                        height: 60.0,
                        child: CircleAvatar(
                          child: Text('R'),
                        ),
                      ),
                    ),
                  ),
                  Container(
                    width: 20.0,
                    height: 20.0,
                    decoration: const BoxDecoration(
                      color: Colors.white,
                      borderRadius: BorderRadius.all(Radius.circular(20)),
                    ),
                    child: IconButton(
                      iconSize: 18.0,
                      onPressed: () {
                        Navigator.pop(context);
                      },
                      icon: const Icon(
                        Icons.close,
                        color: Colors.red,
                      ),
                      tooltip: 'close',
                    ),
                  ),
                ],
              ),
            ),
            const ListTile(
              leading: Icon(Icons.settings),
              title: Text('设置'),
            ),
          ],
        ),
      ),
      appBar: AppBar(
        // leading: const Icon(
        //   Icons.menu,
        //   color: Colors.white38,
        // ),
        // Here we take the value from the MyHomePage object that was created by
        // the App.build method, and use it to set our appbar title.
        title: Text(widget.title),
        backgroundColor: Colors.black87,
      ),
      body: const Call(),
      // floatingActionButton: FloatingActionButton(
      //   onPressed: _incrementCounter,
      //   tooltip: 'Increment',
      //   child: const Icon(Icons.add),
      // ),
      bottomNavigationBar: BottomNavigationBar(
        backgroundColor: themeColor,
        items: [
          BottomNavigationBarItem(
            icon: IconButton(
              onPressed: () {
                Navigator.push(
                  context,
                  MaterialPageRoute(builder: (context) => const Call()),
                );
              },
              icon: const Icon(
                Icons.call,
                color: Colors.white60,
              ),
            ),
            label: '打电话',
          ),
          BottomNavigationBarItem(
            icon: IconButton(
              onPressed: () {
                Navigator.push(
                  context,
                  MaterialPageRoute(builder: (context) => const Person()),
                );
              },
              icon: const Icon(
                Icons.person,
                color: Colors.white60,
              ),
            ),
            label: '联系人',
          ),
          BottomNavigationBarItem(
            icon: IconButton(
              onPressed: () {
                Navigator.push(
                  context,
                  MaterialPageRoute(builder: (context) => const Message()),
                );
              },
              icon: const Icon(
                Icons.message,
                color: Colors.white60,
              ),
            ),
            label: '发短信',
          ),
        ],
      ), // This trailing comma makes auto-formatting nicer for build methods.
    );
  }
}

Future<List> getData() async {
  var url = Uri.parse('https://www.randomuser.me/api/?results=10');
  HttpClient http = HttpClient();
  var request = await http.getUrl(url);
  var response = await request.close();
  var responseBody = await response.transform(utf8.decoder).join();

  List data = json.decode(responseBody);
  return data;
}

class ListBodyContent extends StatelessWidget {
  const ListBodyContent({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      color: Colors.black87,
      child: null,
    );
  }
}

/*

// ignore: prefer_const_constructors
          Container(
            alignment: Alignment.center,
            child: const Icon(
              Icons.account_box,
              color: Colors.white70,
            ),
          ),
          Container(
            alignment: Alignment.center,
            child: Expanded(
              child: Column(
                children: const [
                  Text(
                    '姓名',
                    style:
                        TextStyle(fontSize: 13.0, fontWeight: FontWeight.w800),
                  ),
                  Text('称呼'),
                ],
              ),
            ),
          ),
          Container(
            alignment: Alignment.center,
            child: const Icon(
              Icons.phone,
              color: Colors.white60,
            ),
          ),


*/

class Call extends StatelessWidget {
  const Call({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Center(
      child: Container(
        alignment: Alignment.center,
        color: Colors.black54,
        width: 400,
        height: 400,
        child: Column(
          mainAxisAlignment: MainAxisAlignment.spaceEvenly,
          children: [
            Row(
              mainAxisAlignment: MainAxisAlignment.spaceEvenly,
              children: [
                phoneCallNumber('1'),
                phoneCallNumber('2'),
                phoneCallNumber('3'),
              ],
            ),
            Row(
              mainAxisAlignment: MainAxisAlignment.spaceEvenly,
              children: [
                phoneCallNumber('4'),
                phoneCallNumber('5'),
                phoneCallNumber('6'),
              ],
            ),
            Row(
              mainAxisAlignment: MainAxisAlignment.spaceEvenly,
              children: [
                phoneCallNumber('7'),
                phoneCallNumber('8'),
                phoneCallNumber('9'),
              ],
            ),
            Row(
              mainAxisAlignment: MainAxisAlignment.spaceEvenly,
              children: [
                phoneCallNumber('*'),
                phoneCallNumber('0'),
                phoneCallNumber('#'),
              ],
            ),
          ],
        ),
      ),
    );
  }
}

Widget phoneCallNumber(String number) {
  return Text(
    number,
    style: const TextStyle(
      color: Colors.white,
      fontSize: 22.0,
      fontWeight: FontWeight.w800,
    ),
  );
}

class Person extends StatelessWidget {
  const Person({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return const Text('person page');
  }
}

class Message extends StatelessWidget {
  const Message({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return const Text('message page');
  }
}
