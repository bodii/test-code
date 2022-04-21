import 'package:flutter/material.dart';
import 'package:flutter_in_action_2nd/chapter02/004/random_words_widget.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(title: const Text('home page - package useing')),
        body: Center(
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              const Text('useing english_words package'),
              const RandomWordsWidget(),
              Padding(
                padding: const EdgeInsets.only(top: 20),
                child: IconButton(
                  onPressed: () => Navigator.of(context).popAndPushNamed('/'),
                  icon: const Icon(Icons.refresh),
                ),
              ),
            ],
          ),
        ));
  }
}
