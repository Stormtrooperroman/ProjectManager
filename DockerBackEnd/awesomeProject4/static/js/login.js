$("#send").click(function (e) {
    let name = $("#inputLogin").val();
    let pass = $("#inputPassword").val();
    let send_data = JSON.stringify({
        login: name,
        pass: pass
    })
    $.ajax({
        type: "POST",
        url: "../api/login",
        data: send_data,
        contentType: "application/json",
        dataType: "json",
        success: function () {
            console.log("AHAHHAHAHAHAH")
            location.reload()
        }
    });
    
});