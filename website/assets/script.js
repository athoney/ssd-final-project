function emailChange() {
    email = document.querySelector('#emailCheck').checked;
    // console.log("status: " + email)
    if (email) {
        document.getElementById("form").action = "/login/email";
    } else {
        document.getElementById("form").action = "/login";
    }
}