import 'package:flutter/material.dart';

class UseImageBoxFitWidget extends StatelessWidget {
  const UseImageBoxFitWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    AssetImage imgPath = const AssetImage('assets/images/avatar.png');
    return SingleChildScrollView(
      child: Column(
        children: [
          Image(image: imgPath, height: 50, width: 100, fit: BoxFit.fill),
          Image(image: imgPath, width: 50.0, height: 50.0, fit: BoxFit.contain),
          Image(image: imgPath, width: 100, height: 50.0, fit: BoxFit.cover),
          Image(
              image: imgPath, width: 100.0, height: 50.0, fit: BoxFit.fitWidth),
          Image(
              image: imgPath,
              width: 100.0,
              height: 50.0,
              fit: BoxFit.fitHeight),
          Image(
              image: imgPath,
              width: 100.0,
              height: 50.0,
              fit: BoxFit.scaleDown),
          Image(image: imgPath, height: 100.0, width: 50.0, fit: BoxFit.none),
          Image(
              image: imgPath,
              width: 100,
              color: Colors.blue,
              colorBlendMode: BlendMode.difference),
          Image(
              image: imgPath,
              width: 100.0,
              height: 100.0,
              repeat: ImageRepeat.repeatY),
        ].map((e) {
          return Row(
            children: [
              Padding(
                padding: const EdgeInsets.all(6.0),
                child: SizedBox(
                  width: 100,
                  child: e,
                ),
              ),
              Text(e.fit.toString())
            ],
          );
        }).toList(),
      ),
    );
  }
}
