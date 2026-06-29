import 'package:fl_chart/fl_chart.dart';
import 'package:flutter/material.dart';
import '../theme.dart';

class HistoryChart extends StatelessWidget {
  final List<dynamic> historyData;

  const HistoryChart({Key? key, required this.historyData}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    if (historyData.isEmpty) {
      return const Center(
        child: Text(
          "Tidak ada data grafik.",
          style: TextStyle(color: AquaTheme.textMuted),
        ),
      );
    }

    // Proses data (lama ke baru)
    final data = List.from(historyData.reversed);
    
    List<FlSpot> tempSpots = [];
    List<FlSpot> lightSpots = [];

    for (int i = 0; i < data.length; i++) {
      final item = data[i];
      double temp = (item['Temperature'] as num).toDouble();
      // Skalakan cahaya agar muat di grafik yang sama
      double light = (item['LightIntensity'] as num).toDouble() / 100.0;
      
      tempSpots.add(FlSpot(i.toDouble(), temp));
      lightSpots.add(FlSpot(i.toDouble(), light));
    }

    return Padding(
      padding: const EdgeInsets.all(16.0),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          const Text(
            "Tren Sensor Real-Time",
            style: TextStyle(
              fontSize: 18,
              fontWeight: FontWeight.bold,
              color: AquaTheme.textLight,
            ),
          ),
          const SizedBox(height: 24),
          Expanded(
            child: LineChart(
              LineChartData(
                gridData: FlGridData(
                  show: true,
                  drawVerticalLine: false,
                  getDrawingHorizontalLine: (value) => FlLine(
                    color: AquaTheme.borderColor,
                    strokeWidth: 1,
                    dashArray: [5, 5],
                  ),
                ),
                titlesData: FlTitlesData(
                  show: true,
                  topTitles: const AxisTitles(sideTitles: SideTitles(showTitles: false)),
                  rightTitles: const AxisTitles(sideTitles: SideTitles(showTitles: false)),
                  bottomTitles: const AxisTitles(sideTitles: SideTitles(showTitles: false)),
                  leftTitles: AxisTitles(
                    sideTitles: SideTitles(
                      showTitles: true,
                      reservedSize: 40,
                      getTitlesWidget: (value, meta) {
                        return Text(
                          value.toInt().toString(),
                          style: const TextStyle(color: AquaTheme.textMuted, fontSize: 12),
                        );
                      },
                    ),
                  ),
                ),
                borderData: FlBorderData(show: false),
                lineBarsData: [
                  LineChartBarData(
                    spots: tempSpots,
                    isCurved: true,
                    color: AquaTheme.primaryGlow,
                    barWidth: 3,
                    isStrokeCapRound: true,
                    dotData: const FlDotData(show: false),
                    belowBarData: BarAreaData(
                      show: true,
                      color: AquaTheme.primaryGlow.withOpacity(0.1),
                    ),
                  ),
                  LineChartBarData(
                    spots: lightSpots,
                    isCurved: true,
                    color: AquaTheme.warning,
                    barWidth: 3,
                    isStrokeCapRound: true,
                    dotData: const FlDotData(show: false),
                  ),
                ],
              ),
            ),
          ),
          const SizedBox(height: 8),
          const Row(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              Icon(Icons.circle, color: AquaTheme.primaryGlow, size: 10),
              SizedBox(width: 4),
              Text("Suhu (°C)", style: TextStyle(color: AquaTheme.textMuted, fontSize: 12)),
              SizedBox(width: 16),
              Icon(Icons.circle, color: AquaTheme.warning, size: 10),
              SizedBox(width: 4),
              Text("Cahaya (x100 Lux)", style: TextStyle(color: AquaTheme.textMuted, fontSize: 12)),
            ],
          )
        ],
      ),
    );
  }
}
