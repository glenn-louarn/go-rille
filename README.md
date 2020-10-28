# go-rille

### Contributor
* Christophe Los Santos Neto

* Malcom Hamelin

* Khuong Vo

* Glenn Louarn

## Installation et démarrage

### Front 

Se placer dans le dossier front et lancer les commandes suivantes : 

* npm install

* npm run serve

Pour build le projet et mettre les fichiers sur un serveur :

* npm run build

### Back
Serveur redis déjà lancé à distance à l'adresse 93.23.6.57:6379
Accessible aussi avec un client redis-client. 

Lancer dans l'ordre en se placant à la racine du projet :
* cmd/api/go_api.exe
* external/go_sub.exe
* internal/go_pub.exe
## Documentation de l'API
Documentation disponible sur Swagger avec la commande suivante toujours 
en se placant à la racine du projet : 
* swagger serve -F=swagger docs\swagger.yaml
## Structure des fichiers
* cmd : Api de l'application
* external : Sub
* front : Projet en Vue.js contenant le front-end
* internal : Pub des valeurs des mesures des capteurs
npm install

npm run serve

npm run build

### Back
Lancer le server redis<br/>
go build 


## Structure des fichiers
* cmb : Api de l'application
* external : Sub
* front : Projet en Vue.js contenant le front-end
* internal : Pub Les capteurs 
