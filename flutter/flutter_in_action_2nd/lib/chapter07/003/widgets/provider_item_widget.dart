import 'package:flutter/material.dart';
import 'cart_model.dart';
import 'change_notifier_provider.dart';
import 'item.dart';
import 'consumer.dart';

class ProviderItemWidget extends StatefulWidget {
  const ProviderItemWidget({Key? key}) : super(key: key);

  @override
  _ProviderItemWidgetState createState() => _ProviderItemWidgetState();
}

class _ProviderItemWidgetState extends State<ProviderItemWidget> {
  @override
  Widget build(BuildContext context) {
    return Center(
      child: ChangeNotifierProvider<CartModel>(
        data: CartModel(),
        child: Builder(
          builder: (context) {
            return Column(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                Consumer<CartModel>(
                    builder: (context, cart) =>
                        Text('总价：' + cart!.totalPrice.toStringAsFixed(1))),
                Builder(
                  builder: (context) {
                    debugPrint('add item Button build');
                    return ElevatedButton(
                      child: const Text('添加商品'),
                      onPressed: () {
                        ChangeNotifierProvider.of<CartModel>(context,
                                listen: false)
                            .add(Item('', 20.0, 1));
                      },
                    );
                  },
                ),
              ],
            );
          },
        ),
      ),
    );
  }
}
