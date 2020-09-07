# golang-cleanarch-multiple-datastore
Golang CleanArch With Multiple Data Store

<a href="https://goreportcard.com/report/github.com/opannapo/golang-cleanarch-multiple-datastore" target="_blank">
<img src="https://goreportcard.com/badge/github.com/opannapo/golang-cleanarch-multiple-datastore" 
alt="IMAGE ALT TEXT HERE" width="10%" height="10%" border="2" />
</a>

| K | V
| ------ | ------ | 
| Project Structure | Package, Manual DI, Etc
| Firebase Firestore | Realtime Database
| MySql | Database
| Redis | Database / Cache
| Gin | Http Handler, Routing, Endpoint
| Gorm | ORM, Relationship Model, Query
| Viper | App Configuration
| Jwt | Auth, Token, 
| Middleware | Middleware




##Project Structure
```
-root project
	└ app 			 		» Application Project
 	└ config 				» Configuration Files (JSON, Firebase, etc...)
 	└ db 					» Migration Script


- app 						» Application Project 
	└ apis 			 		» Application Layer / Application Business Rules
		└ endpoints 			» Request-Response Controller 
		└ middleware 			» Middleware 
		
 	└ entities 				» Domain Layer
	
 	└ injection				» Injection
 		└ repositories 			» Repository types
 		└ services 			» Service types
		
 	└ repository				» Data Layer
 		└ firestore 			» Firestore Repository
 		└ mysql				» Mysql Repository  
 		└ redis 			» Redis Repository
		
 	└ services				» Enterprise Business Rules
```
