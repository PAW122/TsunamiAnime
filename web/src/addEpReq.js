document.addEventListener('DOMContentLoaded', function() {
    const form = document.getElementById('episode-form');

    form.addEventListener('submit', function(event) {
        event.preventDefault(); // Zapobiegaj domyślnej akcji wysyłania formularza

        // Pobierz wartości z pól formularza
        const title = document.getElementById('title').value;
        const episodeNumber = parseInt(document.getElementById('episodeNumber').value);
        const episodeTitle = document.getElementById('episodeTitle').value;
        const episode_relise_date = document.getElementById('episode_relise_date').value;
        const hosting = document.getElementById('hosting').value;
        const quality = document.getElementById('quality').value;
        const audio_Language = document.getElementById('audio_Language').value;
        const text_Language = document.getElementById('text_Language').value;
        const Ai_translation = document.getElementById('Ai_translation').checked;
        const link_embed = document.getElementById('link_embed').value;
        const translation = document.getElementById('translation').value;

        // Przygotuj dane do wysłania na serwer
        const requestData = {
            title: title,
            episodeNumber: episodeNumber,
            episodeTitle: episodeTitle,
            episode_relise_date: episode_relise_date,
            hosting: hosting,
            quality: quality,
            audio_Language: audio_Language,
            text_Language: text_Language,
            Ai_translation: Ai_translation,
            link_embed: link_embed,
            translation: translation
        };

        console.log(requestData)

        // Wyślij żądanie POST na serwer z danymi formularza
        fetch('http://localhost:8080/submit_episode_request', { // Zmień na właściwy adres URL
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(requestData),
        })
        .then(response => {
            if (!response.ok) {
                throw new Error('Error during submission');
            }
        })
        .then(data => {
            // Obsłuż odpowiedź od serwera
            console.log(data);
            alert('Episode request submitted successfully!');
            // Możesz dodać więcej logiki tutaj, np. przekierowanie użytkownika
        })
        .catch(error => {
            console.error('Error:', error);
            alert('Error during submission. Please try again later.');
        });
    });
});
