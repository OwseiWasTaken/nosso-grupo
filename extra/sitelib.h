//start

String site(){

String HTML = "";


//pre

	HTML+="<!DOCTYPE html>";
	HTML+="<html lang=\"pt-br\">";
	HTML+="<head>";
	HTML+="	<meta charset=\"UTF-8\">";
	HTML+="	<meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge\">";
	HTML+="	<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">";
	HTML+="	<title>Locust</title>";
	HTML+="	<link rel=\"shortcut icon\" href=\"media/favicon.ico\">";
	HTML+="	<link rel=\"preconnect\" href=\"https://fonts.googleapis.com\">";
	HTML+="	<link rel=\"preconnect\" href=\"https://fonts.gstatic.com\" crossorigin>";
	HTML+="	<link href=\"https://fonts.googleapis.com/css2?family=Noto+Sans&display=swap\" rel=\"stylesheet\">";
	HTML+="	";
	HTML+="	<style>";
	HTML+="		body {";
	HTML+="			background-color: rgb(62, 231, 62);";
	HTML+="		}";
	HTML+="		nav {";
	HTML+="			padding-left: 30px;";
	HTML+="			padding-right: 30px;";
	HTML+="			margin:20px;";
	HTML+="			height: 130px;";
	HTML+="			background-color:white;";
	HTML+="			display: flex;";
	HTML+="			justify-content: space-between;";
	HTML+="			align-items: center;";
	HTML+="			border-radius: 20px;";
	HTML+="			box-shadow:3px 2px 5px darkgreen;";
	HTML+="		}";
	HTML+="		.navimg {";
	HTML+="			width: 150px;";
	HTML+="			height: auto;";
	HTML+="			margin-left: 15px;";
	HTML+="		}";
	HTML+="		.navlink {";
	HTML+="			color: black;";
	HTML+="			font-size: 22px;";
	HTML+="			text-decoration: none;";
	HTML+="			margin: 15px;";
	HTML+="		}";
	HTML+="		";
	HTML+="		div {";
	HTML+="			text-align: center;";
	HTML+="		}";
	HTML+="	";
	HTML+="		.centerimg {";
	HTML+="			width: 350px;";
	HTML+="			border-radius: 20px;";
	HTML+="			margin: 40px;";
	HTML+="			box-shadow: 6px 5px 8px darkgreen;";
	HTML+="			background-color: white;";
	HTML+="		}";
	HTML+="		a{";
	HTML+="			font-family: 'Noto Sans', sans-serif;";
	HTML+="		}";
	HTML+="	</style>";
	HTML+="</head>";
	HTML+="<body>";
	HTML+="	<nav class=\"navegation\">";
	HTML+="		<a class=\"navlink\" href=\"/nosso-grupo-dev/documentos.html\">Documentos</a>";
	HTML+="		<img class=\"navimg\" src=\"https://raw.githubusercontent.com/OwseiWasTaken/nosso-grupo/dev/media/locust.png\" alt=\"locust\">";
	HTML+="		<a class=\"navlink\" href=\"/nosso-grupo-dev/integrantes.html\">Integrantes</a>";
	HTML+="	</nav>";
	HTML+="	<div>";
	HTML+="		<figure id=\"center\">";
	HTML+="			<img class=\"centerimg\" src=\"https://raw.githubusercontent.com/OwseiWasTaken/nosso-grupo/dev/media/icones-de-localisation-de-la-carte-verte.png\" alt=\"conectado\">";
	HTML+="		</figure>";
	HTML+="	</div>";
	HTML+="</body>";
	HTML+="</html>";

//post

	return HTML;
}

//end
