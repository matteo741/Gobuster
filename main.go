package main

import (
"fmt"
"os"
"flag"
"sync"
"net/http"
"bufio"
)

func main() {

// Afficher la bannière lors du démarrage du tool :

showBanner() 

// Options disponibles :

LibraryPath := flag.String("d", "", "Veuillez entrer le chemin vers votre library.")
QuietMode := flag.Bool("q", false, "Mode silencieux : Permet d'afficher seulement les urls retournant un code HTTP 200.")
Target := flag.String("t", "", "Veuillez entrer l'URL ou l'adresse IP du site web ciblé.")
Workers := flag.Int("w", 1, "Affichage du nombre de Workers")
Help := flag.Bool("h", false, "Afficher l'aide.")

// Analyse des drapeaux : 
flag.Parse()

// Vérification des options : 

if *Help {
    flag.Usage()
    os.Exit(0)
}

if *LibraryPath == "" || *Target == "" {
fmt.Println("Le chemin du dictionnaire et l'URL cible sont requis")
os.Exit(1)
}

// Affichage de la Target, de la Library et des Workers: 

fmt.Printf("\x1b[32mCible: %s\nChemin du dictionnaire: %s\nNombre de workers: %d\x1b[0m\n\n", *Target, *LibraryPath, *Workers)

// Lecture du chemin de la Library : 

paths, err := readLines(*LibraryPath)
if err != nil {
fmt.Printf("Le chemin de la Library n'est pas reconnu. Veuillez entrer le bon chemin: %v\n", err)
os.Exit(1)
}

// Initialisation du WaitGroup (permet d'attendre que tous les Workers se terminent) :

var wg sync.WaitGroup

// Créer un canal (permet la communication entre les Workers) :

pathChan := make(chan string)

// Création de plusieurs Workers : 

for i := 0; i < *Workers; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        for path := range pathChan {
            // Construire l'URL à scanner (le symbole "+" est la concaténation entre les différents éléments de l'URL)

            url := *Target + "/" + path

            // Effectuer une requête HTTP GET vers l'adresse

            resp, err := http.Get(url)
            if err != nil {
                fmt.Printf("\nErreur lors de la requête vers %s: %v\n", url, err)
                continue
            }

            // Afficher le résultat si le mode silencieux est désactivé ou si le statut est 200 OK

            if !*QuietMode || resp.StatusCode == http.StatusOK {
                fmt.Printf("%s - %d\n", url, resp.StatusCode)
            }

            resp.Body.Close()
        }
    }()
}


// Envoyer chaque chemin à check :

for _, path := range paths {
    pathChan <- path
}

// Ferme le canal :

close(pathChan)

// Attendre la fin des Workers :

wg.Wait()
}

// Fonction qui lit les lignes d'un fichier :

func readLines(path string) ([]string, error) {
    // Ouvrir le fichier :
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close() // Fermer le fichier lorsque la fonction se termine

    var lines []string
    // Créer un scanner pour parcourir le fichier ligne par ligne :
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    // Vérifier s'il y a eu des erreurs lors de la lecture du fichier :
    if err := scanner.Err(); err != nil {
        return nil, err
    }
    return lines, nil // Retourner les lignes lues
}

// Bannière

func showBanner() {
fmt.Println(`
 @@@@@@@  @@@@@@@   @@@@@@      @@@ @@@@@@@@ @@@@@@@       @@@@@@@   @@@@@@ 
 @@!  @@@ @@!  @@@ @@!  @@@     @@! @@!        @@!        !@@       @@!  @@@
 @!@@!@!  @!@!!@!  @!@  !@!     !!@ @!!!:!     @!!        !@! @!@!@ @!@  !@!
 !!:      !!: :!!  !!:  !!! .  .!!  !!:        !!:        :!!   !!: !!:  !!!
  :        :   : :  : :. :  ::.::   : :: :::    :          :: :: :   : :. :                                                                                
`)
}