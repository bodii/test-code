import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:food_delivery_app/configs/colors.dart';
import 'package:food_delivery_app/controllers/home_search_controller.dart';
import 'package:get/get.dart';

class HomeAppBarWidget extends StatelessWidget {
  const HomeAppBarWidget({
    Key? key,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: const EdgeInsets.only(top: 25),
      padding: EdgeInsets.only(
        top: MediaQuery.of(context).padding.top,
        left: 25,
        right: 25,
      ),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: [
          GestureDetector(
            child: Container(
              padding: const EdgeInsets.all(8),
              decoration: BoxDecoration(
                shape: BoxShape.circle,
                color: Colors.grey[100],
              ),
              child: const Icon(Icons.arrow_back),
            ),
            onTap: null,
          ),
          const SearchValueWidget(),
        ],
      ),
    );
  }
}

class SearchValueWidget extends StatefulWidget {
  const SearchValueWidget({
    Key? key,
  }) : super(key: key);

  @override
  State<SearchValueWidget> createState() => _SearchValueWidgetState();
}

class _SearchValueWidgetState extends State<SearchValueWidget> {
  final toSearch = Get.put(HomeSearchController());
  final TextEditingController _inputController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Row(
      children: [
        Offstage(
          offstage: toSearch.getVal(),
          child: Container(
            padding: const EdgeInsets.all(8),
            decoration: BoxDecoration(
              shape: BoxShape.circle,
              color: Colors.grey[100],
            ),
            child: GestureDetector(
              child: const Icon(Icons.search),
              onTap: () {
                setState(() {
                  toSearch.change();
                });
              },
            ),
          ),
        ),
        Offstage(
          offstage: !toSearch.getVal(),
          child: Container(
            width: 160,
            height: 40,
            alignment: Alignment.topLeft,
            decoration: BoxDecoration(
              borderRadius: BorderRadius.circular(20),
              color: Colors.grey[100],
            ),
            child: GestureDetector(
              child: Row(
                children: [
                  Container(
                    alignment: Alignment.centerLeft,
                    margin: const EdgeInsets.only(left: 18),
                    width: 125,
                    height: 40,
                    child: TextField(
                      controller: _inputController,
                      textAlign: TextAlign.left,
                      cursorColor: Colors.black,
                      style: const TextStyle(height: 1, fontSize: 17),
                      cursorHeight: 3.0,
                      decoration: const InputDecoration(
                        hintText: 'Search',
                        hintStyle: TextStyle(),
                        border: InputBorder.none,
                        suffixIcon: Icon(
                          Icons.search,
                          color: Colors.black,
                        ),
                      ),
                    ),
                  ),
                  Container(
                    alignment: Alignment.topRight,
                    child: GestureDetector(
                      child: const Icon(
                        Icons.cancel,
                        size: 17,
                        color: kPrimaryColor,
                      ),
                      onTap: () {
                        setState(() {
                          toSearch.change();
                        });
                      },
                    ),
                  ),
                ],
              ),
              onTap: () {
                setState(() {
                  toSearch.change();
                });
              },
            ),
          ),
        ),
      ],
    );
  }
}
