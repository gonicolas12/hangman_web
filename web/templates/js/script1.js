function playAudio(audioID) {
    var audio = document.getElementById(audioID);
    audio.play();
}

document.addEventListener("DOMContentLoaded", function() {
    var audio = document.getElementById("myAudio");
    var musicImage = document.getElementById("musicImage");

    
    toggleButton.addEventListener("click", function() {
        
        if (audio.paused) {
            
            audio.play();
            
            musicImage.src = "./image/musicOn";
            musicImage.alt = "Music On";
        } else {
            
            audio.pause();
        
            musicImage.src = "./image/musicOff.png";
            musicImage.alt = "Music Off";
        }
    });
});