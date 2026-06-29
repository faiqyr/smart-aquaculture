import 'dart:async';
import 'package:flutter/material.dart';
import 'package:lucide_icons/lucide_icons.dart';
import '../services/api_service.dart';
import '../services/fcm_service.dart';
import '../theme.dart';
import '../widgets/metric_card.dart';
import '../widgets/history_chart.dart';

class DashboardScreen extends StatefulWidget {
  const DashboardScreen({Key? key}) : super(key: key);

  @override
  State<DashboardScreen> createState() => _DashboardScreenState();
}

class _DashboardScreenState extends State<DashboardScreen> {
  Map<String, dynamic>? latestData;
  List<dynamic> historyData = [];
  bool isFeeding = false;
  Timer? _timer;

  @override
  void initState() {
    super.initState();
    FCMService.init(); // Meminta izin & mendafarkan token FCM
    _fetchData();
    // Polling every 5 seconds
    _timer = Timer.periodic(const Duration(seconds: 5), (timer) {
      _fetchData();
    });
  }

  @override
  void dispose() {
    _timer?.cancel();
    super.dispose();
  }

  Future<void> _fetchData() async {
    final latest = await ApiService.getLatest();
    final history = await ApiService.getHistory();

    if (mounted) {
      setState(() {
        if (latest != null) latestData = latest;
        if (history.isNotEmpty) historyData = history;
      });
    }
  }

  Future<void> _handleFeed() async {
    setState(() => isFeeding = true);
    
    final success = await ApiService.sendFeedCommand();
    
    if (mounted) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(
          content: Text(
            success ? 'Perintah Pakan Berhasil Dikirim!' : 'Gagal mengirim perintah pakan.',
            style: const TextStyle(color: Colors.white, fontWeight: FontWeight.bold),
          ),
          backgroundColor: success ? Colors.green.shade600 : Colors.red.shade600,
          behavior: SnackBarBehavior.floating,
          shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(10)),
        ),
      );
    }
    
    // Simulate delay to prevent spamming
    await Future.delayed(const Duration(seconds: 2));
    if (mounted) {
      setState(() => isFeeding = false);
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: SafeArea(
        child: RefreshIndicator(
          onRefresh: _fetchData,
          color: AquaTheme.primaryGlow,
          backgroundColor: AquaTheme.cardColor,
          child: SingleChildScrollView(
            physics: const AlwaysScrollableScrollPhysics(),
            child: Padding(
              padding: const EdgeInsets.all(16.0),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  const SizedBox(height: 16),
                  const Text(
                    "Smart Aquaculture",
                    style: TextStyle(
                      fontSize: 32,
                      fontWeight: FontWeight.w900,
                      color: AquaTheme.primaryGlow,
                    ),
                  ),
                  const Text(
                    "Dashboard Otomasi Kolam Ikan",
                    style: TextStyle(
                      fontSize: 16,
                      color: AquaTheme.textMuted,
                    ),
                  ),
                  const SizedBox(height: 32),
                  
                  // Metric Cards
                  MetricCard(
                    title: "Suhu Air",
                    value: latestData != null ? latestData!['Temperature'].toString() : '--',
                    unit: "°C",
                    icon: LucideIcons.thermometer,
                  ),
                  MetricCard(
                    title: "Kelembaban",
                    value: latestData != null ? latestData!['Humidity'].toString() : '--',
                    unit: "%",
                    icon: LucideIcons.droplet,
                  ),
                  MetricCard(
                    title: "Intensitas Cahaya",
                    value: latestData != null ? latestData!['LightIntensity'].toString() : '--',
                    unit: "Lux",
                    icon: LucideIcons.sun,
                  ),
                  
                  const SizedBox(height: 24),
                  
                  // Chart
                  Container(
                    height: 400,
                    decoration: BoxDecoration(
                      color: AquaTheme.cardColor,
                      borderRadius: BorderRadius.circular(20),
                      border: Border.all(color: AquaTheme.borderColor),
                      boxShadow: [
                        BoxShadow(
                          color: Colors.black.withOpacity(0.2),
                          blurRadius: 10,
                          offset: const Offset(0, 5),
                        )
                      ],
                    ),
                    child: HistoryChart(historyData: historyData),
                  ),
                  
                  const SizedBox(height: 100), // Space for FAB
                ],
              ),
            ),
          ),
        ),
      ),
      floatingActionButtonLocation: FloatingActionButtonLocation.centerFloat,
      floatingActionButton: AnimatedContainer(
        duration: const Duration(milliseconds: 300),
        height: 65,
        width: 200,
        child: FloatingActionButton.extended(
          onPressed: isFeeding ? null : _handleFeed,
          backgroundColor: isFeeding ? Colors.grey.shade700 : AquaTheme.primaryGlow,
          elevation: isFeeding ? 0 : 15,
          icon: isFeeding 
            ? const SizedBox(
                width: 24, 
                height: 24, 
                child: CircularProgressIndicator(color: AquaTheme.background, strokeWidth: 3)
              )
            : const Icon(LucideIcons.fish, color: AquaTheme.background),
          label: Text(
            isFeeding ? "Memproses..." : "Beri Pakan",
            style: const TextStyle(
              color: AquaTheme.background, 
              fontWeight: FontWeight.bold,
              fontSize: 18,
            ),
          ),
        ),
      ),
    );
  }
}
