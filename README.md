# Projet Go

Le but de ce projet va être de réaliser un programme similaire à gobuster. GoBuster est un outil
permettant d’identifier les fichiers et répertoires cachés sur un serveur web. Il fonctionne en
effectuant des requêtes HTTP sur un serveur web, en testant différentes URL et en analysant les
codes de réponse pour déterminer si les fichiers ou répertoires existent (HTTP 200, 301, 404 etc.).

# Getting Start :

Pour utiliser cet outil il faut :
  - Télécharger le script main.go ainsi que la library "wordlist.txt" et le mettre dans un répertoire.
  - Aller sur le répertoire et mettre les droits sur le fichier via la commande "chmod".
  - Utiliser la commande "go run main.go" ou "go build main.go" pour démarrer l'outil.

# Exemples :

![Exemple Screenshot](Images/Exemple-Add.PNG)
![Exemple Screenshot](Images/Exemple2.PNG)

# License :

Ce projet est sous la licence MIT, celle-ci est disponible dans les fichiers du répertoire [MIT License](./LICENSE.txt).

# Fonctionalités :

  - On retrouve plusieurs options :

  ![Exemple Screenshot](Images/Exemple2.PNG)
  
  - Vous pouvez manipuler la stack : afficher, supprimer et dupliquer.
  - Avec la commande 'help', afficher les commandes disponibles ainsi que les opérations.

![Exemple Screenshot](Images/Help.PNG)

# Avertissements : 

Les commandes sont sensibles à la casse. La commande pour faire une addition est 'add', si vous écrivez 'Add' ou encore 'ADD' il y aura une erreur !
