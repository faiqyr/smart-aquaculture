import 'dart:convert';
import 'package:http/http.dart' as http;
import 'package:flutter/foundation.dart'; // Digunakan untuk cek Web dan TargetPlatform

class ApiService {
  // Gunakan 10.0.2.2 untuk Android Emulator, localhost untuk Web/iOS
  static String get baseUrl {
    if (kIsWeb) {
      return 'http://localhost:3000/api/v1'; // Browser web menggunakan localhost host
    } else if (defaultTargetPlatform == TargetPlatform.android) {
      return 'http://10.0.2.2:3000/api/v1'; // Emulator Android
    }
    return 'http://localhost:3000/api/v1'; // iOS Simulator dll.
  }

  static Future<Map<String, dynamic>?> getLatest() async {
    try {
      final response = await http.get(Uri.parse('$baseUrl/monitoring/latest')).timeout(const Duration(seconds: 3));
      if (response.statusCode == 200) {
        final json = jsonDecode(response.body);
        if (json['success'] == true) return json['data'];
      }
    } catch (e) {
      print('Error fetching latest: $e');
    }
    return null;
  }

  static Future<List<dynamic>> getHistory() async {
    try {
      final response = await http.get(Uri.parse('$baseUrl/monitoring/history')).timeout(const Duration(seconds: 3));
      if (response.statusCode == 200) {
        final json = jsonDecode(response.body);
        if (json['success'] == true && json['data'] != null) {
          return json['data'];
        }
      }
    } catch (e) {
      print('Error fetching history: $e');
    }
    return [];
  }

  static Future<bool> sendFeedCommand() async {
    try {
      final response = await http.post(Uri.parse('$baseUrl/control/feed')).timeout(const Duration(seconds: 5));
      if (response.statusCode == 200) {
        return true;
      }
    } catch (e) {
      print('Error sending feed command: $e');
    }
    return false;
  }
}
