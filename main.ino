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
#include "Locust.h"
#include "sitelib.h"
//INVALID check



//COMP
//#include <string.h>

const int tempoDelay = 5000;
const int minDelay = 3000;
const int ledBlink = 5;

#define posX 21
#define posY 0


SSD1306Brzo display(0x3c, D3, D5);// Conexão da tela OLED
SoftwareSerial saidaGps (3,1);//RX e TX GPS (conexão)
ESP8266WebServer Locust (80); //Nome do Server
TinyGPS gps1;// nome do GPS
void imagem (){
   display.drawXbm(posX, posY,Locust_Logo_width,Locust_logo_height,Locust_Logo_bits);
   display.display();
}

void dPrintln(){
  display.display();
  delay (tempoDelay);
  display.clear();
}

void dString(){
   display.display();
   delay (minDelay);
   display.clear();
}

//CODE
void handleIntegrantes (){
  Locust.send(200,"text/html", integrantes());
}

void handleRoot() { //função que retorna as informações para o site
  /*
  COMP
  if strcmp(Locust.uri(), "/integrantes.html") {
    Locust.send(200,"text/html", integrantes());
    // 'return' para finalizar a função
    return
  }
  */

  //(Abaixo está o código do gps)
  display.setLogBuffer(10, 40);

  bool recebido = true;

  while (saidaGps.available()) {
    char cIn = saidaGps.read();
    recebido = gps1.encode(cIn);
  }



  // latitude longitude e idade da informação
  long latitude, longitude;
  unsigned long idadeInfo;
  gps1.get_position(&latitude, &longitude, &idadeInfo);

  if (!recebido){
     pinMode(ledBlink, OUTPUT);
     digitalWrite(ledBlink, HIGH);
     delay(1000);
     digitalWrite(ledBlink, LOW);
     delay(1000);
  }

  if (recebido) {
      display.drawString (0,0,"---------------------");
      dString();
  if (latitude != TinyGPS::GPS_INVALID_F_ANGLE) {
      display.drawString (0,0,"Latitude: ");
      dString();
      display.println (float(latitude)/ 100000, 6);
      display.drawLogBuffer(2,1);
      dPrintln();
   }
  if (longitude != TinyGPS::GPS_INVALID_F_ANGLE) {
      display.drawString(0,0,"Longitude: ");
      dString();
      display.println (float (longitude) / 100000, 6);
      display.drawLogBuffer (2,2);
      dPrintln();

   }
  if (idadeInfo!= TinyGPS::GPS_INVALID_AGE) {
     display.drawString(0,0,"idade da informacao(ms): ");
     dString();
     display.println(idadeInfo);
     display.drawLogBuffer (2,3);
     dPrintln();
   }
    display.drawString (0,0,"---------------------");
    dString();

  //mais informações como dia e hora
  //TODO(1): UTF-3: corrigir o fuso
  int ano;
  byte mes, dia, hora, minuto, segundo,centesimo;
  gps1.crack_datetime(&ano, &mes, &dia, &hora, &minuto, &segundo, &centesimo, & idadeInfo);

  //altitude
  float altitudeGPS;
  altitudeGPS = gps1.f_altitude();

  if (altitudeGPS != TinyGPS::GPS_INVALID_ALTITUDE) {
     display.drawString(0,0,"Altitude (cm): ");
     dString();
     display.println(altitudeGPS);
     display.drawLogBuffer (2,4);
     dPrintln();
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

  if (satelites !=TinyGPS::GPS_INVALID_SATELLITES) {
    display.drawString(0,0,"Satelites: ");
    dString();
    display.println(satelites);
    display.drawLogBuffer (2,5);
    dPrintln();
  }
  if (precisao != TinyGPS::GPS_INVALID_HDOP) {
    display.drawString(0,0,"Precisao (centesimos de segundo): ");
    dString();
    display.println(precisao);
    display.drawLogBuffer (2,6);
    dPrintln();
  }
  Locust.send(200,"text/html", site());
  }
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

  Serial.begin(115200);
  Serial.println();
  Serial.println();


  // Initialising the UI will init the display too.
  display.init();

  display.flipScreenVertically();
  display.setFont(ArialMT_Plain_16);


  imagem();
  delay(tempoDelay);
  display.clear();

  //delay(tempoDelay); // a imagem rodará por 5 segundos, 1000 = 1 segundo
  WiFi.mode(WIFI_STA); //permite que o ESP8266 se conecte a uma rede Wi-Fi
  WiFiManager fGen;
  fGen.resetSettings();
  bool bAp;
  bAp = fGen.autoConnect("LOCUST"); // Mostra o nome da rede (tipo um "Conecta Senac")

  if (!bAp) {
      display.drawString(0,0,"A conexao falhou");
      dString();
  } else {
      display.drawString(0,0,"Conectado");
      dString();
  }
  delay(tempoDelay);

  while (WiFi.status() != WL_CONNECTED) {
    delay(tempoDelay);
    display.drawString(0,0,"...");
    display.display();
    delay (10);
    display.clear();
  }

 // int IP = (WiFi.localIP()); // depois testar se essa variável não vai interferir no resultado :V
  display.drawString(0,0,"Conectado ao  ");
  dString();
  display.drawString(0,0,"Endereço IP: ");
  dString();
  //display.println(IP);
  display.setLogBuffer(5, 30);
  display.println(WiFi.localIP());
  display.drawLogBuffer(0,0);
  display.display();
  delay(tempoDelay);
  display.clear();
  if (MDNS.begin("esp8266")) {
    display.drawString(0,0,"MDNS iniciado");
    dString();
  }
  int setupLongitude = 0;

  Locust.on("/", handleRoot);

  //Locust.on("/integrantes",handleIntegrantes);

  Locust.onNotFound(handleIntegrantes);

  Locust.begin();
  display.drawString(0,0,"Server iniciado");
  dString();
  display.println(WiFi.localIP());//mudar depois
  display.drawLogBuffer (1,0);
  display.display();
  delay(tempoDelay);
  display.clear();
}

void loop() {
  Locust.handleClient();
}
