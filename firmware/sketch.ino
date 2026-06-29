#include <WiFi.h>
#include <PubSubClient.h>
#include <DHTesp.h>
#include <ESP32Servo.h>
#include <ArduinoJson.h>

// Konfigurasi Pin
const int DHT_PIN = 15;
const int LDR_PIN = 34;
const int SERVO_PIN = 13;
const int LED_PIN = 2;

// Konfigurasi WiFi Wokwi
const char* ssid = "Wokwi-GUEST";
const char* password = "";

// Konfigurasi MQTT (Sesuaikan dengan backend jika menggunakan broker lain)
const char* mqtt_server = "broker.emqx.io"; 
const int mqtt_port = 1883;
const char* topic_telemetry = "kolam/aquarium/telemetry";
const char* topic_control = "kolam/aquarium/control";

// Inisialisasi Objek
WiFiClient espClient;
PubSubClient client(espClient);
DHTesp dht;
Servo myServo;

// Variabel Timer (Non-blocking)
unsigned long lastMsg = 0;
const long interval = 5000; // 5 detik

// Variabel State Servo (Non-blocking)
bool isFeeding = false;
unsigned long feedStartTime = 0;
const long feedDuration = 2000; // Tahan di 90 derajat selama 2 detik

void setup_wifi() {
  delay(10);
  Serial.println();
  Serial.print("Connecting to ");
  Serial.println(ssid);

  WiFi.mode(WIFI_STA);
  WiFi.begin(ssid, password);

  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
  }

  Serial.println("");
  Serial.println("WiFi connected");
  Serial.println("IP address: ");
  Serial.println(WiFi.localIP());
}

void callback(char* topic, byte* payload, unsigned int length) {
  Serial.print("Message arrived [");
  Serial.print(topic);
  Serial.print("] ");
  
  String message = "";
  for (int i = 0; i < length; i++) {
    message += (char)payload[i];
  }
  Serial.println(message);

  // Parsing JSON masuk
  StaticJsonDocument<200> doc;
  DeserializationError error = deserializeJson(doc, message);

  if (error) {
    Serial.print("deserializeJson() failed: ");
    Serial.println(error.c_str());
    return;
  }

  const char* command = doc["command"];
  if (String(command) == "feed") {
    Serial.println("Command 'feed' diterima! Menggerakkan servo ke 90 derajat...");
    myServo.write(90);
    isFeeding = true;
    feedStartTime = millis();
  }
}

void reconnect() {
  // Loop until we're reconnected
  while (!client.connected()) {
    Serial.print("Attempting MQTT connection...");
    String clientId = "ESP32Client-Aqua-";
    clientId += String(random(0xffff), HEX);
    
    if (client.connect(clientId.c_str())) {
      Serial.println("connected");
      // Resubscribe
      client.subscribe(topic_control);
      Serial.println("Subscribed to control topic.");
    } else {
      Serial.print("failed, rc=");
      Serial.print(client.state());
      Serial.println(" try again in 5 seconds");
      delay(5000);
    }
  }
}

void setup() {
  Serial.begin(115200);
  
  // Setup Sensor & Aktuator
  dht.setup(DHT_PIN, DHTesp::DHT22);
  pinMode(LDR_PIN, INPUT);
  pinMode(LED_PIN, OUTPUT);
  
  myServo.attach(SERVO_PIN);
  myServo.write(0); // Posisi awal servo (0 derajat)

  setup_wifi();
  client.setServer(mqtt_server, mqtt_port);
  client.setCallback(callback);
}

void loop() {
  if (!client.connected()) {
    reconnect();
  }
  client.loop();

  unsigned long now = millis();

  // Logika Aktuator Servo (Non-Blocking)
  if (isFeeding && (now - feedStartTime >= feedDuration)) {
    myServo.write(0); // Kembalikan ke posisi 0
    isFeeding = false;
    Serial.println("Proses feeding selesai. Servo kembali ke 0 derajat.");
  }

  // Logika Pengiriman Telemetri (Tiap 5 Detik)
  if (now - lastMsg > interval) {
    lastMsg = now;

    // Membaca Sensor
    TempAndHumidity data = dht.getTempAndHumidity();
    int ldrValue = analogRead(LDR_PIN);
    
    // Logika Lokal: LED menyala jika gelap (nilai analog LDR tinggi > 2000)
    if (ldrValue > 2000) {
      digitalWrite(LED_PIN, HIGH);
    } else {
      digitalWrite(LED_PIN, LOW);
    }

    if (dht.getStatus() != 0) {
      Serial.println("Error membaca sensor DHT22!");
      return;
    }

    // Bungkus Data ke JSON
    StaticJsonDocument<200> doc;
    doc["temperature"] = data.temperature;
    doc["humidity"] = data.humidity;
    doc["light"] = ldrValue;

    char jsonBuffer[256];
    serializeJson(doc, jsonBuffer);

    Serial.print("Publishing telemetry: ");
    Serial.println(jsonBuffer);
    client.publish(topic_telemetry, jsonBuffer);
  }
}
