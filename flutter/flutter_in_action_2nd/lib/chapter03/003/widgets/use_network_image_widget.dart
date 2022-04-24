import 'package:flutter/material.dart';

class UseNetworkImageWidget extends StatelessWidget {
  const UseNetworkImageWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Row(
      mainAxisAlignment: MainAxisAlignment.spaceAround,
      children: [
        const Image(
          image: NetworkImage(
              'https://cliply.co/wp-content/uploads/2020/08/442008112_GLANCING_AVATAR_3D_400px.gif'),
          width: 100.0,
        ),
        Image.network(
          'https://cliply.co/wp-content/uploads/2020/08/442008112_GLANCING_AVATAR_3D_400px.gif',
          width: 100.0,
        ),
      ],
    );
  }
}
