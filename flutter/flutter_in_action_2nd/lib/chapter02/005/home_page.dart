import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:flutter/services.dart' show rootBundle;

class HomePage extends StatefulWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  State<HomePage> createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('test loading json assets'),
      ),
      body: FutureBuilder(
        future: loadJsonAsset(context),
        builder: (BuildContext context, AsyncSnapshot snapshot) {
          if (snapshot.hasData) {
            Map<String, dynamic> employee =
                json.decode(snapshot.data.toString())['employee'];

            return Column(
              children: [
                Text('name: ' + employee['name'].toString()),
                Text('salary: ' + employee['salary'].toString()),
                Text('married: ' + employee['married'].toString()),
              ],
            );
          } else {
            return const Text('null');
          }
        },
      ),
    );
  }
}

Future<String> loadJsonAsset(BuildContext context) async {
  return await DefaultAssetBundle.of(context)
      .loadString('assets/configs/employee.json');
}
