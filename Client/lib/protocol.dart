import 'dart:math';
import 'dart:typed_data';

import 'utils.dart';

class Protocol3Vil {
//   Map processFlag(Uint16List byteFlag) {
//     // TODO: Implement this
//     var flags = {};
//     return flags;
//   }

//   int processSize(Uint16List byteSize) {}

//   Map processRequest() {
//     var byteFlag = obtainFlag();
//     var flags = processFlag(byteFlag);

//     var data = obtainData();

//     return {'flags': flags, 'data': data};
//   }

  Uint8List createPacket(Type packetType, Uint8List data,
      {bool permitted = true, bool internalErr = true}) {
    var buffer = BytesBuilder();

    if ([PacketTypes.data, PacketTypes.hash].contains(packetType)) {
      int accum = 0;

      while (accum < data.length) {
        // Craft component byte Flag: packetType
        var byteFlag = packetType.code << Shifts.shifts[packetType.type]!;

        // Craft component byte Flag: sucessive
        byteFlag += data.length - accum > MAX_SIZE_PACKET ? 1 : 0;

        // Craft byte Size
        var byteSize = min(data.length - accum, MAX_SIZE_PACKET);

        var packetBuff = convInt2Byte(byteFlag, LEN_FLAG) +
            convInt2Byte(byteSize, LEN_SIZE) +
            data.getRange(accum, accum + byteSize).toList();

        buffer.add((Uint8List.fromList(packetBuff)));

        accum += byteSize;
      }
    } else if (packetType == PacketTypes.response) {
      var byteFlag = 0;

      // Craft component byte Flag: responseType
      Type responseType = ResponseTypes.internalErr;
      if (permitted) {
        responseType = ResponseTypes.permitted;
      } else if (!permitted) {
        responseType = ResponseTypes.notPermitted;
      } else if (!internalErr) {
        throw UnimplementedError();
      }
      byteFlag += responseType.code << Shifts.shifts[responseType.type]!;

      // Craft component byte Flag: packetType
      byteFlag += packetType.code << Shifts.shifts[packetType.type]!;

      // Craft byte Size
      var byteSize = 0;

      var packetBuff = convInt2Byte(byteFlag, 1) + convInt2Byte(byteSize, 2);

      buffer.add(Uint8List.fromList(packetBuff));
    }

    return buffer.toBytes();
  }
}
