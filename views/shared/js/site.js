$(document).ready(function() {
    $("#home-link").hover(function() {
        $("#fish-pic").addClass("h-swimming")
    },
    function() {
        $("#fish-pic").removeClass("h-swimming")
    });
});