import 'dart:core';
import 'dart:io';

import 'package:eviltrans_fe/data/media.dart';
import 'package:eviltrans_fe/utils.dart';
import 'package:flutter/material.dart';

class Viewer extends StatefulWidget {
  const Viewer({super.key});

  @override
  State<Viewer> createState() => _ViewerState();
}

class _ViewerState extends State<Viewer> {
  final List<String> tabsName = ["Image", "Video"];
  Future<Map<MediaType, List<String>>>? resources;

  @override
  void initState() {
    resources = getResources();

    super.initState();
  }

  Widget getGridView(MediaType mediaType) {
    return FutureBuilder(
      future: resources,
      builder: (context, snapshot) => snapshot.hasData
          ? GridView.count(
              crossAxisCount: 2,
              children: snapshot.data![mediaType]!
                  .map((e) => Container(
                        padding: const EdgeInsets.all(10),
                        margin: const EdgeInsets.all(10),
                        decoration: BoxDecoration(
                          color: Colors.amberAccent,
                          borderRadius: BorderRadius.circular(10),
                        ),
                        child: Image.file(File(e)),
                      ))
                  .toList(),
            )
          : const SizedBox(
              width: 60,
              height: 60,
              child: CircularProgressIndicator(),
            ),
    );
  }

  @override
  Widget build(BuildContext context) {
    return DefaultTabController(
      length: tabsName.length,
      child: Scaffold(
        appBar: AppBar(
          leading: IconButton(
            icon: const Icon(Icons.menu),
            onPressed: () {
              // TODO: HoangLe [Mar-07]: Implement this
            },
          ),
          title: const Text(
            "Viewer",
            style: TextStyle(
              fontWeight: FontWeight.bold,
              fontSize: 20,
            ),
          ),
          actions: [
            IconButton(
              onPressed: () {},
              icon: const Icon(Icons.settings),
            )
          ],
          bottom: TabBar(
            tabs: tabsName.map((e) => Tab(text: e)).toList(),
          ),
        ),
        body: Container(
          padding: const EdgeInsets.all(10),
          child: TabBarView(
            children: [
              getGridView(MediaType.image),
              getGridView(MediaType.thumbnail),
              // GridView.count(
              //   crossAxisCount: 2,
              //   children: imagePaths!
              //       .map(
              //         (e) => Container(
              //           padding: const EdgeInsets.all(10),
              //           margin: const EdgeInsets.all(10),
              //           decoration: BoxDecoration(
              //             color: Colors.amberAccent,
              //             borderRadius: BorderRadius.circular(10),
              //           ),
              //           child: Image(
              //             image: AssetImage(e),
              //           ),
              //         ),
              //       )
              //       .toList(),
              // ),
            ],
          ),
        ),
      ),
    );
  }
}
