import 'package:flutter/material.dart';

class MyClipper extends CustomClipper<Rect> {
  @override
  Rect getClip(Size size) => const Rect.fromLTWH(10.0, 15.0, 40.0, 30.0);

  @override
  bool shouldReclip(covariant CustomClipper<Rect> oldClipper) => false;
}

class UseCustomClipper01Widget extends StatelessWidget {
  const UseCustomClipper01Widget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    Image avatar = Image.asset('assets/images/avatar.png');
    return SizedBox(
      width: 40,
      height: 40,
      child: DecoratedBox(
        decoration: const BoxDecoration(color: Colors.red),
        child: ClipRect(
          clipper: MyClipper(),
          child: avatar,
        ),
      ),
    );
  }
}
