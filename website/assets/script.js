

document.addEventListener("DOMContentLoaded", function () {
    var myModal = new bootstrap.Modal(document.getElementById("myModal"));
    myModal.show();
});

document.addEventListener("DOMContentLoaded", function () {
    convertToHTML()
});

function convertToHTML() {
    breaches = document.getElementById("breaches")
    var children = breaches.children;
    for (var i = 0; i < children.length; i++) {
        // Do stuff
        if (children[i].className == "breach") {
            var parser = new DOMParser();
            var html = parser.parseFromString(children[i].innerHTML, 'text/html');
            children[i].innerHTML = html.body.innerText;
        }
    }
}