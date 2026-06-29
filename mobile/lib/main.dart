import 'package:flutter/material.dart';
import 'package:firebase_core/firebase_core.dart';
import 'package:firebase_messaging/firebase_messaging.dart';
import 'screens/dashboard_screen.dart';
import 'services/fcm_service.dart';
import 'theme.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  
  // Inisialisasi Firebase (harus ada file google-services.json)
  try {
    await Firebase.initializeApp();
    FirebaseMessaging.onBackgroundMessage(firebaseMessagingBackgroundHandler);
  } catch (e) {
    print("Firebase init error (Mungkin belum ada google-services.json): $e");
  }
  
  runApp(const SmartAquacultureApp());
}

class SmartAquacultureApp extends StatelessWidget {
  const SmartAquacultureApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Smart Aquaculture',
      debugShowCheckedModeBanner: false,
      theme: AquaTheme.darkTheme,
      home: const DashboardScreen(),
    );
  }
}
