function togglepopup(){
    document.getElementById("popup-1").classList.toggle("active");
}

function toggleCredits() {
    var popup = document.getElementById("popup-credits");
    popup.classList.toggle("active");
}

function playAudio(audioID) {
    var audio = document.getElementById(audioID);
    audio.play();
}

document.addEventListener("DOMContentLoaded", function() {
    var audio = document.getElementById("myAudio");
    var musicImage = document.getElementById("musicImage");

    toggleButton.addEventListener("click", function() {
        // Vérification de l'état de lecture audio actuel
        if (audio.paused) {
            
            audio.play();
        
            musicImage.src = "./image/musicOn.png";
            musicImage.alt = "Music On";
        } else {
            // Si la lecture est en cours, mettre en pause l'audio
            audio.pause();
            // Changer l'image en musique désactivée
            musicImage.src = "./image/musicOff.png";
            musicImage.alt = "Music Off";
        }
    });
});


  