import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:frontend/models/rest/user.go.dart';
import 'package:frontend/service/rest.dart';
import 'package:http/http.dart' as http;

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Demo',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: const MyHomePage(title: 'Flutter Demo Home Page'),
    );
  }
}

class MyHomePage extends StatefulWidget {
  const MyHomePage({super.key, required this.title});
  final String title;
  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  int _counter = 0;
  late List<User> _user = [
    User(
      id: 0,
      name: "name",
    )
  ];

  void _getUser() async {
    String uri = "http://localhost:4716";
    var response = await http.get(
      Uri.parse(uri),
      headers: {
        'Content-Type': 'application/json; charset=UTF-8',
      },
    );
    Iterable it = [json.decode(response.body)];
    setState(() {
      _user = List<User>.from(
        it.map(
          (e) => User.fromJson(json.decode(e.toString())),
        ),
      );
    });
    print("Body ${it.length} ");
    print("User" + _user.length.toString());
  }

  void _incrementCounter() {
    setState(() {
      _counter++;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(widget.title),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            Text(
              'User is ' + json.encode(_user),
            ),
            const Text(
              'You have pushed the button this many times:',
            ),
            Text(
              '$_counter',
              style: Theme.of(context).textTheme.headlineMedium,
            ),
          ],
        ),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          _incrementCounter();
          _getUser();
        },
        tooltip: 'Increment',
        child: const Icon(Icons.add),
      ),
    );
  }
}
