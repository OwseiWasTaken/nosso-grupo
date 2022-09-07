//Bibliotecas utilizadas no projeto profissional.
//Parte inical no qual crio um pequeno servidor "Beta" para o Locust.
//IMPORTANTE: o código está pronto, a compilação funciona, mas não testei no hardware ainda
//TODO(1): testar no hardware
//então ele ainda pode ser modificado, falta a parte de HTML,CSS e JS a partir da linha 106

#include <SoftwareSerial.h>
#include <ESP8266WiFi.h>
#include <WiFiManager.h>
#include <ESP8266mDNS.h>
#include <WiFiClient.h>
#include <brzo_i2c.h>
#include <TinyGPS.h>
#include <Wire.h>

//TODO(1): upload esses arquivos p/ o repo
// botar as dependencies no github
// https://docs.github.com/en/code-security/supply-chain-security/understanding-your-software-supply-chain/about-the-dependency-graph#supported-package-ecosystems
#include "SSD1306Brzo.h"
#include "SSD1306Wire.h"
#include "Locust.h"
#include "sitelib.h"
//INVALID check

const int tempoDelay = 0;
bool ShowInvalid = false;

SSD1306Brzo display(0x3c, D3, D5);// Conexão da tela OLED
SoftwareSerial saidaGps (3,1);//RX e TX GPS (conexão)
ESP8266WebServer Locust (80); //Nome do Server
TinyGPS gps1;// nome do GPS

#define DEMO_DURATION 3000 // Variável para a função da logo
typedef void (*Demo)(void);

int demoMode = 0;
int counter = 1;

void drawImageDemo(){
  //Imagem do Locust
  display.drawXbm(34, 14,Locust_Logo_width,Locust_logo_height,Locust_Logo_bits);
}

Demo demos[] = {drawImageDemo};// se quiserem aqui da para colocar mais funções de display :)
int demoLength = (sizeof(demos) / sizeof(Demo));
long timeSinceLastModeSwitch = 0;

//CODE
void handleRoot() { //função que retorna as informações para o site

  //(Abaixo está o código do gps)

  bool recebido = true;

  while (saidaGps.available()) {
    char cIn = saidaGps.read();
    recebido = gps1.encode(cIn);
  }

  if (!recebido) {
   ShowInvalid = true;
  }

  // latitude longitude e idade da informação
  long latitude, longitude;
  unsigned long idadeInfo;
  gps1.get_position(&latitude, &longitude, &idadeInfo);

  if (ShowInvalid) {
    Serial.println ("---------------------");
    if (latitude != TinyGPS::GPS_INVALID_F_ANGLE) {
      Serial.print("Latitude: ");
      Serial.println (float(latitude)/ 100000, 6);
    }
    if (longitude != TinyGPS::GPS_INVALID_F_ANGLE) {
      Serial.print("Longitude: ");
      Serial.println (float (longitude) / 100000, 6);
    }
    if (idadeInfo!= TinyGPS::GPS_INVALID_AGE) {
      Serial.print("idade da informacao(ms): ");
      Serial.println(idadeInfo);
    }
    Serial.println ("---------------------");
  }

  //mais informações como dia e hora
  //TODO(1): UTF-3: corrigir o fuso
  int ano;
  byte mes, dia, hora, minuto, segundo,centesimo;
  gps1.crack_datetime(&ano, &mes, &dia, &hora, &minuto, &segundo, &centesimo, & idadeInfo);

  //altitude
  float altitudeGPS;
  altitudeGPS = gps1.f_altitude();

  if (ShowInvalid) {
    if (altitudeGPS != TinyGPS::GPS_INVALID_ALTITUDE) {
      Serial.print("Altitude (cm): ");
      Serial.println(altitudeGPS);
    }
  }

  //Velocidade
  float velocidade;
  velocidade = gps1.f_speed_kmph();

  //sentido (para onde me desloco)
  unsigned long sentido;
  sentido = gps1.course();

  //satelites e precisão (quantos satelites estão conectados)
  unsigned short satelites;
  unsigned long precisao;
  satelites = gps1.satellites();
  precisao = gps1.hdop();

  if (ShowInvalid) {
    if (satelites !=TinyGPS::GPS_INVALID_SATELLITES) {
      Serial.print("Satelites: ");
      Serial.println(satelites);
    }
    if (precisao != TinyGPS::GPS_INVALID_HDOP) {
      Serial.print("Precisao (centesimos de segundo): ");
      Serial.println(precisao);
    }
  }

  Locust.send(200,"text/html", site());
  }

void handleNotFound() {
  String msg = "Arquivo não encontrado\n\n";
  msg += "URI: ";
  msg += Locust.uri();
  msg += "\nMethod: ";
  msg += (Locust.method() == HTTP_GET) ? "GET":"POST";
  msg += "\nArguments: ";
  msg += Locust.args();
  msg += "\n";
  for (uint8_t i=0; i<Locust.args(); i++) {
    msg += " " + Locust.argName(i) + ": " = Locust.arg(i) + "\n";
  }
  Locust.send(404, "text/plain", msg);
}

void setup(void) {
  //Segunda parte do programa, um gerenciador de redes.
  saidaGps.begin(9600);//velocidade de comunicação do GPS

  display.init();// inicializa o display (tela)
  display.flipScreenVertically();
  display.setFont(ArialMT_Plain_10);

  drawImageDemo();// chama a função da imagem do Locust
  delay(tempoDelay); // a imagem rodará por 5 segundos, 1000 = 1 segundo
  WiFi.mode(WIFI_STA); //permite que o ESP8266 se conecte a uma rede Wi-Fi
  Serial.begin(115200);
  WiFiManager fGen; 
  fGen.resetSettings();
  bool bAp;
  bAp = fGen.autoConnect("LOCUST"); // Mostra o nome da rede (tipo um "Conecta Senac")

  if (!bAp) {
      display.drawString(0,0,"A conexão falhou");
  } else {
      display.drawString(0,0,"Conectado");
  }
  delay(tempoDelay);

  while (WiFi.status() != WL_CONNECTED) {
    delay(tempoDelay);
    display.drawString(0,0,".");
  }

  int IP = (WiFi.localIP()); // depois testar se essa variável não vai interferir no resultado :V
  display.drawString(0,0,"Conectado ao  ");
  delay(tempoDelay);
  display.drawString(0,0,"Endereço IP: ");
  delay(tempoDelay);
  //display.println(IP);
  Serial.println(IP);
  delay(tempoDelay);
  if (MDNS.begin("esp8266")) {
    display.drawString(0,0,"MDNS responder iniciado");
    delay (tempoDelay);
  }
  int setupLongitude = 0;

  Locust.on("/", handleRoot);

  Locust.on("/inline", [](){
    Locust.send(200, "text/plain", "this works as well");
  });

  Locust.onNotFound(handleNotFound);

  Locust.begin();
  display.drawString(0,0,"Server HTTP iniciado");
  delay(tempoDelay);
  Serial.println(IP);//mudar depois
  delay(tempoDelay);
}

void loop() {
  Locust.handleClient();
}
