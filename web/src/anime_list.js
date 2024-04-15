fetch("/all_titles")
    .then(response => response.json())
    .then(data => {
        const animeList = data.anime_list; // Pobierz listę tytułów anime

        // Posortuj listę tytułów anime malejąco względem roku i miesiąca
        animeList.sort((a, b) => {
            // Porównaj rok i miesiąc
            if (a.Season.y !== b.Season.y) {
                return b.Season.y - a.Season.y; // Sortuj malejąco względem roku
            } else {
                return b.Season.m - a.Season.m; // Jeśli rok jest taki sam, sortuj malejąco względem miesiąca
            }
        });

        // Utwórz kontener dla tytułów anime
        const titlesContainer = document.getElementById("titles-container");

        // Iteruj po każdym tytule anime i utwórz kafelki
        animeList.forEach(animeTitle => {
            // Utwórz kafelek dla tytułu anime
            const titleTile = document.createElement("div");
            titleTile.classList.add("title-tile");
            
            // Utwórz kontener dla obrazu i opisu
            const contentContainer = document.createElement("div");
            contentContainer.classList.add("content-container");
            
            // Utwórz element obrazu
            const imageElement = document.createElement("img");
            imageElement.src = animeTitle.Img;
            imageElement.alt = animeTitle.Title;
            
            // Utwórz element opisu
            const descriptionElement = document.createElement("p");
            descriptionElement.textContent = animeTitle.Description;
            
            // Dodaj elementy obrazu i opisu do kontenera
            contentContainer.appendChild(imageElement);
            contentContainer.appendChild(descriptionElement);

            // Ustaw zawartość kafelka
            titleTile.appendChild(contentContainer);
            titleTile.innerHTML += `
                <h2>${animeTitle.Title}</h2>
                <p>Categories: ${animeTitle.Categories}</p>
                <p>Rating: ${animeTitle.Rating.overall}</p>
            `;

            // Utwórz kontener dla źródeł
            const srcContainer = document.createElement("div");
            srcContainer.classList.add("src-container");

            // Dodaj kontener źródeł do kafelka tytułu anime
            titleTile.appendChild(srcContainer);
            titleTile.setAttribute("data-id", animeTitle.Id);

            // Dodaj odnośnik do strony /watch/id dla kafelka
            titleTile.addEventListener("click", function() {
                // Pobierz id z atrybutu data-id
                const id = this.getAttribute("data-id");
                // Przejdź do odpowiedniego adresu URL
                window.location.href = `/watch/${id}`;
            });

            // Dodaj kafelek do kontenera tytułów
            titlesContainer.appendChild(titleTile);
        });
    })
    .catch(error => console.error('Error:', error));
