import 'package:flutter/material.dart';

class TapBoxA extends StatefulWidget {
  const TapBoxA({Key? key}) : super(key: key);

  @override
  _TapBoxAState createState() => _TapBoxAState();
}

class _TapBoxAState extends State<TapBoxA> {
  bool _active = false;

  void _handleTap() {
    setState(() {
      _active = !_active;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      width: double.infinity,
      height: double.infinity,
      color: Colors.white,
      child: GestureDetector(
        child: Container(
          margin: const EdgeInsets.all(100),
          width: 100.0,
          height: 100.0,
          child: Center(
            child: Text(
              _active ? 'Active' : 'Inactive',
              style: const TextStyle(
                fontSize: 32.0,
                color: Colors.white,
              ),
            ),
          ),
          color: _active ? Colors.lightGreen[700] : Colors.grey[600],
        ),
        onTap: _handleTap,
      ),
    );
  }
}
