document.addEventListener('DOMContentLoaded', function() {
    const form = document.getElementById('episode-form');

    form.addEventListener('submit', function(event) {
        event.preventDefault(); // Zapobiegaj domyślnej akcji wysyłania formularza

        // Pobierz wartości z pól formularza
        const title = document.getElementById('title').value;
        const description = document.getElementById('description').value;
        const tags = document.getElementById('tags').value;
        const studio = document.getElementById('studio').value;
        const season = document.getElementById('season').value;
        const img = document.getElementById('img').value;
        // Pobierz kolejne wartości z pól formularza

        // Przygotuj dane do wysłania na serwer
        const requestData = {
            title: title,
            description: description,
            tags: tags,
            studio: studio,
            season: season,
            img: img,
            // Dodaj kolejne dane o odcinku
            // ...
        };

        // Wyślij żądanie POST na serwer z danymi formularza
        fetch('/submit_episode_request', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(requestData),
        })
        .then(response => {
            if (response.ok) {
                return response.json();
            } else {
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
