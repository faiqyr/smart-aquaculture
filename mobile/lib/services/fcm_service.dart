import 'dart:convert';
import 'package:firebase_messaging/firebase_messaging.dart';
import 'package:http/http.dart' as http;
import 'api_service.dart';

class FCMService {
  static final FirebaseMessaging _firebaseMessaging = FirebaseMessaging.instance;

  static Future<void> init() async {
    // Meminta izin notifikasi (untuk iOS, Android 13+)
    NotificationSettings settings = await _firebaseMessaging.requestPermission(
      alert: true,
      badge: true,
      sound: true,
    );

    if (settings.authorizationStatus == AuthorizationStatus.authorized) {
      print('Izin notifikasi diberikan');
      
      // Ambil Token FCM
      String? token = await _firebaseMessaging.getToken();
      if (token != null) {
        print("FCM Token didapatkan: $token");
        await _sendTokenToBackend(token);
      }

      // Dengarkan jika ada token baru
      _firebaseMessaging.onTokenRefresh.listen(_sendTokenToBackend);
    } else {
      print('Izin notifikasi ditolak');
    }
  }

  static Future<void> _sendTokenToBackend(String token) async {
    try {
      final response = await http.post(
        Uri.parse('${ApiService.baseUrl}/fcm/token'),
        headers: {'Content-Type': 'application/json'},
        body: jsonEncode({'token': token}),
      );
      if (response.statusCode == 200) {
        print("Token berhasil didaftarkan ke backend!");
      }
    } catch (e) {
      print("Gagal mendaftarkan token ke backend: $e");
    }
  }
}

// Handler untuk pesan saat aplikasi ditutup (Background)
Future<void> firebaseMessagingBackgroundHandler(RemoteMessage message) async {
  print("Menangani pesan di background: ${message.messageId}");
}
