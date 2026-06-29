import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';

class AquaTheme {
  // Colors
  static const Color background = Color(0xFF0a192f);
  static const Color cardColor = Color(0xFF112240);
  static const Color borderColor = Color(0xFF233554);
  static const Color primaryGlow = Color(0xFF64ffda);
  static const Color textLight = Color(0xFFe6f1ff);
  static const Color textMuted = Color(0xFF8892b0);
  static const Color warning = Color(0xFFf59e0b);

  static ThemeData get darkTheme {
    return ThemeData(
      brightness: Brightness.dark,
      scaffoldBackgroundColor: background,
      primaryColor: primaryGlow,
      textTheme: GoogleFonts.interTextTheme().apply(
        bodyColor: textLight,
        displayColor: textLight,
      ),
      cardTheme: CardThemeData(
        color: cardColor,
        elevation: 10,
        shape: RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(20),
          side: const BorderSide(color: borderColor, width: 1),
        ),
      ),
    );
  }
}
