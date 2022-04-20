import 'package:flutter/material.dart';

class TapBoxB extends StatelessWidget {
  const TapBoxB({Key? key, this.active = false, required this.onChanged})
      : super(key: key);

  final bool active;
  final ValueChanged<bool> onChanged;

  void _handleTap() {
    onChanged(!active);
  }

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: _handleTap,
      child: Container(
        child: Center(
          child: Text(
            active ? 'Active' : 'Inactive',
            style: const TextStyle(
              fontSize: 32.0,
              color: Colors.white,
            ),
          ),
        ),
        width: 200.0,
        height: 200.0,
        margin: const EdgeInsets.all(150),
        color: active ? Colors.lightGreen[700] : Colors.grey[600],
      ),
    );
  }
}
