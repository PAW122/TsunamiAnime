package modules

import (
	"database/sql"
	"fmt"

	// "log"
	"TsunamiSteram/types"

	_ "github.com/mattn/go-sqlite3"
)

func SaveToDatabase(episode types.Submit_episode_request) error {
	// Otwórz lub stwórz nową bazę danych SQLite
	db, err := sql.Open("sqlite3", "./db/tsunami_stream.db")
	if err != nil {
		return err
	}
	defer db.Close()

	// Utwórz tabelę episodes jeśli nie istnieje
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS episodes (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT,
			episode_number INTEGER,
			episode_title TEXT,
			episode_release_date TEXT,
			hosting TEXT,
			quality TEXT,
			audio_language TEXT,
			text_language TEXT,
			ai_translation BOOLEAN,
			link_embed TEXT,
			translation TEXT
		)
	`)
	if err != nil {
		return err
	}

	// Wstaw dane odcinka do bazy danych
	_, err = db.Exec(`
		INSERT INTO episodes (title, episode_number, episode_title, episode_release_date, hosting, quality, audio_language, text_language, ai_translation, link_embed, translation)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, episode.Title, episode.EpisodeNumber, episode.EpisodeTitle, episode.Episode_relise_date, episode.Hosting, episode.Quality, episode.Audio_Language, episode.Text_Language, episode.Ai_translation, episode.Link_embed, episode.Translation)
	if err != nil {
		return err
	}

	fmt.Println("Episode data saved to database successfully")
	return nil
}

func ReadFromDatabase() ([]types.Submit_episode_request, error) {
	// Otwórz lub stwórz nową bazę danych SQLite
	db, err := sql.Open("sqlite3", "./db/tsunami_stream.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Wykonaj zapytanie SQL, aby pobrać żądania odcinków
	rows, err := db.Query("SELECT title, episode_number, episode_title, episode_release_date, hosting, quality, audio_language, text_language, ai_translation, link_embed, translation FROM episodes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Przeiteruj przez wyniki zapytania i zapisz je do tablicy struktur Submit_episode_request
	var episodeRequests []types.Submit_episode_request
	for rows.Next() {
		var episode types.Submit_episode_request
		err := rows.Scan(&episode.Title, &episode.EpisodeNumber, &episode.EpisodeTitle, &episode.Episode_relise_date, &episode.Hosting, &episode.Quality, &episode.Audio_Language, &episode.Text_Language, &episode.Ai_translation, &episode.Link_embed, &episode.Translation)
		if err != nil {
			return nil, err
		}
		episodeRequests = append(episodeRequests, episode)
	}

	return episodeRequests, nil
}
