import 'dart:io';

import 'dart:typed_data';
import 'utils.dart';

void sendData(Uint8List data) async {
  // connect to the socket server
  final socket = await Socket.connect(HOST, PORT);
  print('Connected to: ${socket.remoteAddress.address}:${socket.remotePort}');

  // listen for responses from the server
  socket.listen(
    // handle data from the server
    (Uint8List data) {
      final serverResponse = String.fromCharCodes(data);
      print('Server: $serverResponse');
    },

    // handle errors
    onError: (error) {
      print(error);
      socket.destroy();
    },

    // handle server ending connection
    onDone: () {
      print('Server left.');
      socket.destroy();
    },
  );

  socket.add(data);

  await socket.close();
}

void readData(Socket socket, path) async {
  File file;

  if (path != null) {
    file = File(path);

    var futureBytes = file.readAsBytes();
    futureBytes.then((Uint8List value) {
      print("Total bytes: ${value.length}");
      // var tmp = Uint8List(8)
      //   ..buffer.asByteData().setInt64(0, value.length, Endian.big);
      // socket.add(tmp);

      socket.add(value);
    });
  }
}
