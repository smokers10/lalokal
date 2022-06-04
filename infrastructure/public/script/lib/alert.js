function setup_alert(message, type) {
    // remove alert
    var el = document.getElementsByClassName("alert")
    for (let i = 0; i < el.length; i++) {
        el[i].remove()
    }

    if (type == "loading") {
        $("#alert-place").append(`
            <div class="alert alert-secondary" role="alert">
                <div class="spinner-border spinner-border-sm" role="status"> </div>
                Tunggu sebentar...
            </div>
        `)
    }

    if (type == "success") {
        $("#alert-place").append(`
            <div class="alert alert-success" role="alert">
                ${message}
            </div>
        `)
    }

    if (type == "failed") {
        $("#alert-place").append(`
            <div class="alert alert-danger" role="alert">
                ${message}
            </div>
        `)
    }
}
