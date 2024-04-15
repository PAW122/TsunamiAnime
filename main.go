package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Struktura reprezentująca pojedynczy tytuł anime
type Src struct {
	EpisodeNumber  int    `json:"episodeNumber"`
	Data           string `json:"data"`
	Name           string `json:"Name"`
	Quality        string `json:"quality"`
	AudioLang      string `json:"Audio_lang"`
	TextLang       string `json:"Text_lang"`
	Link           string `json:"link"`
	Ai_translation bool   `json:"Ai_translation"`
}

// Struktura reprezentująca pojedynczy tytuł anime
type AnimeTitle struct {
	ID          string `json:"id"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Categories  string `json:"Categories"`
	Img         string `json:"Img"`
	Season      struct {
		Y int `json:"y"`
		M int `json:"m"`
	} `json:"Season"`
	Rating struct {
		Overall string `json:"overall"`
	} `json:"Rating"`
	Src []Src `json:"src"`
}

// Struktura reprezentująca listę tytułów anime
type AnimeList struct {
	AnimeList []AnimeTitle `json:"anime_list"`
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/listPage.html") // Serwowanie pliku listPage.html
	})

	http.HandleFunc("/addEpisodeRequest", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/addEpReq.html")
	})

	// Endpoint do wczytywania listy tytułów anime
	http.HandleFunc("/titles/", func(w http.ResponseWriter, r *http.Request) {
		// Pobierz ID tytułu z adresu URL
		id := r.URL.Path[len("/titles/"):]

		// Wczytaj dane z pliku JSON w katalogu /db
		data, err := ioutil.ReadFile("./db/anime_list.json")
		if err != nil {
			http.Error(w, "Błąd odczytu danych", http.StatusInternalServerError)
			fmt.Println("Błąd odczytu danych:", err)
			return
		}

		// Wyszukaj odpowiedni tytuł anime w danych JSON
		var animeData map[string]interface{}
		if err := json.Unmarshal(data, &animeData); err != nil {
			http.Error(w, "Błąd dekodowania danych JSON", http.StatusInternalServerError)
			fmt.Println("Błąd dekodowania danych JSON:", err)
			return
		}

		animeList, ok := animeData["anime_list"].([]interface{})
		if !ok {
			http.Error(w, "Błąd przetwarzania danych JSON", http.StatusInternalServerError)
			fmt.Println("Błąd przetwarzania danych JSON")
			return
		}

		var targetAnime map[string]interface{}
		for _, anime := range animeList {
			if a, ok := anime.(map[string]interface{}); ok {
				if a["Id"].(string) == id {
					targetAnime = a
					break
				}
			}
		}

		// Jeśli nie znaleziono tytułu, zwróć 404
		if targetAnime == nil {
			http.NotFound(w, r)
			return
		}

		// Zwróć dane tytułu anime jako odpowiedź HTTP
		jsonResponse, err := json.Marshal(targetAnime)
		if err != nil {
			http.Error(w, "Błąd kodowania danych JSON", http.StatusInternalServerError)
			fmt.Println("Błąd kodowania danych JSON:", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	})

	http.HandleFunc("/watch/", func(w http.ResponseWriter, r *http.Request) {
		// Pobierz ID tytułu anime z adresu URL
		id := r.URL.Path[len("/watch/"):] // Pobierz ostatnią część ścieżki

		// Pobierz nazwę pliku HTML
		filename := "./web/watchPage.html"

		// Wczytaj zawartość pliku HTML
		htmlData, err := ioutil.ReadFile(filename)
		if err != nil {
			http.Error(w, "Błąd odczytu danych", http.StatusInternalServerError)
			fmt.Println("Błąd odczytu danych:", err)
			return
		}

		// Ustaw nagłówki odpowiedzi
		w.Header().Set("Content-Type", "text/html")

		// Zastąp placeholder "{id}" w pliku HTML rzeczywistym ID
		htmlData = bytes.ReplaceAll(htmlData, []byte("{id}"), []byte(id))

		// Zwróć zawartość pliku HTML jako odpowiedź
		w.Write(htmlData)
	})

	// Dodaj endpoint "/all_titles", który zwraca całą listę tytułów anime
	http.HandleFunc("/all_titles", func(w http.ResponseWriter, r *http.Request) {
		// Ustaw nagłówek Content-Type na application/json
		w.Header().Set("Content-Type", "application/json")

		data, err := ioutil.ReadFile("./db/anime_list.json")
		if err != nil {
			http.Error(w, "Błąd odczytu danych", http.StatusInternalServerError)
			fmt.Println("Błąd odczytu danych:", err)
			return
		}

		// Wyszukaj odpowiedni tytuł anime w danych JSON
		var animeData map[string]interface{}
		if err := json.Unmarshal(data, &animeData); err != nil {
			http.Error(w, "Błąd dekodowania danych JSON", http.StatusInternalServerError)
			fmt.Println("Błąd dekodowania danych JSON:", err)
			return
		}

		// Serializuj listę tytułów anime do formatu JSON
		titlesJSON, err := json.Marshal(animeData)
		if err != nil {
			http.Error(w, "Błąd serializacji danych", http.StatusInternalServerError)
			fmt.Println("Błąd serializacji danych:", err)
			return
		}

		// Zwróć listę tytułów anime jako odpowiedź
		w.Write(titlesJSON)
	})

	// Serwowanie plików statycznych
	fs := http.FileServer(http.Dir("./web"))
	http.Handle("/web/", http.StripPrefix("/web/", fs))

	// Uruchomienie serwera na porcie 8080
	fmt.Println("Serwer uruchomiony na http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Błąd podczas uruchamiania serwera:", err)
	}
}
